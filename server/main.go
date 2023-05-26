package main

import (
	"AirGo/global"
	"AirGo/initialize"
	"AirGo/model"
	"fmt"
)

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
	Addr string
}

func (p Student) GetStructData() interface{} {
	return p
}

func main() {
	global.VP = initialize.InitViper() // 初始化Viper
	global.DB = initialize.Gorm()      // gorm连接数据库
	initialize.InitCasbin()            //casbin，生成casbin_rule 表
	if global.DB != nil && !global.DB.Migrator().HasTable(&model.User{}) {
		fmt.Println("未找到sys_user库表,开始建表")
		initialize.RegisterTables()      //创建table
		initialize.InsertInto(global.DB) //导入数据
	}
	initialize.InitServer()     //全局系统配置
	initialize.InitTheme()      //全局主题
	initialize.InitLocalCache() //local cache
	//initialize.InitBase64Captcha() //Base64Captcha
	initialize.InitCrontab()      //定时任务
	initialize.InitAlipayClient() //alipay
	initialize.InitRouter()       //初始总路由

}
