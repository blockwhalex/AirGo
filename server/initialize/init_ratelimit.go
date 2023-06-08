package initialize

import (
	"AirGo/global"
	"github.com/yudeguang/ratelimit"
	"time"
)

func InitRatelimit() {
	global.RateLimit.IPRole = ratelimit.NewRule()
	global.RateLimit.IPRole.AddRule(time.Second*60, global.Server.RateLimitParams.IPRoleParam)
	global.RateLimit.VisitRole = ratelimit.NewRule()
	global.RateLimit.VisitRole.AddRule(time.Second*60, global.Server.RateLimitParams.VisitParam)
}
