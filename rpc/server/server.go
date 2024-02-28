package server

import (
	"context"
	"github.com/lankaiyun/kaiyunchain/common"
	"github.com/lankaiyun/kaiyunchain/core"
	"github.com/lankaiyun/kaiyunchain/db"
	"strconv"

	"github.com/cockroachdb/pebble"
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
	var temp *core.Block
	temp = lastBlock
	var bs []*pb.GetAllBlockResp_Block
	for i := 0; i <= int(lastBlock.Header.Height.Int64()); i++ {
		bs = append(bs, &pb.GetAllBlockResp_Block{
			Height:   temp.Header.Height.String(),
			Time:     common.TimestampToTime(temp.Header.Time),
			Txs:      strconv.Itoa(len(temp.Body.Txs)),
			Reward:   temp.Header.Reward.String(),
			Coinbase: temp.Header.Coinbase.Hex(),
		})
		if i == int(lastBlock.Header.Height.Int64()) {
			break
		}
		prevBlockBytes := db.Get(temp.Header.PrevBlockHash.Bytes(), s.ChainDbObj)
		temp = core.DeserializeBlock(prevBlockBytes)
	}
	return &pb.GetAllBlockResp{Block: bs}, nil
}
