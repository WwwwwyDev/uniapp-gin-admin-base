package system

import (

	"gin-server/global"
	"gin-server/model/system"
	"gin-server/util"

)

func Register(u system.SysUser) (err error, userInter system.SysUser) {
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = util.MD5V([]byte(u.Salt+u.Password))
	err = global.DB.Create(&u).Error
	return err, u
}

func Login(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	err = global.DB.Where("username = ?", u.Username).First(&user).Error
	return err, user
}