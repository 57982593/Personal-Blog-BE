package main

import (
	"context"
	"fmt"
)

type Student struct {
	age  int
	name string
}

func main() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	c := favContextKey("color")
	ctx := context.WithValue(context.Background(), k, "Go")
	ctx2 := context.WithValue(context.Background(), c, "red")

	f(ctx2, c)
	f(ctx, k)
	var a *Student
	a = new(Student)
	a.age = 10
	a.name = "test"
	a1 := new([]int)
	a2 := new(map[string]string)
	a3 := new(channel)
	b := make([]int, 5)
	// b[0] = 1
	fmt.Printf("slice = %v map = %v\n", a1, a2)
	fmt.Printf("make = %v\n", b)
}
