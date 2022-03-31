// 邮件相关配置信息
package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {
	config.Add("cfg.email", func() map[string]interface{} {
		return map[string]interface{}{

			// 使用支持 ESMTP 的 SMTP 服务器发送邮件
			"driver": "smtp", // 目前仅支持 smtp 驱动

			"smtp": map[string]interface{}{
				"host":       config.Get("Email.Host", "localhost"), // SMTP 服务器地址
				"port":       config.Get("Email.Port", 25),          // SMTP 服务器端口
				"username":   config.Get("Email.UserName", ""),      // 账号
				"password":   config.Get("Email.Password", ""),      // 密码
				"encryption": config.Get("Email.Encryption", "ssl"), // 加密类型，ssl 或 tls
			},

			// 发件人信息
			"form": map[string]interface{}{
				"address": config.Get("Email.FromAddress", ""), // 发件者地址
				"name":    config.Get("Email.FromName", ""),    // 发送者名称
			},
		}
	})
}
