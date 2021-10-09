package controller

import (
	"context"
	"goTestProject/learn/grpc/data"
	"goTestProject/learn/grpc/proto"
)

type Server struct {
	proto.UnimplementedRootServer
}

func (s *Server)CreateUser (ctx context.Context, in *proto.CreateUserReq) (*proto.CreateUserRes, error) {
	account := in.Account
	password := in.Password
	if err := data.InsertUser(account, password); err != nil {
		return nil, err
	}
	return &proto.CreateUserRes{Msg: "ok"}, nil
}

