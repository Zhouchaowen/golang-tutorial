package main

import (
	"fmt"
)

/*
	1.init() 函数
*/

// init 函数会在 main 函数之前执行，而且无需调用就会执行
func init() {
	fmt.Println("Golang Tutorial")
}

func main() {
	fmt.Println("Hello World")
}
