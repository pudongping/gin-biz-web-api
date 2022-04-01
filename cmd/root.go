// document link： https://cobra.dev/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"gin-biz-web-api/global"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/helper/arrayx"
)

// RegisterGlobalFlags 注册全局选项 flag
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	// 定义全局变量，使用【持久标识】以便所有的子命令都可以使用
	envUsage := `Project run environment, support "local、dev、test、prod", eg: "--env=test" will use test_config.yaml configuration file`
	portUsage := `Http server run port, eg: "--port=8081" will use 8081 port run http server`
	ginRunModeUsage := `Gin framework application startup mode, support "debug、release、test", suggestion use "release" mode`
	configPathUsage := `Specify the path to the configuration file. Separate multiple paths with english commas (,) eg: "etc/,configs/"`

	rootCmd.PersistentFlags().StringVarP(&global.Env, "env", "e", "", envUsage)
	rootCmd.PersistentFlags().StringVarP(&global.Port, "port", "p", "", portUsage)
	rootCmd.PersistentFlags().StringVar(&global.GinRunMode, "mode", "", ginRunModeUsage)
	rootCmd.PersistentFlags().StringVarP(&global.ConfigPath, "config_path", "c", "etc/", configPathUsage)
}

// RegisterDefaultCmd 注册没有命令参数时默认启动的命令
func RegisterDefaultCmd(rootCmd *cobra.Command, subCmd *cobra.Command) {
	argsWithoutFirstElement := os.Args[1:] // 移除第一个参数
	cmd, _, err := rootCmd.Find(argsWithoutFirstElement)
	// 取出第一个参数
	firstArg := arrayx.ArrayFirstElementString(argsWithoutFirstElement)

	// 只有没有参数时、或者参数不为 `-h` 和 `--help` 时才设置 root cmd 的参数
	if err == nil && cmd.Use == rootCmd.Use && firstArg != "-h" && firstArg != "--help" {
		args := append([]string{subCmd.Use}, argsWithoutFirstElement...)
		console.Warning("default run command: %v", arrayx.Array2Str(args, " "))
		// 重新设置参数，比如：`server --port=8081 --mode=test`
		rootCmd.SetArgs(args)
	}

}
