// Package crontab
// https://github.com/robfig/cron
// document link: https://pkg.go.dev/github.com/robfig/cron
package crontab

import (
	"time"

	"gin-biz-web-api/global"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/logger"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var task *cron.Cron

func addSchedule() {

	// @daily 或者 @midnight 每天 0 点执行清理日志
	clearLogsCrontabEntryID, err := task.AddJob("@daily", clearLogsCrontab{})
	ifError(err, int(clearLogsCrontabEntryID))

}

func Run() {

	console.Info("Crontab Start ...")
	l := lg{}

	chinaTimezone, _ := time.LoadLocation(config.GetString("cfg.app.timezone"))
	task = cron.New(
		cron.WithLocation(chinaTimezone), // 设置时区
		cron.WithSeconds(),               // 支持秒级颗粒度
		cron.WithChain( // job 中间件
			cron.Recover(l), // 捕捉内部 job 产生的 panic
		),
		cron.WithLogger(l), // 自定义日志
	)
	global.Crontab = task

	addSchedule()

	task.Start()
}

type lg struct {
}

func (l lg) Info(msg string, keysAndValues ...interface{}) {
	logger.InfoJSON("【定时任务：】", msg, keysAndValues)
}

func (l lg) Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Error("【定时任务出错：】", zap.Error(err), zap.Any(msg, keysAndValues))
}

func ifError(err error, entryID int) {
	if err != nil {
		logger.Error(
			"加入定时任务失败：",
			zap.String("job", "clearLogsCrontab"),
			zap.Int("entryID", entryID),
			zap.Error(err),
		)
	} else {
		logger.Info(
			"加入定时任务成功：",
			zap.String("job", "clearLogsCrontab"),
			zap.Int("entryID", entryID),
		)
	}
}
