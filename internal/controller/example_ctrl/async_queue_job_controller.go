package example_ctrl

import (
	"time"

	"gin-biz-web-api/global"
	"gin-biz-web-api/job"
	"gin-biz-web-api/pkg/logger"
	"gin-biz-web-api/pkg/responses"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type AsyncQueueJobController struct {
}

func (ctrl *AsyncQueueJobController) Job(c *gin.Context) {
	response := responses.New(c)

	logger.Info("开始投递异步任务")

	task, err := job.NewFooTask(job.FooPayload{
		Arg1: "hello Alex",
		Arg2: 18,
	})
	if err != nil {
		response.ToErrorResponse(err, "异步任务投递失败")
		return
	}
	info, err := global.QueueJobClient.Enqueue(
		task,
		asynq.ProcessIn(10*time.Second), // 10 秒钟后执行
	)
	if err != nil {
		response.ToErrorResponse(err)
		return
	}

	logger.Info(
		"队列信息",
		zap.String("info id", info.ID),
		zap.String("info queue", info.Queue),
	)

	logger.Info("结束投递异步任务")

	response.ToResponse(nil)
}
