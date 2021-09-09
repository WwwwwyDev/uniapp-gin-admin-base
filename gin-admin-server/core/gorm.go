package core

import (
	"fmt"
	"gin-server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func Gorm() *gorm.DB {
	gormCfg := global.CONFIG.Mysql
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		gormCfg.User,
		gormCfg.Password,
		gormCfg.Host,
		gormCfg.Port,
		gormCfg.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		global.LOG.Error(err)
		os.Exit(0)
	}else{
		global.LOG.Info("mysql "+gormCfg.Host +":" + strconv.Itoa(gormCfg.Port) +" connect successfully")
	}

	return db
}
