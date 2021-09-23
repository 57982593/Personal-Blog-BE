package main

import "fmt"

func deferTest()  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(456)
}

func main()  {
	deferTest()
}
