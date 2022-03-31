package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"gin-biz-web-api/global"
	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/file"
	"gin-biz-web-api/pkg/helper/arrayx"
)

var MakeCmd = &cobra.Command{
	Use:   "make",
	Short: "Generate file or code",
}

// eg: go run main.go make model users
var makeModelStructCmd = &cobra.Command{
	Use:   "model",
	Short: "Generate model struct. Pay attention to, this command only applies to mysql databases.",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1), // 有且仅有一个参数
}

func init() {
	MakeCmd.AddCommand(makeModelStructCmd)
}

// runMakeModel sql 表结构生成模型结构体
func runMakeModel(cmd *cobra.Command, args []string) {

	platform := app.GetOSName()
	// document link: https://github.com/pudongping/go-tour
	script := global.RootPath + "/scripts/" + platform + "/go-tour"
	if _, ok := file.IsExists(script); !ok {
		console.Exit("[%s] script file does not exist!", script)
	}

	tableName := args[0] // 表名作为唯一参数

	params := []string{
		"sql",
		"struct",
		"--username",
		config.GetString("cfg.database.mysql.username"),
		"--password",
		config.GetString("cfg.database.mysql.password"),
		"--db",
		config.GetString("cfg.database.mysql.database"),
		"--table",
		tableName,
	}

	//  /Users/pudongping/glory/codes/golang/gin-biz-web-api/scripts/darwin/go-tour sql struct --username root --password 123456 --db gin_biz_web_api --table users
	console.Warning("Now execute script: %v %v", script, arrayx.Array2Str(params, " "))

	result, err := exec.Command(script, params...).Output()
	console.ExitIf(err)
	fmt.Println(string(result))
}
