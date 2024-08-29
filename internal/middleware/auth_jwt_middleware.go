// 授权中间件
package middleware

import (
	"gin-biz-web-api/constant"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/model"
	"gin-biz-web-api/pkg/database"
	"gin-biz-web-api/pkg/errcode"
	"gin-biz-web-api/pkg/jwt"
	"gin-biz-web-api/pkg/responses"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		response := responses.New(c)

		// 自动获取 token，并解析 token
		claims, err := jwt.NewJWT().ParseToken(c)

		// jwt 解析失败
		if err != nil {
			response.ToErrorResponse(errcode.BadRequest.WithDetails(err.Error()), err.Error())
			c.Abort() // 终止后续中间件和处理函数的执行
			return
		}

		// jwt 解析成功，设置用户信息
		var user model.User
		database.DB.First(&user, claims.UserID)
		if user.ID == 0 {
			response.ToErrorResponse(errcode.Unauthorized, "找不到对应用户")
			c.Abort()
			return
		}

		// 将用户信息存入 gin.context 上下文中，方便后续直接从上下文中拿到用户信息
		c.Set(constant.CurrentUserID, user.GetStringID())
		c.Set(constant.CurrentUserInfo, user)

		c.Next() // 继续执行后续中间件和处理函数
	}
}
