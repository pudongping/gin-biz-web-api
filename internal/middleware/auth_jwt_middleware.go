// 授权中间件
package middleware

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/model/user_model"
	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/jwt"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 自动获取 token，并解析 token
		claims, err := jwt.NewJWT().ParseToken(c)

		// jwt 解析失败
		if err != nil {
			app.NewResponse(c).ToErrorResponse(errcode.BadRequest.WithDetails(err.Error()))
			return
		}

		// jwt 解析成功，设置用户信息
		user := user_model.GetOne(claims.UserID)
		if user.ID == 0 {
			app.NewResponse(c).ToErrorResponse(errcode.NotFound.Msgf("用户"))
			return
		}

		// 将用户信息存入 gin.context 上下文中，方便后续直接从上下文中拿到用户信息
		c.Set("current_user_id", user.GetStringID())
		c.Set("current_user_info", user)

		c.Next()
	}
}
