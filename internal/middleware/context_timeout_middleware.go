package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// ContextTimeout 上下文超时时间
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 调用 context.WithTimeout() 方法设置当前 context 的超时时间
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		// 将设置了超时时间的 context 重新赋予给了 gin.Context
		// 这样，在当前请求运行到指定的时间后，在使用了该 context 的运行流程就会
		// 针对 context 所提供的超时时间进行处理，并在指定的时间进行取消行为
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
