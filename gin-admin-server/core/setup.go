package core

import (
	"gin-server/global"
	"time"
)

func Setup(env string){
	global.TIME = time.Now()
	global.VIPER = Viper(env) // 初始化Viper配置
	global.LOG = Logrus()        //初始化日志系统
	global.DB = Gorm() //初始化数据库连接池
	global.REDIS = Redis() //初始化redis
}
