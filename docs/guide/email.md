# 发送邮件

> 使用了扩展包 [go-gomail/gomail](https://github.com/go-gomail/gomail)

## 修改配置文件信息

```yaml

// 修改 etc/config.yaml 中的 Email 配置
// 更多配置详见：config/config.go 文件
Email:
  Host: smtp.qq.com
  Port: 25
  UserName: xxxx@qq.com
  Password: xxxx
  FromAddress: xxx@qq.com
  FromName: alex

```

## 发送邮件目前提供了两种方式

1. 使用封装好的邮件包发送

```go

to := []string{
    "276558492@qq.com",
}

subject := "hello Alex"

body := fmt.Sprintf("<h1> 你好 %v </h1>", 123)

// 发送邮件
err := email.NewMailer().Send(to, subject, body)
fmt.Println(err)

```

2. 直接使用 smtp 邮件驱动发送（实质性第一种方式也是默认采用的第二种方式发送的邮件）

> 这种方式比较适合多种邮件配置信息时使用

```go

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
fmt.Println(err)

```

## 示例文件

参考 `internal/controller/example_ctrl/email_controller.go`