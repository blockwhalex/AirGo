package initialize

import (
	"AirGo/global"
	"AirGo/service"
	"fmt"
)

// 系统配置
func InitTheme() {
	res, err := service.GetThemeConfig()
	if err != nil {
		fmt.Println("系统配置获取失败")
		return
	}
	global.Theme = *res
	//fmt.Println("global.theme:", global.Theme)
}
