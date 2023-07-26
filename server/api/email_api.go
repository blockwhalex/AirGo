package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/utils/encode_plugin"
	"AirGo/utils/mail_plugin"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

// 邮箱验证码
func GetMailCode(ctx *gin.Context) {
	var u model.UserRegisterEmail
	err := ctx.ShouldBind(&u)
	if err != nil {
		global.Logrus.Error("邮箱验证码参数错误", err.Error())
		response.Fail("邮箱验证码参数错误"+err.Error(), nil, ctx)
		return
	}
	_, ok := global.LocalCache.Get(u.UserName + "emailcode")
	if ok {
		response.Fail("邮箱验证码获取频繁", nil, ctx)
		return
	}
	//用户是否存在且是否有效
	//user, err := service.FindUserByEmail(&model.User{UserName: u.UserName})
	//if err == gorm.ErrRecordNotFound {
	//	response.Fail("用户不存在", nil, ctx)
	//	return
	//}
	//生成验证码
	randomStr := encode_plugin.RandomString(4) //4位随机数
	var wg sync.WaitGroup
	wg.Add(3)
	//验证码存入local cache
	go func(wg *sync.WaitGroup) {
		global.LocalCache.Set(u.UserName+"emailcode", randomStr, 60*time.Second) //过期
		wg.Done()
	}(&wg)
	//发送邮件
	go func(wg *sync.WaitGroup) {
		err = mail_plugin.SendEmail(global.EmailDialer, u.UserName, randomStr, global.Server.Email.EmailContent)
		if err != nil {
			global.Logrus.Error("验证码获取失败:", err.Error())
			//response.OK("验证码获取失败"+err.Error(), nil, ctx)
		}
		wg.Done()
	}(&wg)
	go func(ctx *gin.Context, wg *sync.WaitGroup) {
		response.OK("邮箱验证码已发送，请注意检查邮箱", nil, ctx)
		wg.Done()
	}(ctx, &wg)
	wg.Wait()
}
