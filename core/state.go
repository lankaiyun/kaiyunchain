package core

import (
	"bytes"
	"encoding/gob"
	"math/big"

	"github.com/lankaiyun/kaiyunchain/common"
)

type State struct {
	Nonce        uint64   // 交易次数
	Balance      *big.Int // 余额
	ContractCode []byte   // 合约结构体bytes
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
