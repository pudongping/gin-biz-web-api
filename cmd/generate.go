package cmd

import (
	"github.com/spf13/cobra"

	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/helper/strx"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate keys",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// 生成 jwt 的密钥 key
var genJwtKeyCmd = &cobra.Command{
	Use:   "jwt-key",
	Short: "generate JWT's secret key",
	Run: func(cmd *cobra.Command, args []string) {
		console.Success("JWT Key is : %v", strx.StrRandomString(32))
	},
}

func init() {
	GenerateCmd.AddCommand(genJwtKeyCmd)
}
