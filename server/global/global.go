package global

import (
	"AirGo/config"
	"AirGo/model"
	"github.com/casbin/casbin/v2"
	"github.com/mojocn/base64Captcha"
	alipay "github.com/smartwalle/alipay/v3"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	//GVA_DBList map[string]*gorm.DB
	//GVA_REDIS  *redis.Client
	CONFIG             config.Server          //全局配置（本地）
	VP                 *viper.Viper           //
	LocalCache         local_cache.Cache      //本地kv cache
	AlipayClient       *alipay.Client         //alipay
	Casbin             *casbin.CachedEnforcer //casbin
	Server             model.Server           //全局配置（数据库）
	Theme              model.Theme            //全局主题配置
	Base64Captcha      *base64Captcha.Captcha //Base64Captcha
	Base64CaptchaStore base64Captcha.Store    //Base64CaptchaStore
	//Base64CaptchaDriver *base64Captcha.DriverDigit //Base64CaptchaDriver
)
