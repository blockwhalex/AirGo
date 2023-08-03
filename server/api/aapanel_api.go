package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// sspanel 响应模板
//
//	{
//		"ret": 1,
//		"data": {
//			"node_group": 0,        //节点群组
//			"node_class": 0,        //节点等级
//			"node_speedlimit": 0,   //节点限速/Mbps
//			"traffic_rate": 1,      //倍率
//			"mu_only": 1,           //只启用单端口多用户
//			"sort": 11,             //类型sort==11V2Ray sort==15>V2Ray vless  sort==12>V2Ray中转    sort==0>Shadowsocks   sort==1>VPN/Radius基础   sort==2>SSH   sort==5>Anyconnect   sort==9>Shadowsocks 单端口多用户   sort==10>Shadowsocks中转  sort==13>Shadowsocks V2Ray-Plugin&Obfs  sort==14>Trojan
//			"server": "plmoknijb.f3322.net;5566;0;ws;;path=/;host=tms.dingtalk.com",
//			"type": "ss-panel-v3-mod_Uim" //显示与隐藏
//		}
//	}
//
// 当前节点设置
func SSNodeInfo(ctx *gin.Context) {
	//验证key
	if global.Server.System.MuKey != ctx.Query("muKey") {
		return
	}
	nodeID := ctx.Param("nodeID")
	nodeIDInt, _ := strconv.Atoi(nodeID)
	//设置节点在线
	go func(nodeID string) {
		_, ok := global.LocalCache.Get(nodeID + "status")
		if !ok {
			//nodeStatus := vStatus.(model.NodeStatus)
			var nodeStatus model.NodeStatus
			nodeStatus.ID = nodeIDInt
			nodeStatus.Status = true
			global.LocalCache.Set(nodeID+"status", nodeStatus, time.Minute)
		}
	}(nodeID)
	nodeInfo, err := service.SSNodeInfo(nodeIDInt)
	if err != nil {
		global.Logrus.Error("当前节点设置,error:", err.Error())
		response.SSUsersFail(ctx)
		return
	}
	response.SSUsersOK(nodeInfo, ctx)
}

// 可连接的用户
func SSUsers(ctx *gin.Context) {
	//验证key
	if global.Server.System.MuKey != ctx.Query("muKey") {
		return
	}
	//节点号 是否启用？
	nodeID := ctx.Query("node_id")
	nodeIDInt, _ := strconv.Atoi(nodeID)
	//节点属于哪些goods
	goods, err := service.FindGoodsByNodeID(nodeIDInt)
	if err != nil {
		return
	}
	//goods属于哪些用户
	users, err := service.FindUsersByGoods(&goods)
	if err != nil {
		global.Logrus.Error("可连接的用户,error:", err.Error())
		response.SSUsersFail(ctx)
		return
	}
	response.SSUsersOK(users, ctx)
}

// 用户的流量使用情况: {"data":[{"user_id":1,"u":445782,"d":1757834}]}
// 上报用户的流量使用情况
func SSUsersTraffic(ctx *gin.Context) {
	//验证key
	if global.Server.System.MuKey != ctx.Query("muKey") {
		return
	}
	//节点号 是否启用？
	node_id_str := ctx.Query("node_id")
	node_id, _ := strconv.Atoi(node_id_str)
	var trafficReq model.TrafficReq
	err := ctx.ShouldBind(&trafficReq)
	if err != nil {
		return
	}
	response.SSUsersOK("ok", ctx)

	var userIds []int
	var userArr []model.User
	var trafficLog = model.TrafficLog{
		NodeID: node_id,
	}
	for _, v := range trafficReq.Data {
		//每个用户流量
		var user model.User
		userIds = append(userIds, v.UserID)
		user.ID = v.UserID
		user.SubscribeInfo.U = v.U
		user.SubscribeInfo.D = v.D
		userArr = append(userArr, user)
		//该节点总流量
		trafficLog.D = trafficLog.U + v.U
		trafficLog.U = trafficLog.D + v.D
	}
	// 统计status
	go func(id, userAmount int, u, d int64) {
		var nodeStatus = model.NodeStatus{
			ID:         id,
			UserAmount: userAmount,
			U:          float64(u),
			D:          float64(d),
			LastTime:   time.Now(),
			Status:     true,
		}
		var duration float64 = 60
		cacheStatus, ok := global.LocalCache.Get(strconv.Itoa(id) + "status")
		if ok && cacheStatus != nil {
			oldStatus := cacheStatus.(model.NodeStatus)
			duration = nodeStatus.LastTime.Sub(oldStatus.LastTime).Seconds()
		}
		nodeStatus.D, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", nodeStatus.D/1024/1024/duration*8), 64) //Mbps
		nodeStatus.U, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", nodeStatus.U/1024/1024/duration*8), 64)
		global.LocalCache.Set(strconv.Itoa(id)+"status", nodeStatus, 2*time.Minute)

	}(node_id, len(userIds), trafficLog.U, trafficLog.D)
	//插入流量统计统计
	err = service.NewTrafficLog(&trafficLog)
	if err != nil {
		global.Logrus.Error("插入流量统计统计error:", err)
		return
	}
	//更新用户流量信息
	if len(userArr) == 0 {
		return
	}
	err = service.UpdateUserTrafficInfo(userArr, userIds)
	if err != nil {
		global.Logrus.Error("更新用户流量信息error:", err)
		return
	}
}

// 上报用户的当前在线IP
func SSUsersAliveIP(ctx *gin.Context) {
	//body:{\"data\":[{\"user_id\":1,\"ip\":\"39.67.5.148\"}]}"
	//body, err := ioutil.ReadAll(ctx.Request.Body)
	//global.Logrus.Error("上报用户的当前在线IP body:", string(body), err)
}
