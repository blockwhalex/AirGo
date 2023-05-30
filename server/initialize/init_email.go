package initialize

import (
	"AirGo/global"
	"AirGo/utils/mail_plugin"
)

func InitEmailDialer() {
	d := mail_plugin.InitEmailDialer()
	if d != nil {
		global.EmailDialer = d
	}
}
