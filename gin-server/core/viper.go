package core

import (
	"fmt"
	"gin-server/global"
	"github.com/spf13/viper"
	"os"
)

func Viper(env string) *viper.Viper {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := viper.New()
	config.AddConfigPath(path+"/config")     //设置读取的文件路径
	config.SetConfigName("config-"+env) //设置读取的文件名
	config.SetConfigType("yml")   //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := config.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return config
}
