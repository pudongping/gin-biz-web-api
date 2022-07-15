package job

import (
	"context"
	"encoding/json"

	"gin-biz-web-api/pkg/logger"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const TypeFoo = "foo"

// DefaultQueueName 使用此框架启动多个项目时，防止其他项目将此项目中的队列消费，因此建议不同项目使用不同的队列名称
// 一定要注意，此名称一定要存在于 cfg.queue_job.config_opt.queues 键中
const DefaultQueueName = "gin-biz-web-api"

type FooPayload struct {
	Arg1 string
	Arg2 int
}

func NewFooTask(params FooPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeFoo, payload, asynq.Queue(DefaultQueueName)), nil
}

func HandFooTask(ctx context.Context, t *asynq.Task) error {
	var payload FooPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return errors.Errorf("json.Unmarshal failed: %v: %v", err, asynq.SkipRetry)
	}

	logger.Info(
		"开始处理 HandFooTask 任务",
		zap.String("Arg1", payload.Arg1),
		zap.Int("Arg2", payload.Arg2),
	)

	return nil
}
