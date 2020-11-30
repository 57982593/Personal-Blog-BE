package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
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

	port := viper.GetInt("webport")
	fmt.Println(port)
	//content, err := ioutil.ReadFile(config)
	//fmt.Println(content)
	//fmt.Println(err)
	//http.HandleFunc("/write", user.Write)
	//err := http.ListenAndServe(port, nil)
	//if (err) != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}
