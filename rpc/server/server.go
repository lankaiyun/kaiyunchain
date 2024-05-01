package server

import (
	"context"
	"errors"
	"github.com/cockroachdb/pebble"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/config"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/crypto/ecdsa"
	"github.com/lankaiyun/kaiyunchain/crypto/keccak256"
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/mpt"
	"github.com/lankaiyun/kaiyunchain/rpc/pb"
	"github.com/lankaiyun/kaiyunchain/wallet"
	"log"
	"math/big"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	MptDbObj      *pebble.DB
	ChainDbObj    *pebble.DB
	TxDbObj       *pebble.DB
	ContractDbObj *pebble.DB
	pb.UnimplementedRpcServer
}

func (s *Server) GetLatestBlockHeight(ctx context.Context, req *pb.GetLatestBlockHeightReq) (*pb.GetLatestBlockHeightResp, error) {
	return &pb.GetLatestBlockHeightResp{Height: core.GetLastBlock(s.ChainDbObj).Header.Height.String()}, nil
}

func (s *Server) GetAllBlock(ctx context.Context, req *pb.GetAllBlockReq) (*pb.GetAllBlockResp, error) {
	lastBlock := core.GetLastBlock(s.ChainDbObj)
	high := int(lastBlock.Header.Height.Int64())
	begin, _ := strconv.Atoi(req.GetBegin())
	end, _ := strconv.Atoi(req.GetEnd())
	var bs []*pb.GetAllBlockResp_Block
	for i := high - begin; i >= high-end; i-- {
		block := core.GetBlock(big.NewInt(int64(i)), s.ChainDbObj)
		bs = append(bs, &pb.GetAllBlockResp_Block{
			Height:   block.Header.Height.String(),
			Time:     common.TimestampToTime(block.Header.Time),
			TxNum:    strconv.Itoa(len(block.Body.Txs)),
			Reward:   block.Header.Reward.String(),
			Coinbase: block.Header.Coinbase.Hex(),
		})
	}
	return &pb.GetAllBlockResp{Blocks: bs}, nil
}

func (s *Server) GetBlock(ctx context.Context, req *pb.GetBlockReq) (*pb.GetBlockResp, error) {
	lastBlock := core.GetLastBlock(s.ChainDbObj)
	lastBlockHeight := int(lastBlock.Header.Height.Int64())
	height, err := strconv.Atoi(req.GetHeight())
	if err != nil {
		return nil, err
	}
	if height < 0 || height > lastBlockHeight {
		return nil, errors.New("out of block range")
	}
	block := core.GetBlock(big.NewInt(int64(height)), s.ChainDbObj)
	return &pb.GetBlockResp{
		Nonce:          block.Header.Nonce.String(),
		Time:           common.TimestampToTime(block.Header.Time),
		TxNum:          strconv.Itoa(len(block.Body.Txs)),
		Reward:         block.Header.Reward.String(),
		Difficulty:     block.Header.Difficulty.String(),
		Coinbase:       block.Header.Coinbase.Hex(),
		BlockHash:      block.Header.BlockHash.Hex(),
		PrevBlockHash:  block.Header.PrevBlockHash.Hex(),
		StateTreeRoot:  block.Header.StateTreeRoot.Hex(),
		MerkleTreeRoot: block.Header.MerkleTreeRoot.Hex(),
	}, nil
}

func (s *Server) GetLatestTxNum(ctx context.Context, req *pb.GetLatestTxNumReq) (*pb.GetLatestTxNumResp, error) {
	return &pb.GetLatestTxNumResp{Num: core.GetTxNum(s.TxDbObj).String()}, nil
}

