package service

import (
	"AirGo/global"
	"AirGo/model"
	"strconv"
	"time"
)

// 根据node name 模糊查询节点
func GetNodeByName(name string) ([]model.Node, error) {
	var nodes []model.Node
	err := global.DB.Where("name like ?", ("%" + name + "%")).Find(&nodes).Error
	return nodes, err

}

// 查询全部节点
func GetAllNode() (*[]model.Node, error) {
	var nodes []model.Node
	err := global.DB.Find(&nodes).Error
	if err != nil {
		return nil, err
	}
	return &nodes, nil
}

// 新建节点
func NewNode(node *model.Node) error {
	//node.ID = 0 //清空节点id 防止插入失败
	err := global.DB.Create(&node).Error
	return err
}

// 删除节点
func DeleteNode(node *model.Node) error {
	//删除关联
	err := global.DB.Where("node_id = ?", node.ID).Delete(&model.GoodsAndNodes{}).Error
	if err != nil {
		return err
	}
	err = global.DB.Where(&model.Node{ID: node.ID}).Delete(&model.Node{}).Error
	return err
}

// 更新节点
func UpdateNode(node *model.Node) error {
	return global.DB.Save(&node).Error
}

// 插入节点流量统计
func NewTrafficLog(t *model.TrafficLog) error {
	return global.DB.Create(&t).Error
}

// 查询节点流量
func GetNodeTraffic(params model.QueryParamsWithDate) model.NodesWithTotal {
	//var nodeArr []model.Node
	var nodeArr model.NodesWithTotal
	var startTime, endTime time.Time
	//时间格式转换
	if len(params.Date) == 2 {
		startTime, _ = time.ParseInLocation("2006-01-02 15:04:05", params.Date[0], time.Local)
		endTime, _ = time.ParseInLocation("2006-01-02 15:04:05", params.Date[1], time.Local)
	} else {
		//默认前1个月数据
		endTime = time.Now().Local()
		startTime = endTime.AddDate(0, 0, -30)
	}
	if params.Search != "" {
		err := global.DB.Model(&model.Node{}).Count(&nodeArr.Total).Where("name LIKE ?", "%"+params.Search+"%").Limit(params.PageSize).Offset((params.PageNum-1)*params.PageSize).Preload("TrafficLogs", global.DB.Where("created_at > ? and created_at < ?", startTime, endTime)).Find(&nodeArr.NodeList).Error
		if err != nil {
			global.Logrus.Error("查询节点流量error:", err.Error())
			return model.NodesWithTotal{}
		}
	} else {
		err := global.DB.Model(&model.Node{}).Count(&nodeArr.Total).Limit(params.PageSize).Offset((params.PageNum-1)*params.PageSize).Preload("TrafficLogs", global.DB.Where("created_at > ? and created_at < ?", startTime, endTime)).Find(&nodeArr.NodeList).Error
		if err != nil {
			global.Logrus.Error("查询节点流量error:", err.Error())
			return model.NodesWithTotal{}
		}
	}
	for i1, _ := range nodeArr.NodeList {
		for _, v := range nodeArr.NodeList[i1].TrafficLogs {
			nodeArr.NodeList[i1].TotalUp = nodeArr.NodeList[i1].TotalUp + v.U
			nodeArr.NodeList[i1].TotalDown = nodeArr.NodeList[i1].TotalDown + v.D
		}
		//nodeArr[i1].TrafficLogs=[]model.TrafficLog{} //清空traffic
	}
	return nodeArr
}

// 获取 node status
func GetNodesStatus() *[]model.NodeStatus {
	var nodesIds []model.Node
	global.DB.Model(&model.Node{}).Select("id", "name", "traffic_rate").Find(&nodesIds)
	var nodestatusArr []model.NodeStatus
	for _, v := range nodesIds {
		var nodeStatus = model.NodeStatus{}
		vStatus, ok := global.LocalCache.Get(strconv.Itoa(v.ID) + "status")
		if !ok {
			nodeStatus.ID = v.ID
			nodeStatus.Name = v.Name
			nodeStatus.TrafficRate = v.TrafficRate
			nodeStatus.Status = false
			nodeStatus.D = 0
			nodeStatus.U = 0
			nodestatusArr = append(nodestatusArr, nodeStatus)
		} else {
			nodeStatus = vStatus.(model.NodeStatus)
			nodeStatus.Name = v.Name
			nodeStatus.TrafficRate = v.TrafficRate
			if time.Now().Sub(nodeStatus.LastTime).Seconds() > 60 {
				nodeStatus.UserAmount = 0
				nodeStatus.Status = false
			} else {
				nodeStatus.Status = true
			}
			nodestatusArr = append(nodestatusArr, nodeStatus)
		}
	}
	return &nodestatusArr
}
