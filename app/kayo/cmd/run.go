package cmd

import (
	"github.com/lankaiyun/kaiyunchain/db"
	"github.com/lankaiyun/kaiyunchain/rpc/pb"
	"github.com/lankaiyun/kaiyunchain/rpc/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	runCmd.Flags().StringP("rpc.port", "r", "8545", "rpc service port")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start a kaiyunchain node",
	Long:  "Start and connect to the kaiyunchain main network.",
	Run: func(cmd *cobra.Command, args []string) {
		rpcPort, _ := cmd.Flags().GetString("rpc.port")
		listen, _ := net.Listen("tcp", ":"+rpcPort)
		grpcServer := grpc.NewServer()
		pb.RegisterRpcServer(grpcServer, &server.Server{
			MptDbObj:   db.GetDbObj(db.MptDataPath),
			ChainDbObj: db.GetDbObj(db.ChainDataPath),
			TxDbObj:    db.GetDbObj(db.TxDataPath),
		})
		err := grpcServer.Serve(listen)
		if err != nil {
			log.Panic("Failed to Server: ", err)
			return
		}
	},
}
