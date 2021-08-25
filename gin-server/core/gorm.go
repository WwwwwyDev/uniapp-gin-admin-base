package core

import (
	"fmt"
	"gin-server/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	root := global.CONFIG.Database
	if root.Type == "mysql"{
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			root.User,
			root.Password,
			root.Host,
			root.Port,
			root.Name,
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println(err)
		}

		return db
	}
	if root.Type == "postgresql"{
		dsn := fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			root.User,
			root.Password,
			root.Name,
			root.Host,
			root.Port,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println(err)
		}

		return db
	}
	return nil
}
