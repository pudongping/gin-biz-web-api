// 图像验证码相关配置
package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {
	config.Add("cfg.captcha", func() map[string]interface{} {
		return map[string]interface{}{

			// 验证码图片长度
			"height": config.Get("Captcha.Height", 80),

			// 验证码图片宽度
			"width": config.Get("Captcha.Width", 240),

			// 验证码的长度
			"length": config.Get("Captcha.Length", 6),

			// 数字的最大倾斜角度
			"max_skew": config.Get("Captcha.MaxSkew", 0.7),

			// 图片背景里的混淆点数量
			"dot_count": config.Get("Captcha.DotCount", 80),

			// 过期时间，单位是：分钟
			"expire_time": config.Get("Captcha.ExpireTime", 15),

			// local 环境下的过期时间，方便本地开发调试，单位是：分钟
			"local_expire_time": 10080,

			// local 环境下，使用此 key 可跳过验证，方便测试
			"local_key": "captcha_skip_local",
		}
	})
}
