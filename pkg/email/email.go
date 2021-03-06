// 邮件发送
package email

import (
	"sync"

	"gin-biz-web-api/pkg/config"
)

type Mailer struct {
	Driver
}

var (
	once           sync.Once
	internalMailer *Mailer
)

// NewMailer 单例模式获取驱动
func NewMailer() *Mailer {
	once.Do(func() {

		internalMailer = &Mailer{
			// 使用 email.SMTP 驱动绑定驱动
			Driver: &SMTP{&MailInfo{
				Host:        config.GetString("cfg.email.smtp.host"),
				Port:        config.GetInt("cfg.email.smtp.port"),
				Username:    config.GetString("cfg.email.smtp.username"),
				Password:    config.GetString("cfg.email.smtp.password"),
				Encryption:  config.GetString("cfg.email.smtp.encryption"),
				FromAddress: config.GetString("cfg.email.form.address"),
				FromName:    config.GetString("cfg.email.form.name"),
			}},
		}

	})

	return internalMailer
}

// SendMail 发送邮件信息
// to 收件人邮箱地址数组
// subject 邮件主题
// body 邮件内容
func (m *Mailer) SendMail(to []string, subject, body string) error {
	return m.Driver.Send(to, subject, body)
}
