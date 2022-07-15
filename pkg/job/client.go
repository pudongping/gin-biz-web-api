// Package job document link: https://github.com/hibiken/asynq
package job

import (
	"github.com/hibiken/asynq"
)

var Client *asynq.Client

func NewAsynqClient(redisAddr, redisUserName, redisPassword string, redisDB int) *asynq.Client {
	Client = asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Username: redisUserName,
		Password: redisPassword,
		DB:       redisDB,
	})
	// defer func(Client *asynq.Client) {
	// 	err := Client.Close()
	// 	if err != nil {
	// 		logger.Error("关闭 asynq.Client 出错", zap.Error(err))
	// 	}
	// }(Client)

	return Client
}
