package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认数据库
			"driver": config.Get("DB.Driver", "mysql"),

			"mysql": map[string]interface{}{

				// 数据库连接信息
				"host":     config.Get("DB.Host", "127.0.0.1"),
				"port":     config.Get("DB.Port", 3306),
				"database": config.Get("DB.Database", ""),
				"username": config.Get("DB.Username", ""),
				"password": config.Get("DB.Password", ""),
				"charset":  config.Get("DB.Charset", ""),

				// 连接池配置
				"max_open_connections": config.Get("DB.MaxOpenConnections", 25),  // 最大连接数
				"max_idle_connections": config.Get("DB.MaxIdleConnections", 100), // 最大空闲连接数
				"max_life_seconds":     config.Get("DB.MaxLifeSeconds", 5*60),    // 每个链接的过期时间

			},
		}
	})
}
