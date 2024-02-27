package console

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/fatih/color"
	kayoCommon "github.com/lankaiyun/kaiyunchain/app/kayo/internal/common"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/peterh/liner"
	"strconv"
	"strings"
)

func TxPool(txDB *pebble.DB, line *liner.State) {
	_, loc := core.TxIsFull(txDB)
	if loc[0] == 0 {
		color.Yellow("There is no transaction in the Txpool.")
		fmt.Println()
		return
	}
	color.Blue("Welcome to the Txpool Mode!")
	fmt.Printf("There are now %d txs in the pool.\n", loc[0])
	fmt.Printf("You can enter 0 ~ %d to see the info of tx.\n", loc[0]-1)
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
			i, _ := strconv.Atoi(input)
			if !kayoCommon.IsInteger(input) || i < 0 || i > int(loc[0])-1 {
				color.Yellow("Your input is not valid!")
				continue
			}
			txBytes := db.Get([]byte{byte(i)}, txDB)
			tx := core.DeserializeTx(txBytes)
			fmt.Printf("TxHash: %x\n", tx.Hash())
			fmt.Printf("From: %s\n", tx.From.Hex())
			fmt.Printf("To: %s\n", tx.To.Hex())
			fmt.Println("Value: ", tx.Value, "kyc")
			fmt.Println("State: pending")
			fmt.Println("Time: ", common.TimestampToTime(int64(tx.Time)))
			fmt.Println()
		}
	}
}
