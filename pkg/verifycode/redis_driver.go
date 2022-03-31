// 数字验证码 redis 存储驱动
package verifycode

import (
	"time"

	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/redis"
)

// RedisDriver 实现了 verifycode.Driver interface
type RedisDriver struct {
	KeyPrefix string // 缓存前缀名称
	Group     string // 这里对应的是 config/redis.go 配置文件中的缓存相关的 key
}

// Set 实现了 verifycode.Driver interface 的 Set 方法
func (r *RedisDriver) Set(key, value string) bool {

	expireTime := time.Minute * time.Duration(config.GetInt64("cfg.verify_code.expire_time"))

	// 方便本地开发调试，如果不需要，可以将这段代码删掉
	if app.IsLocal() {
		expireTime = time.Minute * time.Duration(config.GetInt64("cfg.verify_code.local_expire_time"))
	}

	return redis.Set(r.KeyPrefix+key, value, expireTime, r.Group)
}

// Get 实现了 verifycode.Driver interface 的 Get 方法
func (r *RedisDriver) Get(key string, isClear bool) string {
	key = r.KeyPrefix + key
	val := redis.Get(key, r.Group)

	if isClear {
		redis.Del(key, r.Group)
	}

	return val
}

// Verify 实现 verifycode.Driver interface 的 Verify 方法
func (r *RedisDriver) Verify(key, answer string, isClear bool) bool {
	return answer == r.Get(key, isClear)
}
