// 验证码存储驱动
package captcha

import (
	"errors"
	"time"

	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/redis"
)

// RedisDriver 实现 base64Captcha.Store interface
type RedisDriver struct {
	RedisClient *redis.RdsClient
	KeyPrefix   string // 缓存前缀名称
	Group       string // 这里对应的是 config/redis.go 配置文件中的缓存相关的 key
}

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (r *RedisDriver) Set(key, value string) error {

	expireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))

	// 方便本地开发调试
	if app.IsLocal() {
		expireTime = time.Minute * time.Duration(config.GetInt64("captcha.local_expire_time"))
	}

	if ok := redis.Set(r.KeyPrefix+key, value, expireTime, r.Group); !ok {
		return errors.New("无法存储图片验证码结果")
	}

	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (r *RedisDriver) Get(key string, isClear bool) string {
	key = r.KeyPrefix + key
	val := redis.Get(key, r.Group)

	if isClear {
		redis.Del(key, r.Group)
	}

	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (r *RedisDriver) Verify(key, answer string, isClear bool) bool {
	return answer == r.Get(key, isClear)
}
