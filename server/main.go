package main

import (
	"AirGo/global"
	"AirGo/initialize"
	"AirGo/model"
)

func main() {
	initialize.InitLogrus()            //logrus
	global.VP = initialize.InitViper() //初始化Viper
	global.DB = initialize.Gorm()      //gorm连接数据库
	if global.DB != nil {
		if !global.DB.Migrator().HasTable(&model.User{}) {
			global.Logrus.Info("未找到sys_user库表,开始建表并初始化数据...")
			initialize.InitCasbin()          //加载casbin,生成casbin_rule 表
			initialize.RegisterTables()      //创建table
			initialize.InsertInto(global.DB) //导入数据
			initialize.InitCasbin()          //重新加载casbin
		} else {
			initialize.RegisterTables() //AutoMigrate 自动迁移 schema
			initialize.InitCasbin()     //加载casbin
		}
	}
	initialize.InitServer()     //全局系统配置
	initialize.InitTheme()      //全局主题
	initialize.InitLocalCache() //local cache
	//initialize.InitBase64Captcha() //Base64Captcha
	initialize.InitCrontab()      //定时任务
	initialize.InitAlipayClient() //alipay
	initialize.InitEmailDialer()  //gomail Dialer
	initialize.InitWebsocket()    //websocket
	initialize.InitRatelimit()    //限流
	initialize.InitRouter()       //初始总路由
	//var userArr []model.User
	//global.DB.Find(&userArr)
	//for k, _ := range userArr {
	//	userArr[k].InvitationCode = encode_plugin.RandomString(8)
	//	service.SaveUser(&userArr[k])
	//}

}
