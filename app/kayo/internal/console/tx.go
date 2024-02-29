package console

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/fatih/color"
	localCommon "github.com/lankaiyun/kaiyunchain/app/kayo/internal/common"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/crypto/ecdsa"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/mpt"
	"github.com/lankaiyun/kaiyunchain/wallet"
	"github.com/peterh/liner"
)

func Transaction(account string, txDbObj, mptDbObj *pebble.DB, w *wallet.Wallet, line *liner.State) {
	// Look tx number
	ok, loc := core.TxIsFull(txDbObj)
	if ok {
		color.Yellow("The current txpool is full!")
		fmt.Println()
		return
	}
	// Prompt
	color.Blue("Welcome to Tx Mode!")
	fmt.Println("To exit, press ctrl-d or input quit")
	// Get mpt
	mptBytes := db.Get(common.Latest, mptDbObj)
	trie := mpt.Deserialize(mptBytes)
	accByte := common.Hex2Bytes(account[2:])
	stateByte, _ := trie.Get(accByte)
	state := core.DeserializeState(stateByte)
	balance := state.Balance
	var finish bool
	for {
		fmt.Println("Which account do you want to transfer?")
		if to, err := line.Prompt("> "); err != nil {
			fmt.Println()
			break
		} else {
			if strings.Compare("quit", to) == 0 {
				fmt.Println()
				break
			}
			if !IsAccountExist(to) {
				color.Red("The account you entered does net exist!")
				fmt.Println()
				continue
			}
			if strings.Compare(account, to) == 0 {
				color.Red("Do not transfer to yourself!")
				fmt.Println()
				continue
			}
			toB := common.Hex2Bytes(to[2:])
			fmt.Println()
			for {
				fmt.Println("How many kyc do you want to transfer?")
				if input, err := line.Prompt("> "); err != nil {
					fmt.Println()
					finish = true
					break
				} else {
					if strings.Compare("quit", input) == 0 {
						fmt.Println()
						finish = true
						break
					}
					if !(localCommon.IsPositiveInteger(input)) {
						color.Red("illegal!")
						fmt.Println()
						continue
					}
					value, _ := strconv.Atoi(input)
					valueBig := big.NewInt(int64(value))
					if balance.Cmp(valueBig) == -1 {
						color.Red("Your balance is insufficient!")
						fmt.Println()
						continue
					}
					state.Balance = balance.Sub(balance, valueBig)
					state.Nonce += 1
					trie.Update(accByte, core.Serialize(state))
					state2Bytes, _ := trie.Get(toB)
					state2 := core.DeserializeState(state2Bytes)
					state2.Balance = state2.Balance.Add(state2.Balance, valueBig)
					trie.Update(toB, core.Serialize(state2))
					db.Set(common.Latest, mpt.Serialize(trie.Root), mptDbObj)
					// Build tx
					core.NewTx(common.BytesToAddress(accByte), common.BytesToAddress(toB), valueBig, time.Now().Unix(), ecdsa.EncodePubKey(w.PubKey), loc, w, txDbObj)
					// Prompt
					times := strings.Split(common.GetCurrentTime(), " ")
					color.Green("INFO [%s|%s] Successful transaction!", times[0], times[1])
					fmt.Println()
					finish = true
					break
				}
			}
		}
		if finish {
			break
		}
	}
}
