package system

import (
	"errors"
	"gin-server/model/system"
	"regexp"
)

func IsValid4sysUser(u system.SysUser) error {
	reUsername := regexp.MustCompile("^.{4,16}$")
	if !reUsername.MatchString(u.Username) {
		return errors.New("用户名不合法,长度为4-16位")
	}
	rePassword := regexp.MustCompile("^.{6,18}$")
	if !rePassword.MatchString(u.Password) {
		return errors.New("密码不合法,长度为6-18位")
	}
	reNickName := regexp.MustCompile("^.{2,20}$")
	if !reNickName.MatchString(u.NickName) {
		return errors.New("昵称不合法,长度为2-20位")
	}
	rePhone := regexp.MustCompile("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$")
	if u.Phone != "" && !rePhone.MatchString(u.Phone) {
		return errors.New("手机号不合法")
	}
	reEmail := regexp.MustCompile("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")
	if u.Email != "" && !reEmail.MatchString(u.Email) {
		return errors.New("邮箱不合法")
	}
	return nil
}
