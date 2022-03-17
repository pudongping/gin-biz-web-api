// redis 驱动
package cache

import (
	"time"

	"github.com/spf13/cast"

	"gin-biz-web-api/pkg/logger"
	"gin-biz-web-api/pkg/redis"
)

// RedisDriver 实现 cache.Store interface 的 redis 驱动
type RedisDriver struct {
	RedisClient *redis.RdsClient
	KeyPrefix   string
	Group       string
}

// NewRedisDriver 初始化缓存 redis 驱动实例
func NewRedisDriver(rdsClient *redis.RdsClient, keyPrefix, group string) *RedisDriver {
	rs := &RedisDriver{}
	rs.RedisClient = rdsClient // 取出缓存 redis 客户端实例对象
	rs.KeyPrefix = keyPrefix   // 缓存前缀名称
	rs.Group = group           // 这里对应的是 config/redis.go 配置文件中的缓存相关的 key
	return rs
}

func (r *RedisDriver) Set(key string, value interface{}, expiration time.Duration) bool {
	return redis.Set(r.KeyPrefix+key, value, expiration, r.Group)
}

func (r *RedisDriver) Get(key string) string {
	return redis.Get(r.KeyPrefix+key, r.Group)
}

func (r *RedisDriver) Exists(key string) bool {
	return redis.Exists(r.KeyPrefix+key, r.Group)
}

func (r *RedisDriver) Forget(key string) bool {
	return redis.Del(r.KeyPrefix+key, r.Group)
}

func (r *RedisDriver) Forever(key string, value interface{}) bool {
	return redis.Set(r.KeyPrefix+key, value, 0, r.Group)
}

// Flush 这是一个危险操作，会清空 cache 所在的 redis 数据库中所有的数据
func (r *RedisDriver) Flush() bool {
	return redis.FlushDB(r.Group)
}

func (r *RedisDriver) IsAlive() error {
	return r.RedisClient.Ping()
}

func (r *RedisDriver) Increment(parameters ...interface{}) bool {
	return redis.Increment(r.fetchKey(parameters...))
}

func (r *RedisDriver) Decrement(parameters ...interface{}) bool {
	return redis.Decrement(r.fetchKey(parameters...))
}

// fetchKey 递增或者递减时，得到 redis.Increment() 或者 redis.Decrement() 的参数
func (r *RedisDriver) fetchKey(parameters ...interface{}) (string, int64, string) {
	if len(parameters) < 1 {
		logger.FatalString("Cache", "fetchKey", "参数不能为空")
		return "", 0, r.Group
	}

	key := r.KeyPrefix + cast.ToString(parameters[0])

	switch len(parameters) {
	case 1:
		// 当只有一个参数时，第一个参数为 key，默认步长为 1
		return key, 1, r.Group
	case 2:
		// 当有两个参数时，第一个参数为 key，第二个参数为步长
		return key, cast.ToInt64(parameters[1]), r.Group
	default:
		// 超过了两个参数时，直接报错
		logger.FatalString("Cache", "fetchKey", "参数过多")
		return "", 0, r.Group
	}

}
