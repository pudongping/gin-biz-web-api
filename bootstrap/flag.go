package bootstrap

import (
	"flag"

	"gin-biz-web-api/global"
)

// SetupFlag 初始化加载命令行参数
func SetupFlag() {
	flag.StringVar(&global.Env, "env", "", "项目运行环境，支持 local、dev、test、prod")
	flag.StringVar(&global.Port, "port", "", "http 服务启动端口")
	flag.StringVar(&global.GinRunMode, "mode", "", "gin 框架应用程序的启动模式")
	flag.StringVar(&global.ConfigPath, "config_path", "config/", "指定要使用的配置文件路径，多个使用英文逗号分割")
	flag.BoolVar(&global.IsVersion, "version", false, "是否查看编译后的二进制文件信息")
	flag.Parse()
}
