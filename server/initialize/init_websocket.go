package initialize

import (
	"AirGo/global"
	"AirGo/utils/websocket_plugin"
)

func InitWebsocket() {
	global.WsManager = websocket_plugin.NewManager()
	global.WsManager.NewClientManager()
}
