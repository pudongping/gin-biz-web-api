// [juju/ratelimit 令牌桶限流器分析](https://maratrix.cn/post/2021/01/10/juju-ratelimit-read/)
package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type TokenBucketMethodLimiter struct {
	*TokenBucketLimiter
}

func NewTokenBucketMethodLimiter() TokenBucketLimiterInterface {
	l := &TokenBucketLimiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}
	return TokenBucketMethodLimiter{TokenBucketLimiter: l}
}

// Key 获取对应的限流器的键值对名称
func (t TokenBucketMethodLimiter) Key(c *gin.Context) string {
	// 只有请求地址，不带参数 eg：`/api/user`
	return c.Request.URL.Path
}

// GetBucket 通过 key 获取指定的令牌桶
func (t TokenBucketMethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := t.TokenBucketLimiter.limiterBuckets[key]
	return bucket, ok
}

// AddBuckets 新增多个令牌桶
func (t TokenBucketMethodLimiter) AddBuckets(rules ...TokenBucketLimiterRule) TokenBucketLimiterInterface {

	for _, rule := range rules {
		if _, ok := t.TokenBucketLimiter.limiterBuckets[rule.Key]; !ok {
			// ratelimit.NewBucketWithQuantum 初始化限流对象
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInterval, // 间隔多久时间放 N 个令牌
				rule.Capacity,     // 令牌桶的容量
				rule.Quantum,      // 每次到达间隔时间后所放的具体令牌数量
			)
			t.TokenBucketLimiter.limiterBuckets[rule.Key] = bucket
		}
	}

	return t
}
