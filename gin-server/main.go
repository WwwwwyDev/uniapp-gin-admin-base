package main

import (
	"fmt"
	"gin-server/core"
	"gin-server/global"
	"gin-server/initialize"
	"github.com/gin-gonic/gin"
)

func init() {
	core.Setup("dev")
	initialize.Setup()
}

func main() {
	serverCfg := global.CONFIG.Server
	gin.SetMode(serverCfg.RunMode)
	endPoint := fmt.Sprintf(":%d", serverCfg.HttpPort)
	app := initialize.RegisterRouter()
	app.Run(endPoint)
}
