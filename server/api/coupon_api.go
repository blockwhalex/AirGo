package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
)

// 新建折扣
func NewCoupon(ctx *gin.Context) {
	var coupon model.Coupon
	err := ctx.ShouldBind(&coupon)
	if err != nil {
		global.Logrus.Error("新建折扣参数错误", err.Error())
		response.Fail("新建折扣参数错误"+err.Error(), nil, ctx)
		return
	}
	//fmt.Println("新建折扣", coupon)
	err = service.NewCoupon(coupon)
	if err != nil {
		global.Logrus.Error("新建折扣错误", err.Error())
		response.Fail("新建折扣错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("新建折扣成功", nil, ctx)

}

// 删除折扣
func DeleteCoupon(ctx *gin.Context) {
	var coupon model.Coupon
	err := ctx.ShouldBind(&coupon)
	if err != nil {
		global.Logrus.Error("删除折扣参数错误", err.Error())
		response.Fail("删除折扣参数错误"+err.Error(), nil, ctx)
		return
	}
	err = service.DeleteCoupon(coupon)
	if err != nil {
		global.Logrus.Error("删除折扣错误", err.Error())
		response.Fail("删除折扣错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("删除折扣成功", nil, ctx)

}

// 更新折扣
func UpdateCoupon(ctx *gin.Context) {
	var coupon model.Coupon
	err := ctx.ShouldBind(&coupon)
	if err != nil {
		global.Logrus.Error("更新折扣参数错误", err.Error())
		response.Fail("更新折扣参数错误"+err.Error(), nil, ctx)
		return
	}
	err = service.UpdateCoupon(coupon)
	if err != nil {
		global.Logrus.Error("更新折扣错误", err.Error())
		response.Fail("更新折扣错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("更新折扣成功", nil, ctx)

}

// 获取折扣列表
func GetCoupon(ctx *gin.Context) {
	couponArr, err := service.GetCoupon()
	if err != nil {
		global.Logrus.Error("获取折扣列表错误", err.Error())
		response.Fail("获取折扣列表错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("获取折扣列表成功", couponArr, ctx)
}

// 验证折扣是否有效
func VerifyCoupon(ctx *gin.Context) {

}
