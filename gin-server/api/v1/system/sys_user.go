package system

import (
	"errors"
	"gin-server/global"
	model "gin-server/model/system"
	Req "gin-server/model/system/request"
	service "gin-server/service/system"
	"gin-server/util"
	"gin-server/util/app"
	"gin-server/util/captcha"
	"gin-server/util/e"
	"gin-server/util/valid"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

// @Tags SysUser
// @Summary 用户注册
// @Produce  application/json
// @Param data body Req.Register true "必填：用户名, 密码, 昵称；非必填：头像链接，电子邮箱，联系电话"
// @Success 200 {string} string "{"code": 20000,"data": {},"msg": "密码正确"}"
// @Router /api/v1/sysUser/register [post]
func Register(c *gin.Context) {
	var r Req.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.LOG.Error(err)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	u := model.SysUser{Username: r.Username, Password: r.Password, Salt: uuid.NewV4().String(), NickName: r.NickName, Email: r.Email, Phone: r.Phone, HeaderImg: r.HeaderImg}
	err1 := valid.IsValid4sysUser(u)
	if err1 != nil {
		app.Error(c, e.ERROR_REGISTER_FAIL, err1, err1.Error())
		return
	}
	err2, _ := service.Register(u)
	if err2 != nil {
		global.LOG.Error(err2)
		app.Error(c, e.ERROR, err2, err2.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "注册成功")
}

// @Tags SysUser
// @Summary 用户登录
// @Produce  application/json
// @Param data body Req.Login true "用户名, 密码, 验证码，验证码ID"
// @Success 200 {string} string "{"code": 20000,"data":{},"msg":"登陆成功"}"
// @Router /api/v1/sysUser/login [post]
func Login(c *gin.Context) {
	var r Req.Login
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.LOG.Error(err)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	store := captcha.NewDefaultRedisStore()
	getCaptcha := store.Get(r.CaptchaId, true)
	if getCaptcha == ""{
		app.Error(c, e.ERROR_LOGIN_FAIL, errors.New("验证码失效"), "验证码失效")
		return
	}
	if r.Captcha != getCaptcha {
		app.Error(c, e.ERROR_LOGIN_FAIL, errors.New("验证码错误"), "验证码错误")
		return
	}
	u := model.SysUser{Username: r.Username, Password: r.Password}
	err1, ur := service.Login(u)
	if err1 != nil {
		global.LOG.Error(err1)
		app.Error(c, e.ERROR_LOGIN_FAIL, err1, "没有此用户")
		return
	}
	rp := util.MD5V([]byte(ur.Salt + u.Password))
	if rp != ur.Password {
		app.Error(c, e.ERROR_LOGIN_FAIL, errors.New("密码错误"), "密码错误")
		return
	}
	signedtoken, secret := util.GenToken(ur.ID)
	global.REDIS.Set("SYSUSER_"+signedtoken, secret, time.Second*1800)
	app.OK(c, map[string]interface{}{"token": signedtoken}, "登录成功")
}
