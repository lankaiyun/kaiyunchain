package console

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/fatih/color"
	kayoCommon "github.com/lankaiyun/kaiyunchain/app/kayo/internal/common"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/crypto/ecdsa"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/mpt"
	"github.com/lankaiyun/kaiyunchain/rlp"
	"github.com/lankaiyun/kaiyunchain/wallet"
	"github.com/peterh/liner"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func Transaction(acc string, txDB, mptDB *pebble.DB, w *wallet.Wallet, line *liner.State) {
	ok, loc := core.TxIsFull(txDB)
	if ok {
		color.Yellow("The current txpool is full!")
		fmt.Println()
		return
	}
	color.Blue("Welcome to Transaction Mode!")
	fmt.Println("To exit, press ctrl-d or input quit")
	mptBytes := db.Get([]byte("latest"), mptDB)
	var e []interface{}
	err := rlp.DecodeBytes(mptBytes, &e)
	if err != nil {
		log.Panic("Failed to DecodeBytes:", err)
	}
	trie := mpt.NewTrieWithDecodeData(e)
	accByte := common.Hex2Bytes(acc[2:])
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
			if strings.Compare(acc, to) == 0 {
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
					if !(kayoCommon.IsPositiveInteger(input)) {
						color.Red("The amount you entered is illegal!")
						fmt.Println()
						continue
					}
					amount, _ := strconv.ParseInt(input, 10, 64)
					bigAmount := big.NewInt(amount)
					if balance.Cmp(bigAmount) == -1 {
						color.Red("Your balance is insufficient!")
						fmt.Println()
						continue
					}
					state.Balance = balance.Sub(balance, bigAmount)
					state.Nonce += 1
					trie.Update(accByte, state.Serialize())
					state2Bytes, _ := trie.Get(toB)
					state2 := core.DeserializeState(state2Bytes)
					state2.Balance = state2.Balance.Add(state2.Balance, bigAmount)
					trie.Update(toB, state2.Serialize())
					db.Set([]byte("latest"), mpt.Serialize(trie.Root), mptDB)
					tx := core.NewTransaction(bigAmount, uint64(time.Now().Unix()),
						common.BytesToAddress(accByte), common.BytesToAddress(toB), ecdsa.EncodePubKey(w.PubKey))
					txHash := tx.Hash()
					tx.TxHash.SetBytes(txHash)
					sign := w.Sign(txHash)
					tx.Signature = sign
					core.PushTxToPool(loc, tx, txDB)
					times := strings.Split(common.GetCurrentTime(), " ")
					color.Green("INFO [%s|%s] Successful transaction!", times[0], times[1])
					fmt.Println("TransactionHash: ", common.Bytes2Hex(tx.Hash()))
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
