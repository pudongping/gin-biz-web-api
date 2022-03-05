package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {

	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认使用的 redis 配置信息
			"default": map[string]interface{}{
				"host":     config.Get("Redis.Host", "127.0.0.1"),
				"port":     config.Get("Redis.Port", 6379),
				"username": config.Get("Redis.Username", ""),
				"password": config.Get("Redis.Password", ""),
				"db":       config.Get("Redis.DB", 0),
			},

			// 缓存专用的 redis 配置信息
			"cache": map[string]interface{}{
				"host":     config.Get("Cache.Host", "127.0.0.1"),
				"port":     config.Get("Cache.Port", 6379),
				"username": config.Get("Cache.Username", ""),
				"password": config.Get("Cache.Password", ""),
				"db":       config.Get("Cache.DB", 1),
			},
		}
	})

}
