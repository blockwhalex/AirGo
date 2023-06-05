package initialize

import (
	"AirGo/global"
	"AirGo/service"
)

// 系统配置
func InitServer() {
	res, err := service.GetSetting()
	if err != nil {
		global.Logrus.Error("系统配置获取失败", err.Error())
		return
	}
	global.Server = *res
}
