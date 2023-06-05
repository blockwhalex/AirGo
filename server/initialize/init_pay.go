package initialize

import (
	"AirGo/global"
	"AirGo/utils/alipay_plugin"
)

func InitAlipayClient() {
	client, err := alipay_plugin.InitAlipayClient()
	if err != nil {
		global.Logrus.Error("alipay client err:", err)
		return
	}

	global.AlipayClient = client
}
