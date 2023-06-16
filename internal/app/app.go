package app

import (
	"net"

	api "github.com/rlapenok/dumb_base/grpc_generate/proto"
	"github.com/rlapenok/dumb_base/internal/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func StartGrpcServer() {

	logrus.Info("Start Dumb_base grpc server on localhost:8080")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatal("failed to listen: ", err)

	}

	my_server := server.NewMyServer()

	grpc := grpc.NewServer()
	api.RegisterApiServer(grpc, &my_server)
	grpc.Serve(lis)

}
