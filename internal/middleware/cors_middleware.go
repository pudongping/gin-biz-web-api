// 跨域中间件
package middleware

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/responses"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")
		c.Header("Access-Control-Max-Age", "3628800")
		c.Header("Access-Control-Allow-Headers", "DNT,Keep-Alive,User-Agent,Cache-Control,Content-Type,Authorization,token")

		if c.Request.Method == "OPTIONS" {
			responses.New(c).ToResponse(nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
