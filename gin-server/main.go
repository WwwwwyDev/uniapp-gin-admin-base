package main

import (
	"fmt"
	"gin-server/core"
	"gin-server/global"
	"github.com/briandowns/spinner"
	"time"
)

func init() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)  // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(2 * time.Second)                                  // Run for some time to simulate work
	global.TIME = time.Now()
	global.VIPER = core.Viper("dev") // 初始化Viper配置
	global.LOG = core.Logrus()        //初始化日志系统
	global.DB = core.Gorm() //初始化数据库连接池
	if global.DB == nil{
		fmt.Print("不支持此类型数据库")
		return
	}
	s.Stop()
}

func main() {

	//global.GVA_DB = initialize.Gorm() // gorm连接数据库
	//initialize.Timer()
	//if global.GVA_DB != nil {
	//	initialize.MysqlTables(global.GVA_DB) // 初始化表
	//	// 程序结束前关闭数据库链接
	//	db, _ := global.GVA_DB.DB()
	//	defer db.Close()
	//}
	//core.RunWindowsServer()
}
