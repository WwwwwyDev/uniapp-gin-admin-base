package configModel

type Config struct {
	Mysql   Mysql
	Redis   Redis
	Runtime Runtime
	Server  Server
	Jwt     Jwt
	Captcha Captcha
}