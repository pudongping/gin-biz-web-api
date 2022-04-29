package bootstrap

import (
	"strings"

	_ "gin-biz-web-api/config"
	"gin-biz-web-api/global"
	pkgConfig "gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
)

// setupConfig 初始化配置文件信息
func setupConfig() {

	console.Info("init config ...")

	// 通过匿名加载的方式自动加载了 config 包中所有的 init 函数

	// 加载配置文件
	pkgConfig.NewConfig(global.Env, strings.Split(global.ConfigPath, ",")...)

}
