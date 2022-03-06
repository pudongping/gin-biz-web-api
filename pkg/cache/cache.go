// package cache 缓存工具类，可以缓存各种类型包括 struct 对象
package cache

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/spf13/cast"

	"gin-biz-web-api/pkg/logger"
)

type CacheService struct {
	Driver Store
}

var once sync.Once
var Cache *CacheService

// InitWithCacheStore 将缓存驱动绑定到 cache.Cache 对象变量上
func InitWithCacheStore(storeDriver Store) {
	once.Do(func() {
		Cache = &CacheService{Driver: storeDriver}
	})
}

func Set(key string, obj interface{}, expireTime time.Duration) {
	b, err := json.Marshal(&obj)
	logger.LogErrorIf(err)
	Cache.Driver.Set(key, string(b), expireTime)
}

func Get(key string) interface{} {
	strValue := Cache.Driver.Get(key)
	var wanted interface{}
	err := json.Unmarshal([]byte(strValue), &wanted)
	logger.LogErrorIf(err)
	return wanted
}

func Exists(key string) bool {
	return Cache.Driver.Exists(key)
}

// GetObject 从缓存中获取一个对象实例（第二个参数应该传地址）
// 用法如下：
// model := user.User{}
// cache.GetObject("key", &model)
func GetObject(key string, wanted interface{}) {
	val := Cache.Driver.Get(key)
	if len(val) > 0 {
		err := json.Unmarshal([]byte(val), &wanted)
		logger.LogErrorIf(err)
	}
}

func GetString(key string) string {
	return cast.ToString(Get(key))
}

func GetBool(key string) bool {
	return cast.ToBool(Get(key))
}

func GetInt(key string) int {
	return cast.ToInt(Get(key))
}

func GetInt32(key string) int32 {
	return cast.ToInt32(Get(key))
}

func GetInt64(key string) int64 {
	return cast.ToInt64(Get(key))
}

func GetUint(key string) uint {
	return cast.ToUint(Get(key))
}

func GetUint32(key string) uint32 {
	return cast.ToUint32(Get(key))
}

func GetUint64(key string) uint64 {
	return cast.ToUint64(Get(key))
}

func GetFloat64(key string) float64 {
	return cast.ToFloat64(Get(key))
}

func GetTime(key string) time.Time {
	return cast.ToTime(Get(key))
}

func GetDuration(key string) time.Duration {
	return cast.ToDuration(Get(key))
}

func GetIntSlice(key string) []int {
	return cast.ToIntSlice(Get(key))
}

func GetStringSlice(key string) []string {
	return cast.ToStringSlice(Get(key))
}

func GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(Get(key))
}

func GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(Get(key))
}

func GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(Get(key))
}

func Forget(key string) {
	Cache.Driver.Forget(key)
}

func Forever(key string, value interface{}) {
	Cache.Driver.Forever(key, value)
}

// Flush 这是一个危险操作，会清空 cache 所在的 redis 数据库中所有的数据
func Flush() {
	Cache.Driver.Flush()
}

func IsAlive() error {
	return Cache.Driver.IsAlive()
}

func Increment(parameters ...interface{}) {
	Cache.Driver.Increment(parameters...)
}

func Decrement(parameters ...interface{}) {
	Cache.Driver.Decrement(parameters...)
}
