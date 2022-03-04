package global

var (
	Env        string // 当前项目运行环境，一般为 local dev test prod
	Port       string // http 服务启动端口
	GinRunMode string // gin 框架应用程序的启动模式
	ConfigPath string // 指定要使用的配置文件路径，多个使用英文逗号分割
	IsVersion  bool   // 是否查看编译后的二进制文件信息
)
