package middleware

import (
	"AirGo/global"
	"github.com/gin-gonic/gin"
)

func RateLimitIP() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		if ok := global.RateLimit.IPRole.AllowVisitByIP4(ip); !ok {
			global.Logrus.Error(ip+"访问量超出,其剩余访问次数情况如下:", global.RateLimit.IPRole.RemainingVisits(ip))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
