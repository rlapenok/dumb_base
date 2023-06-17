package app

import (
	"net"
	"os"

	api "github.com/rlapenok/dumb_base/grpc_generate/proto"
	"github.com/rlapenok/dumb_base/internal/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func StartGrpcServer() {
	port, exist := os.LookupEnv("PORT")
	if !exist {
		logrus.Error("Not found PORT in .env=>On default PORT=8080")
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logrus.Fatal("failed to listen: ", err)

	}
	logrus.Info("Start Dumb_base grpc server on localhost:8080")
	my_server := server.NewMyServer()

	grpc := grpc.NewServer()
	api.RegisterApiServer(grpc, &my_server)
	grpc.Serve(lis)

}
