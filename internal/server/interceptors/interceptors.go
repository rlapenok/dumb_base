package interceptors

import (
	"context"
	"strings"

	api "github.com/rlapenok/dumb_base/api/proto"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func InputMessageTypeSwithAndCheckLenKey(req interface{}, ctx context.Context) (interface{}, error) {
	switch req := req.(type) {
	case *api.Req:
		logrus.Info("Incoming  type grpcRequest:Req")
		return req, nil
	case *api.NewKey:
		logrus.Info("Incoming  type grpcRequest:NewKey")
		new_key := strings.TrimSpace(req.Key)
		if len(new_key) != 64 {
			logrus.Warn("Key length must be 64")
			_, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				logrus.Error("Cannot parse ctx")
			}
			err := status.Error(codes.InvalidArgument, "Key length must be 64")
			new_md := metadata.Pairs("error", err.Error())
			grpc.SendHeader(ctx, new_md)
			return &api.RespUpdateKey{Result: api.RespUpdateKey_ERR}, err
		}
		req.Key = new_key
		return req, nil
	default:
		logrus.Info("Incoming  type grpcRequest:Unknown")
		err := status.Error(codes.InvalidArgument, "Не поддерживается такой тип сообщений")
		new_md := metadata.Pairs("error", err.Error())
		grpc.SendHeader(ctx, new_md)
		return nil, err
	}
}

func LoggingClientIpNetworkInterceptor(ctx context.Context) {
	peer, err := peer.FromContext(ctx)
	if !err {
		logrus.Warn("couldn't parse client IP address,Network")
	}
	logrus.Info("Clinet IP: ", peer.Addr.String())
	logrus.Info("Clinet Network: ", peer.Addr.Network())
}

// Interceptor wich loging callable Api
func GetIpNetworkInterceptor(info *grpc.UnaryServerInfo) {
	logrus.Info("CallableApi: ", info.FullMethod)
}

func ComplexUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		new_req, err := InputMessageTypeSwithAndCheckLenKey(req, ctx)
		if err != nil {
			return new_req, nil
		}
		//Call inner interceptor
		GetIpNetworkInterceptor(info)
		//Get metadata from context
		LoggingClientIpNetworkInterceptor(ctx)
		return handler(ctx, new_req)

	}
}
