package initialize

import (
	middleWare "gin-server/middleWare"
	"github.com/gin-gonic/gin"
)

//无需认证
func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/v1")
}


func RegisterRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleWare.Cors())
	g := r.Group("/api")
	sysNoCheckRoleRouter(g)
	return r
}
