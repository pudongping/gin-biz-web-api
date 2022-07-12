// Package crontab
// https://github.com/robfig/cron
// document link: https://pkg.go.dev/github.com/robfig/cron
package crontab

import (
	"time"

	"gin-biz-web-api/pkg/logger"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var Task *cron.Cron

func NewTask(timezone string) *cron.Cron {
	l := lg{}

	chinaTimezone, _ := time.LoadLocation(timezone)
	Task = cron.New(
		cron.WithLocation(chinaTimezone), // 设置时区
		cron.WithSeconds(),               // 支持秒级颗粒度
		cron.WithChain( // job 中间件
			cron.Recover(l), // 捕捉内部 job 产生的 panic
		),
		cron.WithLogger(l), // 自定义日志
	)

	return Task
}

type lg struct {
}

func (l lg) Info(msg string, keysAndValues ...interface{}) {
	logger.InfoJSON("【定时任务：】", msg, keysAndValues)
}

func (l lg) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Error("【定时任务出错：】", zap.Error(err), zap.Any(msg, keysAndValues))
}
