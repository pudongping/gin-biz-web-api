package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/logger"
)

// RunServer 启动服务
func RunServer() {

	console.Info("run server ...")

	// 设置 gin 框架的运行模式
	gin.SetMode(config.GetString("app.gin_run_mode"))
	// gin 实例
	router := gin.New()
	// 初始化路由绑定
	setupRoute(router)
	// 运行服务器
	srv := initServer(router)
	console.Success("App Server is running at: http://0.0.0.0:%d", config.GetInt("app.port"))
	// 优雅的重启和停止
	gracefulShutdown(srv)
}

// gracefulShutdown 优雅的重启和停止
func gracefulShutdown(srv *http.Server) {
	// 优雅的重启和停止
	// see gin web framework document examples : https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.ErrorString("Server", "gracefulShutdown", err.Error())
			console.Exit("server.ListenAndServe err: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal)
	// 接受 syscall.SIGINT 和 syscall.SIGTERM 信号
	// kill 不加参数发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，因此不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	console.Warning("Shutting down server...")
	logger.WarnString("Server", "gracefulShutdown", "正在关闭服务器……")

	// 最大时间控制，用于通知该服务端它有 5 秒的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.FatalString("Server", "gracefulShutdown", err.Error())
	}

	console.Warning("Server exiting")
	logger.WarnString("Server", "gracefulShutdown", "服务已经退出")

}

// initServer 初始化服务器
func initServer(router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           fmt.Sprintf(":%d", config.GetInt("app.port")), // 服务启动的端口
		Handler:        router,
		ReadTimeout:    time.Second * time.Duration(config.GetInt64("app.read_timeout")),  // 允许读取的最大时间
		WriteTimeout:   time.Second * time.Duration(config.GetInt64("app.write_timeout")), // 允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                                                           // 请求头的最大字节数
	}
}
