package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	database "goTestProject/init"
	"goTestProject/model"
	pb "goTestProject/proto/root"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)
const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedRootServer
}


func (s *server)GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserReply, error)  {
	user := ctx.Value("user").(*model.User)
	userInfo := pb.User{
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
	return &pb.GetUserReply{User: &userInfo}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	config := "../goTestProject/config/admin-local.yaml"
	viper.SetConfigFile(config)
	content, err := ioutil.ReadFile(config)
	if err != nil {
		panic(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		panic(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	database.InitDatabase()
	fmt.Println("Service started Successfully!")
	s := grpc.NewServer()
	pb.RegisterRootServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

