package main

import "fmt"

func main() {
	// make 创建管道类型；可以创建多个管道；<- 赋值给管道； var := <- channel 接收值；管道取值遵循先进先出原则；如果没有使用缓冲区则为阻塞管道；
	// go 开启线程，可以通过api 限制线程数量；
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
			fmt.Println(ok)
		} else {
			break
		}
	}
	fmt.Println("main结束")
}
