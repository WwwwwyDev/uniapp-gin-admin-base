package v1

import (
	"gin-server/util"
	"gin-server/util/app"
	"github.com/gin-gonic/gin"
)

func TestJwt(c *gin.Context) {
	get, _ := c.Get("claims")
	get = get.(*util.Claims)
	app.OK(c, get,"OK")
}
