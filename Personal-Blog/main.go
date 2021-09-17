package main

import (
	"fmt"
	"github.com/spf13/viper"
	"goTestProject/Personal-Blog/initialization"
	"goTestProject/Personal-Blog/router"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configFilePath := "./dev.yaml"
	viper.SetConfigFile(configFilePath)
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		panic(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	port := fmt.Sprint(":",viper.Get("port"))
	initialization.Mysql()
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders: []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization"},
		ExposeHeaders: []string{"Content-Length, Access-Control-Allow-Origin, " +
			"Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Router(engine)
	_ = engine.Run(port)
}
