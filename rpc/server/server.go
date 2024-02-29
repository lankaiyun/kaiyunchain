package server

import (
	"context"
	"math/big"
	"strconv"

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
	for i := high - begin; i >= high-end; i-- {
		block := core.GetBlock(big.NewInt(int64(i)), s.ChainDbObj)
		bs = append(bs, &pb.GetAllBlockResp_Block{
			Height:   block.Header.Height.String(),
			Time:     common.TimestampToTime(block.Header.Time),
			Txs:      strconv.Itoa(len(block.Body.Txs)),
			Reward:   block.Header.Reward.String(),
			Coinbase: block.Header.Coinbase.Hex(),
		})
	}
	return &pb.GetAllBlockResp{Block: bs}, nil
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
					TxHash:      block.Body.Txs[j].TxHash.Hex()[2:],
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
