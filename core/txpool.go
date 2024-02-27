package core

import (
	"github.com/cockroachdb/pebble"
	"github.com/lankaiyun/kaiyunchain/db"
)

const PoolSize = 50

func TxIsFull(txDB *pebble.DB) (bool, []byte) {
	for i := 0; i < PoolSize; i++ {
		b := db.Get([]byte{byte(i)}, txDB)
		if b == nil {
			return false, []byte{byte(i)}
		} else {
			tx := DeserializeTx(b)
			if tx.State == 1 {
				return false, []byte{byte(i)}
			}
		}
	}
	return true, []byte{byte(PoolSize)}
}

func PushTxToPool(loc []byte, tx *Transaction, txDB *pebble.DB) {
	db.Set(loc, tx.Serialize(), txDB)
}
