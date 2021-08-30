package request

type Register struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	NickName  string `json:"nickName"`  //昵称
	HeaderImg string `json:"headerImg"` //头像链接
	Email     string `json:"email"`     //电子邮箱
	Phone     string `json:"phone"`     //联系电话
}

type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type ChangePassword struct {
	Username  string `json:"username"`  // 用户名
	OldPassword  string `json:"oldPassword"`  // 旧密码
	NewPassword   string `json:"newPassword"`   // 新密码
}