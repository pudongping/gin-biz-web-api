package bootstrap

import (
	"fmt"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/redis"
)

// setupRedis 初始化 redis
func setupRedis() {

	console.Info("init redis ...")

	// 初始化配置信息组
	rdsConfigs := make(redis.RdsConfigs)
	// 获取 config/redis.go 中的所有配置信息组
	configs := config.GetStringMapString("cfg.redis")

	for group := range configs {
		rdsConfigs[group] = &redis.RdsClientConfig{
			Addr: fmt.Sprintf(
				"%v:%v",
				config.GetString("cfg.redis."+group+".host"),
				config.GetString("cfg.redis."+group+".port")),
			Username: config.GetString("cfg.redis." + group + ".username"),
			Password: config.GetString("cfg.redis." + group + ".password"),
			DB:       config.GetInt("cfg.redis." + group + ".db"),
		}
	}

	// 连接 redis
	redis.ConnectRedis(rdsConfigs)

}
