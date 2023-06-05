package websocket_plugin

import (
	"AirGo/model"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 消息体的定义
type WsMessage struct {
	Type int         `json:"type"`
	Data interface{} `json:"data"`
}

// 客户端结构体定义
type Client struct {
	ID            string
	IpAddress     string
	WsSocket      *websocket.Conn //WsSocket则是用来记录客户端和服务端直接到WebSocket连接
	ClientChannel chan []byte
	ExpireTime    time.Duration // 一段时间没有接收到心跳则过期
	QuitChanel    chan bool
}

// 客户端管理
type ClientManager struct {
	Clients        map[string]*Client // 记录在线用户
	MapLock        sync.RWMutex       //map读写锁
	Broadcast      chan []byte        //触发消息广播
	OnlineChannel  chan *Client       // 新用户登陆
	OfflineChannel chan *Client       // 用户退出
}

func NewManager() *ClientManager {
	return &ClientManager{
		Clients:        make(map[string]*Client), // 记录在线用户
		Broadcast:      make(chan []byte, 100),   //触发消息广播
		OnlineChannel:  make(chan *Client, 100),  // 触发新用户登陆
		OfflineChannel: make(chan *Client, 100),  // 触发用户退出
	}
}

// 启动Manager，管理所有client，并进行相应handler（广播，下线等）
func (m *ClientManager) NewClientManager() {
	//该goroutine负责处理用户上线
	go func() {
		for {
			select {
			case oneClient := <-m.OnlineChannel:
				m.MapLock.Lock()
				m.Clients[oneClient.ID] = oneClient
				m.MapLock.Unlock()
				// 如果有新用户连接则发送消息给他
				//count := len(m.Clients)
				//resp, _ := json.Marshal(&WsMessage{Type: 1, Data: count})
				//log.Printf("当前在线%v人", count)
				//m.Clients[oneClient.ID].ClientChannel <- resp
			}
		}
	}()
	//该goroutine负责处理用户下线
	go func() {
		for {
			select {
			case oneClient := <-m.OfflineChannel:
				err := oneClient.WsSocket.Close() //关闭该client连接
				if err != nil {
					fmt.Println("oneClient.WsSocket.Close() err:", err)
					continue
				}
				oneClient.QuitChanel <- true   //通知write 关闭
				time.Sleep(time.Second)        //延时1秒关闭chanel
				close(oneClient.ClientChannel) //关闭client chanel
				close(oneClient.QuitChanel)    //关闭client quit chanel
				m.MapLock.Lock()
				delete(m.Clients, oneClient.ID) //deleted client
				m.MapLock.Unlock()
				//data := "系统通知:" + offlineMsg
				//resp, _ := json.Marshal(&WsMessage{Type: 1, Data: data})
				//m.Broadcast <- resp
			}
		}

	}()
	//该goroutine负责广播
	go func() {
		for {
			select {
			// 只要有一方发消息就广播
			case msg := <-m.Broadcast:
				for _, oneClient := range m.Clients {
					oneClient.ClientChannel <- msg
				}
			}
		}
	}()
}

// 读取客户端发送过来的消息,根据不同的type发送到对应的channel
func (c *Client) Read(manager *ClientManager, f func() *[]model.NodeStatus) {
	// 把当前客户端注销
	defer func() {
		//_ = c.WsSocket.Close()
		fmt.Println("defer close ws read")
		manager.OfflineChannel <- c //通知manager 下线该client
	}()
	for {
		_, data, err := c.WsSocket.ReadMessage() //阻塞
		if err != nil {
			fmt.Println("closed network connection，close ws read chanel:", err.Error())
			return
		}
		var wsmsg WsMessage
		err = json.Unmarshal(data, &wsmsg)
		if err != nil {
			fmt.Println("json.Unmarshal err:", err)
			continue
		}
		switch wsmsg.Type {
		case 8:
			return
		case 9:
			// 利用心跳监测
			resp, _ := json.Marshal(&WsMessage{Type: 10, Data: "pong"})
			c.ClientChannel <- resp
		case 1:
			// 推送TextMessage
			data := f()
			resp, _ := json.Marshal(&WsMessage{Type: 1, Data: data})
			c.ClientChannel <- resp

			//case 3:
			//	// 广播消息
			//	resp, _ := json.Marshal(&WsMessage{Type: 4, Data: "你好，这是 广播模式"})
			//	Manager.Broadcast <- resp
		}
	}
}

// Write 把对应消息写回客户端
func (c *Client) Write(manager *ClientManager) {
	defer func() {
		//_ = c.WsSocket.Close()
		fmt.Println("defer close ws write")
		//manager.OfflineChannel <- c
	}()
	for {
		select {
		case msg, ok := <-c.ClientChannel:
			if !ok {
				continue
			}
			err := c.WsSocket.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println("c.WsSocket.WriteMessage err:", err)
				return
			}
		case <-c.QuitChanel:
			return
		case <-time.After(c.ExpireTime):
			manager.OfflineChannel <- c
			return
		}

	}
}
