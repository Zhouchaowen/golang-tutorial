package main

import (
	"fmt"
	"golang-tutorial/3-function/test" // 引用 test 包, 会先执行 test 包的init函数
)

// init 函数会在 main 函数之前执行，而且无需调用就会执行
func init() {
	fmt.Println("Golang Tutorial")
}

func main() {
	fmt.Println("Hello World")

	// 应用 test 包的 Print() 函数
	test.Print()
}
