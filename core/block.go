package core

import (
	"bytes"
	"encoding/gob"
	"math/big"

	"github.com/cockroachdb/pebble"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/consensus/pow"
	"github.com/lankaiyun/kaiyunchain/crypto/keccak256"
	"github.com/lankaiyun/kaiyunchain/db"
)

type Block struct {
	Header *BlockHeader
	Body   *BlockBody
}

type BlockHeader struct {
	BlockHash      common.Hash
	PrevBlockHash  common.Hash
	StateTreeRoot  common.Hash
	MerkleTreeRoot common.Hash
	Coinbase       common.Address
	Height         *big.Int
	Difficulty     *big.Int
	Nonce          *big.Int
	Reward         *big.Int
	GasLimit       uint64
	GasUsed        uint64
	Time           int64
}

type BlockBody struct {
	Txs []*Tx
}

func NewGenesisBlock(chainDbObj, txDbObj *pebble.DB) {
	height := common.Big0
	block := &Block{
		&BlockHeader{
			Height:     height,
			Time:       common.GetCurrentTimestamp(),
			Difficulty: common.BigInitDifficulty,
			Nonce:      common.Big0,
			Reward:     common.Big0,
		},
		&BlockBody{},
	}
	// Set hash
	blockHash := keccak256.Keccak256(Serialize(block))
	block.Header.BlockHash.SetBytes(blockHash)
	// Store
	db.Set(height.Bytes(), Serialize(block), chainDbObj)
	db.Set(common.Latest, height.Bytes(), chainDbObj)
	db.Set(common.Difficulty, common.BigInitDifficulty.Bytes(), chainDbObj)
	db.Set(common.TxNum, common.Big0.Bytes(), txDbObj)
}

func NewBlockAllin(reward *big.Int, coinbase common.Address, txs []*Tx, chainDbObj, mptDbObj, txDbObj *pebble.DB) {
	// Get last block
	lastBlock := GetLastBlock(chainDbObj)
	// Create new Block
	height := new(big.Int).Add(lastBlock.Header.Height, common.Big1)
	block := &Block{
		&BlockHeader{
			Height:        height,
			Time:          common.GetCurrentTimestamp(),
			Coinbase:      coinbase,
			Reward:        reward,
			PrevBlockHash: lastBlock.Header.BlockHash,
		},
		&BlockBody{
			Txs: txs,
		},
	}
	// Store mpt
	mptBytes := db.Get(common.Latest, mptDbObj)
	db.Set(block.Header.Height.Bytes(), mptBytes, mptDbObj)
	block.Header.StateTreeRoot.SetBytes(keccak256.Keccak256(mptBytes))
	// Building MerkleTree
	if len(txs) != 0 {
		merkleTree := NewMerkleTree(txs)
		block.Header.MerkleTreeRoot.SetBytes(merkleTree.RootNode.Hash)
		txNumBytes := db.Get(common.TxNum, txDbObj)
		db.Set(common.TxNum, new(big.Int).Add(new(big.Int).SetBytes(txNumBytes), big.NewInt(int64(len(txs)))).Bytes(), txDbObj)
	}
	// Mine
	nonce, difficulty := pow.Pow(lastBlock.Header.Difficulty, Serialize(block))
	block.Header.Difficulty = difficulty
	block.Header.Nonce = nonce
	// Set hash
	blockHash := keccak256.Keccak256(Serialize(block))
	block.Header.BlockHash.SetBytes(blockHash)
	// Store
	db.Set(height.Bytes(), Serialize(block), chainDbObj)
	db.Set(common.Latest, height.Bytes(), chainDbObj)
	db.Set(common.Difficulty, difficulty.Bytes(), chainDbObj)
}

func GetLastBlock(chainDbObj *pebble.DB) *Block {
	lastBlockHeight := db.Get(common.Latest, chainDbObj)
	lastBlockBytes := db.Get(lastBlockHeight, chainDbObj)
	return DeserializeBlock(lastBlockBytes)
}

func GetBlock(height *big.Int, chainDbObj *pebble.DB) *Block {
	blockBytes := db.Get(height.Bytes(), chainDbObj)
	return DeserializeBlock(blockBytes)
}

func NewBlockByNoMine(reward, nonce, diff, prevNum *big.Int, prevHash []byte, coinbase common.Address, chainDbObj, mptDBObj *pebble.DB, txs []*Tx) {

}

func DeserializeBlock(bs []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(bs))
	_ = decoder.Decode(&block)
	return &block
}
