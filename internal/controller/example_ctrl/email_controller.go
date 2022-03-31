package example_ctrl

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/email"
	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/responses"
	"gin-biz-web-api/pkg/verifycode"
)

type EmailController struct {
}

// SendEmail 发送邮件
// curl --location --request POST 'localhost:3000/api/example/send-email'
func (ctrl *EmailController) SendEmail(c *gin.Context) {
	response := responses.New(c)

	to := []string{
		"276558492@qq.com",
	}

	subject := "hello Alex"

	body := fmt.Sprintf("<h1> 你好 %v </h1>", 123)

	// 发送邮件
	err := email.NewMailer().Send(to, subject, body)

	if err != nil {
		response.ToErrorResponse(errcode.InternalServerError.WithDetails(err.Error()), "邮件发送失败")
		return
	}

	response.ToResponse(nil)
}

// SendMailer 使用 email 驱动发送邮件
// curl --location --request POST 'localhost:3000/api/example/send-mailer'
func (ctrl *EmailController) SendMailer(c *gin.Context) {
	response := responses.New(c)

	to := []string{
		"276558492@qq.com",
	}

	subject := "邮箱发送测试"

	body := fmt.Sprintf("<h1> Hello World! </h1></br><span color=\"red;\">Alex!</span>")

	mail := email.SMTP{MailInfo: &email.MailInfo{
		Host:        config.GetString("cfg.email.smtp.host"),
		Port:        config.GetInt("cfg.email.smtp.port"),
		Username:    config.GetString("cfg.email.smtp.username"),
		Password:    config.GetString("cfg.email.smtp.password"),
		Encryption:  config.GetString("cfg.email.smtp.encryption"),
		FromAddress: config.GetString("cfg.email.form.address"),
		FromName:    config.GetString("cfg.email.form.name"),
	}}

	err := mail.Send(to, subject, body)

	if err != nil {
		response.ToErrorResponse(errcode.InternalServerError.WithDetails(err.Error()), "邮件发送失败")
		return
	}

	response.ToResponse(nil)
}

// SendEmailVerifyCode 发送邮件验证码
// curl --location --request POST 'localhost:3000/api/example/send-email-verify-code'
func (ctrl *EmailController) SendEmailVerifyCode(c *gin.Context) {
	response := responses.New(c)

	err := verifycode.NewVerifyCode().SendEmailVerifyCode("276558492@qq.com")

	if err != nil {
		response.ToErrorResponse(errcode.InternalServerError.WithDetails(err.Error()), "邮件验证码发送失败")
		return
	}

	response.ToResponse(nil)
}
