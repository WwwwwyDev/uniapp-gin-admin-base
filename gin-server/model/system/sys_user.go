package system

import (
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	Username  string `json:"userName" gorm:"type:varchar(191);comment:用户登录名;unique;not null"` // 用户登录名
	Password  string `json:"password"  gorm:"type:varchar(191);comment:用户登录密码;not null"`      // 用户登录密码
	Salt      string `json:"salt"  gorm:"type:varchar(191);comment:密码盐;not null"`             // 密码盐
	NickName  string `json:"nickName" gorm:"type:varchar(191);comment:用户昵称;not null"`         // 用户昵称
	Email     string `json:"email" gorm:"comment:用户邮箱"`
	Phone     string `json:"phone" gorm:"comment:用户电话"`
	HeaderImg string `json:"headerImg" gorm:"default:http://uk3b3.cn/YGrF8;comment:用户头像"` // 用户头像
	IsDel     bool   `json:"isDel" gorm:"type:bool;default:false;comment:是否删除"`           // 是否删除
}
