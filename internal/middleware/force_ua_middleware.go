package middleware

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/errcode"
)

// ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 User-Agent 标头信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			app.NewResponse(c).ToErrorResponse(errcode.BadRequest, "请求必须附带 User-Agent 标头")
			return
		}

		c.Next()
	}
}
