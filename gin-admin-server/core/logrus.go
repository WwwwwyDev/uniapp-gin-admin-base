package core

import (
	"fmt"
	"gin-server/global"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)


func Logrus() *logrus.Logger{
	logrusCfg := global.CONFIG.Runtime.Log

	customFormatter := new(logrus.JSONFormatter)
	customFormatter.TimestampFormat = logrusCfg.TimestampFormat // 时间格式
	customFormatter.DisableTimestamp = logrusCfg.DisableTimestamp   // 禁止显示时间

	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")
	filename := year+month+day + ".log"
	path := logrusCfg.LogRootPath + filename

	f, err:=os.OpenFile(path,os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil{
		fmt.Println(err)
		os.Exit(0)
	}
	mw := io.MultiWriter(os.Stdout,f)
	log := logrus.New()
	log.SetFormatter(customFormatter)
	log.SetOutput(mw)
	log.SetLevel(logrus.DebugLevel)
	return log
}
