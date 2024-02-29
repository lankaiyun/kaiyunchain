package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	localCommon "github.com/lankaiyun/kaiyunchain/app/kayo/internal/common"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/mpt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Complete the initialization work of the kaiyunchain",
	Long: `This command will generate the KaiYunChainData directory
in the current path, which stores the important data of
the blockchain, please do not modify.`,
	Run: func(cmd *cobra.Command, args []string) {
		if initDir() {
			// Get and Close DbObj
			chainDbObj := db.GetDbObj(db.ChainDataPath)
			defer db.CloseDbObj(chainDbObj)
			mptDbObj := db.GetDbObj(db.MptDataPath)
			defer db.CloseDbObj(mptDbObj)
			txDbObj := db.GetDbObj(db.TxDataPath)
			defer db.CloseDbObj(txDbObj)
			// Blockchain initialization
			core.NewGenesisBlock(chainDbObj, txDbObj)
			// Mpt initialization
			trie := mpt.NewTrie()
			_ = trie.Put([]byte("hello"), []byte("world"))
			db.Set(common.Latest, mpt.Serialize(trie.Root), mptDbObj)
			// Prompt
			time := strings.Split(common.GetCurrentTime(), " ")
			color.Green("INFO [%s|%s] Initialization is successful!", time[0], time[1])
			fmt.Println("The data directory is generated for you in the current directory.")
			fmt.Println()
		} else {
			color.Red("The initialization is done!")
			fmt.Println()
		}
	},
}

func initDir() bool {
	if localCommon.IsInitDir() {
		// 已经初始化
		return false
	} else {
		// 还未初始化
		_ = os.Mkdir("./KaiYunChainData", 0777)
		_ = os.Mkdir(db.ChainDataPath, 0777)
		_ = os.Mkdir(db.KeystoreDataPath, 0777)
		_ = os.Mkdir(db.MptDataPath, 0777)
		_ = os.Mkdir(db.LogDataPath, 0777)
		_ = os.Mkdir(db.TxDataPath, 0777)
		return true
	}
}
