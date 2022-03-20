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

// IsDebug 是否为调试模式
func IsDebug() bool {
	return config.GetBool("app.debug")
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

// RemoveQueryKey 移除 uri 中的某个参数
// eg：
// query = `aa=11&bb=22`
// keys = []string{aa}
// return = `bb=22`
func RemoveQueryKey(query string, keys []string) string {
	// 切割出所有的参数
	l := strings.Split(query, "&")
	var n []string

	for _, v := range l {
		for _, key := range keys {
			if !strings.HasPrefix(v, key) {
				n = append(n, v)
			}
		}
	}

	// 组合新参数
	var s string
	for _, v := range n {
		s = s + "&" + v
	}

	// 去除掉最前面的 `&`
	s = strings.TrimPrefix(s, "&")

	return s
}
