package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// M个接收者和一个发送者 see https://gfw.go101.org/article/channel-closing.html
func main() {
	rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要
	log.SetFlags(0)

	// ...
	const Max = 1000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	// ...
	dataCh := make(chan int)

	// 发送者
	go func() {
		for {
			if value := rand.Intn(Max); value == 0 {
				// 此唯一的发送者可以安全地关闭此数据通道。
				fmt.Printf("value: %d", 100)
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	// 接收者
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			// 接收数据直到通道dataCh已关闭
			// 并且dataCh的缓冲队列已空。
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}
