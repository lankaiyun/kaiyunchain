package console

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/cockroachdb/pebble"
	"github.com/fatih/color"
	localCommon "github.com/lankaiyun/kaiyunchain/app/kayo/internal/common"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/peterh/liner"
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
			inputNum, _ := strconv.Atoi(input)
			inputNumBig := big.NewInt(int64(inputNum))
			if !localCommon.IsInteger(input) || inputNum < 0 || height.Cmp(big.NewInt(int64(inputNum))) == -1 {
				color.Yellow("Your input is not valid!")
				continue
			}
			block := core.GetBlock(inputNumBig, chainDbObj)
			// Prompt
			color.Green("BlockHeaderInformation")
			fmt.Printf("Height: %d\n", block.Header.Height)
			fmt.Printf("Nonce: %d\n", block.Header.Nonce)
			fmt.Printf("Difficulty: %d\n", block.Header.Difficulty)
			fmt.Printf("Time: %s\n", common.TimestampToTime(int64(block.Header.Time)))
			fmt.Printf("Coinbase: %s\n", block.Header.Coinbase.Hex())
			fmt.Printf("BlockHash: %x\n", block.Header.BlockHash.Bytes())
			fmt.Printf("PrevBlockHash: %x\n", block.Header.PrevBlockHash.Bytes())
			fmt.Printf("StateTreeRoot: %x\n", block.Header.StateTreeRoot.Bytes())
			fmt.Printf("MerkleTreeRoot: %x\n", block.Header.MerkleTreeRoot.Bytes())
			fmt.Println()
			if block.Header.Height.String() == "0" {
				continue
			}
			if len(block.Body.Txs) > 0 {
				color.Green("BlockBodyInformation")
				for i := 0; i < len(block.Body.Txs); i++ {
					color.Green("Tx%d", i)
					fmt.Printf("TxHash: %x\n", block.Body.Txs[i].TxHash)
					fmt.Printf("From: %s\n", block.Body.Txs[i].From.Hex())
					fmt.Printf("To: %s\n", block.Body.Txs[i].To.Hex())
					fmt.Println("Value: ", block.Body.Txs[i].Value, "skc")
					fmt.Println("State: stored")
					fmt.Println("Time: ", common.TimestampToTime(block.Body.Txs[i].Time))
					fmt.Println()
				}
			}
		}
	}
}
