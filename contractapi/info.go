package contractapi

import (
	"github.com/lankaiyun/kaiyunchain/common"
	"math/big"
)

type ContractInfo struct {
}

type GlobalInfo struct {
	BlockHash  common.Hash
	Coinbase   common.Address
	Height     *big.Int
	Difficulty *big.Int
	Time       int64
	Sender     common.Address
	GasPrice   uint
}
