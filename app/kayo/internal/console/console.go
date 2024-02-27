package console

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/wallet"
	"github.com/spf13/cobra"
)

func Start(cmd *cobra.Command) {
	acc, _ := cmd.Flags().GetString("account")
	pass, _ := cmd.Flags().GetString("password")
	if acc == "" || pass == "" {
		color.Red("Console command must specify account and password!")
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
	fmt.Println()
	Interface(acc, w)
}

func Interface(acc string, w *wallet.Wallet) {
	line := GetLiner()
	txDbObj := db.GetDbObj(db.TxDataPath)
	defer db.CloseDbObj(txDbObj)
	mptDbObj := db.GetDbObj(db.MptDataPath)
	defer db.CloseDbObj(mptDbObj)
	chainDbObj := db.GetDbObj(db.ChainDataPath)
	defer db.CloseDbObj(chainDbObj)
	for {
		i := MainPrompt(acc)
		if i == 0 || i == -1 {
			color.Blue("bye")
			fmt.Println()
			break
		} else if i == 1 {
			ShowBalance(acc, mptDbObj)
		} else if i == 2 {
			Transaction(acc, txDbObj, mptDbObj, w, line)
		} else if i == 3 {
			TxPool(txDbObj, line)
		} else if i == 4 {
			Mine(acc, txDbObj, chainDbObj, mptDbObj)
		} else if i == 5 {
			Blockchain(chainDbObj, line)
		} else if i == 6 {
			Mpt(chainDbObj, mptDbObj, line)
		}
	}
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
