# 异步队列任务

> 使用了依赖包 [hibiken/asynq](https://github.com/hibiken/asynq)

## 配置

配置文件位于 `config/queue_job.go` 文件，由于 [hibiken/asynq](https://github.com/hibiken/asynq) 包依赖于 `redis`
因此，必须能够正常连接 `redis`。配置均遵循 `hibiken/asynq` 包要求，更多配置，请查阅 `hibiken/asynq` 文档。

## 使用示例

- 定义处理任务逻辑方法

在 `job` 目录下编写 `foo_job.go` 文件

```go

package job

import (
	"context"
	"encoding/json"

	"gin-biz-web-api/pkg/logger"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// 定义类型名称
const TypeFoo = "foo"

// DefaultQueueName 使用此框架启动多个项目时，防止其他项目将此项目中的队列消费，因此建议不同项目使用不同的队列名称
// 一定要注意，此名称一定要存在于 cfg.queue_job.config_opt.queues 键中
const DefaultQueueName = "gin-biz-web-api"

// 定义有效载荷，执行任务时，所需要的参数，有效载荷必须是可序列化的
type FooPayload struct {
	Arg1 string
	Arg2 int
}

func NewFooTask(params FooPayload) (*asynq.Task, error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeFoo,
		payload,
		asynq.Queue(DefaultQueueName),
		asynq.MaxRetry(3), // 任务执行失败后，最多重试 3 次
		// asynq.ProcessIn(20*time.Second),  // 这里的优先级低于 global.QueueJobClient.Enqueue() 中的参数优先级
	), nil
}

// 定义具体的消费行为
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

```

- 投递任务

这里写了一个示例接口，可以参考 `/api/example/job` 接口，例如：

> 可以尝试直接请求接口 `curl --location --request GET '0.0.0.0:3000/api/example/job'`

```go

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

```

- 注册处理程序

需要在 `bootstrap/queue_job.go` 文件中的 `addQueueJob()` 方法中注册处理程序，例如：

```go

func addQueueJob(mux *asynq.ServeMux) {
	mux.HandleFunc(job.TypeFoo, job.HandFooTask)
}

```