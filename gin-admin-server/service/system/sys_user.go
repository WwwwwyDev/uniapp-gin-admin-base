package system

import (
	"gin-server/global"
	model "gin-server/model/system"
	"gin-server/util"
)

func AddSysUser(u model.SysUser) (err error) {
	u.Password = util.MD5V([]byte(u.Salt + u.Password))
	err = global.DB.Create(&u).Error
	return err
}

func GetSysUserByUsername(username string) (err error, userInter model.SysUser) {
	var user model.SysUser
	err = global.DB.Where("username = ?", username).First(&user).Error
	return err, user
}

func GetSysUserById(id uint) (err error, userInter model.SysUser) {
	var user model.SysUser
	err = global.DB.Where("id = ?", id).First(&user).Error
	return err, user
}

func UpdateSysUser(u model.SysUser) (err error) {
	u.Password = util.MD5V([]byte(u.Salt + u.Password))
	err = global.DB.Save(&u).Error
	return err
}

func DeleteSysUser(u model.SysUser) (err error) {
	err = global.DB.Delete(&u).Error
	return err
}
