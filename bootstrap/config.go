package bootstrap

import (
	"strings"

	"gin-biz-web-api/config"
	"gin-biz-web-api/global"
	pkgConfig "gin-biz-web-api/pkg/config"
)

// setupConfig 初始化配置文件信息
func setupConfig() {

	// 触发加载 config 包的所有 init 函数
	config.Initialize()

	// 加载配置文件
	pkgConfig.NewConfig(global.Env, strings.Split(global.ConfigPath, ",")...)

}
