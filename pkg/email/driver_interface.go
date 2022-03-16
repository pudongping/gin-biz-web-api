package email

// MailInfo 邮件驱动所必须要的配置信息
type MailInfo struct {
	Host        string // SMTP 服务器地址
	Port        int    // SMTP 服务器端口
	Username    string // 账号
	Password    string // 密码
	Encryption  string // 加密类型，ssl 或 tls
	FromAddress string // 发件人地址
	FromName    string // 发件人名称
}

// Driver 邮件驱动接口
type Driver interface {

	// Send 发送邮件信息
	// to 收件人邮箱地址数组
	// subject 邮件主题
	// body 邮件内容
	Send(to []string, subject, body string) error
}
