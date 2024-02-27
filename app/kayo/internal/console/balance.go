package console

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/mpt"
	"github.com/lankaiyun/kaiyunchain/rlp"
	"log"
)

func ShowBalance(acc string, dbObj *pebble.DB) {
	mptBytes := db.Get([]byte("latest"), dbObj)
	var e []interface{}
	err := rlp.DecodeBytes(mptBytes, &e)
	if err != nil {
		log.Panic("Failed to DecodeBytes:", err)
	}
	trie := mpt.NewTrieWithDecodeData(e)
	stateB, _ := trie.Get(common.Hex2Bytes(acc[2:]))
	state := core.DeserializeState(stateB)
	fmt.Println(state.Balance.String(), "kyc")
	fmt.Println()
}
