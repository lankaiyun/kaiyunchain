package cmd

import (
	"github.com/lankaiyun/kaiyunchain/app/kayo/internal/account"
	"github.com/spf13/cobra"
)

func init() {
	accountCmd.AddCommand(newCmd)
	accountCmd.AddCommand(listCmd)
}

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts",
	Long: `Account command can create a new account and list all
your existing accounts.`,
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "New account",
	Long:  "New Command can create a new account for you.",
	Run: func(cmd *cobra.Command, args []string) {
		account.NewAccount()
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all accounts",
	Long:  "List command can list all your accounts.",
	Run: func(cmd *cobra.Command, args []string) {
		account.ListAccount()
	},
}
