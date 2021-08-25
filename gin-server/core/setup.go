package core

import (
	"fmt"
	"gin-server/global"
	"os"
	"time"
)

func Setup(){
	global.TIME = time.Now()
	global.VIPER = Viper("dev") // 初始化Viper配置
	global.LOG = Logrus()        //初始化日志系统
	global.DB = Gorm() //初始化数据库连接池
	if global.DB == nil{
		fmt.Print("不支持此类型数据库")
		os.Exit(1)
	}
}
