package middleware

import (
	"AirGo/utils/jwt_plugin"
	"AirGo/utils/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token
		token := c.GetHeader("Authorization")
		//判断
		if token == "" {
			response.Fail("未携带token", nil, c)
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer ") {
			//去掉bearer
			token = token[7:]
		}
		claims, err := jwt_plugin.ParseToken(token)
		if err != nil {
			response.Fail(err.Error(), nil, c)
			c.Abort()
			return
		}
		//log.Println("token解析后 claims.ID：", claims.ID)
		//设置user id
		c.Set("uID", claims.ID)
		c.Set("uName", claims.UserName)
		c.Next()

	}

}
