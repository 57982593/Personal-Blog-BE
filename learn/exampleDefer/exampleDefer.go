package main

import "fmt"


type People struct {
	id int
	name string
}
type Dog struct {
	dogName string
}
type Animal struct {
	Dog
	People
}

func deferTest(p *People, t *[]interface{})  {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	p.id = 2
	c := append(*t, "a")
	fmt.Println(456)
	fmt.Println(*p)
	fmt.Println(c)
}

func main()  {
	a := Animal{Dog{"dog"}, People{1,"test"}}

	p := People{1, "test"}
	var t = []interface{}{123, "abc"}
	fmt.Println(t[1])
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) { println(i) }(i)
	}
	deferTest(&p, &t)
	fmt.Println(p)
	fmt.Println(&t)
	fmt.Println(a)
}
