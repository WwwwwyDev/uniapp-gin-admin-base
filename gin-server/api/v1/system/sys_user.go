package system

import (
	"errors"
	"gin-server/global"
	model "gin-server/model/system"
	req "gin-server/model/system/request"
	service "gin-server/service/system"
	"gin-server/util"
	"gin-server/util/app"
	"gin-server/util/e"
	"gin-server/util/valid/system"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"time"
)

// @Tags SysUser
// @Summary 用户注册
// @Produce  application/json
// @Param data body req.Register true "必填：用户名, 密码, 昵称；非必填：头像链接，电子邮箱，联系电话"
// @Success 200 {string} string "{"code": 20000,"data": {},"msg": "注册成功"}"
// @Router /api/v1/sysUser/register [post]
func Register(c *gin.Context) {
	var r req.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.LOG.Error(err)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	u := model.SysUser{Username: r.Username, Password: r.Password, Salt: uuid.NewV4().String(), NickName: r.NickName, Email: r.Email, Phone: r.Phone, HeaderImg: r.HeaderImg}
	err1 := system.IsValid4sysUser(u)
	if err1 != nil {
		app.Error(c, e.ERROR_REGISTER_FAIL, err1, err1.Error())
		return
	}
	err2:= service.AddSysUser(u)
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
// @Param data body req.Login true "用户名, 密码, 验证码，验证码ID"
// @Success 200 {string} string "{"code": 20000,"data":{},"msg":"登陆成功"}"
// @Router /api/v1/sysUser/login [post]
func Login(c *gin.Context) {

	var r req.Login
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.LOG.Error(err)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	//store := captcha.NewDefaultRedisStore()
	//getCaptcha := store.Get(r.CaptchaId, true)
	//if getCaptcha == ""{
	//	app.Error(c, e.ERROR_LOGIN_FAIL, errors.New("验证码失效"), "验证码失效")
	//	return
	//}
	//if r.Captcha != getCaptcha {
	//	app.Error(c, e.ERROR_LOGIN_FAIL, errors.New("验证码错误"), "验证码错误")
	//	return
	//}
	err1, ur := service.GetSysUserByUsername(r.Username)
	if err1 != nil {
		app.Error(c, e.ERROR_LOGIN_FAIL, err1, "没有此用户")
		return
	}
	rp := util.MD5V([]byte(ur.Salt + r.Password))
	if rp != ur.Password {
		app.Error(c, e.ERROR_LOGIN_FAIL, errors.New("密码错误"), "密码错误")
		return
	}
	signedtoken, secret := util.GenToken(ur.ID,c.ClientIP())
	global.REDIS.Set("SYSUSER_"+signedtoken, secret, time.Second*1800)
	app.OK(c, map[string]interface{}{"token": signedtoken}, "登录成功")
}

// @Tags SysUser
// @Summary 修改密码
// @Produce  application/json
// @Param data body req.ChangePassword true "用户名, 旧密码, 新密码"
// @Success 200 {string} string "{"code": 20000,"data": {},"msg": "修改密码成功"}"
// @Router /api/v1/sysUser/changePassword [post]
func ChangePassword(c *gin.Context){
	var r req.ChangePassword
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.LOG.Error(err)
		app.Error(c, e.ERROR, err, err.Error())
		return
	}
	err1, ur := service.GetSysUserByUsername(r.Username)
	if err1 != nil {
		app.Error(c, e.ERROR_CHANGEPASSWORD_FAIL, err1, "没有此用户")
		return
	}
	rp := util.MD5V([]byte(ur.Salt + r.OldPassword))
	if rp != ur.Password {
		app.Error(c, e.ERROR_CHANGEPASSWORD_FAIL, errors.New("旧密码错误"), "旧密码错误")
		return
	}
	ur.Password = r.NewPassword
	err2 := system.IsValid4sysUser(ur)
	if err2 != nil {
		app.Error(c, e.ERROR_CHANGEPASSWORD_FAIL, err2, err2.Error())
		return
	}
	err3 := service.UpdateSysUser(ur)
	if err3 != nil {
		global.LOG.Error(err3)
		app.Error(c, e.ERROR, err3, err3.Error())
		return
	}
	app.OK(c, map[string]interface{}{}, "修改密码成功")
}
