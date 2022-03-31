package app

import (
	"strings"
	"time"

	"github.com/shirou/gopsutil/host"

	"gin-biz-web-api/pkg/config"
)

// IsLocal 是否为本地环境
func IsLocal() bool {
	return strings.ToLower(config.GetString("cfg.app.env")) == "local"
}

// IsDev 是否为开发环境
func IsDev() bool {
	return strings.ToLower(config.GetString("cfg.app.env")) == "dev"
}

// IsTest 是否为测试环境
func IsTest() bool {
	return strings.ToLower(config.GetString("cfg.app.env")) == "test"
}

// IsProd 是否为生产环境
func IsProd() bool {
	return strings.ToLower(config.GetString("cfg.app.env")) == "prod"
}

// IsDebug 是否为调试模式
func IsDebug() bool {
	return config.GetBool("cfg.app.debug")
}

// TimeNowInTimezone 获取当前时区的时间
func TimeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(config.GetString("cfg.app.timezone"))
	return time.Now().In(chinaTimezone)
}

// TimeParseInTimezone 解析格式化的字符串并返回它表示的时间值
// eg：
// layout := "2006-01-02 15:04:05"
// inputTime := "2029-09-04 12:02:33"
// output is："2029-09-04 12:02:33"
func TimeParseInTimezone(layout, inputTime string) string {
	chinaTimezone, _ := time.LoadLocation(config.GetString("cfg.app.timezone"))
	t, _ := time.ParseInLocation(layout, inputTime, chinaTimezone)
	return time.Unix(t.Unix(), 0).In(chinaTimezone).Format(layout)
}

// URL 拼接站点的 url
func URL(path string) string {
	return config.GetString("cfg.app.url") + path
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

// GetOSName 获取系统名称，eg：darwin （全小写）
func GetOSName() string {
	platform, _, _, err := host.PlatformInformation()
	if err != nil {
		return ""
	}
	return platform
}
