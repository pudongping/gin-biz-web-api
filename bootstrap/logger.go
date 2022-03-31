package bootstrap

import (
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/logger"
)

// setupLogger 初始化 Logger
func setupLogger() {

	console.Info("init logger ...")

	logger.InitLogger(
		config.GetString("cfg.log.filename"),
		config.GetInt("cfg.log.max_size"),
		config.GetInt("cfg.log.max_backup"),
		config.GetInt("cfg.log.max_age"),
		config.GetBool("cfg.log.compress"),
		config.GetString("cfg.log.type"),
		config.GetString("cfg.log.level"),
	)

}
