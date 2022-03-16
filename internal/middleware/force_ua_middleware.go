package middleware

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/responses"
)

// ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 User-Agent 标头信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			responses.New(c).ToErrorResponse(errcode.BadRequest, "请求必须附带 User-Agent 标头")
			return
		}

		c.Next()
	}
}