func (s *Server) GetAllTx(ctx context.Context, req *pb.GetAllTxReq) (*pb.GetAllTxResp, error) {
	lastBlock := core.GetLastBlock(s.ChainDbObj)
	high := int(lastBlock.Header.Height.Int64())
	begin, _ := strconv.Atoi(req.GetBegin())
	end, _ := strconv.Atoi(req.GetEnd())
	var txs []*pb.GetAllTxResp_Tx
	var count int
	for i := high; i >= 0; i-- {
		block := core.GetBlock(big.NewInt(int64(i)), s.ChainDbObj)
		for j := len(block.Body.Txs) - 1; j >= 0; j-- {
			if count >= begin && count < end {
				txs = append(txs, &pb.GetAllTxResp_Tx{
					TxHash:      block.Body.Txs[j].TxHash.Hex(),
					From:        block.Body.Txs[j].From.Hex(),
					To:          block.Body.Txs[j].To.Hex(),
					Value:       block.Body.Txs[j].Value.String(),
					Time:        common.TimestampToTime(block.Body.Txs[j].Time),
					BelongBlock: block.Header.Height.String(),
				})
			}
			count++
			if count >= end {
				break
			}
		}
		if count >= end {
			break
		}
	}
	return &pb.GetAllTxResp{Txs: txs}, nil
}

func (s *Server) GetTx(ctx context.Context, req *pb.GetTxReq) (*pb.GetTxResp, error) {
	txHash := req.GetTxHash()
	tx := core.GetTx(txHash, s.TxDbObj)
	if tx == nil {
		return nil, errors.New("tx get failed")
	}
	return &pb.GetTxResp{
		TxHash:      tx.TxHash.Hex(),
		From:        tx.From.Hex(),
		To:          tx.To.Hex(),
		Value:       tx.Value.String(),
		Time:        common.TimestampToTime(tx.Time),
		BelongBlock: tx.BelongBlock.String(),
	}, nil
}

func (s *Server) NewAccount(ctx context.Context, req *pb.NewAccountReq) (*pb.NewAccountResp, error) {
	w := wallet.NewWallet()
	path := db.KeystoreDataPath + "/" + w.Address.Hex()
	w.StoreKey(path, req.GetPassword())

	mptBytes := db.Get(common.Latest, s.MptDbObj)
	trie := mpt.Deserialize(mptBytes)
	state := core.NewState()
	err := trie.Put(w.Address.Bytes(), core.Serialize(state))
	if err != nil {
		log.Panic("Failed to Put:", err)
	}
	db.Set(common.Latest, mpt.Serialize(trie.Root), s.MptDbObj)
	return &pb.NewAccountResp{Account: w.Address.Hex()}, nil
}

func IsAccountExist(address string) bool {
	files, _ := filepath.Glob(db.KeystoreDataPath + "/*")
	for i := 0; i < len(files); i++ {
		if strings.Compare(files[i][25:], address) == 0 {
			return true
		}
	}
	return false
}

func (s *Server) GetBalance(ctx context.Context, req *pb.GetBalanceReq) (*pb.GetBalanceResp, error) {
	addr := req.GetAddress()
	if !IsAccountExist(addr) {
		return nil, errors.New("账号不存在")
	}

	mptBytes := db.Get(common.Latest, s.MptDbObj)
	trie := mpt.Deserialize(mptBytes)
	stateBytes, _ := trie.Get(common.Hex2Bytes(addr[2:]))
	state := core.DeserializeState(stateBytes)
	return &pb.GetBalanceResp{Balance: state.Balance.String()}, nil
}

