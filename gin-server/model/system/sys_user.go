package system

import "gin-server/global"

type SysUser struct {
	global.MODEL
	Username  string `json:"userName" gorm:"comment:用户登录名;not null"`    // 用户登录名
	Password  string `json:"password"  gorm:"comment:用户登录密码;not null"`  // 用户登录密码
	Salt      string `json:"salt"  gorm:"comment:密码盐;not null"`         // 密码盐
	NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Email     string `json:"nickName" gorm:"default:123456@qq.com;comment:用户邮箱"`
	Phone     string `json:"nickName" gorm:"default:8888888888;comment:用户电话"`
	HeaderImg string `json:"headerImg" gorm:"default:http://uk3b3.cn/YGrF8;comment:用户头像"` // 用户头像
	Is_del    bool
}
