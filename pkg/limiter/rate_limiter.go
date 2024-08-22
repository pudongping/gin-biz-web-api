// 处理请求限流逻辑
// document link：https://github.com/ulule/limiter-examples/blob/master/gin/main.go
package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	limiterLib "github.com/ulule/limiter/v3"
	limiterRedis "github.com/ulule/limiter/v3/drivers/store/redis"

	"gin-biz-web-api/pkg/logger"
	"gin-biz-web-api/pkg/redis"
)

// GetKeyIP 获取客户端的 ip 地址（针对全局做频率限制）
func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

// GetKeyRouteWithIP 获取客户端请求的地址加 ip 地址（针对单个路由进行频率限制）
func GetKeyRouteWithIP(c *gin.Context) string {
	return c.ClientIP() + "|" + routeToKeyString(c.FullPath())
}

// routeToKeyString 将 url 中所有的分隔符替换成以 `-` 作为分割
func routeToKeyString(urlPath string) string {
	urlPath = strings.ReplaceAll(urlPath, "/", "-") // 将 url 中所有的 `/` 替换成 `-`
	urlPath = strings.ReplaceAll(urlPath, ":", "-") // 将 url 中所有的 `:` 替换成 `-`
	return urlPath
}

// CheckRate 检查请求是否超频
// key 为缓存 key
// formatted 为频率限制格式
// * "S": second
// * "M": minute
// * "H": hour
// * "D": day
//
// Examples:
//
// * 5 reqs/second: "5-S"
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"
func CheckRate(c *gin.Context, key, formatted string) (limiterLib.Context, error) {
	var ctx limiterLib.Context

	// 实例化 limiterLib.Rate 对象
	rate, err := limiterLib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogErrorIf(err)
		return ctx, err
	}

	store, err := limiterRedis.NewStoreWithOptions(
		redis.Instance().Client, // 将我们自定义的 redis 客户端连接绑定到 limiter 存储驱动上
		limiterLib.StoreOptions{
			Prefix: redis.GenNamespace("limiter"), // 设置限流器的缓存前缀
		},
	)
	if err != nil {
		logger.LogErrorIf(err)
		return ctx, err
	}

	instance := limiterLib.New(store, rate)

	// 获取限流处理结果
	if c.GetBool("rate-limiter-once") {
		// Peek() 取结果，不增加访问次数
		return instance.Peek(c, key)
	} else {
		// 确保多个路由组里调用 LimitIP 进行限流时，只增加一次访问次数
		c.Set("rate-limiter-once", true)

		// Get() 取结果且增加访问次数
		return instance.Get(c, key)
	}

}
