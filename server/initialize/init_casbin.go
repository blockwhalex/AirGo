package initialize

import (
	"AirGo/global"
	"AirGo/utils/casbin_plugin"
)

func InitCasbin() {
	global.Casbin = casbin_plugin.Casbin()
}
