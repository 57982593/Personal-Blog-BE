package controller

import (
	"context"
	"goTestProject/grpc/model"
	"goTestProject/proto"
)

type Server struct {
	proto.UnimplementedRootServer
}

func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "hello " + in.Name}, nil
}

func (s *Server) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserReply, error) {
	user := model.Get(15)
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
func (s *Server) GetUserList(ctx context.Context, in *proto.GetUserListRequest) (*proto.GetUserListReply, error) {
	users := model.GetUserList()
	return &proto.GetUserListReply{Users: users}, nil
}
