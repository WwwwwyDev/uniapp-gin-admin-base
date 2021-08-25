package initialize

import (
	"gin-server/global"
	"gin-server/model/system"
	"os"
)

func RegisterTables() {
	err := global.DB.AutoMigrate(
		system.SysUser{},
	)
	if err != nil {
		global.LOG.Error(err)
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
