package server

import (
	"context"
	"errors"
	"github.com/cockroachdb/pebble"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/rpc/pb"
	"math/big"
	"strconv"
)

type Server struct {
	MptDbObj   *pebble.DB
	ChainDbObj *pebble.DB
	TxDbObj    *pebble.DB
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
