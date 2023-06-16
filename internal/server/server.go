package server

import (
	"context"
	"sync"

	api "github.com/rlapenok/dumb_base/grpc_generate/proto"
	database "github.com/rlapenok/dumb_base/internal/data_base"
)

type MyServer struct {
	db    database.DataBase
	ctx   context.Context
	mutex sync.Mutex
}

func NewMyServer() MyServer {

	return MyServer{db: database.New(), ctx: context.Background(), mutex: sync.Mutex{}}
}

func (server *MyServer) GetKeys(ctx context.Context, req *api.Req) (*api.Resp, error) {
	server.mutex.Lock()
	keys := server.db.GetKeys(&server.mutex)
	resp := api.Resp{Keys: keys}
	return &resp, nil
}
