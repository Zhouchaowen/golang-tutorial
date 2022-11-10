package main

import (
	"fmt"
)

type Handler func(x, y int) int

func compute(x, y int, handler Handler) int {
	x = x * 10
	y = y * 10
	return handler(x, y)
}

func main() {
	// 函数也可以当做类型,可以像其它值一样传递
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	Multi := func(x, y int) int {
		return x * y
	}
	fmt.Println(compute(1, 2, Multi))
}
