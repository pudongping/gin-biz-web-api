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
		cfgPrefix := "cfg.redis." + group + "."
		rdsConfigs[group] = &redis.RdsClientConfig{
			Addr: fmt.Sprintf(
				"%v:%v",
				config.GetString(cfgPrefix+"host"),
				config.GetString(cfgPrefix+"port")),
			Username: config.GetString(cfgPrefix + "username"),
			Password: config.GetString(cfgPrefix + "password"),
			DB:       config.GetInt(cfgPrefix + "db"),
		}
	}

	// 连接 redis
	redis.ConnectRedis(rdsConfigs)

}
