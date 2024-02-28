package core

import (
	"bytes"
	"encoding/gob"
	"math/big"

	"github.com/lankaiyun/kaiyunchain/common"
)

type State struct {
	Nonce    uint64      // The number of transactions completed by the wallet
	Balance  *big.Int    // Current Account Balance
	Storage  common.Hash // Contract Storage Trie Hash
	CodeHash []byte      // Contract Code Hash
}

func NewState() *State {
	return &State{Balance: common.Big0}
}

func DeserializeState(bs []byte) *State {
	var state State
	decoder := gob.NewDecoder(bytes.NewReader(bs))
	_ = decoder.Decode(&state)
	return &state
}
