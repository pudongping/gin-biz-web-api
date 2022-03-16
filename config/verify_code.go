// 验证码相关配置信息
package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {
	config.Add("verify_code", func() map[string]interface{} {
		return map[string]interface{}{

			// 验证码的长度
			"length": config.Get("VerifyCode.Length", 6),

			// 过期时间，单位是：分钟
			"expire_time": config.Get("VerifyCode.ExpireTime", 15),

			// local 环境下的过期时间，方便本地开发调试，单位是：分钟
			"local_expire_time": 10080,
			// local 环境下的验证码，默认为：123456
			"local_code": 123456,
		}
	})
}
