package console

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/fatih/color"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/crypto/ecdsa"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/mpt"
	"github.com/lankaiyun/kaiyunchain/rlp"
	"github.com/lankaiyun/kaiyunchain/wallet"
	"log"
	"math/big"
	"strings"
)

func Mine(acc string, txDbObj, chainDbObj, mptDbObj *pebble.DB) {
	var txs []*core.Transaction
	_, loc := core.TxIsFull(txDbObj)
	for i := 0; i < int(loc[0]); i++ {
		txBytes := db.Get([]byte{byte(i)}, txDbObj)
		tx := core.DeserializeTx(txBytes)
		if !wallet.Verity(tx.Hash(), tx.Signature, ecdsa.DecodePubKey(tx.PubKey)) {
			color.Red("Transaction verification failed!")
			fmt.Println()
			return
		}
		tx.State = 1
		txs = append(txs, tx)
		db.Set([]byte{byte(i)}, tx.Serialize(), txDbObj)
	}
	accBytes := common.Hex2Bytes(acc[2:])
	// Update state tree
	mptBytes := db.Get([]byte("latest"), mptDbObj)
	var e []interface{}
	err := rlp.DecodeBytes(mptBytes, &e)
	if err != nil {
		log.Panic("Failed to DecodeBytes: ", err)
	}
	trie := mpt.NewTrieWithDecodeData(e)
	stateBytes, _ := trie.Get(accBytes)
	state := core.DeserializeState(stateBytes)
	i, _ := new(big.Int).SetString("100", 10)
	state.Balance = state.Balance.Add(state.Balance, i)
	trie.Update(accBytes, state.Serialize())
	db.Set([]byte("latest"), mpt.Serialize(trie.Root), mptDbObj)
	// Create block
	core.NewBlock(i, common.BytesToAddress(accBytes), chainDbObj, mptDbObj, txs)
	// Prompt
	times := strings.Split(common.GetCurrentTime(), " ")
	fmt.Println()
	color.Green("INFO [%s|%s] A block was successfully mined!", times[0], times[1])
	fmt.Println("Account", acc, "will be awarded 100 kyc.")
	fmt.Println()
}
