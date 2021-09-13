package main

type Person struct {
	name string
	age int
	height int
}

//func main() {
//	p := Person{age: 1}
//	p.setAge(10)
//	fmt.Println("p.age=",p.age)
//}
// * 代表传入的是指针地址 没有 * 表示传入的是值
func (p *Person)setAge(age int) {
	p.age = age
}
