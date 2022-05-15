package auth_request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"

	"gin-biz-web-api/pkg/validator"
)

// SignupUsingEmailRequest 通过邮箱注册的请求信息
type SignupUsingEmailRequest struct {
	Email           string `form:"email" json:"email,omitempty" valid:"email"`
	VerifyCode      string `form:"verify_code" json:"verify_code,omitempty" valid:"verify_code"`
	Account         string `form:"account" json:"account,omitempty" valid:"account"`
	Password        string `form:"password" json:"password,omitempty" valid:"password"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm,omitempty" valid:"password_confirm"`
}

// SignupUsingEmail 通过邮箱注册验证器方法
func SignupUsingEmail(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"account":          []string{"required", "alpha_num", "between:3,20", "not_exists:users,account"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"account": []string{
			"required:账号为必填项",
			"alpha_num:账号格式错误，只允许数字和英文",
			"between:账号长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码为必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validator.ValidateStruct(data, rules, messages)

	req := data.(*SignupUsingEmailRequest)
	errs = validator.ValidationPasswordConfirm(req.Password, req.PasswordConfirm, errs) // 验证两次密码是否一致
	errs = validator.ValidationVerifyCode(req.Email, req.VerifyCode, errs)              // 验证邮件验证码是否正确

	return errs
}
