package main

import (
	"fmt"
	"reflect"
	"runtime"
)

/*
	需求：给add()函数添加调用日志
*/

// 缺陷：
// 1.和业务代码耦合
// 2.每个类似业务都要添加额外日志打印逻辑
func addX(x, y int) int {
	fmt.Printf("\tCalling addX function")
	result := x + y
	fmt.Printf("\tCalling addX function returned %v \n", result)
	return result
}

func add(x, y int) int {
	return x + y
}

// 通过函数实现 日志修饰器
// 缺陷：
// 1.每个调用的地方都需要包裹一层
func logAdd(x, y int, fn func(x, y int) int) int {
	fmt.Printf("\tlogAdd Calling function %v \n", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
	result := fn(x, y)
	fmt.Printf("\tlogAdd Function returned %v \n", result)
	return result
}

func Steps3() {
	logAdd(1, 2, add) // 调用add时包裹一层
}

// 通过闭包实现 日志修饰器
// 缺陷：
// 1.入参格式和返回格式需要固定
func logWrapper(fn func(x, y int) int) func(x, y int) int {
	return func(x, y int) int {
		fmt.Printf("\tlogWrapper Calling function %v \n", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name())
		result := fn(x, y)
		fmt.Printf("\tlogWrapper Function returned %v \n", result)
		return result
	}
}

func Steps3Plus() {
	wrappedAdd := logWrapper(add)
	wrappedAdd(2, 3)
}
