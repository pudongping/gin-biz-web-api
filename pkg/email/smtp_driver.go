// smtp 邮件发送驱动
package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

// SMTP smtp 驱动服务器发送邮件
type SMTP struct {
	*MailInfo
}

// SendMail 发送邮件信息
// to 收件人邮箱地址数组
// subject 邮件主题
// body 邮件内容
func (s *SMTP) Send(to []string, subject, body string) error {
	m := gomail.NewMessage() // 创建一个消息实例

	// m.SetHeader("From", s.FromAddress) // 发件人
	m.SetAddressHeader("From", s.FromAddress, s.FromName) // 发件人加发件人名称
	m.SetHeader("To", to...)                              // 收件人
	m.SetHeader("Subject", subject)                       // 邮件主题
	m.SetBody("text/html", body)                          // 邮件正文
	// m.SetBody("text/plain", body)       // 邮件正文

	dialer := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password) // 创建一个新的 SMTP 拨号实例
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: "ssl" == s.Encryption}

	return dialer.DialAndSend(m) // 打开与 SMTP 服务器的连接并发送电子邮件
}
