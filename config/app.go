package config

import (
	"github.com/gin-gonic/gin"

	"gin-biz-web-api/global"
	"gin-biz-web-api/pkg/config"
)

func init() {
	config.Add("cfg.app", func() map[string]interface{} {

		var (
			env, port, ginRunMode string
		)

		if global.Env == "" {
			env = config.GetString("App.Env", "local")
		} else {
			env = global.Env
		}

		if global.Port == "" {
			port = config.GetString("App.Port", "8000")
		} else {
			port = global.Port
		}

		if global.GinRunMode == "" {
			ginRunMode = config.GetString("App.GinRunMode", gin.ReleaseMode)
		} else {
			ginRunMode = global.GinRunMode
		}

		return map[string]interface{}{
			// 应用名称，使用 redis 和 jwt 时会用到
			"name": config.Get("App.Name", "gin-biz-web-api"),

			// 是否进入调试模式
			"debug": config.Get("App.Debug", false),

			// 当前项目的访问地址
			"url": config.Get("App.Url", "http://localhost"),

			// 项目运行环境，支持 local、dev、test、prod
			"env": env,

			// http 服务端口
			"port": port,

			// 设置 gin 的运行模式，支持 debug, release, test
			// release 会屏蔽调试信息，官方建议生产环境中使用
			// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
			// 故此设置为 release，有特殊情况手动改为 debug 即可
			"gin_run_mode": ginRunMode,

			// 允许读取的最大持续时间（单位：s）
			"read_timeout": config.Get("App.ReadTimeout", 60),

			// 允许写入的最大持续时间（单位：s）
			"write_timeout": config.Get("App.WriteTimeout", 60),

			// 设置时区，使用 jwt 包时会用到
			"timezone": config.Get("App.Timezone", "Asia/Shanghai"),

			// 上下文超时时间（单位：s）
			"context_timeout": config.Get("App.ContextTimeout", 60),
		}
	})
}
