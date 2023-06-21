package main

import "fmt"

/*
	1.defer+recovery()异常处理
*/

func Steps1() {
	defer func() {
		if err := recover(); err != nil {
			// 捕捉错误 run err: runtime error: integer divide by zero
			fmt.Println("【run err】:", err)
		}
	}()

	a := 10
	b := 0
	_ = a / b // 发生 panic
	fmt.Println("return Steps1")
}

// 捕捉处理异常 recover
func main() {
	Steps1()
	fmt.Println("return main")
}
