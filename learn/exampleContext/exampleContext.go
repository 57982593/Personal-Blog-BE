package main

import (
	"context"
	"fmt"
	"time"
)
func main() {
	//ctx, clear := context.WithCancel(context.Background())//基础用法
	//ctx, clear := context.WithCancel(context.WithValue(context.Background(), "test", "哈哈"))//传值用法
	ctx, clear := context.WithDeadline(context.Background(), time.Now().Add(time.Second*8))//规定时间内结束
	//flag := make(chan bool)
	defer clear()
	msg := make(chan int)
	go son(ctx, msg)
	for i := 0; i < 10; i++ {
		msg <- i
	}
	//flag <- true

	time.Sleep(time.Second)
	fmt.Println("主进程结束")
}

func son(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case <-ctx.Done():
			fmt.Printf("我结束了\n")
			return
		case m := <-msg:
			fmt.Printf("接收到值 %d\n", m)
		}
	}

}
