package util

import (
	"gin-server/global"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type Claims struct {
	UserID  uint `json:"userID"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID uint) (string, string) {
	// 创建一个我们自己的声明
	c := Claims{
		userID, // 自定义字段
		jwt.StandardClaims{
			Issuer:    global.CONFIG.Jwt.Issuer,              // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	secret := uuid.NewV4().String()
	signedtoken, err := token.SignedString([]byte(secret))
	if err != nil {
		global.LOG.Error(err)
	}
	return signedtoken, secret
}

// ParseToken 解析JWT
func ParseToken(tokenString string, secret string) (*Claims){
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(secret), nil
	})
	if err != nil {
		global.LOG.Error(err)
		return nil
	}
	return token.Claims.(*Claims)
}
