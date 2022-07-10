package crontab

import (
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/logger"
)

type clearLogsCrontab struct {
}

// Run 按日期轮转日志文件
func (c clearLogsCrontab) Run() {
	if "daily" != config.GetString("cfg.log.type") {
		return
	}

	_ = logger.Rotate(
		config.GetInt64("cfg.log.max_size"),
		config.GetInt64("cfg.log.max_age"),
	)
}
