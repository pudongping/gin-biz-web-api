// jwt 相关配置信息
package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{

			// jwt 加密 key
			"key": config.Get("JWT.Key"),

			// 过期时间，单位是分钟，一般不超过两个小时
			"expire_time": config.Get("JWT.ExpireTime", 120),

			// local 环境下的过期时间，方便本地开发调试
			"local_expire_time": 86400,

			// 允许刷新时间，单位分钟，86400 为两个月，从 Token 的签名时间算起
			"max_refresh_time": config.Get("JWT.MaxRefreshTime", 86400),
		}
	})
}
