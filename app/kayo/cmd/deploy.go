package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/wallet"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	deployCmd.Flags().StringP("go", "l", "", ".go file")
	deployCmd.Flags().StringP("account", "a", "", "an account")
	deployCmd.Flags().StringP("password", "p", "", "password used to unlock the account")
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy contract",
	Long:  "Deploy smart contracts on the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		Deploy(cmd)
	},
}

func Deploy(cmd *cobra.Command) {
	file, _ := cmd.Flags().GetString("go")
	acc, _ := cmd.Flags().GetString("account")
	pass, _ := cmd.Flags().GetString("password")
	if acc == "" || pass == "" || file == "" {
		color.Red("Console command must specify account, password and contract file!")
		fmt.Println()
		return
	}
	isHas := strings.HasSuffix(file, ".go")
	if !isHas {
		color.Red("The filetype must be an .go file!")
		fmt.Println()
		return
	}
	bs, err := os.ReadFile(file)
	if err != nil {
		color.Red("File read failure: %s", err)
		fmt.Println()
		return
	}
	if !IsAccountExist(acc) {
		color.Red("The account you entered does net exist!")
		fmt.Println()
		return
	}
	p := db.KeystoreDataPath + "/" + acc
	w := wallet.LoadWallet(p, pass, acc)
	if w == nil {
		color.Red("The password you entered does not match!")
		fmt.Println()
		return
	}
	fmt.Println(string(bs))
}

func IsAccountExist(address string) bool {
	files, _ := filepath.Glob(db.KeystoreDataPath + "/*")
	for i := 0; i < len(files); i++ {
		if strings.Compare(files[i][25:], address) == 0 {
			return true
		}
	}
	return false
}
