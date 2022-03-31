package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/internal/controller"
	"gin-biz-web-api/internal/controller/auth_ctrl"
	"gin-biz-web-api/internal/controller/example_ctrl"
	"gin-biz-web-api/internal/middleware"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/limiter"
)

// 实例化针对方法的令牌桶，并添加令牌桶规则
var methodTokenBucketLimiters = limiter.NewTokenBucketMethodLimiter().AddBuckets(
	// 这里可以理解为：
	// 当访问 `/api/test` 路由时，第一次访问，在一分钟内，最多可以访问 3 次
	// （因为在 middleware.LimitMethodTokenBucket 中间件中）每次 TakeAvailable 了 1 次
	// 当超过一分钟后，会往令牌桶中加 1 个令牌，也就是超过一分钟后，这时候只允许访问一次，访问多次时，会报错
	// 如果超过一分钟后，你不访问，当达到三分钟以上时，你又可以访问 3 次（因为每增加一分钟，
	// 会往令牌桶中增加 1 个令牌，超过 3 分钟，则增加了 3 个令牌，此时不会增加 4 个令牌，因为桶的容积为 3）
	limiter.TokenBucketLimiterRule{
		Key:          "/api/test", // 自定义键值对名称
		FillInterval: time.Minute, // 间隔多久时间放 N 个令牌
		Capacity:     3,           // 令牌桶的容量
		Quantum:      1,           // 每次到达间隔时间后所放的具体令牌数量
	},
	limiter.TokenBucketLimiterRule{
		Key:          "abc", // 默认采用的是路由地址作为 key，如果自己自定义了 key，那么则需要将自定义 key 传入中间件中
		FillInterval: time.Second * 5,
		Capacity:     5,
		Quantum:      1,
	},
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
	r.StaticFS(config.GetString("cfg.upload.static_fs_relative_path"), http.Dir(config.GetString("cfg.upload.save_path")))
}

func apiTest(api *gin.RouterGroup) {
	testGroup := api.Group("/test")

	testCtrl := new(controller.TestController)
	testGroup.GET("", middleware.LimitMethodTokenBucket(methodTokenBucketLimiters), testCtrl.Test)         // 测试
	testGroup.GET("/tt", middleware.LimitMethodTokenBucket(methodTokenBucketLimiters, "abc"), testCtrl.Tt) // 测试
	testGroup.POST("", testCtrl.Test)                                                                      // 测试

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
