// document link: https://redis.uptrace.dev/guide/server.html
package redis

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/logger"
)

// rdsClient Redis Client 客户端
type rdsClient struct {
	Client  *redis.Client
	Context context.Context
}

// RdsClientConfig redis 连接配置信息
type RdsClientConfig struct {
	Addr, Username, Password string
	DB                       int
}

// RdsConfigs 配置信息组
type RdsConfigs map[string]*RdsClientConfig

// once 确保全局的 Redis 对象只实例一次（单例模式）
var once sync.Once

// rdsCollections 用于保存多个 redis 实例对象的集合
var rdsCollections map[string]*rdsClient

// Instance 通过组名得到一个 redis 客户端实例对象
func Instance(group ...string) *rdsClient {
	if len(group) > 0 {
		if rds, ok := rdsCollections[group[0]]; ok {
			return rds
		}
		console.Exit("The Redis instance object named [%s] group could not be found!", group[0])
	}

	// 默认使用 config/redis.go 中 default 组下的配置信息
	return rdsCollections["default"]
}

// ConnectRedis 连接 redis 并设置全局的 redis 对象实例集合
func ConnectRedis(configs RdsConfigs) {
	once.Do(func() {

		if rdsCollections == nil {
			rdsCollections = make(map[string]*rdsClient)
		}

		for group, config := range configs {
			// 实例化多个 redis 对象到实例集合中
			rdsCollections[group] = NewClient(
				config.Addr,
				config.Username,
				config.Password,
				config.DB,
			)
		}

	})
}

// NewClient 创建一个新的 redis 连接
func NewClient(address, username, password string, db int) *rdsClient {

	// 初始化自定的 RdsClient 实例
	rds := &rdsClient{}
	// 使用默认的 context
	rds.Context = context.Background()

	// 使用 redis 类库里面的 NewClient 初始化连接
	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 测试一下连接是否正常
	err := rds.Ping()
	logger.LogErrorIf(err)
	console.ExitIf(err)

	return rds
}

// Ping 用以测试 redis 连接是否正常
func (r rdsClient) Ping() error {
	_, err := r.Client.Ping(r.Context).Result()
	return err
}

// genNamespace 获取 redis 的命名空间
func genNamespace() string {
	return config.GetString("app.name") + ":"
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func Set(key string, value interface{}, expiration time.Duration, group ...string) bool {
	instance := Instance(group...)
	key = genNamespace() + key
	if err := instance.Client.Set(instance.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func Get(key string, group ...string) string {
	instance := Instance(group...)
	key = genNamespace() + key
	result, err := instance.Client.Get(instance.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// Exists 判断一个 key 是否存在，内部错误和 redis.Nil 都会返回 false
func Exists(key string, group ...string) bool {
	if result := Get(key, group...); "" == result {
		return false
	}
	return true
}

// Del 删除存储在 redis 里的数据
// 原本可以支持多个 key 传参，但因这里需要兼容配置组，因此只允许一个一个的 key 删除
func Del(key string, group ...string) bool {
	instance := Instance(group...)
	key = genNamespace() + key
	if err := instance.Client.Del(instance.Context, key).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

// FlushDB 清空指定 redis db 里的所有数据
func FlushDB(group ...string) bool {
	instance := Instance(group...)
	if err := instance.Client.FlushDB(instance.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return false
}

// Increment 递增
func Increment(parameters ...interface{}) bool {
	if len(parameters) < 1 {
		return false
	}
	namespace := genNamespace()
	key := namespace + cast.ToString(parameters[0])
	switch len(parameters) {
	case 1:
		// 只有一个参数时，第一个参数为 key
		if err := Instance().Client.Incr(Instance().Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		// 有两个参数时，第一个参数为 key，第二个参数为步长 int64
		value := cast.ToInt64(parameters[1])
		if err := Instance().Client.IncrBy(Instance().Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 3:
		// 有三个参数时，第一个参数为 key，第二个参数为步长 int64，第三个参数为配置组名称
		value := cast.ToInt64(parameters[1])
		group := cast.ToString(parameters[2])
		instance := Instance(group)
		if err := instance.Client.IncrBy(instance.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Increment", "参数不能为空或者过多")
		return false
	}
	return true
}

// Decrement 递减
func Decrement(parameters ...interface{}) bool {
	if len(parameters) < 1 {
		return false
	}
	namespace := genNamespace()
	key := namespace + cast.ToString(parameters[0])
	switch len(parameters) {
	case 1:
		// 只有一个参数时，第一个参数为 key
		if err := Instance().Client.Decr(Instance().Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 2:
		// 有两个参数时，第一个参数为 key，第二个参数为步长 int64
		value := cast.ToInt64(parameters[1])
		if err := Instance().Client.DecrBy(Instance().Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 3:
		// 有三个参数时，第一个参数为 key，第二个参数为步长 int64，第三个参数为配置组名称
		value := cast.ToInt64(parameters[1])
		group := cast.ToString(parameters[2])
		instance := Instance(group)
		if err := instance.Client.DecrBy(instance.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "参数过多")
		return false
	}
	return true
}
