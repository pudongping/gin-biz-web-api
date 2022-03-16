// 数字验证码处理逻辑
package verifycode

import (
	"sync"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/helper/strx"
	"gin-biz-web-api/pkg/logger"
)

type VerifyCode struct {
	Driver Driver
}

var (
	once               sync.Once // once 确保 internalVerifyCode 对象只初始化一次
	internalVerifyCode *VerifyCode
)

// NewVerifyCode 单例模式获取
func NewVerifyCode() *VerifyCode {
	once.Do(func() {

		// 这里的 driver 也可以绑定其他的驱动
		internalVerifyCode = &VerifyCode{Driver: &RedisDriver{
			KeyPrefix: config.GetString("app.name") + ":verify_code:",
			Group:     "default", // 使用默认的 config/redis.go 中的 default 配置连接
		}}

	})

	return internalVerifyCode
}

// generateVerifyCode 生成随机数字验证码
func (v *VerifyCode) generateVerifyCode(key string) string {

	// 生成指定长度的数字随机码作为验证码
	code := strx.StrRandomNumber(config.GetInt("verify_code.length"))

	logger.DebugJSON("VerifyCode", "生成的验证码", map[string]string{key: code})

	// 将 key 和 code 进行保存，方便后续验证
	v.Driver.Set(key, code)

	return code
}

// CheckVerifyCode 检查验证码是否正确
func (v *VerifyCode) CheckVerifyCode(key, answer string) bool {
	logger.DebugJSON("VerifyCode", "检查验证码", map[string]string{key: answer})

	return v.Driver.Verify(key, answer, false)
}