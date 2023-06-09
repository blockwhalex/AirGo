package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/alipay_plugin"
	"AirGo/utils/other_plugin"
	"AirGo/utils/response"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	alipay "github.com/smartwalle/alipay/v3"
)

// 订单预处理
func PreHandleOrder(ctx *gin.Context) *model.Orders {
	uID, _ := ctx.Get("uID")
	uName, _ := ctx.Get("uName")
	uIDInt := uID.(int)
	uNameStr := uName.(string)

	var receiveOrder model.Orders
	err := ctx.ShouldBind(&receiveOrder) //前端传过来 goods_id
	if err != nil {
		global.Logrus.Error("订单预处理参数错误", err.Error())
		response.Fail("订单预处理参数错误"+err.Error(), nil, ctx)
		return nil
	}
	//通过商品id查找商品
	goods, err := service.FindGoodsByGoodsID(receiveOrder.GoodsID)
	if err != nil {
		global.Logrus.Error("通过商品id查找商品error:", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//response.Fail("商品不存在"+err.Error(), nil, ctx)
			return nil
		}
	}
	//构造系统订单参数
	uIDStr := other_plugin.Sup(uIDInt, 5)
	//uIDStr := strconv.Itoa(preOrder.UserID)
	receiveOrder.GoodsID = goods.ID
	receiveOrder.OutTradeNo = strconv.FormatInt(time.Now().UnixNano(), 10) + uIDStr
	receiveOrder.Subject = goods.Subject
	receiveOrder.TotalAmount = goods.TotalAmount
	receiveOrder.Price = goods.TotalAmount //优惠码
	receiveOrder.UserID = uIDInt
	receiveOrder.UserName = uNameStr
	receiveOrder.PayType = receiveOrder.PayType //添加付款方式

	return &receiveOrder
}

// 获取订单
func GetOrderInfo(ctx *gin.Context) {
	order := PreHandleOrder(ctx)
	if order == nil {
		response.Fail("获取订单详情错误", nil, ctx)
		return
	}
	response.OK("订单详情", order, ctx)
}

// 支付预创建
func PreCreatePay(ctx *gin.Context) {
	order := PreHandleOrder(ctx)
	if order == nil {

		response.Fail("获取订单错误", nil, ctx)
		return
	}
	//创建系统订单
	order, err := service.CreateOrder(order)
	if err != nil {
		global.Logrus.Error("创建系统订单error:", err.Error())
		response.Fail("创建系统订单error:"+err.Error(), nil, ctx)
		return
	}

	response.OK("创建系统订单成功", order, ctx)

}

func Purchase(ctx *gin.Context) {
	uID, _ := ctx.Get("uID")
	uIDInt := uID.(int)
	var receiveOrder model.Orders
	err := ctx.ShouldBind(&receiveOrder)
	if err != nil || receiveOrder.OutTradeNo == "" {
		response.Fail("订单参数获取错误", nil, ctx)
		return
	}
	//根据订单号查询订单
	receiveOrder.UserID = uIDInt //确认user id
	sysOrder, err := service.GetOrderByOrderID(&receiveOrder)
	if err != nil {
		global.Logrus.Error("根据订单号查询订单error:", err.Error())
		if err == gorm.ErrRecordNotFound {
			response.Fail("订单不存在"+err.Error(), nil, ctx)
			return
		} else {
			response.Fail("订单查询错误"+err.Error(), nil, ctx)
			return
		}
	}
	sysOrder.PayType = receiveOrder.PayType
	//如果价格=0，跳过支付
	if sysOrder.TotalAmount == "0" {
		sysOrder.TradeStatus = "completed"    //更新数据库订单状态,自定义结束状态completed
		service.UpdateOrder(sysOrder)         //更新数据库状态
		service.UpdateUserSubscribe(sysOrder) //更新用户订阅信息
		response.OK("购买成功", nil, ctx)
		return
	}
	//判断支付方式
	switch sysOrder.PayType {
	case "alipay":
		res, err := alipay_plugin.TradePreCreatePay(global.AlipayClient, sysOrder)
		//fmt.Println("AlipayTradePreCreatePay res:", res)
		if err != nil || res.QRCode == "" {
			response.Fail("alipay TradePreCreatePay err"+err.Error(), nil, ctx)
			return
		}
		sysOrder.QRCode = res.QRCode
		sysOrder.TradeStatus = "WAIT_BUYER_PAY"
		service.UpdateOrder(sysOrder)                                     //更新数据库状态
		response.OK("alipay TradePreCreatePay success:", res.QRCode, ctx) //返回用户qrcode
		//5分钟等待付款
		go PollAliPay(sysOrder)
	}

}
func PollAliPay(order *model.Orders) {
	t := time.NewTicker(10 * time.Second)
	//defer t.Stop()
	i := 0
	for {
		if i == 18 { // 18*10s 3分钟
			if order.TradeNo != "" {
				res, _ := alipay_plugin.TradeClose(global.AlipayClient, order) //超时，取消订单
				//fmt.Println("支付宝取消订单结果:", res)
				global.Logrus.Error("支付宝取消订单结果:", res)
			}
			order.TradeStatus = "TRADE_CLOSED" //更新数据库订单状态(超时已取消)
			service.UpdateOrder(order)         //更新数据库状态
			t.Stop()
			return
		}
		<-t.C
		rsp, _ := alipay_plugin.TradeQuery(global.AlipayClient, order)
		//fmt.Println("支付宝TradeQuery rsp.Content.TradeStatus:", rsp.TradeStatus)
		if rsp.TradeStatus == "TRADE_SUCCESS" || rsp.TradeStatus == "TRADE_FINISHED" { //交易结束
			//fmt.Println("支付宝支付成功")
			order.TradeStatus = "TRADE_SUCCESS"
			order.BuyerLogonId = rsp.BuyerLogonId
			order.ReceiptAmount = rsp.ReceiptAmount
			order.BuyerPayAmount = rsp.BuyerPayAmount
			service.UpdateOrder(order)         //更新数据库状态
			service.UpdateUserSubscribe(order) //更新用户订阅信息
			t.Stop()
			return
		}
		if rsp.TradeStatus == "WAIT_BUYER_PAY" && order.TradeStatus != "WAIT_BUYER_PAY" {
			order.TradeNo = rsp.TradeNo
			service.UpdateOrder(order) //更新数据库状态
		}
		i++
	}
}

