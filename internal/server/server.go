package server

import (
	"context"
	"log"
	"net"

	api "github.com/rlapenok/dumb_base/grpc_generate/proto"
	database "github.com/rlapenok/dumb_base/internal/data_base"
	"google.golang.org/grpc"
)

type MyServer struct {
	db database.DataBase
}

// mustEmbedUnimplementedApiServer implements api.ApiServer.
func (server *MyServer) mustEmbedUnimplementedApiServer() {
	panic("unimplemented")
}

func NewMyServer() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	grpc_server := grpc.NewServer()
	db := database.New()
	my_server := MyServer{db: db}
	api.RegisterApiServer(grpc_server, my_server)
	grpc_server.Serve(lis)
}

func (server *MyServer) GetKeys(ctx context.Context, req *api.Req) (*api.Resp, error) {

	keys := server.db.GetKeys()
	resp := api.Resp{Keys: keys}
	return &resp, nil
}
