package api

import (
	"AirGo/global"
	"AirGo/service"
	"AirGo/utils/response"
	"AirGo/utils/websocket_plugin"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}

// websocket im 测试
func WebSocketMsg(ctx *gin.Context) {
	//uID, ok := ctx.Get("uID")
	//if !ok || uID == nil {
	//	response.Fail("uID参数错误", nil, ctx)
	//	return
	//}
	//uIDInt := uID.(int)
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		global.Logrus.Error("websocket upgrade err:", err)
		response.Fail("websocket err"+err.Error(), nil, ctx)
		return
	}
	//defer conn.Close()
	//client
	client := &websocket_plugin.Client{
		//ID:            strconv.Itoa(uIDInt),
		ID:            ctx.ClientIP(),
		WsSocket:      conn,
		ClientChannel: make(chan []byte),
		ExpireTime:    5 * time.Second, //5秒过期时间
		QuitChanel:    make(chan bool),
	}
	global.WsManager.OnlineChannel <- client
	go client.Read(global.WsManager, service.GetNodesStatus)
	go client.Write(global.WsManager)
	time.Sleep(time.Second)
}
