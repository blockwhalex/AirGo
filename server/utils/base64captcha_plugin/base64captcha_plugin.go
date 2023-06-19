package base64captcha_plugin

import (
	"AirGo/global"
	"github.com/mojocn/base64Captcha"
)

// base64 验证码
func InitBase64Captcha() *base64Captcha.Captcha {
	//var store = base64Captcha.DefaultMemStore
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, global.Base64CaptchaStore)
	//生成captcha
	return cp
}

//生成captcha
//id, b64s, err := cp.Generate()
//	model.Result(200, "验证码获取成功", gin.H{
//		"captchaId":     id,
//		"picPath":       b64s,
//		"captchaLength": global.GVA_CONFIG.Captcha.KeyLong,
//		"openCaptcha":   oc,
//	}, c)

//store.Verify(l.CaptchaId, l.Captcha, true)
