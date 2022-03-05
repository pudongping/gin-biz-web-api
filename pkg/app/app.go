package app

import (
	"strings"
	"time"

	"gin-biz-web-api/pkg/config"
)

// IsLocal 是否为本地环境
func IsLocal() bool {
	return strings.ToLower(config.GetString("app.env")) == "local"
}

// IsDev 是否为开发环境
func IsDev() bool {
	return strings.ToLower(config.GetString("app.env")) == "dev"
}

// IsTest 是否为测试环境
func IsTest() bool {
	return strings.ToLower(config.GetString("app.env")) == "test"
}

// IsProd 是否为生产环境
func IsProd() bool {
	return strings.ToLower(config.GetString("app.env")) == "prod"
}

// TimeNowInTimezone 获取当前时区的时间
func TimeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	return time.Now().In(chinaTimezone)
}

// URL 拼接站点的 url
func URL(path string) string {
	return config.GetString("app.url") + path
}
