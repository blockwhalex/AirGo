package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
)

func UpdateThemeConfig(ctx *gin.Context) {

	var theme model.Theme
	err := ctx.ShouldBind(&theme)
	if err != nil {
		response.Fail("", nil, ctx)
		return
	}
	err = service.UpdateThemeConfig(&theme)
	if err != nil {
		global.Logrus.Error("设置主题 err:", err)
		response.Fail("主题设置错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("主题设置成功", nil, ctx)

}

func GetSetting(ctx *gin.Context) {
	res, err := service.GetSetting()
	if err != nil {
		response.Fail("系统设置获取错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("系统设置获取成功", res, ctx)

}

func UpdateSetting(ctx *gin.Context) {
	var setting model.Server
	err := ctx.ShouldBind(&setting)
	if err != nil {
		response.Fail("更改系统设置参数错误"+err.Error(), nil, ctx)
		return
	}
	err = service.UpdateSetting(&setting)
	if err != nil {
		response.Fail("更改系统设置错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("更改系统设置成功", nil, ctx)

}
