package example_request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"

	"gin-biz-web-api/pkg/validator"
)

type VerifyCaptchaCodeRequest struct {
	CaptchaID     string `form:"captcha_id" json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `form:"captcha_answer" json:"captcha_answer,omitempty" valid:"captcha_answer"`
}

// VerifyCaptchaCode 验证图片验证码的结果，返回长度等于 0 即表示通过
func VerifyCaptchaCode(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs := validator.ValidateStruct(data, rules, messages)

	req := data.(*VerifyCaptchaCodeRequest)
	errs = validator.ValidateCaptcha(req.CaptchaID, req.CaptchaAnswer, errs)

	return errs
}
