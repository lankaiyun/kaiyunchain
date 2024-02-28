package pow

import (
	"crypto/sha256"
	"github.com/lankaiyun/kaiyunchain/common"
	"math/big"
	"time"
)

func Pow(diff *big.Int, data []byte) (*big.Int, *big.Int) {
	nonce := new(big.Int)
	begin := time.Now().UnixNano()
	for !Mine(diff, nonce, data) {
		nonce.Add(nonce, common.Big1)
	}
	end := time.Now().UnixNano()
	consumedTime := (end - begin) / 1e6
	if consumedTime < 60000 {
		diff.Add(diff, big.NewInt(16384))
	} else {
		diff.Sub(diff, big.NewInt(16384))
	}
	return nonce, diff
}

func Mine(difficulty, nonce *big.Int, data []byte) bool {
	temp := append(data, nonce.Bytes()...)
	rand := common.Bytes2BigInt(sha256.Sum256(temp))
	cons := new(big.Int).Exp(common.Big2, common.Big256, nil) // 2**256
	target := new(big.Int).Div(cons, difficulty)
	return target.Cmp(rand) > 0
}
