package initialize

import (
	"AirGo/global"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 读取配置文件，并转换成 struct结构
func InitViper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml") //config路径
	v.SetConfigType("yaml")        //设置文件的类型
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig() //监听

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil { //解析到全局配置
		fmt.Println(err)
	}
	return v
}
