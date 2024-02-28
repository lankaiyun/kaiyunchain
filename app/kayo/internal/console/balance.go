package console

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/mpt"
)

func ShowBalance(account string, dbObj *pebble.DB) {
	mptBytes := db.Get(common.Latest, dbObj)
	trie := mpt.Deserialize(mptBytes)
	stateBytes, _ := trie.Get(common.Hex2Bytes(account[2:]))
	state := core.DeserializeState(stateBytes)
	fmt.Println(state.Balance.String(), "kyc")
	fmt.Println()
}
