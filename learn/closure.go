package main

import "fmt"

// 闭包内的数据不会被回收释放
func a() func() int  {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

//func main() {
//	c := a()
//	d := a()
//	c()
//	c()
//	c()
//	d()
//	d()
//	d()
//}
