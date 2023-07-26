package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
)

// 主题配置
func GetThemeConfig(ctx *gin.Context) {
	//local cache
	//theme, ok := global.LocalCache.Get("theme")
	//if ok && theme != nil {
	//	response.OK("主题获取成功", theme, ctx)
	//	return
	//}
	//theme, err := service.GetThemeConfig()
	//if err != nil {
	//	response.Fail("主题获取错误"+err.Error(), nil, ctx)
	//	return
	//}
	//global.LocalCache.SetNoExpire("theme", theme)
	//response.OK("主题获取成功", theme, ctx)
	response.OK("主题获取成功", global.Theme, ctx)

}
func UpdateThemeConfig(ctx *gin.Context) {

	var theme model.Theme
	err := ctx.ShouldBind(&theme)
	if err != nil {
		global.Logrus.Error("主题设置参数错误:", err)
		response.Fail("主题设置参数错误"+err.Error(), nil, ctx)
		return
	}
	err = service.UpdateThemeConfig(&theme)
	if err != nil {
		global.Logrus.Error("设置主题 error:", err)
		response.Fail("主题设置错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("主题设置成功", nil, ctx)

}

// 获取系统设置
func GetSetting(ctx *gin.Context) {
	res, err := service.GetSetting()
	if err != nil {
		global.Logrus.Error("系统设置获取错误:", err.Error())
		response.Fail("系统设置获取错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("系统设置获取成功", res, ctx)

}

// 获取公共系统设置
func GetPublicSetting(ctx *gin.Context) {
	res, err := service.GetPublicSetting()
	if err != nil {
		global.Logrus.Error("系统设置获取错误:", err.Error())
		response.Fail("系统设置获取错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("系统设置获取成功", res, ctx)
}

func UpdateSetting(ctx *gin.Context) {
	var setting model.Server
	err := ctx.ShouldBind(&setting)
	if err != nil {
		global.Logrus.Error("更改系统设置参数错误:", err.Error())
		response.Fail("更改系统设置参数错误"+err.Error(), nil, ctx)
		return
	}
	err = service.UpdateSetting(&setting)
	if err != nil {
		global.Logrus.Error("更改系统设置错误:", err.Error())
		response.Fail("更改系统设置错误"+err.Error(), nil, ctx)
		return
	}
	response.OK("更改系统设置成功", nil, ctx)

}
