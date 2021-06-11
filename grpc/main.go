package main

import (
	"fmt"
	"github.com/spf13/viper"
	database "goTestProject/init"
	"goTestProject/grpc/controller"
	"goTestProject/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	lis, err := net.Listen("tcp", ":9060")
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
	proto.RegisterRootServer(s, &controller.Server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

