package server

import (
	"context"
	"math/big"
	"strconv"
	"strings"

	"github.com/cockroachdb/pebble"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/rpc/pb"
)

type Server struct {
	MptDbObj   *pebble.DB
	ChainDbObj *pebble.DB
	TxDbObj    *pebble.DB
	pb.UnimplementedRpcServer
}

func (s *Server) GetAllBlock(ctx context.Context, req *pb.GetAllBlockReq) (*pb.GetAllBlockResp, error) {
	lastBlock := core.GetLastBlock(s.ChainDbObj)
	high := int(lastBlock.Header.Height.Int64())
	begin, _ := strconv.Atoi(req.GetBegin())
	end, _ := strconv.Atoi(req.GetEnd())
	var bs []*pb.GetAllBlockResp_Block
	if end == -1 {
		// Get all block
		for i := high; i >= 0; i-- {
			block := core.GetBlock(big.NewInt(int64(i)), s.ChainDbObj)
			bs = append(bs, &pb.GetAllBlockResp_Block{
				Height:   block.Header.Height.String(),
				Time:     common.TimestampToTime(block.Header.Time),
				TxNum:    strconv.Itoa(len(block.Body.Txs)),
				Reward:   block.Header.Reward.String(),
				Coinbase: block.Header.Coinbase.Hex(),
			})
		}
	} else {
		// Get block by begin and end
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
	}
	return &pb.GetAllBlockResp{Blocks: bs}, nil
}

func (s *Server) GetBlock(ctx context.Context, req *pb.GetBlockReq) (*pb.GetBlockResp, error) {
	height, _ := strconv.Atoi(req.GetHeight())
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

func (s *Server) GetAllTx(ctx context.Context, req *pb.GetAllTxReq) (*pb.GetAllTxResp, error) {
	lastBlock := core.GetLastBlock(s.ChainDbObj)
	high := int(lastBlock.Header.Height.Int64())
	begin, _ := strconv.Atoi(req.GetBegin())
	end, _ := strconv.Atoi(req.GetEnd())
	var txs []*pb.GetAllTxResp_Tx
	var count int
	if end == -1 {
		for i := high; i >= 0; i-- {
			block := core.GetBlock(big.NewInt(int64(i)), s.ChainDbObj)
			for j := len(block.Body.Txs) - 1; j >= 0; j-- {
				txs = append(txs, &pb.GetAllTxResp_Tx{
					TxHash:      block.Body.Txs[j].TxHash.Hex(),
					From:        block.Body.Txs[j].From.Hex(),
					To:          block.Body.Txs[j].To.Hex(),
					Value:       block.Body.Txs[j].Value.String(),
					Time:        common.TimestampToTime(block.Body.Txs[j].Time),
					BelongBlock: block.Header.Height.String(),
				})
			}
		}
	} else {
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
	}
	return &pb.GetAllTxResp{Txs: txs}, nil
}

func (s *Server) GetTx(ctx context.Context, req *pb.GetTxReq) (*pb.GetTxResp, error) {
	lastBlock := core.GetLastBlock(s.ChainDbObj)
	high := int(lastBlock.Header.Height.Int64())
	txHash := req.GetTxHash()
	for i := high; i >= 0; i-- {
		block := core.GetBlock(big.NewInt(int64(i)), s.ChainDbObj)
		for j := len(block.Body.Txs) - 1; j >= 0; j-- {
			if strings.Compare(txHash, block.Body.Txs[j].TxHash.Hex()) == 0 {
				return &pb.GetTxResp{
					TxHash:      block.Body.Txs[j].TxHash.Hex(),
					From:        block.Body.Txs[j].From.Hex(),
					To:          block.Body.Txs[j].To.Hex(),
					Value:       block.Body.Txs[j].Value.String(),
					Time:        common.TimestampToTime(block.Body.Txs[j].Time),
					BelongBlock: block.Header.Height.String(),
				}, nil
			}
		}
	}
	return nil, nil
}
