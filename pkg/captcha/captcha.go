package captcha

import (
	"sync"

	"github.com/mojocn/base64Captcha"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/redis"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// once 确保 internalCaptcha 对象只初始化一次
var once sync.Once

// internalCaptcha 内部使用的 Captcha 对象
var internalCaptcha *Captcha

// NewCaptcha 单例模式获取
func NewCaptcha() *Captcha {
	once.Do(func() {

		// 使用 redis 作为存储图像验证码的存储驱动
		store := RedisDriver{
			RedisClient: redis.Instance(), // 使用默认的 config/redis.go 中的 default 配置连接
			KeyPrefix:   "captcha:",
			Group:       "default", // 使用默认的 config/redis.go 中的 default 配置连接
		}

		// 配置 base64Captcha 驱动信息
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("cfg.captcha.height"),       // 验证码图片长度
			config.GetInt("cfg.captcha.width"),        // 验证码图片宽度
			config.GetInt("cfg.captcha.length"),       // 验证码的长度
			config.GetFloat64("cfg.captcha.max_skew"), // 数字的最大倾斜角度
			config.GetInt("cfg.captcha.dot_count"),    // 图片背景里的混淆点数量
		)

		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}

		// 实例化 base64Captcha 并赋值给内部使用的 internalCaptcha 对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)

	})

	return internalCaptcha
}

// GenerateCaptcha 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id, answer string, isClear ...bool) bool {
	var clear bool
	if len(isClear) > 0 {
		clear = isClear[0]
	} else {
		clear = false
	}

	// 第三个参数是验证后是否删除，我们默认选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, answer, clear)
}
