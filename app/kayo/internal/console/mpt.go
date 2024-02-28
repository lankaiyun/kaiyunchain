package console

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/fatih/color"
	kayoCommon "github.com/lankaiyun/kaiyunchain/app/kayo/internal/common"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/mpt"
	"github.com/peterh/liner"
	"math/big"
	"path/filepath"
	"strconv"
	"strings"
)

func Mpt(chainDbObj, mptDB *pebble.DB, line *liner.State) {
	lastBlock := core.GetLastBlock(chainDbObj)
	height := lastBlock.Header.Height
	if height.String() == "0" {
		color.Yellow("Current only genesis block, this mode is not allowed.")
		fmt.Println()
		return
	}
	color.Blue("Welcome to the mpt Mode!")
	fmt.Printf("You can enter 1 ~ %s to see state of different blocks.\n", height.String())
	fmt.Println()
	fmt.Println("To exit, press ctrl-d or input quit")
	for {
		if input, err := line.Prompt("> "); err != nil {
			fmt.Println()
			break
		} else {
			if strings.Compare("quit", input) == 0 {
				fmt.Println()
				break
			}
			inputNum, _ := strconv.Atoi(input)
			if !kayoCommon.IsInteger(input) || inputNum < 1 || height.Cmp(big.NewInt(int64(inputNum))) == -1 {
				color.Yellow("Your input is not valid!")
				continue
			}
			mptBytes := db.Get([]byte{byte(inputNum)}, mptDB)
			fmt.Println("->", inputNum)
			trie := mpt.Deserialize(mptBytes)
			files, _ := filepath.Glob(db.KeystoreDataPath + "/*")
			for i := 0; i < len(files); i++ {
				stateBytes, _ := trie.Get(common.Hex2Bytes(files[i][27:]))
				if stateBytes == nil {
					continue
				}
				state := core.DeserializeState(stateBytes)
				fmt.Println(files[i][25:], "{nonce:", state.Nonce, "|balance:", state.Balance, "}")
			}
		}
	}
}
