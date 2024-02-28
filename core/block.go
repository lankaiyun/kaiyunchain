package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
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

func NewGenesisBlock(dbObj *pebble.DB) {
	block := &Block{
		&BlockHeader{
			Height:     common.Big0,
			Time:       common.GetCurrentTimestamp(),
			Difficulty: common.BigInitDifficulty,
			Nonce:      common.Big0,
			Reward:     common.Big0,
		},
		&BlockBody{},
	}
	blockHash := keccak256.Keccak256(Serialize(block))
	block.Header.BlockHash.SetBytes(blockHash)
	db.Set(blockHash, Serialize(block), dbObj)
	db.Set(common.Latest, blockHash, dbObj)
	db.Set(common.Difficulty, common.InitDifficulty, dbObj)
}

func NewBlockByMine(reward *big.Int, coinbase common.Address, txs []*Tx, chainDbObj, mptDbObj *pebble.DB) {
	// Get last block
	lastBlock := GetLastBlock(chainDbObj)
	// Create new Block
	height := new(big.Int).Add(lastBlock.Header.Height, common.Big1)
	block := &Block{
		&BlockHeader{
			Height:   height,
			Time:     common.GetCurrentTimestamp(),
			Coinbase: coinbase,
			Reward:   reward,
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
	}
	block.Header.PrevBlockHash.SetBytes(lastBlock.Header.BlockHash.Bytes())
	fmt.Println("Mining is underway now, please wait patiently.")
	nonce, difficulty := pow.Pow(lastBlock.Header.Difficulty, Serialize(block))
	block.Header.Difficulty = difficulty
	block.Header.Nonce = nonce
	// End
	blockHash := keccak256.Keccak256(Serialize(block))
	block.Header.BlockHash.SetBytes(blockHash)
	db.Set(blockHash, Serialize(block), chainDbObj)
	db.Set(common.Latest, blockHash, chainDbObj)
	db.Set(common.Difficulty, difficulty.Bytes(), chainDbObj)
}

func GetLastBlock(dbObj *pebble.DB) *Block {
	lastBlockHash := db.Get(common.Latest, dbObj)
	lastBlockBytes := db.Get(lastBlockHash, dbObj)
	return DeserializeBlock(lastBlockBytes)
}

func NewBlockByNoMine(reward, nonce, diff, prevNum *big.Int, prevHash []byte, coinbase common.Address, chainDbObj, mptDBObj *pebble.DB, txs []*Tx) {
	// Create new Block
	number := new(big.Int).Add(prevNum, common.Big1)
	block := &Block{
		&BlockHeader{
			Height:   number,
			Time:     common.GetCurrentTimestamp(),
			Coinbase: coinbase,
			Reward:   reward,
		},
		&BlockBody{
			Txs: txs,
		},
	}
	// Store current mpt
	mptBytes := db.Get(common.Latest, mptDBObj)
	db.Set(number.Bytes(), mptBytes, mptDBObj)
	block.Header.StateTreeRoot.SetBytes(keccak256.Keccak256(mptBytes))
	// Building MerkleTree
	merkleTree := NewMerkleTree(txs)
	block.Header.MerkleTreeRoot.SetBytes(merkleTree.RootNode.Hash)

	block.Header.PrevBlockHash.SetBytes(prevHash)
	block.Header.Difficulty = diff
	block.Header.Nonce = nonce

	blockHash := keccak256.Keccak256(Serialize(block))
	block.Header.BlockHash.SetBytes(blockHash)
	db.Set(blockHash, Serialize(block), chainDbObj)
	db.Set(common.Latest, blockHash, chainDbObj)
}

func DeserializeBlock(bs []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(bs))
	_ = decoder.Decode(&block)
	return &block
}
