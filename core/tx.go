package core

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/cockroachdb/pebble"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/crypto/keccak256"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/wallet"
)

type Tx struct {
	TxHash      common.Hash
	From        common.Address
	To          common.Address
	Value       *big.Int
	BelongBlock *big.Int
	Time        int64
	PubKey      []byte
	Signature   []byte
	Type        int
	// 0 普通交易
	// 1 创建合约交易
	// 2 调用合约交易
	State int
	// 0 represent not yet included in the blockchain
	// 1 represent already included in the blockchain
}

func NewTx(from, to common.Address, value, belongBlock *big.Int, time int64, pubKey, loc []byte, w *wallet.Wallet, txDbObj *pebble.DB) {
	tx := &Tx{
		From:        from,
		To:          to,
		Value:       value,
		Time:        time,
		PubKey:      pubKey,
		BelongBlock: belongBlock,
	}
	tx.TxHash.SetBytes(keccak256.Keccak256(Serialize(tx)))
	tx.Signature = w.Sign(tx.TxHash.Bytes())
	db.Set(loc, Serialize(tx), txDbObj)
	db.Set(tx.TxHash.Bytes(), Serialize(tx), txDbObj)
}

func NewNewContractTx(from, to common.Address, value, belongBlock *big.Int, time int64, pubKey, loc []byte, w *wallet.Wallet, txDbObj *pebble.DB) {
	tx := &Tx{
		From:        from,
		To:          to,
		Value:       value,
		Time:        time,
		PubKey:      pubKey,
		BelongBlock: belongBlock,
		Type:        1,
	}
	tx.TxHash.SetBytes(keccak256.Keccak256(Serialize(tx)))
	tx.Signature = w.Sign(tx.TxHash.Bytes())
	db.Set(loc, Serialize(tx), txDbObj)
	db.Set(tx.TxHash.Bytes(), Serialize(tx), txDbObj)
}

func NewCallContractTx(from, to common.Address, value, belongBlock *big.Int, time int64, pubKey, loc []byte, w *wallet.Wallet, txDbObj *pebble.DB) {
	tx := &Tx{
		From:        from,
		To:          to,
		Value:       value,
		Time:        time,
		PubKey:      pubKey,
		BelongBlock: belongBlock,
		Type:        2,
	}
	tx.TxHash.SetBytes(keccak256.Keccak256(Serialize(tx)))
	tx.Signature = w.Sign(tx.TxHash.Bytes())
	db.Set(loc, Serialize(tx), txDbObj)
	db.Set(tx.TxHash.Bytes(), Serialize(tx), txDbObj)
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

func GetTxNum(txDbObj *pebble.DB) *big.Int {
	txNumBytes := db.Get(common.TxNum, txDbObj)
	return new(big.Int).SetBytes(txNumBytes)
}

func GetTx(txHash string, txDbObj *pebble.DB) *Tx {
	txHashBytes, err := hex.DecodeString(txHash[2:])
	if err != nil {
		fmt.Println("解码失败:", err)
		return nil
	}
	txBytes := db.Get(txHashBytes, txDbObj)
	if txBytes == nil {
		return nil
	}
	return DeserializeTx(txBytes)
}
