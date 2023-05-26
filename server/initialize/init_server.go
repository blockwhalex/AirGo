package initialize

import (
	"AirGo/global"
	"AirGo/service"
	"fmt"
)

// 系统配置
func InitServer() {
	res, err := service.GetSetting()
	if err != nil {
		fmt.Println("系统配置获取失败")
		return
	}
	global.Server = *res
}
