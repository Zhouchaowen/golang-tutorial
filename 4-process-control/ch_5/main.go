package main

import "fmt"

// 捕捉处理异常 recover
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("run err:", err)
		}
	}()

	a := 10
	b := 0
	i := a / b
	fmt.Println("return")
}
