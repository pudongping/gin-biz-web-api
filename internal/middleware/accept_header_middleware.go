// 设置 http 请求头信息
package middleware

import (
	"github.com/gin-gonic/gin"
)

func AcceptHeader() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Request.Header.Set("Accept", "application/json; charset=utf-8")

		c.Next()
	}
}
