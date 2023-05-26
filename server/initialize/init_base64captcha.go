package initialize

import (
	"AirGo/global"
	"AirGo/utils/base64captcha_plugin"
	"github.com/mojocn/base64Captcha"
)

func InitBase64Captcha() {
	global.Base64CaptchaStore = base64Captcha.DefaultMemStore
	global.Base64Captcha = base64captcha_plugin.InitBase64Captcha()
}
