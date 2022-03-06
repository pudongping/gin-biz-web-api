package bootstrap

import (
	"gin-biz-web-api/pkg/cache"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/redis"
)

// setupCache 初始化缓存
func setupCache() {

	console.Info("init cache ...")

	group := "cache"
	// 实例化 redis 缓存驱动
	redisStoreDriver := cache.NewRedisDriver(redis.Instance(group), "cache:", group)
	// 将 redis 缓存驱动绑定到 cache.Cache 对象上
	cache.InitWithCacheStore(redisStoreDriver)
}
