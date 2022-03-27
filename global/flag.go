package global

var (
	Env        string // 当前项目运行环境，一般为 local dev test prod
	Port       string // http 服务启动端口，比如：8081
	GinRunMode string // gin 框架应用程序的启动模式，比如：debug、release、test
	ConfigPath string // 指定要使用的配置文件路径，多个使用英文逗号分割，比如：etc/,configs/
)
