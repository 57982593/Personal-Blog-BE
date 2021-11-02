package main

import "fmt"

//go:generate go env -w GO111MODULE=on

var s1 []int
//func main() {
//	a0 := [3]int{1,2,3}
//	s0 := &a0
//	s0[0] = 0
//	s1 = append(s1, 1,2)
//	test(&s1)
//	fmt.Println("s0=",s0)
//	fmt.Println("a0=",a0)
//}
func main()  {
	slice := []int{}
	fmt.Println(slice)
}
func test(s *[]int)  {
	*s = append(*s, 456)
	fmt.Println(s)
}
