package bootstrap

import (
	"context"
	"time"

	"gin-biz-web-api/global"
	"gin-biz-web-api/job"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	jobPkg "gin-biz-web-api/pkg/job"
	"gin-biz-web-api/pkg/logger"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

func setupQueueJob() {

	console.Info("Queue Job Start ...")

	redisHost := config.GetString("cfg.queue_job.redis.host")
	redisPort := config.GetString("cfg.queue_job.redis.port")
	redisUsername := config.GetString("cfg.queue_job.redis.username")
	redisPassword := config.GetString("cfg.queue_job.redis.password")
	redisDB := config.GetInt("cfg.queue_job.redis.db")
	redisAddr := redisHost + ":" + redisPort

	client := jobPkg.NewAsynqClient(redisAddr, redisUsername, redisPassword, redisDB)
	global.QueueJobClient = client

	server := jobPkg.NewAsynqServer(
		redisAddr,
		redisUsername,
		redisPassword,
		redisDB,
		config.GetInt("cfg.queue_job.config_opt.concurrency"),
		config.Get("cfg.queue_job.config_opt.queues").(map[string]int),
	)
	global.QueueJobServer = server

	mux := asynq.NewServeMux()
	mux.Use(jobLoggingMiddleware)

	addQueueJob(mux)

	go func(mux *asynq.ServeMux, server *asynq.Server) {
		if err := server.Run(mux); err != nil {
			logger.Error("Queue Job Server Failed", zap.Error(err))
			console.Exit("Queue Job Server Failed %v", err)
		}
	}(mux, server)

}

// addQueueJob 添加异步队列任务
func addQueueJob(mux *asynq.ServeMux) {
	mux.HandleFunc(job.TypeFoo, job.HandFooTask)
}

// jobLoggingMiddleware 异步任务执行日志中间件
func jobLoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
		start := time.Now()
		logger.Info(
			"Start processing",
			zap.String("Type", task.Type()),
			zap.ByteString("Payload", task.Payload()),
		)
		err := h.ProcessTask(ctx, task)
		if err != nil {
			return err
		}
		logger.Info(
			"Finished processing",
			zap.String("Type", task.Type()),
			zap.Duration("Elapsed Time", time.Since(start)),
		)
		return nil
	})
}
