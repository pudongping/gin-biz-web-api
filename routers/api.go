package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/internal/controller"
	"gin-biz-web-api/internal/controller/auth_ctrl"
	"gin-biz-web-api/internal/controller/example_ctrl"
	"gin-biz-web-api/internal/middleware"
	"gin-biz-web-api/pkg/config"
)

// RegisterAPIRoutes 注册 api 相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 设置静态资源访问
	setStaticURL(r)

	var api *gin.RouterGroup
	api = r.Group("/api")

	// 全局限流中间件
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）
	api.Use(middleware.LimitIP("200-H"))

	// 测试
	apiTest(api)
	// 授权相关
	apiAuth(api)
	// 示例文件
	apiExample(api)

}

// setStaticURL 设置静态资源访问
func setStaticURL(r *gin.Engine) {
	// 设置文件服务去提供静态资源的访问：https://gin-gonic.com/zh-cn/docs/examples/serving-static-files
	// eg：
	// 需要访问 `public/uploads/image/2022/03/19/c20ad4d76fe97759aa27a0c99bff6710-20220319023344.jpg` 文件时
	// 则访问地址为：http://localhost:3000/static/image/2022/03/19/c20ad4d76fe97759aa27a0c99bff6710-20220319023344.jpg
	r.StaticFS(config.GetString("upload.static_fs_relative_path"), http.Dir(config.GetString("upload.save_path")))
}

func apiTest(api *gin.RouterGroup) {
	testGroup := api.Group("/test")

	testCtrl := new(controller.TestController)
	testGroup.GET("", testCtrl.Test)  // 测试
	testGroup.POST("", testCtrl.Test) // 测试

}

func apiAuth(api *gin.RouterGroup) {
	authGroup := api.Group("/auth")
	authGroup.Use(middleware.LimitIP("60-H"))
	{
		registerCtrl := new(auth_ctrl.RegisterController)
		authGroup.POST("/register/using-email", middleware.LimitRoute("30-H"), registerCtrl.SignupUsingEmail) // 使用邮箱注册用户

		userCtrl := new(auth_ctrl.UserController)
		authGroup.GET("/user", userCtrl.Index)                       // 用户列表
		authGroup.GET("/me", middleware.AuthJWT(), userCtrl.Profile) // 用户个人信息
	}
}

func apiExample(api *gin.RouterGroup) {
	exampleGroup := api.Group("/example")
	{
		captchaCtrl := new(example_ctrl.CaptchaController)
		exampleGroup.GET("/show-captcha", captchaCtrl.ShowCaptcha)               // 显示图像验证码
		exampleGroup.POST("/verify-captcha-code", captchaCtrl.VerifyCaptchaCode) // 验证图像验证码

		emailCtrl := new(example_ctrl.EmailController)
		exampleGroup.POST("/send-email", emailCtrl.SendEmail)                       // 发送邮件
		exampleGroup.POST("/send-mailer", emailCtrl.SendMailer)                     // 使用 email 驱动发送邮件
		exampleGroup.POST("/send-email-verify-code", emailCtrl.SendEmailVerifyCode) // 发送邮件验证码

		uploadCtrl := new(example_ctrl.UploadController)
		exampleGroup.POST("/upload-file", uploadCtrl.UploadFile)     // 上传文件
		exampleGroup.POST("/upload-avatar", uploadCtrl.UploadAvatar) // 上传用户头像

		pagerCtrl := new(example_ctrl.PagerController)
		exampleGroup.GET("/pager", pagerCtrl.Pager) // 数据分页演示
	}
}
