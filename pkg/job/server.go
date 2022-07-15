package job

import (
	"github.com/hibiken/asynq"
)

var Server *asynq.Server

func NewAsynqServer(
	redisAddr,
	redisUserName,
	redisPassword string,
	redisDB,
	configConcurrency int,
	configQueues map[string]int,
) *asynq.Server {

	Server = asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     redisAddr,
			Username: redisUserName,
			Password: redisPassword,
			DB:       redisDB,
		},
		asynq.Config{
			Concurrency: configConcurrency,
			Queues:      configQueues,
		})

	return Server
}
