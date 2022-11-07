package main

import "fmt"

// 定义一个 print 函数, 它接受 0 个参数，返回 0 个值
func print() {
	fmt.Println("hello world")
}

// 定义一个 add 函数
// 它接受两个 int 类型的参数,返回 一个 int 类型的值 （参数的声明请看下一章）
func add(x int, y int) int {
	return x + y
}

// 定义一个 swap 函数
// 它接受两个 int 类型的参数,返回 两个 int 类型的值
func swap(x, y int) (int, int) {
	return y, x
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量
func swap2(x, y int) (a int, b int) {
	a = y
	b = x
	return
}

func main() {
	print()
	fmt.Println(add(1, 2))
	fmt.Println(swap(1, 2))
}
