package main

import (
	"fmt"
)

/* 函数章节讲解
1.声明函数变量
2.传递函数变量
*/

/*
	定义函数变量
	var variableName = func

*/

func compute(x, y int, handler func(x, y int) int) int {
	x = x * 10
	y = y * 10
	return handler(x, y)
}

// 函数也可以当做类型,可以像其它值一样传递
func main() {
	var add = func(x, y int) int {
		return x + y
	}
	fmt.Println("add", add(1, 2))

	Multi := func(x, y int) int {
		return x * y
	}
	fmt.Println("Multi", compute(1, 2, Multi))
}
