package cmd

import (
	"github.com/lankaiyun/kaiyunchain/app/kayo/internal/console"
	"github.com/spf13/cobra"
)

func init() {
	consoleCmd.Flags().StringP("account", "a", "", "an account for mining and transferring skc")
	consoleCmd.Flags().StringP("password", "p", "", "password used to unlock the account")
}

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Start an interactive environment",
	Long: `The kayo is an interactive shell for blockchain
runtime environment witch user can interactive
with kaiyunchain blockchain.`,
	Run: func(cmd *cobra.Command, args []string) {
		console.Start(cmd)
	},
}