func (s *Server) NewTx(ctx context.Context, req *pb.NewTxReq) (*pb.NewTxResp, error) {
	from := req.From
	to := req.To
	amount := req.Amount
	pass := req.Pass
	if !IsAccountExist(from) {
		return &pb.NewTxResp{Result: "账号不存在！"}, nil
	}
	if !IsAccountExist(to) {
		return &pb.NewTxResp{Result: "账号不存在！"}, nil
	}
	p := db.KeystoreDataPath + "/" + from
	w := wallet.LoadWallet(p, pass, from)
	if w == nil {
		return &pb.NewTxResp{Result: "密码错误！"}, nil
	}
	ok, loc := core.TxIsFull(s.TxDbObj)
	if ok {
		return &pb.NewTxResp{Result: "交易池已满！"}, nil
	}
	mptBytes := db.Get(common.Latest, s.MptDbObj)
	trie := mpt.Deserialize(mptBytes)
	accByte := common.Hex2Bytes(from[2:])
	stateByte, _ := trie.Get(accByte)
	state := core.DeserializeState(stateByte)
	balance := state.Balance
	val, _ := new(big.Int).SetString(amount, 10)
	if balance.Cmp(val) == -1 {
		return &pb.NewTxResp{Result: "余额不足！"}, nil
	}
	state.Balance = balance.Sub(balance, val)
	state.Nonce += 1
	trie.Update(accByte, core.Serialize(state))
	toByte := common.Hex2Bytes(from[2:])
	state2Bytes, _ := trie.Get(toByte)
	state2 := core.DeserializeState(state2Bytes)
	state2.Balance = state2.Balance.Add(state2.Balance, val)
	trie.Update(toByte, core.Serialize(state2))
	db.Set(common.Latest, mpt.Serialize(trie.Root), s.MptDbObj)
	// Get belong block
	lastBlock := core.GetLastBlock(s.ChainDbObj)
	height := new(big.Int).Add(lastBlock.Header.Height, common.Big1)
	// Build tx
	core.NewTx(common.BytesToAddress(accByte), common.BytesToAddress(toByte), val, height, time.Now().Unix(), ecdsa.EncodePubKey(w.PubKey), loc, w, s.TxDbObj)
	return &pb.NewTxResp{Result: "交易成功！"}, nil
}

func (s *Server) TxPool(ctx context.Context, req *pb.TxPoolReq) (*pb.TxPoolResp, error) {
	_, loc := core.TxIsFull(s.TxDbObj)
	if loc[0] == 0 {
		return nil, nil
	}
	var txs []*pb.TxPoolResp_Tx
	for i := 0; i < int(loc[0]); i++ {
		txBytes := db.Get([]byte{byte(i)}, s.TxDbObj)
		tx := core.DeserializeTx(txBytes)
		txs = append(txs, &pb.TxPoolResp_Tx{
			TxHash: tx.TxHash.Hex(),
			From:   tx.From.Hex(),
			To:     tx.To.Hex(),
			Value:  tx.Value.String(),
			Time:   common.TimestampToTime(tx.Time),
		})
	}
	return &pb.TxPoolResp{Txs: txs}, nil
}

func (s *Server) Deploy(ctx context.Context, req *pb.DeployReq) (*pb.DeployResp, error) {
	account := req.Account
	password := req.Password
	contractBs := req.Contract

	if !IsAccountExist(account) {
		return &pb.DeployResp{Result: "账号不存在！"}, nil
	}

	p := db.KeystoreDataPath + "/" + account
	w := wallet.LoadWallet(p, password, account)
	if w == nil {
		return &pb.DeployResp{Result: "密码错误！"}, nil
	}

	ok, loc := core.TxIsFull(s.TxDbObj)
	if ok {
		return &pb.DeployResp{Result: "The current txpool is full!"}, nil
	}

	mptBytes := db.Get(common.Latest, s.MptDbObj)
	trie := mpt.Deserialize(mptBytes)

	accBs := common.Hex2Bytes(account[2:])
	accStateBs, _ := trie.Get(accBs)
	accState := core.DeserializeState(accStateBs)

	balance := accState.Balance
	valueBig := big.NewInt(int64(config.DeployCostInt))
	if balance.Cmp(valueBig) == -1 {
		return &pb.DeployResp{Result: "Your balance is insufficient!"}, nil
	}

	nounceStr := strconv.Itoa(int(accState.Nonce))
	state := core.NewState()
	state.ContractCode = contractBs
	contractAddress := keccak256.Keccak256(contractBs, []byte(account), []byte(nounceStr))[:20]
	t := common.Address{}
	t.SetBytes(contractAddress)
	accState.Balance = balance.Sub(balance, valueBig)
	accState.Nonce += 1

	lastBlock := core.GetLastBlock(s.ChainDbObj)
	height := new(big.Int).Add(lastBlock.Header.Height, common.Big1)

	core.NewNewContractTx(common.BytesToAddress(accBs), common.BytesToAddress(t.Bytes()), valueBig, height, time.Now().Unix(), ecdsa.EncodePubKey(w.PubKey), loc, w, s.TxDbObj)
	trie.Put(t.Bytes(), core.Serialize(state))
	trie.Put(accBs, core.Serialize(accState))
	db.Set(common.Latest, mpt.Serialize(trie.Root), s.MptDbObj)
	return &pb.DeployResp{Result: "1", ContractAddress: t.Hex()}, nil
}

