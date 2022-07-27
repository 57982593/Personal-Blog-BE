package main

import (
	"fmt"
	"time"
)

func test()  {
	for i := 1; i < 10;i++ {
		fmt.Println("test():",i)
		time.Sleep(time.Second)
	}
}

//func main() {
//	num := runtime.NumCPU()
//	fmt.Println(num)
//	go test()
//
//	for i := 1; i < 10;i++ {
//		fmt.Println("main():",i)
//		time.Sleep(time.Second)
//	}
	//main(): 1
	//test(): 1
	//test(): 2
	//main(): 2
	//main(): 3
	//test(): 3
	//test(): 4
	//main(): 4
	//main(): 5
	//test(): 5
	//test(): 6
	//main(): 6
	//main(): 7
	//test(): 7
	//test(): 8
	//main(): 8
	//main(): 9
	//test(): 9

//}

func Hello(ch chan int)  {
	fmt.Println("hello everybody , I'm asong")
	ch <- 1
}

func main()  {
	ch := make(chan int)
	go Hello(ch)
	test := <-ch
	fmt.Println(test)
	fmt.Println("Golang梦工厂")
}
