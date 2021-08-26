package main

import (
	"gin-server/core"
	"gin-server/initialize"
)

func init() {
	core.Setup("dev")
	initialize.RegisterTables()
}

func main() {
	//result, _ := global.REDIS.Set("test", "2343", time.).Result()
	//fmt.Print(result)
	initialize.Run()
}
