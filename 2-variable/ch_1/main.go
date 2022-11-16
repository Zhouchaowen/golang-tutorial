package main

import "fmt"

/*
	通过 var 定义变量
	var variableName T
	var variableName T = Value
	var variableName = Value
	variableName := Value
*/

// var 语句可以声明全局变量
var aa int64

// 可以在 var 中定义多个全局变量
var (
	bb int8
	cc int16
	dd bool
	ee string
)

func main() {
	// var 语句用于声明一个变量列表,默认值为对应零值
	var a int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	var b float32 // float64
	var c bool
	var d string
	var e byte // 等同于 uint8
	var f rune // 等同于 int32,表示一个 Unicode 码点
	var g interface{}

	// 多变量声明
	var h, i string

	// 没有明确初始值的变量声明会被赋予它们的 零值
	// 零值是:
	//    数值类型为 0
	//    布尔类型为 false
	//    字符串为 ""（空字符串）

	// 打印对应零值
	fmt.Println("int zero value: ", a)
	fmt.Println("int64 zero value: ", aa)
	fmt.Println("float32 zero value: ", b)
	fmt.Println("bool zero value: ", c)
	fmt.Println("string zero value: ", d)
	fmt.Println("byte zero value: ", e)
	fmt.Println("rune zero value: ", f)
	fmt.Println("interface zero value: ", g)

	fmt.Println("string zero value: ", h)
	fmt.Println("string zero value: ", i)
}
