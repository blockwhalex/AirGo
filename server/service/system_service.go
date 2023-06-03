package service

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/utils/alipay_plugin"
	"AirGo/utils/mail_plugin"
)

func GetThemeConfig() (*model.Theme, error) {
	var theme model.Theme
	err := global.DB.First(&theme, 1).Error
	return &theme, err
}

func UpdateThemeConfig(theme *model.Theme) error {
	global.Theme = *theme
	return global.DB.Debug().Model(&model.Theme{ID: 1}).Save(&theme).Error
}

// 获取系统配置
func GetSetting() (*model.Server, error) {
	var setting model.Server
	err := global.DB.First(&setting).Error
	return &setting, err
}

// 修改系统配置
func UpdateSetting(setting *model.Server) error {
	//修改theme中的字段
	global.Theme.EnableEmailCode = setting.System.EnableEmailCode
	err := global.DB.Save(&global.Theme).Error
	if err != nil {
		return err
	}
	// 修改系统配置
	err = global.DB.Save(&setting).Error
	if err != nil {
		return err
	}
	//重新加载系统配置
	global.Server = *setting
	//pay
	client, err := alipay_plugin.InitAlipayClient()
	if err == nil {
		global.AlipayClient = client
	}
	//email
	d := mail_plugin.InitEmailDialer()
	if d != nil {
		global.EmailDialer = d
	}
	return nil
}