func PollWeChatPay() {

}
func AlipayNotify(ctx *gin.Context) {
	noti, _ := global.AlipayClient.GetTradeNotification(ctx.Request)
	if noti == nil {
		return
	}
	//根据回调参数更新数据库订单
	var sysOrder = &model.Orders{
		//OutTradeNo:     noti.OutTradeNo,
		//TradeNo:        noti.TradeNo,
		BuyerLogonId:   noti.BuyerLogonId,
		TradeStatus:    string(noti.TradeStatus),
		TotalAmount:    noti.TotalAmount,
		ReceiptAmount:  noti.ReceiptAmount,
		BuyerPayAmount: noti.BuyerPayAmount,
		//Subject:        noti.Subject, // 商品的标题
	}
	err := service.UpdateOrder(sysOrder)
	if err != nil && noti.TradeStatus == "TRADE_SUCCESS" {
		global.Logrus.Error("更新数据库订单错误", err.Error())
		return
	}
	// 确认收到通知消息
	alipay.ACKNotification(ctx.Writer)
	//更新用户订阅信息
	service.UpdateUserSubscribe(sysOrder) //更新用户订阅信息
}

// 查询全部已启用商品
func GetAllEnabledGoods(ctx *gin.Context) {
	goodsArr, err := service.GetAllEnabledGoods()
	if err != nil {
		global.Logrus.Error("查询全部商品错误:", err.Error())
		response.Fail("查询全部商品错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("查询全部商品成功", goodsArr, ctx)
}

// 查询全部商品
func GetAllGoods(ctx *gin.Context) {
	goodsArr, err := service.GetAllGoods()
	if err != nil {
		global.Logrus.Error("查询全部商品错误:", err.Error())
		response.Fail("查询全部商品错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("查询全部商品成功", goodsArr, ctx)

}

// 新建商品
func NewGoods(ctx *gin.Context) {
	var goods model.Goods
	err := ctx.ShouldBind(&goods)
	if err != nil {
		global.Logrus.Error("新建商品参数err", err)
		response.Fail("新建商品参数错误"+err.Error(), nil, ctx)
		return
	}
	err = service.NewGoods(&goods)
	if err != nil {
		global.Logrus.Error("新建商品错误:", err.Error())
		response.Fail("新建商品错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("新建商品成功", nil, ctx)
}

// 删除商品
func DeleteGoods(ctx *gin.Context) {
	var goods model.Goods
	err := ctx.ShouldBind(&goods)
	if err != nil {
		global.Logrus.Error("删除商品参数错误:", err.Error())
		response.Fail("删除商品参数错误"+err.Error(), nil, ctx)
		return
	}
	err = service.DeleteGoods(&goods)
	if err != nil {
		global.Logrus.Error("删除商品错误:", err.Error())
		response.Fail("删除商品错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("删除商品成功", nil, ctx)

}

// 更新商品
func UpdateGoods(ctx *gin.Context) {
	var goods model.Goods
	err := ctx.ShouldBind(&goods)
	if err != nil {
		global.Logrus.Error("更新商品参数错误:", err.Error())
		response.Fail("更新商品参数错误"+err.Error(), nil, ctx)
		return
	}
	err = service.UpdateGoods(&goods)
	if err != nil {
		global.Logrus.Error("更新商品错误:", err.Error())
		response.Fail("更新商品错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("更新商品成功", nil, ctx)
}
