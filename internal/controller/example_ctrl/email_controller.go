package example_ctrl

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/email"
	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/responses"
)

type EmailController struct {
}

// SendEmail 发送邮件
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
func (ctrl *EmailController) SendMailer(c *gin.Context) {
	response := responses.New(c)

	to := []string{
		"276558492@qq.com",
	}

	subject := "邮箱发送测试"

	body := fmt.Sprintf("<h1> Hello World! </h1></br><span color=\"red;\">Alex!</span>")

	mail := email.SMTP{MailInfo: &email.MailInfo{
		Host:        config.GetString("email.smtp.host"),
		Port:        config.GetInt("email.smtp.port"),
		Username:    config.GetString("email.smtp.username"),
		Password:    config.GetString("email.smtp.password"),
		Encryption:  config.GetString("email.smtp.encryption"),
		FromAddress: config.GetString("email.form.address"),
		FromName:    config.GetString("email.form.name"),
	}}

	err := mail.Send(to, subject, body)

	if err != nil {
		response.ToErrorResponse(errcode.InternalServerError.WithDetails(err.Error()), "邮件发送失败")
		return
	}

	response.ToResponse(nil)
}
