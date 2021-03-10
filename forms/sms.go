package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	//注册发送短信验证码和动态验证码登录发送验证码 1-注册 2-登录
	Type uint `form:"type" json:"type" binding:"required,oneof=1 2"`
}
