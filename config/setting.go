package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

func Init() {

	//从配置文件中加载配置文件信息
	// 指定配置⽂文件名称(不不需要带后缀)
	viper.SetConfigName("db")
	// 指定配置⽂文件类型
	viper.SetConfigType("yaml")
	// 指定查找配置⽂文件的路路径(这⾥里里使⽤用相对路路径)
	viper.AddConfigPath("./config/")

	//监控并重新读取配置⽂文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置⽂文件发⽣生变更更之后会调⽤用的回调函数
		// 当配置文件修改之后要把变更后的配置信息更新到全局变量Conf里
		viper.Unmarshal(&Conf)
		fmt.Println("Config file changed:", e.Name)
	})

	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.Unmarshal(&Conf)
	fmt.Println(viper.Get("mysql.db"))
	fmt.Println(viper.Get("redis"))
}
