package casbin_plugin

import (
	"AirGo/global"
	"AirGo/model"
	"errors"
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Casbin() *casbin.CachedEnforcer {
	text := `
	[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act 
	`
	//请求的角色匹配规则的角色，请求的路径匹配规则的路径，请求的方法匹配规则的方法
	//keyMatch2	一个URL 路径，例如 /alice_data/resource1,它返回一个布尔值表示 url 是否匹配

	m, err := casbinModel.NewModelFromString(text)
	if err != nil {
		fmt.Println("casbin加载模型失败!")
		return nil
	}
	a, _ := gormadapter.NewAdapterByDB(global.DB)
	cachedEnforcer, _ := casbin.NewCachedEnforcer(m, a) //生成casbin实施实例
	cachedEnforcer.SetExpireTime(60 * 60)
	err = cachedEnforcer.LoadPolicy()
	if err != nil {
		fmt.Println("cachedEnforcer.LoadPolicy err:", err)
	}
	return cachedEnforcer
}

// 更新casbin权限
func UpdateCasbinPolicy(casbinInfo *model.CasbinInfo) error {
	roleIDInt := strconv.Itoa(casbinInfo.RoleID)
	ClearCasbin(0, roleIDInt)
	rules := [][]string{}
	for _, v := range casbinInfo.CasbinItems {
		rules = append(rules, []string{roleIDInt, v.Path, v.Method})
	}
	//fmt.Println("需要更新的casbin rules:", rules) //不能重复，负责casbin出错！
	//[[2 /user/login POST] [2 /user/getSub GET] [2 /casbin/getPolicyByRoleIds GET] [2 /casbin/updateCasbinPolicy POST] ]
	success, _ := global.Casbin.AddPolicies(rules)
	if !success {
		return errors.New("casbin添加失败")
	}
	err := global.Casbin.InvalidateCache()
	if err != nil {
		return err
	}
	//重新加载全局casbin
	global.Casbin = Casbin()
	return nil
}

// 更新casbin权限
func UpdateCasbinPolicyNew(casbinData *model.CasbinData) error {
	roleID := strconv.Itoa(casbinData.RoleID)
	var list []gormadapter.CasbinRule
	global.DB.Model(&gormadapter.CasbinRule{}).Where("v0 = 1 and v1 in (?)", casbinData.Data).Find(&list)

	ClearCasbin(0, roleID)
	rules := [][]string{}
	for _, v := range list {
		rules = append(rules, []string{roleID, v.V1, v.V2})
	}
	// fmt.Println("casbin list", list)
	// fmt.Println("casbin rules", rules)
	success, err := global.Casbin.AddPolicies(rules)
	if !success {
		fmt.Println("casbin添加失败 err", success, err)
		return errors.New("casbin添加失败")
	}
	err = global.Casbin.InvalidateCache()
	if err != nil {
		return err
	}
	//重新加载全局casbin
	global.Casbin = Casbin()
	return nil
}

// API更新随动
func UpdateCasbinApi(oldPath, oldMethod, newPath, newMethod string) error {
	err := global.DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	err = global.Casbin.InvalidateCache()
	if err != nil {
		return err
	}
	return err
}

// 获取权限列表
func GetPolicyPathByRoleId(casbinInfo *model.CasbinInfo) *model.CasbinInfo {
	roleID := strconv.Itoa(casbinInfo.RoleID)
	list := global.Casbin.GetFilteredPolicy(0, roleID)
	for _, v := range list {
		casbinInfo.CasbinItems = append(casbinInfo.CasbinItems, model.CasbinItem{
			Path:   v[1],
			Method: v[2],
		})
	}
	return casbinInfo
}

// 获取全部权限
func GetAllPolicy() *model.CasbinInfo {
	var casbinInfo model.CasbinInfo
	list := global.Casbin.GetFilteredPolicy(0, "1") //超级管理员为全部api
	fmt.Println("获取全部权限 list:", list)
	for _, v := range list {
		casbinInfo.CasbinItems = append(casbinInfo.CasbinItems, model.CasbinItem{
			Path:   v[1],
			Method: v[2],
		})
	}
	return &casbinInfo

}

// ClearCasbin
func ClearCasbin(v int, p ...string) bool {
	success, _ := global.Casbin.RemoveFilteredPolicy(v, p...)
	return success
}
