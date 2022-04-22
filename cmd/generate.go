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
// eg：go run main.go generate jwt-key
// output："JWT Key is : FKCCWxhIKXmGKQEaQDrcqbLdkKjKvqRZ"
var genJwtKeyCmd = &cobra.Command{
	Use:     "jwt-key",
	Short:   "generate JWT's secret key",
	Example: "go run main.go generate jwt-key",
	Run: func(cmd *cobra.Command, args []string) {
		console.Success("JWT Key is : %v", strx.StrRandomOptionalString(32, strx.LowerCase+strx.UpperCase))
	},
}

func init() {
	GenerateCmd.AddCommand(genJwtKeyCmd)
}
