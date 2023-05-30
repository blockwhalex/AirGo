package initialize

import (
	"AirGo/global"
	"AirGo/utils/alipay_plugin"
	"log"
)

func InitAlipayClient() {
	client, err := alipay_plugin.InitAlipayClient()
	if err != nil {
		log.Println("alipay client err:", err)
		return
	}

	log.Println("alipay client init:", err)
	global.AlipayClient = client
}
