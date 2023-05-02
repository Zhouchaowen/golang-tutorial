package main

import (
	"fmt"
	"math/rand"
)

/*
	1.声明函数变量
	2.传递函数变量
*/

/*
	定义函数变量
	var variableName = func
*/

// compute有名函数接收2个int型参数和一个匿名函数(该匿名函数接收2个int类型参数返回一个int类型值)
func compute(x, y int, handler func(x, y int) int) int {
	x = x * 10
	y = y * 10
	result := handler(x, y) // 调用匿名函数 handler
	return result
	//return handler(x, y) //可以直接这样调用
}

func genNum() func() int {
	return func() int {
		return rand.Intn(999999)
	}
}

// 函数也可以当做类型,可以像其它值一样传递
func main() {
	var printConsole = func() {
		fmt.Println("golang tutorial")
	}
	printConsole()

	// 1 将一个匿名函数赋值给变量 add
	var add = func(x, y int) int {
		return x + y
	}
	result := add(1, 2) // 调用函数变量
	fmt.Println("add result", result)

	// 2
	Multi := func(x, y int) int {
		return x * y
	}
	result = compute(1, 2, Multi) // 调用函数compute并传递2个int参数和函数变量参数Multi
	fmt.Println("Multi result", result)

	num := genNum()
	fmt.Println("genNum result", num())
	fmt.Println("genNum result", num())
	fmt.Println("genNum result", num())
}

/*
func(x, y int) int {
		return x + y
} // 被叫做匿名函数，下一小节介绍。
*/
