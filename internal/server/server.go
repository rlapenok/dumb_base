package server

import (
	"context"

	api "github.com/rlapenok/dumb_base/api/proto"
	database "github.com/rlapenok/dumb_base/internal/data_base"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type MyServer struct {
	db database.DataBase
}

func NewMyServer() MyServer {

	return MyServer{db: database.New()}
}

func (server *MyServer) GetKeys(ctx context.Context, req *api.Req) (*api.Resp, error) {
	var resp api.Resp
	logrus.Info("In GetKeys()")
	channel := make(chan string)
	go server.db.GetKeys(channel)
	for keys := range channel {
		resp = api.Resp{Keys: keys}
	}
	return &resp, nil
}

func (server *MyServer) UpdateKeys(ctx context.Context, key *api.NewKey) (*api.RespUpdateKey, error) {
	logrus.Info("in UpdateKeys()")
	var err error
	var resp api.RespUpdateKey
	channel := make(chan error)
	go server.db.UpdateKeys(key.Key, channel)

	for result := range channel {
		if result != nil {
			my_err := status.Error(codes.Unknown, result.Error())
			new_md := metadata.Pairs("error", my_err.Error())
			grpc.SendHeader(ctx, new_md)
			resp = api.RespUpdateKey{Result: api.RespUpdateKey_ERR}
			err = my_err
		}
		resp = api.RespUpdateKey{Result: api.RespUpdateKey_OK}
		err = nil
	}
	return &resp, err
}
