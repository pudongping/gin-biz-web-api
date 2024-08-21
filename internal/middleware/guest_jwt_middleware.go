package middleware

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/jwt"
	"gin-biz-web-api/pkg/responses"
)

// GuestJWT 强制使用游客模式访问
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		j := jwt.NewJWT()

		if token, _ := j.GetToken(c); len(token) > 0 {

			// 尝试解析 token，如果 token 解析成功，则证明已经登录成功
			if _, err := j.ParseToken(c); err == nil {
				responses.New(c).ToErrorResponse(errcode.BadRequest, "您已经登录，请使用游客身份进行访问")
				c.Abort()
				return
			}

		}

		c.Next()
	}
}
