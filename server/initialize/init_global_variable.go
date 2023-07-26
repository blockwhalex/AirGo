package initialize

import (
	"AirGo/global"
	"AirGo/service"
	"AirGo/utils/base64captcha_plugin"
	"AirGo/utils/casbin_plugin"
	"AirGo/utils/logrus_plugin"
	"AirGo/utils/mail_plugin"
	time_duration "AirGo/utils/time_plugin"
	"AirGo/utils/websocket_plugin"
	"github.com/mojocn/base64Captcha"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/yudeguang/ratelimit"
	"time"
)

func InitGlobalVariable() {

}
func InitBase64Captcha() {
	global.Base64CaptchaStore = base64Captcha.DefaultMemStore
	global.Base64Captcha = base64captcha_plugin.InitBase64Captcha()
}

func InitLogrus() {
	global.Logrus = logrus_plugin.InitLogrus()
}

func InitTheme() {
	res, err := service.GetThemeConfig()
	if err != nil {
		global.Logrus.Error("系统配置获取失败", err.Error())
		return
	}
	global.Theme = *res
}

// 系统配置
func InitServer() {
	res, err := service.GetSetting()
	if err != nil {
		global.Logrus.Error("系统配置获取失败", err.Error())
		return
	}
	global.Server = *res
}
func InitWebsocket() {
	global.WsManager = websocket_plugin.NewManager()
	global.WsManager.NewClientManager()
}

func InitRatelimit() {
	global.RateLimit.IPRole = ratelimit.NewRule()
	global.RateLimit.IPRole.AddRule(time.Second*60, global.Server.RateLimitParams.IPRoleParam)
	global.RateLimit.VisitRole = ratelimit.NewRule()
	global.RateLimit.VisitRole.AddRule(time.Second*60, global.Server.RateLimitParams.VisitParam)
}
func InitEmailDialer() {
	d := mail_plugin.InitEmailDialer()
	if d != nil {
		global.EmailDialer = d
	}
}
func InitLocalCache() {
	//判断有没有设置时间
	dr, err := time_duration.ParseDuration(global.Server.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	//初始化local cache配置
	global.LocalCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr), //设置默认的超时时间
	)
}
func InitCasbin() {
	global.Casbin = casbin_plugin.Casbin()
}
