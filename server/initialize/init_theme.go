package initialize

import (
	"AirGo/global"
	"AirGo/service"
)

// 系统配置
func InitTheme() {
	res, err := service.GetThemeConfig()
	if err != nil {
		global.Logrus.Error("系统配置获取失败", err.Error())
		return
	}
	global.Theme = *res
}
