package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// Handler 是通过type给func(x, y int) int起的别名(简化开发)
type Handler func(x, y int) int

// add 也是一个Handler
func add(x, y int) int {
	return x + y
}

// 入参为Handler返回为也为Handler
func logWrapper(handler Handler) Handler {
	return func(x, y int) int {
		fmt.Printf("\tlogWrapper Calling function %v \n", runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())
		result := handler(x, y)
		fmt.Printf("\tlogWrapper Function returned %v \n", result)
		return result
	}
}

func countWrapper(handler Handler) Handler {
	countX := 0
	return func(x, y int) int {
		fmt.Printf("  {\n")
		fmt.Printf("\tcountWrapper Calling function %v \n", runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())
		result := handler(x, y)
		fmt.Printf("\tcountWrapper Calling function count %v \n", countX)
		fmt.Printf("  }\n")
		countX++
		return result
	}
}

func main() {
	// 修饰器函数countWrapper包裹logWrapper，logWrapper包裹add
	wrappedCount := countWrapper(logWrapper(add))
	wrappedCount(2, 3)
	wrappedCount(3, 3)
	wrappedCount(4, 3)
}
