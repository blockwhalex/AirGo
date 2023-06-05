package middleware

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/utils/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	//"AirGo/utils/casbin_plugin"
)

// Casbin 拦截器
func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从middleware获取user id
		uID, ok := c.Get("uID")
		if !ok {
			response.Fail("接收参数错误", nil, c)
			c.Abort()
			return
		}
		uIdNew, _ := uID.(int)
		//请求的PATH
		path := c.Request.URL.Path
		obj := strings.TrimSpace(path)
		//请求方法
		act := c.Request.Method
		// 获取用户的角色组
		var u model.User
		err := global.DB.Preload("RoleGroup").First(&u, uIdNew).Error
		if err != nil {
			response.Fail("用户查询错误", nil, c)
			c.Abort()
			return
		}
		//
		var roleIds []int
		status := false
		for _, v := range u.RoleGroup {
			roleIds = append(roleIds, v.ID)
			roleID := strconv.Itoa(v.ID)
			success, err := global.Casbin.Enforce(roleID, obj, act) // 判断策略中是否存在
			if err != nil {
				global.Logrus.Error("权限casbin err:", err)
				c.Abort()
				return
			}
			if success {
				status = true
				break
			}
		}
		if !status {
			response.Fail(obj+"权限不足", nil, c)
			c.Abort()
			return
		}
		//将角色组写入 gin.Context
		c.Set("roleIds", roleIds)
		c.Next()

		// // log.Println("casbin sub:", sub) //casbin sub: 888
		// // log.Println("casbin obj:", obj) // casbin obj: /user/getUserInfo
		// // log.Println("casbin act:", act) //casbin act: GET

	}
}
