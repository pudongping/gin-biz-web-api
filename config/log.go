package config

import (
	"gin-biz-web-api/pkg/config"
)

func init() {
	config.Add("cfg.log", func() map[string]interface{} {
		return map[string]interface{}{

			// 日志级别，必须是以下这些选项：
			// "debug" —— 信息量大，一般调试时打开。
			// "info" —— 业务级别的运行日志。
			// "warn" —— 感兴趣、需要引起关注的信息。
			// "error" —— 记录错误信息。Panic 或者 Error。一般生产环境使用的等级。
			// 以上级别从低到高，level 值设置的级别越高，记录到日志的信息就越少
			// 开发时推荐使用 "debug" 或者 "info" ，生产环境下使用 "error"
			"level": config.Get("Log.Level", "debug"),

			// 日志记录类型，可选参数有：
			// "single" 独立的一个日志文件
			// "daily" 按照日期为单位记录日志，即每天一个日志文件
			"type": config.Get("Log.Type", "single"),

			/* ------------------ 滚动日志配置 ------------------ */
			// document link：https://github.com/natefinch/lumberjack/tree/v2.1
			// 日志文件名称
			"filename": config.Get("Log.Filename", "logs.log"),
			// 每个日志文件保存的最大尺寸 单位：M
			"max_size": config.Get("Log.MaxSize", 64),
			// 最多保存日志文件数，0 为不限，MaxAge 到了还是会删
			"max_backup": config.Get("Log.MaxBackup", 5),
			// 最多保存多少天，7 表示一周前的日志会被删除，0 表示不删
			"max_age": config.Get("Log.MaxAge", 30),
			// 是否压缩，压缩日志不方便查看，我们设置为 false（压缩可节省空间）
			"compress": config.Get("Log.Compress", false),
		}
	})
}
