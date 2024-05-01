package contractapi

import (
	"github.com/lankaiyun/kaiyunchain/rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetRpcClient() pb.RpcClient {
	conn, err := grpc.Dial("localhost:8545", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc.Dial err", err)
	}
	return pb.NewRpcClient(conn)
}
