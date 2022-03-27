// gin web server
package cmd

import (
	"github.com/spf13/cobra"

	"gin-biz-web-api/bootstrap"
)

var GinServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start http web server with gin framework",
	Args:  cobra.NoArgs, // 不允许传参
	Run:   runWebServer,
}

// runWebServer 启动 web 服务
func runWebServer(cmd *cobra.Command, args []string) {

	// 启动服务
	bootstrap.RunServer()

}
