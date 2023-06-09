package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/utils/casbin_plugin"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
)

// 更新casbin权限
func UpdateCasbinPolicy(ctx *gin.Context) {
	// body, err := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("body:", string(body), err)

	var casbinInfo model.CasbinInfo
	err := ctx.ShouldBind(&casbinInfo)
	if err != nil {
		response.Fail("更新casbin权限参数error:", err, ctx)
		return
	}
	//fmt.Println("更新casbin权限参:", casbinInfo)
	err = casbin_plugin.UpdateCasbinPolicy(&casbinInfo)
	if err != nil {
		response.Fail("更新casbin权限error:", err, ctx)
		return
	}
	response.OK("更新casbin权限成功", nil, ctx)

}

// 更新casbin权限
func UpdateCasbinPolicyNew(ctx *gin.Context) {
	var data model.CasbinData
	err := ctx.ShouldBind(&data)
	if err != nil {
		global.Logrus.Error("更新casbin权限参数error:", err)
		response.Fail("更新casbin权限参数error:", err, ctx)
		return
	}
	//前端传过来的没有处理，只有method，从数据库查询完整的rules，再更新casbin
	err = casbin_plugin.UpdateCasbinPolicyNew(&data)
	if err != nil {
		global.Logrus.Error("更新casbin权限error:", err)
		response.Fail("更新casbin权限error:", err, ctx)
		return
	}
	response.OK("更新casbin权限成功", nil, ctx)

}

// 获取全部权限
func GetAllPolicy(ctx *gin.Context) {

	res := casbin_plugin.GetAllPolicy()
	response.OK("获取权限列表成功", res, ctx)
}

// 获取权限列表ByRoleIds
func GetPolicyByRoleIds(ctx *gin.Context) {
	// body, err := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("body:", string(body), err)
	var casbinInfo model.CasbinInfo
	err := ctx.ShouldBind(&casbinInfo)
	if err != nil {
		global.Logrus.Error("获取权限列表参数错误", err)
		response.Fail("获取权限列表参数错误", nil, ctx)
		return
	}
	res := casbin_plugin.GetPolicyPathByRoleId(&casbinInfo)
	response.OK("获取权限列表成功", res, ctx)
}
