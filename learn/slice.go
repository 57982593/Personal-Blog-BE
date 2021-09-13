package main

import "fmt"

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
func test(s *[]int)  {
	*s = append(*s, 456)
	fmt.Println(s)
}
