package initialize

import (
	"AirGo/global"
	"AirGo/utils/logrus_plugin"
)

func InitLogrus() {
	global.Logrus = logrus_plugin.InitLogrus()
}
