package bootstrap

import (
	crontabTask "gin-biz-web-api/crontab"
	"gin-biz-web-api/global"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/crontab"
	"gin-biz-web-api/pkg/logger"

	"go.uber.org/zap"
)

// setupCrontab 启动定时任务
func setupCrontab() {

	console.Info("Crontab Start ...")

	task := crontab.NewTask(config.GetString("cfg.app.timezone"))
	global.Crontab = task

	addScheduleTask()

	task.Start()
}

// addScheduleTask 添加计划任务
func addScheduleTask() {

	// @daily 或者 @midnight 每天 0 点执行清理日志
	clearLogsCrontabEntryID, err := global.Crontab.AddJob("@daily", crontabTask.ClearLogsCrontab{})
	ifError(err, int(clearLogsCrontabEntryID))

}

func ifError(err error, entryID int) {
	if err != nil {
		logger.Error(
			"加入定时任务失败：",
			zap.String("task", "ClearLogsCrontab"),
			zap.Int("entryID", entryID),
			zap.Error(err),
		)
	} else {
		logger.Info(
			"加入定时任务成功：",
			zap.String("task", "ClearLogsCrontab"),
			zap.Int("entryID", entryID),
		)
	}
}
