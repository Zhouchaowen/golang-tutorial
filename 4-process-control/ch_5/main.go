package main

import "fmt"

/*
	1.defer+recovery()异常处理
*/

// 捕捉处理异常 recover
func main() {
	defer func() {
		if err := recover(); err != nil {
			// 捕捉错误 run err: runtime error: integer divide by zero
			fmt.Println("run err:", err)
		}
	}()

	a := 10
	b := 0
	_ = a / b
	fmt.Println("return")
}