func (s *Server) GetContract(ctx context.Context, req *pb.GetContractReq) (*pb.GetContractResp, error) {
	addr := req.ContractAddress
	addrBs := common.Hex2Bytes(addr[2:])

	mptBytes := db.Get(common.Latest, s.MptDbObj)
	trie := mpt.Deserialize(mptBytes)

	v, err := trie.Get(addrBs)
	if !err {
		return &pb.GetContractResp{Result: "合约地址不存在！"}, nil
	}
	state := core.DeserializeState(v)
	return &pb.GetContractResp{Result: "1", Contract: state.ContractCode}, nil
}

func (s *Server) Call(ctx context.Context, req *pb.CallReq) (*pb.CallResp, error) {
	account := req.Account
	password := req.Password
	contractBs := req.Contract
	contractAddr := req.ContractAddress

	if !IsAccountExist(account) {
		return &pb.CallResp{Result: "账号不存在！"}, nil
	}

	p := db.KeystoreDataPath + "/" + account
	w := wallet.LoadWallet(p, password, account)
	if w == nil {
		return &pb.CallResp{Result: "密码错误！"}, nil
	}

	mptBytes := db.Get(common.Latest, s.MptDbObj)
	trie := mpt.Deserialize(mptBytes)

	accBs := common.Hex2Bytes(account[2:])
	accStateBs, _ := trie.Get(accBs)
	accState := core.DeserializeState(accStateBs)

	contrAddrBs := common.Hex2Bytes(contractAddr[2:])
	contrAddrStateBs, _ := trie.Get(contrAddrBs)
	contrAddrState := core.DeserializeState(contrAddrStateBs)

	contrAddrState.ContractCode = contractBs

	valueBig := big.NewInt(int64(config.CallCostInt))
	balance := accState.Balance
	if balance.Cmp(valueBig) == -1 {
		return &pb.CallResp{Result: "Your balance is insufficient!"}, nil
	}

	accState.Balance = balance.Sub(balance, valueBig)

	trie.Put(contractBs, core.Serialize(contrAddrState))
	trie.Put(accBs, core.Serialize(accState))
	db.Set(common.Latest, mpt.Serialize(trie.Root), s.MptDbObj)
	return &pb.CallResp{Result: "1"}, nil
}

func (s *Server) Set(ctx context.Context, req *pb.SetReq) (*pb.SetResp, error) {
	k := req.Key
	v := req.Value
	account := req.Account
	password := req.Password
	contractAddr := req.ContractAddress
	kk := append(k, []byte(contractAddr)...)
	if !IsAccountExist(account) {
		return &pb.SetResp{Result: "账号不存在！"}, nil
	}

	p := db.KeystoreDataPath + "/" + account
	w := wallet.LoadWallet(p, password, account)
	if w == nil {
		return &pb.SetResp{Result: "密码错误！"}, nil
	}
	db.Set(kk, v, s.ContractDbObj)

	mptBytes := db.Get(common.Latest, s.MptDbObj)
	trie := mpt.Deserialize(mptBytes)

	accBs := common.Hex2Bytes(account[2:])
	accStateBs, _ := trie.Get(accBs)
	accState := core.DeserializeState(accStateBs)

	valueBig := big.NewInt(int64(config.CallCostInt))
	balance := accState.Balance
	if balance.Cmp(valueBig) == -1 {
		return &pb.SetResp{Result: "Your balance is insufficient!"}, nil
	}

	accState.Balance = balance.Sub(balance, valueBig)
	trie.Put(accBs, core.Serialize(accState))
	db.Set(common.Latest, mpt.Serialize(trie.Root), s.MptDbObj)

	return &pb.SetResp{Result: "1"}, nil
}

func (s *Server) Get(ctx context.Context, req *pb.GetReq) (*pb.GetResp, error) {
	k := req.Key
	contractAddr := req.ContractAddress
	kk := append(k, []byte(contractAddr)...)
	return &pb.GetResp{Result: "1", Value: db.Get(kk, s.ContractDbObj)}, nil
}
