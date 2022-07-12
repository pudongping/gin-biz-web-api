// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"fmt"

	"gin-biz-web-api/global"
	"gin-biz-web-api/pkg/console"
)

func Initialize() {

	fmt.Println(console.Cyan(string(global.LogoContent)))
	console.Info("Initializing ...")

	// 初始化配置文件信息
	setupConfig()

	// 初始化日志
	setupLogger()

	// 初始化数据库
	setupDB()

	// 初始化 redis
	setupRedis()

	// 初始化缓存 cache
	setupCache()

	// 启动 crontab
	setupCrontab()
}
