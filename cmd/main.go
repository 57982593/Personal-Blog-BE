package main

import (
	"fmt"
	database "goTestProject/init"
	"goTestProject/pkg/test"
	"goTestProject/pkg/user"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"github.com/spf13/viper"
)

func main() {
	config := "../config/admin-local.yaml"
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
	http.HandleFunc("/menuList", user.GetMenuList)
	http.HandleFunc("/getRoute", user.GetRoute)
	http.HandleFunc("/createTable", test.CreateTable)
	http.HandleFunc("/eachInsert",test.EachInsert)
	http.HandleFunc("/selectAll", test.SelectAll)
	//并发测试 需要将 go say("并发") 放在 say("测试") 前面，否则会变成顺序执行，具体原因未知
	// go say("并发")
	// say("测试")

	//并发进程间通信 ch <- c 把c 发送至通道 
	//c := <-ch 接收通道消息，并赋值给c 语言
	s := []int{7, 2, 8, -9, 4, 0}
	//声明通道
    c := make(chan int)
    go sum(s[:len(s)/2], c)//从开始到len(s)/2处的数组
    go sum(s[len(s)/2:], c)//从len(s)/2 处到数组结尾
    x, y := <-c, <-c // 从通道 c 中接收

    fmt.Println(x, y, x+y)

	err = http.ListenAndServe(":8088", nil)
	if (err) != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
			fmt.Println("测试%d/n",v)
			sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func say(s string) {	for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
	}
}
