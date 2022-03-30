package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// TokenBucketLimiterInterface 定义当前令牌桶必须要实现的方法
type TokenBucketLimiterInterface interface {
	Key(c *gin.Context) string                                              // 获取对应的限流器的键值对名称
	GetBucket(key string) (*ratelimit.Bucket, bool)                         // 获取令牌桶
	AddBuckets(rules ...TokenBucketLimiterRule) TokenBucketLimiterInterface // 新增多个令牌桶
}

// TokenBucketLimiter 用于存储令牌桶与键值对名称的映射关系，可以存储多个令牌桶
type TokenBucketLimiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// TokenBucketLimiterRule 用于存储令牌桶的一些相应规则属性
type TokenBucketLimiterRule struct {
	Key          string        // 自定义键值对名称
	FillInterval time.Duration // 间隔多久时间放 N 个令牌
	Capacity     int64         // 令牌桶的容量
	Quantum      int64         // 每次到达间隔时间后所放的具体令牌数量
}
