package middleware

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/limiter"
	"gin-biz-web-api/pkg/responses"
)

// LimitMethodTokenBucket 处理令牌桶限流控制
func LimitMethodTokenBucket(methodLimiters limiter.TokenBucketLimiterInterface, k ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 这里的 key，默认取的是路由地址，eg：`/api/user`
		key := methodLimiters.Key(c)

		if len(k) > 0 {
			key = k[0] // 如果指定了 key，则通过指定的 key 去获取对应的令牌桶
		}

		// 从存储的众多令牌桶中取出指定 key 对应的令牌桶
		if bucket, ok := methodLimiters.GetBucket(key); ok {

			// TakeAvailable 它会占用存储桶中立即可用的令牌的数量，返回值为删除的令牌数，如果没有可用的令牌，将会返回 0，也就是已经超出配额了
			count := bucket.TakeAvailable(1)
			if count == 0 {
				// 已经超额
				responses.New(c).ToErrorResponse(errcode.TooManyRequests)
				return
			}
		}

		c.Next()

	}
}
