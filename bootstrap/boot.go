// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"gin-biz-web-api/pkg/console"
)

func init() {

	console.Info("Initializing ...")

	// 初始化加载命令行参数
	setupFlag()

	// 初始化配置文件信息
	setupConfig()

	// 初始化日志
	setupLogger()

	// 初始化数据库
	setupDB()

	// 初始化 redis
	setupRedis()

}

func Initialize() {

}
