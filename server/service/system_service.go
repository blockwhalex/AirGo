package service

import (
	"AirGo/global"
	"AirGo/model"
)

func GetThemeConfig() (*model.Theme, error) {
	var theme model.Theme
	err := global.DB.First(&theme, 1).Error
	return &theme, err
}

func UpdateThemeConfig(theme *model.Theme) error {
	var t model.Theme
	global.DB.First(&t)
	theme.CreatedAt = t.CreatedAt
	//theme.UpdatedAt=t.UpdatedAt
	global.LocalCache.SetNoExpire("theme", theme) //保存theme到local cache
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
	var theme model.Theme
	theme = global.Theme
	theme.EnableEmailCode = setting.System.EnableEmailCode
	global.DB.Save(&theme)
	global.Theme = theme
	// 修改系统配置
	err := global.DB.Save(&setting).Error
	if err != nil {
		return err
	}
	global.Server = *setting
	return nil
}
