package main

import (
	"gin-server/core"
	"github.com/briandowns/spinner"
	"time"
)

func init() {
	s := spinner.New(spinner.CharSets[1], 100*time.Millisecond)
	s.Start()
	core.Setup()
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
