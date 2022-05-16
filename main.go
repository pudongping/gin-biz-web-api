package main

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"gin-biz-web-api/bootstrap"
	"gin-biz-web-api/cmd"
	"gin-biz-web-api/global"
	"gin-biz-web-api/pkg/console"
)

var (
	buildTime    string // 二进制文件编译时间
	buildVersion string // 二进制文件编译版本
	goVersion    string // 打包二进制文件时的 go 版本信息
	gitCommitID  string // 二进制文件编译时的 git 提交版本号
)

var rootCmd = &cobra.Command{
	Short: "Business web api skeleton based on gin framework.",
	Long:  `Default will run "server" command, you can use "-h" flag to see all subcommands.`,
	// 会在 Run 之前执行，并且所有的子命令都会继承并执行
	PersistentPreRun: func(command *cobra.Command, args []string) {
		// 初始化框架
		bootstrap.Initialize()
	},
}

func main() {

	if printBuildInfo() {
		return
	}

	// 设置项目根目录
	global.RootPath = getRootPath()

	// 启动服务
	run()

}

// run 启动服务
func run() {

	// 注册全局参数
	cmd.RegisterGlobalFlags(rootCmd)

	// 注册子命令
	registerChildrenCmd()

	// 注册默认运行的命令
	cmd.RegisterDefaultCmd(rootCmd, cmd.GinServerCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit("Failed to run server with %v ====> %s", os.Args, err.Error())
	}

}

// registerChildrenCmd 注册子命令
// 每写一个子命令需要到此处进行注册
func registerChildrenCmd() {
	rootCmd.AddCommand(
		cmd.GinServerCmd,
		cmd.CacheCmd,
		cmd.GenerateCmd,
		cmd.MakeCmd,
	)
}

// printBuildInfo 打印编译信息
func printBuildInfo() bool {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		console.Info("Build Time:                %s", buildTime)
		console.Info("Build Version:             %s", buildVersion)
		console.Info("Build Go Version:          %s", goVersion)
		console.Info("Build Git Commit Hash ID:  %s", gitCommitID)
		return true
	}

	return false
}

// getRootPath 获取项目根目录
func getRootPath() string {

	// 第一种方式：获取当前执行程序所在的绝对路径
	// 这种仅在 `go build` 时，才可以获取正确的路径
	// 获取当前执行的二进制文件的全路径，包括二进制文件名
	// eg: exePath = "/var/folders/hr/2rqppbcx4kv8_3qc_ky1qcy80000gn/T/go-build265586886/b001/exe/main"
	exePath, err := os.Executable()
	console.ExitIf(err)
	// eg: rootPathByExecutable = "/private/var/folders/hr/2rqppbcx4kv8_3qc_ky1qcy80000gn/T/go-build265586886/b001/exe"
	rootPathByExecutable, _ := filepath.EvalSymlinks(filepath.Dir(exePath))

	// 第二种方式：获取当前执行文件绝对路径
	// 这种方式在 `go run` 和 `go build` 时，都可以获取到正确的路径
	// 但是交叉编译后，执行的结果是错误的结果
	var rootPathByCaller string
	// eg: filename = "/Users/pudongping/glory/codes/golang/gin-biz-web-api/main.go"
	_, filename, _, ok := runtime.Caller(0)
	// eg: rootPathByCaller = "/Users/pudongping/glory/codes/golang/gin-biz-web-api"
	if ok {
		rootPathByCaller = path.Dir(filename)
	}

	// 可以通过 `echo $TMPDIR` 查看当前系统临时目录
	// eg: tmpDir = "/private/var/folders/hr/2rqppbcx4kv8_3qc_ky1qcy80000gn/T"
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())

	// 对比通过 `os.Executable()` 获取到的路径是否与 `TMPDIR` 环境变量设置的路径相同
	// 相同，则说明是通过 `go run` 命令启动的
	// 不同，则是通过 `go build` 命令启动的
	if strings.Contains(rootPathByExecutable, tmpDir) {
		return rootPathByCaller
	}

	return rootPathByExecutable
}
