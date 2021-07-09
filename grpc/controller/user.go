package controller

import (
	"context"
	"fmt"
	"goTestProject/grpc/service"
	"goTestProject/proto"

)

type Server struct {
	proto.UnimplementedRootServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "hello " + in.Name}, nil
}

func (s *Server)GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserReply, error)  {
	userId := in.GetId()
	user := &service.User{}
	if err := user.GetUserInfo(userId); err != nil {
		return nil, fmt.Errorf("get user error")
	}
	userInfo := proto.User{
		Id:          user.Id,
		Account:     user.Account,
		Way:         user.Way,
		Name:        user.Name,
		Email:       user.Email,
		Phone:       user.Phone,
		IsNotify:    user.IsNotify,
		CreateAt:    user.CreateAt,
		UpdateAt:    user.UpdateAt,
		LastLoginAt: user.LastLoginAt,
	}
	return &proto.GetUserReply{User: &userInfo}, nil
}


