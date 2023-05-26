package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/alipay_plugin"
	"AirGo/utils/other_plugin"
	"AirGo/utils/response"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	alipay "github.com/smartwalle/alipay/v3"
)

// 支付预创建
func PreCreatePay(ctx *gin.Context) {
	//res, _ := ioutil.ReadAll(ctx.Request.Body)
	//fmt.Println("支付预创建 preOrder:", string(res))
	uID, _ := ctx.Get("uID")
	uName, _ := ctx.Get("uName")
	uIDInt := uID.(int)
	uNameStr := uName.(string)
	//前端传过来 goods_id
	var preOrder model.Orders
	err := ctx.ShouldBind(&preOrder)
	if err != nil {
		response.Fail("支付预创建参数错误"+err.Error(), nil, ctx)
		return
	}
	//通过商品id查找商品
	goods, err := service.FindGoodsByGoodsID(preOrder.GoodsID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Fail("商品不存在"+err.Error(), nil, ctx)
			return
		} else {
			response.Fail("商品查询错误"+err.Error(), nil, ctx)
			return
		}
	}
	//构造系统订单参数
	uIDStr := other_plugin.Sup(uIDInt, 5)
	//uIDStr := strconv.Itoa(preOrder.UserID)
	preOrder.OutTradeNo = strconv.FormatInt(time.Now().UnixNano(), 10) + uIDStr
	preOrder.Subject = goods.Subject
	preOrder.TotalAmount = goods.TotalAmount
	preOrder.Price = goods.TotalAmount
	preOrder.UserID = uIDInt
	preOrder.UserName = uNameStr
	//创建系统订单
	order, err := service.CreateOrder(&preOrder)
	if err != nil {
		response.Fail("创建系统订单err:"+err.Error(), nil, ctx)
		return
	}
	response.OK("创建系统订单成功", order, ctx)

}

func Purchase(ctx *gin.Context) {
	//判断订单是否合理，未写
	var order model.Orders
	err := ctx.ShouldBind(&order)
	if err != nil || order.OutTradeNo == "" {
		response.Fail("订单参数获取错误"+err.Error(), nil, ctx)
		return
	}
	//判断支付方式
	switch order.PayType {
	case "alipay":
		res, err := alipay_plugin.TradePreCreatePay(global.AlipayClient, &order)
		//fmt.Println("AlipayTradePreCreatePay res:", res)
		if err != nil || res.Content.QRCode == "" {
			response.Fail("alipay TradePreCreatePay err"+err.Error(), nil, ctx)
			return
		}
		//返回用户qrcode
		order.QRCode = res.Content.QRCode
		service.UpdateOrder(&order) //更新数据库状态
		response.OK("alipay TradePreCreatePay success:", res.Content.QRCode, ctx)
		//5分钟等待付款
		go PollAliPay(&order)
	}

}
func PollAliPay(order *model.Orders) {
	t := time.NewTicker(10 * time.Second)
	//defer t.Stop()
	i := 0
	for {
		<-t.C
		rsp, _ := alipay_plugin.TradeQuery(global.AlipayClient, order)
		//fmt.Println("支付宝TradeQuery rsp:", rsp)
		if rsp.Content.TradeStatus == "WAIT_BUYER_PAY" && order.TradeStatus != "WAIT_BUYER_PAY" { //支付宝预创建成功，等待支付，
			order.TradeStatus = "WAIT_BUYER_PAY"
			order.TradeNo = rsp.Content.TradeNo
			err := service.UpdateOrder(order) //更新数据库状态
			if err != nil {
				return
			}
			continue
		} else if rsp.Content.SubCode == "ACQ.TRADE_NOT_EXIST" || rsp.Content.SubCode == "ACQ.SYSTEM_ERROR" {
			continue
		} else if rsp.Content.TradeStatus == "TRADE_SUCCESS" || rsp.Content.TradeStatus == "TRADE_FINISHED" { //交易结束
			fmt.Println("支付宝支付成功")
			order.TradeStatus = "TRADE_SUCCESS" //更新数据库订单状态
			order.BuyerLogonId = rsp.Content.BuyerLogonId
			order.ReceiptAmount = rsp.Content.ReceiptAmount
			order.BuyerPayAmount = rsp.Content.BuyerPayAmount
			service.UpdateOrder(order)         //更新数据库状态
			service.UpdateUserSubscribe(order) //更新用户订阅信息
			return
		}
		if i == 30 {
			if order.TradeNo != "" {
				res, _ := alipay_plugin.TradeClose(global.AlipayClient, order) //超时，取消订单
				fmt.Println("支付宝取消订单结果:", res)
			}
			order.TradeStatus = "TRADE_CLOSED" //更新数据库订单状态(超时已取消)
			service.UpdateOrder(order)         //更新数据库状态
			t.Stop()
			return
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
	//fmt.Println("回调参数更新数据库订单:", sysOrder)
	err := service.UpdateOrder(sysOrder)
	if err != nil && noti.TradeStatus == "TRADE_SUCCESS" {
		log.Println("更新数据库订单错误")
		return
	}
	// 确认收到通知消息
	alipay.ACKNotification(ctx.Writer)
	//更新用户订阅信息
	service.UpdateUserSubscribe(sysOrder) //更新用户订阅信息
}

// 查询商品
func FindGoods(ctx *gin.Context) {

}

// 查询全部商品
func GetAllGoods(ctx *gin.Context) {
	goodsArr, err := service.GetAllGoods()
	if err != nil {
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
		fmt.Println("新建商品参数err", err)
		response.Fail("新建商品参数错误"+err.Error(), nil, ctx)
		return
	}
	fmt.Println("新建商品goods", goods)
	err = service.NewGoods(&goods)
	if err != nil {
		response.Fail("新建商品错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("新建商品成功", nil, ctx)
}

// 删除商品
func DeleteGoods(ctx *gin.Context) {
	// rq, err := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("新建商品参数", string(rq))
	// fmt.Println("新建商品参数err", err)
	var goods model.Goods
	err := ctx.ShouldBind(&goods)
	if err != nil {
		response.Fail("删除商品参数错误"+err.Error(), nil, ctx)
		return
	}
	//fmt.Println("删除商品goods", goods)
	err = service.DeleteGoods(&goods)
	if err != nil {
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
		response.Fail("更新商品参数错误"+err.Error(), nil, ctx)
		return
	}
	fmt.Println("更新商品goods", goods)
	err = service.UpdateGoods(&goods)
	if err != nil {
		response.Fail("更新商品错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("更新商品成功", nil, ctx)
}
