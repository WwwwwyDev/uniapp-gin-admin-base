package initialize

import (
	_ "gin-server/docs"
	"gin-server/global"
	"gin-server/middleware"
	"gin-server/router"
	"gin-server/router/system"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"os"
)

//无需认证
func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/v1")
	system.SysUserRouter(r)
	system.SysCaptchaRouter(r)
}

//认证
func sysCheckRoleRouter(r *gin.RouterGroup) {
	r = r.Group("/v1/c")
	r.Use(middleware.JWTAuth())
	router.TestJwtRouter(r)
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g := r.Group("/api")
	sysNoCheckRoleRouter(g)
	sysCheckRoleRouter(g)
	return r
}
