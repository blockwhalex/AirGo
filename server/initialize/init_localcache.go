package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"AirGo/global"
	time_duration "AirGo/utils/time_plugin"
)

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
