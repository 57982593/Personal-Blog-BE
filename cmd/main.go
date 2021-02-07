package main

import (
	"fmt"
	"github.com/spf13/viper"
	"goTestProject/init"
	"goTestProject/pkg/user"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	config := "D:/goTestProject/config/admin-local.yaml"
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
	http.HandleFunc("/menuList",user.GetMenuList)
	http.HandleFunc("/getRoute",user.GetRoute)
	err = http.ListenAndServe(":8088", nil)
	if (err) != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
