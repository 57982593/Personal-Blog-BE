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
func main() {
	nums := make([]int, 0, 8)
	nums = append(nums, 1, 2, 3, 4, 5)
	nums2 := nums[2:4]
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 5]
	printLenCap(nums2) // len: 2, cap: 6 [3 4]

	nums2 = append(nums2, 50, 60)
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 50]
	printLenCap(nums2) // len: 4, cap: 6 [3 4 50 60]
}

func printLenCap(nums []int) {
	fmt.Print("len: %v, cap: %v", len(nums), cap(nums))
}

func test(s *[]int) {
	*s = append(*s, 456)
	fmt.Println(s)
}
