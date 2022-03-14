package routers

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/internal/controller"
	"gin-biz-web-api/internal/controller/auth_ctrl"
)

// RegisterAPIRoutes 注册 api 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	var api *gin.RouterGroup
	api = r.Group("/api")

	// 授权相关
	apiAuth(api)

	userGroup := api.Group("/user")

	{
		userCtrl := new(controller.UserController)
		userGroup.GET("", userCtrl.Index)
		userGroup.GET("test", userCtrl.Test)
	}

}

func apiAuth(api *gin.RouterGroup) {
	authGroup := api.Group("/auth")

	{
		rgsCtrl := new(auth_ctrl.RegisterController)
		authGroup.POST("/register/using-email", rgsCtrl.SignupUsingEmail) // 使用邮箱注册用户
	}

}
