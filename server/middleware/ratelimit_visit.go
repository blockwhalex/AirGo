package middleware

import (
	"AirGo/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RateLimitVisit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uID, _ := ctx.Get("uID")
		uIDStr := strconv.Itoa(uID.(int))
		if ok := global.RateLimit.VisitRole.AllowVisit(uIDStr); !ok {
			global.Logrus.Error("访问量超出,其剩余访问次数情况如下:", global.RateLimit.VisitRole.RemainingVisits(uIDStr))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
