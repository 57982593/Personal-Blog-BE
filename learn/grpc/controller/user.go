package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"goTestProject/learn/grpc/model"
	"goTestProject/learn/grpc/proto"
)

type Server struct {
	proto.UnimplementedRootServer
}

func (s *Server)GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserReply, error)  {
	userId := in.GetId()
	user := &model.User{}
	if err := user.GetUserInfo(userId); err != nil {
		return nil, fmt.Errorf("get user error")
	}
	jsonStr, err := json.Marshal(user)
	if err != nil {
		fmt.Println("转换失败", err)
	}
	str := string(jsonStr)

	return &proto.GetUserReply{User: str}, nil
}
func (s *Server) GetUserList(ctx context.Context, in *proto.GetUserListRequest)(*proto.GetUserListReply, error) {
	offset := in.GetOffset()
	limit := in.GetLimit()
	user := &model.User{}
	userList, total := user.GetUserList(offset, limit)
	jsonStr, err := json.Marshal(userList)
	if err != nil {
		fmt.Println("转换失败", err)
	}
	str := string(jsonStr)
	return &proto.GetUserListReply{UserList: str, Total: total}, nil
}

func (s *Server)DeleteUser(ctx context.Context, in *proto.DeleteUserRequest)(*proto.DeleteUserRespond, error)  {
	userId := in.GetUserId()
	err := model.DeleteUserInfo(userId)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteUserRespond{Msg: "delete success"},nil
}
func (s *Server) UploadFileStream (ctx context.Context, stream *proto.StreamRequest) error{
	fmt.Println("stream:", stream)
	return nil
}

