package pow

import (
	"crypto/sha256"
	"fmt"
	"github.com/lankaiyun/kaiyunchain/common"
	"math/big"
	"time"
)

func Pow(difficulty *big.Int, data []byte) (*big.Int, *big.Int) {
	fmt.Println("Mining is underway now, please wait patiently.")
	nonce := new(big.Int)
	begin := time.Now().UnixNano()
	for !Mine(difficulty, nonce, data) {
		nonce.Add(nonce, common.Big1)
	}
	end := time.Now().UnixNano()
	consumedTime := (end - begin) / 1e6
	if consumedTime < 60000 {
		difficulty.Add(difficulty, big.NewInt(16384))
	} else {
		difficulty.Sub(difficulty, big.NewInt(16384))
	}
	return nonce, difficulty
}

func Mine(difficulty, nonce *big.Int, data []byte) bool {
	temp := append(data, nonce.Bytes()...)
	rand := common.Bytes2BigInt(sha256.Sum256(temp))
	cons := new(big.Int).Exp(common.Big2, common.Big256, nil) // 2**256
	target := new(big.Int).Div(cons, difficulty)
	return target.Cmp(rand) > 0
}
