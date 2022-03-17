package cache

import (
	"time"
)

type Store interface {

	// Set 存储 key 对应的 value，且设置 expiration 过期时间
	Set(key string, value interface{}, expiration time.Duration) bool

	// Get 获取 key 对应的 value
	Get(key string) string

	// Exists 判断一个 key 是否存在，内部错误和 redis.Nil 都会返回 false
	Exists(key string) bool

	// Forget 删除指定 key 对应的缓存
	Forget(key string) bool

	// Forever 持久化
	Forever(key string, value interface{}) bool

	// Flush 清空数据库
	Flush() bool

	// IsAlive 检查连接是否活跃
	IsAlive() error

	// Increment 递增
	// 只允许传递最多 2 个参数，不传参时，不做任何处理
	// 当参数只有 1 个时，为 key，增加 1。
	// 当参数有 2 个时，第一个参数为 key ，第二个参数为要增加的值，即为步长。（int64 类型）
	Increment(parameters ...interface{}) bool

	// Decrement 递减
	// 只允许传递最多 2 个参数，不传参时，不做任何处理
	// 当参数只有 1 个时，为 key，减去 1。
	// 当参数有 2 个时，第一个参数为 key ，第二个参数为要减去的值，即为步长。（int64 类型）
	Decrement(parameters ...interface{}) bool
}
