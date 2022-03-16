package example_ctrl

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/internal/requests/example_request"
	"gin-biz-web-api/pkg/captcha"
	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/responses"
	"gin-biz-web-api/pkg/validator"
)

type CaptchaController struct {
}

// ShowCaptcha 显示图像验证码
// curl --location --request GET 'localhost:3000/api/example/show-captcha'
func (ctrl *CaptchaController) ShowCaptcha(c *gin.Context) {
	response := responses.New(c)

	// 生成图像验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	if err != nil {
		response.ToErrorResponse(errcode.InternalServerError.WithDetails(err.Error()), "图像验证码生成失败")
		return
	}

	response.ToResponse(gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})

}

// VerifyCaptchaCode 验证图像验证码
// curl --location --request POST 'localhost:3000/api/example/verify-captcha-code' \
// --header 'Content-Type: application/x-www-form-urlencoded' \
// --data-urlencode 'captcha_answer=977448' \
// --data-urlencode 'captcha_id=0uXDqoQmkHYfuctD1tUS'
func (ctrl *CaptchaController) VerifyCaptchaCode(c *gin.Context) {
	response := responses.New(c)

	// 表单验证
	request := example_request.VerifyCaptchaCodeRequest{}
	if ok := validator.BindAndValidate(c, &request, example_request.VerifyCaptchaCode); !ok {
		return
	}

	response.ToResponse(nil)

}
