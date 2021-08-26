package main

import (
	"flag"
	"fmt"
	"goTestProject/grpc/controller"
	database "goTestProject/init"
	"goTestProject/proto"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := "./config/admin-local.yaml"
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
	port := viper.Get("port")
	fmt.Println("Service started Successfully!")
	s := grpc.NewServer()
	proto.RegisterRootServer(s, &controller.Server{})
	reflection.Register(s)
	wrappedGrpc := grpcweb.WrapServer(s,
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	)
	var serveAddr string
	flag.StringVar(&serveAddr,
		"address",
		fmt.Sprintf(":%d", port),
		fmt.Sprintf(":serve address - e.g. :%d", port))
	flag.Parse()

	grpcweb.WithCorsForRegisteredEndpointsOnly(true)

	srv := &http.Server{
		Addr:         serveAddr,
		WriteTimeout: time.Second * 32,
		ReadTimeout:  time.Second * 32,
		IdleTimeout:  time.Second * 60,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if wrappedGrpc.IsGrpcWebRequest(r) {
				// NOT WORKING!
				w.Header().Set("Access-Control-Allow-Origin", "*")
				wrappedGrpc.ServeHTTP(w, r)
				return
			}

			// Fall back to other servers.
			http.DefaultServeMux.ServeHTTP(w, r)
		}),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(fmt.Sprintf("srv.ListenAndServe failed to serve: %v", err))
	}
}
