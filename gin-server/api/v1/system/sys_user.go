package system

import (
	"gin-server/global"
	model "gin-server/model/system"
	Req "gin-server/model/system/request"
	service "gin-server/service/system"
	"gin-server/util"
	"gin-server/util/app"
	"gin-server/util/e"
	"gin-server/util/valid"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

func Register(c *gin.Context) {
	var r Req.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.LOG.Error(err)
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	u := model.SysUser{Username: r.Username, Password: r.Password, Salt: uuid.NewV4().String(), NickName: r.NickName, Email: r.Email, Phone: r.Phone, HeaderImg: r.HeaderImg}
	err1 := valid.IsValid4sysUser(u)
	if err1 != nil {
		global.LOG.Error(err1)
		app.Error(c,e.ERROR,err1,err1.Error())
		return
	}
	err2, _ := service.Register(u)
	if err2 != nil {
		global.LOG.Error(err2)
		app.Error(c,e.ERROR,err2,err2.Error())
		return
	}
	app.OK(c, map[string]interface{}{},"OK")
}

func Login(c *gin.Context) {
	var r Req.Login
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.LOG.Error(err)
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	u := model.SysUser{Username: r.Username, Password: r.Password}
	err1, ur := service.Login(u)
	if err1 != nil {
		global.LOG.Error(err1)
		app.Error(c,e.ERROR,err1,err1.Error())
		return
	}
	rp := util.MD5V([]byte(ur.Salt+u.Password))
	if rp != ur.Password{
		app.OK(c, map[string]interface{}{},"密码错误")
		return
	}
	signedtoken, secret := util.GenToken(ur.ID)
	global.REDIS.Set(signedtoken,secret,time.Second*1800)
	app.OK(c, map[string]interface{}{"token":signedtoken},"密码正确")
}
