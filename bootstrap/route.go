package bootstrap

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/internal/middleware"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/routers"
)

// setupRoute 路由初始化
func setupRoute(router *gin.Engine) {

	console.Info("init route ...")

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册 api 路由
	routers.RegisterAPIRoutes(router)

	// 配置 404 路由
	setup404Handler(router)

}

// registerGlobalMiddleWare 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),                                                                                // gin 框架自身的请求日志中间件（会在控制台中打印出路由请求及请求耗时）
		middleware.AccessLog(),                                                                      // 请求日志中间件
		middleware.Recovery(),                                                                       // 记录 Panic 和 call stack
		middleware.ContextTimeout(time.Duration(config.GetUint("app.context_timeout"))*time.Second), // 上下文超时时间
	)
}

// setup404Handler 配置 404 路由
func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 html 的话
			c.String(http.StatusNotFound, "页面无法找到")
		} else {
			// 默认返回 json 格式
			c.JSON(http.StatusNotFound, gin.H{
				"code": http.StatusNotFound,
				"msg":  "路由未定义",
			})
		}
	})
}
