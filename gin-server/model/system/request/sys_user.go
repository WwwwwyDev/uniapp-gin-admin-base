package request

type Register struct {
	Username  string `json:"userName"`
	Password  string `json:"passWord"`
	NickName  string `json:"nickName"`
	HeaderImg string `json:"headerImg"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}
