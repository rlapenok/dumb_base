package server

import (
	"context"

	api "github.com/rlapenok/dumb_base/grpc_generate/proto"
	database "github.com/rlapenok/dumb_base/internal/data_base"
)

type MyServer struct {
	db database.DataBase
}

func NewMyServer() MyServer {

	return MyServer{db: database.New()}
}

func (server *MyServer) GetKeys(ctx context.Context, req *api.Req) (*api.Resp, error) {
	channel := make(chan string)
	go server.db.GetKeys(channel)
	keys := <-channel
	resp := api.Resp{Keys: keys}
	return &resp, nil
}

func (server *MyServer) UpdateKeys(ctx context.Context, key *api.NewKey) (*api.RespUpdateKey, error) {
	var err error
	var resp api.RespUpdateKey
	channel := make(chan error)
	go server.db.UpdateKeys(key.Key, channel)

	for result := range channel {
		if result != nil {
			convert_to_string := result.Error()
			api.RespUpdateKey_Result_name[1] = convert_to_string
			resp = api.RespUpdateKey{Result: api.RespUpdateKey_ERR}
			err = nil
		} else {
			resp = api.RespUpdateKey{Result: api.RespUpdateKey_OK}
			err = result
		}
	}

	return &resp, err
}
