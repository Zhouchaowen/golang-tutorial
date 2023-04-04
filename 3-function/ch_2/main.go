package main

import (
	"fmt"
)

// compute函数(有名字的函数),接收两个整数，以及一个匿名函数作为参数。
func compute(x, y int, handler func(x, y int) int) int {
	x = x * 10
	y = y * 10
	return handler(x, y)
}

// 匿名函数
func main() {
	// 第1部分 定义匿名函数并赋值给add变量
	var add = func(x, y int) int {
		return x + y
	}
	fmt.Println("add", add(1, 2)) // 调用匿名函数

	// 第2部分 定义匿名函数并赋值给Multi变量
	Multi := func(x, y int) int {
		return x * y
	}
	fmt.Println("Multi", compute(1, 2, Multi)) // 传递匿名函数
}
