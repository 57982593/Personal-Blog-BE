package main

import (
	"fmt"
	"github.com/spf13/viper"
	"goTestProject/database"
	"goTestProject/pkg/user"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	config := "./config/admin-local.yaml"
	viper.SetConfigFile(config)
	content, err := ioutil.ReadFile(config)
	if err != nil {
		panic(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	// Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		panic(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	port := viper.GetInt("port")
	database.InitDatabase()
	fmt.Println(port)
	http.HandleFunc("/write", user.Write)
	err = http.ListenAndServe(":8088", nil)
	if (err) != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
