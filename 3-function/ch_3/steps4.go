package main

import (
	"fmt"
	"reflect"
	"runtime"
)

var count = 0

// 通过函数实现 调用计数修饰器
// 如果先在 logAdd 上在包裹一层统计函数调用次数的日志，代码将需要如下实现将变得比较复杂。
func countAdd(x, y int, fn func(x, y int) int, fnfn func(x, y int, fn func(x, y int) int) int) int {
	fmt.Printf("  {\n")
	fmt.Printf("\tcountAdd Calling function %v \n", runtime.FuncForPC(reflect.ValueOf(fnfn).Pointer()).Name())
	result := fnfn(x, y, fn)
	fmt.Printf("\tcountAdd Calling function count %v \n", count)
	fmt.Printf("  }\n")
	count++
	return result
}

func Steps4() {
	countAdd(1, 2, add, logAdd)
	countAdd(1, 2, add, logAdd)
	countAdd(1, 2, add, logAdd)
}

// 通过闭包实现 调用计数修饰器
func countWrapper(fn func(x, y int) int) func(x, y int) int {
	countX := 0
	return func(x, y int) int {
		fmt.Printf("  {\n")
		fmt.Printf("\tcountWrapper Calling function %v \n", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
		result := fn(x, y)
		fmt.Printf("\tcountWrapper Calling function count %v \n", countX)
		fmt.Printf("  }\n")
		countX++
		return result
	}
}

func Steps4Plus() {
	wrappedAdd := logWrapper(add)
	wrappedCount := countWrapper(wrappedAdd)
	wrappedCount(2, 3)
	wrappedCount(2, 3)
	wrappedCount(2, 3)
}
