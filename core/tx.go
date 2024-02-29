package core

import (
	"bytes"
	"encoding/gob"
	"math/big"

	"github.com/cockroachdb/pebble"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/crypto/keccak256"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/wallet"
)

type Tx struct {
	TxHash    common.Hash
	From      common.Address
	To        common.Address
	Value     *big.Int
	Time      int64
	PubKey    []byte
	Signature []byte
	State     int
	// 0 represent not yet included in the blockchain
	// 1 represent already included in the blockchain
}

func NewTx(from, to common.Address, value *big.Int, time int64, pubKey, loc []byte, w *wallet.Wallet, txDbObj *pebble.DB) {
	tx := &Tx{
		From:   from,
		To:     to,
		Value:  value,
		Time:   time,
		PubKey: pubKey,
	}
	tx.TxHash.SetBytes(keccak256.Keccak256(Serialize(tx)))
	tx.Signature = w.Sign(tx.TxHash.Bytes())
	db.Set(loc, Serialize(tx), txDbObj)
}

const PoolSize = 50

func TxIsFull(dbObj *pebble.DB) (bool, []byte) {
	for i := 0; i < PoolSize; i++ {
		txBytes := db.Get([]byte{byte(i)}, dbObj)
		if txBytes == nil {
			return false, []byte{byte(i)}
		} else {
			tx := DeserializeTx(txBytes)
			if tx.State == 1 {
				return false, []byte{byte(i)}
			}
		}
	}
	return true, []byte{byte(PoolSize)}
}

func DeserializeTx(bs []byte) *Tx {
	var tx Tx
	decoder := gob.NewDecoder(bytes.NewReader(bs))
	_ = decoder.Decode(&tx)
	return &tx
}
