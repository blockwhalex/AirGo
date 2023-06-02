package api

import (
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 获取全部节点
func GetAllNode(ctx *gin.Context) {
	nodes, err := service.GetAllNode()
	if err != nil {
		response.Fail("获取全部节点错误", nil, ctx)
		return
	}
	response.OK("获取全部节点成功", nodes, ctx)

}

// 新建节点
func NewNode(ctx *gin.Context) {
	var node model.Node
	err := ctx.ShouldBind(&node)
	if err != nil {
		response.Fail("新建节点参数错误", nil, ctx)
		return
	}
	fmt.Println("新建节点node", node)
	err = service.NewNode(&node)
	if err != nil {
		response.Fail("新建节点错误", nil, ctx)
		return
	}
	response.OK("新建节点成功", nil, ctx)
}

// 删除节点
func DeleteNode(ctx *gin.Context) {
	var node model.Node
	err := ctx.ShouldBind(&node)
	if err != nil {
		response.Fail("删除节点参数错误", nil, ctx)
		return
	}
	fmt.Println("删除节点node", node)
	err = service.DeleteNode(&node)
	if err != nil {
		response.Fail("删除节点错误", nil, ctx)
		return
	}
	response.OK("删除节点成功", nil, ctx)
}

// 更新节点
func UpdateNode(ctx *gin.Context) {
	var node model.Node
	err := ctx.ShouldBind(&node)
	if err != nil {
		response.Fail("更新节点参数错误", nil, ctx)
		return
	}
	//fmt.Println("更新节点node", node)
	err = service.UpdateNode(&node)
	if err != nil {
		response.Fail("更新节点错误", nil, ctx)
		return
	}
	response.OK("更新节点成功", nil, ctx)

}

// 获取节点流量统计
func GetTraffic(ctx *gin.Context) {

}

// 查询节点流量
func GetNodeTraffic(ctx *gin.Context) {
	var trafficParams model.QueryParamsWithDate
	err := ctx.ShouldBind(&trafficParams)
	if err != nil {
		response.Fail("查询节点流量参数错误"+err.Error(), nil, ctx)
		return
	}
	//fmt.Println("查询节点流量trafficParams:", trafficParams)
	res := service.GetNodeTraffic(trafficParams)
	response.OK("查询节点流量", res, ctx)
}
