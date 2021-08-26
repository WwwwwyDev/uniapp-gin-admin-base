package v1

import (
	"gin-server/util/app"
	"github.com/gin-gonic/gin"
)

func JwtTest(c *gin.Context) {
	get, _ := c.Get("claims")
	app.OK(c, get,"OK")
}
