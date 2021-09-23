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
func main() {
	go test()

	for i := 1; i < 10;i++ {
		fmt.Println("main():",i)
		time.Sleep(time.Second)
	}
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

}
