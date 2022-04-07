// 跨域中间件
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)

		// c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Max-Age", "3628800")
		c.Header("Access-Control-Allow-Headers", "DNT,Keep-Alive,User-Agent,Cache-Control,Content-Type,Authorization,Token,AccessToken,X-Token,X-User-Id")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
