package initialize

import (
	"gin-server/global"
	"gin-server/middleware"
	"gin-server/router"
	"gin-server/router/system"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

//无需认证
func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/v1")
	system.SysUserRouter(r)
}

//认证
func sysCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/v1/c")
	r.Use(middleware.JWTAuth())
	router.TestRouter(r)
}


func RegisterRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	path, err := os.Getwd()
	if err != nil {
		global.LOG.Error(err)
	}
	r.StaticFS("/upload", http.Dir(path+"/runtime/upload"))
	r.StaticFile("/favicon.ico", path+"/favicon.ico")
	g := r.Group("/api")
	sysNoCheckRoleRouter(g)
	sysCheckRoleRouter(g)
	return r
}
