package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
)

// 获取全部订单，分页获取
func GetAllOrder(ctx *gin.Context) {
	var params model.QueryParamsWithDate
	err := ctx.ShouldBind(&params)
	res, err := service.GetAllOrder(&params)
	if err != nil {
		global.Logrus.Error("订单获取错误" + err.Error())
		response.Fail("订单获取错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("全部订单获取成功", res, ctx)
}

// 获取订单统计
func GetMonthOrderStatistics(ctx *gin.Context) {
	var params model.QueryParamsWithDate
	err := ctx.ShouldBind(&params)
	res, err := service.GetMonthOrderStatistics(&params)
	if err != nil {
		global.Logrus.Error("获取订单统计错误" + err.Error())
		response.Fail("获取订单统计错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("获取订单统计成功", res, ctx)
}

// 获取用户订单by user id，显示用户最近10条订单
func GetOrderByUserID(ctx *gin.Context) {
	uID, ok := ctx.Get("uID")
	if !ok {
		response.Fail("订单获取，uID参数错误", nil, ctx)
		return
	}
	uIDInt := uID.(int)
	var params = model.PaginationParams{PageSize: 10} //显示用户最近10条订单
	//err := ctx.ShouldBind(&params)
	res, err := service.GetOrderByUserID(uIDInt, &params)
	if err != nil {
		global.Logrus.Error("获取订单 error:", err)
		response.Fail("订单获取错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("订单获取成功", res, ctx)
}

// 完成未支付订单
func CompletedOrder(ctx *gin.Context) {
	var order model.Orders
	err := ctx.ShouldBind(&order)
	if err != nil {
		global.Logrus.Error("完成未支付订单参数错误:", err)
		response.Fail("完成未支付订单参数错误"+err.Error(), nil, ctx)
		return
	}
	order.TradeStatus = "completed"   //更新数据库订单状态,自定义结束状态completed
	err = service.UpdateOrder(&order) //更新数据库状态
	if err != nil {
		global.Logrus.Error("更新数据库状态错误:", err)
		response.Fail("更新数据库状态错误"+err.Error(), nil, ctx)
		return
	}
	err = service.UpdateUserSubscribe(&order) //更新用户订阅信息
	if err != nil {
		global.Logrus.Error("更新用户订阅信息错误:", err)
		response.Fail("更新用户订阅信息错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("完成未支付订单成功", nil, ctx)

}
