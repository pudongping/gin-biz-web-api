package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {

	config.Add("cfg.queue_job", func() map[string]interface{} {
		return map[string]interface{}{
			"redis": map[string]interface{}{
				"host":     config.Get("QueueJob.Redis.Host", "127.0.0.1"),
				"port":     config.Get("QueueJob.Redis.Port", 6379),
				"username": config.Get("QueueJob.Redis.Username", ""),
				"password": config.Get("QueueJob.Redis.Password", ""),
				"db":       config.Get("QueueJob.Redis.DB", 3),
			},
			"config_opt": map[string]interface{}{
				// 指定使用多少个并发工作进程
				"concurrency": config.Get("QueueJob.ConfigOpt.Concurrency", 10),
				// 可选地指定多个具有不同优先级的队列
				// 由于 asynq 包，不允许修改 redis 队列前缀（默认前缀是 `asynq`），因此
				// 如果是多个项目同时运行时，可以考虑以项目名称作为 key 来使指定项目只消费指定项目中的队列
				// 比如：有两个项目 project1 和 project2
				// 那么，可以设定为
				// project1 中
				// 				"queues": map[string]int{
				//					"project1": 6,
				//				},
				// project2 中
				// 				"queues": map[string]int{
				//					"project2": 6,
				//				},
				// 然后加入队列时：
				// project1 中 asynq.NewTask(taskType, taskPayload, asynq.Queue("project1"))
				// project2 中 asynq.NewTask(taskType, taskPayload, asynq.Queue("project2"))
				"queues": map[string]int{
					"critical":        6,
					"default":         3,
					"low":             1,
					"gin-biz-web-api": 5,
				},
			},
		}
	})

}
