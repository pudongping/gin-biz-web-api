package bootstrap

import (
	"flag"

	"gin-biz-web-api/global"
)

func setupFlag() {
	flag.StringVar(&global.Port, "port", "", "http 服务启动端口")
	flag.StringVar(&global.RunMode, "mode", "", "应用程序的启动模式")
	flag.StringVar(&global.ConfigPath, "config_path", "config/", "指定要使用的配置文件路径，多个使用英文逗号分割")
	flag.BoolVar(&global.IsVersion, "version", false, "是否查看编译后的二进制文件信息")
	flag.Parse()
}
