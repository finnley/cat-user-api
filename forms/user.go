package forms

type PasswordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` // 手机号码格式需要自定义validator
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=4,max=4"` // 验证码
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}
