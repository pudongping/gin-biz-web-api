// 限流中间件，支持全局或者单独路由进行限流
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/limiter"
	"gin-biz-web-api/pkg/logger"
	"gin-biz-web-api/pkg/responses"
)

// LimitIP 针对 IP 进行限流，比较适合做全局限流中间件
// limit 为频率限制格式
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
//
func LimitIP(limit string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 针对 IP 限流
		key := limiter.GetKeyIP(c)
		if ok := handlerLimit(c, key, limit); !ok {
			return
		}

		c.Next()
	}
}

// LimitRoute 用于单独的路由
func LimitRoute(limit string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 针对单个路由，增加访问次数
		c.Set("rate-limiter-once", false)

		// 针对 IP 加路由进行限流
		key := limiter.GetKeyRouteWithIP(c)
		if ok := handlerLimit(c, key, limit); !ok {
			return
		}

		c.Next()

	}
}

// handlerLimit 处理请求频率限制
func handlerLimit(c *gin.Context, key, limit string) bool {

	response := responses.New(c)

	// 获取限流结果
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogErrorIf(err)
		response.ToErrorResponse(errcode.InternalServerError.WithDetails(err.Error()))
		return false
	}

	// 设置返回头信息
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))         // 在指定时间内允许的最大请求次数
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining)) // 在指定时间段内剩下的请求次数
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))         // 距离下次重试请求需要等待的时间点（s）

	// 如果请求已经超额
	if rate.Reached {
		// 提示用户已经超额
		response.ToErrorResponse(errcode.TooManyRequests)
		return false
	}

	return true
}
