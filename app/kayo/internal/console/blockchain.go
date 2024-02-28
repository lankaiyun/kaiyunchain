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
	"math/big"
	"strconv"
	"strings"
)

func Blockchain(chainDbObj *pebble.DB, line *liner.State) {
	lastBlock := core.GetLastBlock(chainDbObj)
	height := lastBlock.Header.Height
	color.Blue("Welcome to the Blockchain Mode!")
	fmt.Printf("There are now %s blocks in blockchain.\n", new(big.Int).Add(height, common.Big1).String())
	fmt.Printf("You can enter 0 ~ %s to see the info of block.\n", height.String())
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
			in, _ := strconv.Atoi(input)
			if !kayoCommon.IsInteger(input) || in < 0 || height.Cmp(big.NewInt(int64(in))) == -1 {
				color.Yellow("Your input is not valid!")
				continue
			}
			var temp *core.Block
			temp = lastBlock
			for i := 0; i < int(height.Int64())-in; i++ {
				prevBlockHash := db.Get(temp.Header.PrevBlockHash.Bytes(), chainDbObj)
				temp = core.DeserializeBlock(prevBlockHash)
			}
			color.Green("BlockHeaderInformation")
			fmt.Printf("Height: %d\n", temp.Header.Height)
			fmt.Printf("Nonce: %d\n", temp.Header.Nonce)
			fmt.Printf("Difficulty: %d\n", temp.Header.Difficulty)
			fmt.Printf("Time: %s\n", common.TimestampToTime(int64(temp.Header.Time)))
			fmt.Printf("Coinbase: %s\n", temp.Header.Coinbase.Hex())
			fmt.Printf("BlockHash: %x\n", temp.Header.BlockHash.Bytes())
			fmt.Printf("PrevBlockHash: %x\n", temp.Header.PrevBlockHash.Bytes())
			fmt.Printf("StateTreeRoot: %x\n", temp.Header.StateTreeRoot.Bytes())
			fmt.Printf("MerkleTreeRoot: %x\n", temp.Header.MerkleTreeRoot.Bytes())
			fmt.Println()
			if temp.Header.Height.String() == "0" {
				continue
			}
			if len(temp.Body.Txs) > 0 {
				color.Green("BlockBodyInformation")
				for i := 0; i < len(temp.Body.Txs); i++ {
					color.Green("Tx%d", i)
					fmt.Printf("TxHash: %x\n", temp.Body.Txs[i].TxHash)
					fmt.Printf("From: %s\n", temp.Body.Txs[i].From.Hex())
					fmt.Printf("To: %s\n", temp.Body.Txs[i].To.Hex())
					fmt.Println("Value: ", temp.Body.Txs[i].Value, "skc")
					fmt.Println("State: stored")
					fmt.Println("Time: ", common.TimestampToTime(int64(temp.Body.Txs[i].Time)))
					fmt.Println()
				}
			}
		}
	}
}
