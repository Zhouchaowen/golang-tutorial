package main

import "fmt"

/*
	通过 var 定义指针变量
	var variableName *T
	var variableName *T = Value
	var variableName = &Value
	variableName := &Value
*/

// Steps1 定义指针变量
func Steps1() {
	var a *int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	var b *float32 // float64
	var c *bool
	var d *string
	var e *byte // 等同于 uint8
	var f *rune // 等同于 int32,表示一个 Unicode 码点
	var g *interface{}

	fmt.Println("\t*int zero value: ", a)
	fmt.Println("\t*float32 zero value: ", b)
	fmt.Println("\t*bool zero value: ", c)
	fmt.Println("\t*string zero value: ", d)
	fmt.Println("\t*byte zero value: ", e)
	fmt.Println("\t*rune zero value: ", f)
	fmt.Println("\t*rune zero value: ", g)
}

var b = 1

func Steps2() {
	// 定义了一个指针变量 a, 指针变量只能存储地址
	var a *int

	fmt.Println("\ta addr:", a) // 打印 a 存储的地址值
	// 取空指针变量存储地址上的值会导致 panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println("a value:", *a) // *a 取出 a 存储的地址上的数据并打印

	fmt.Println("\tb value:", b) // 打印 b 的值

	// & 表示取 b 变量的地址并赋值给 a, 改动 a 就相当于改动 b
	a = &b
	fmt.Println("\ta addr:", a)   // 打印 a 存储的地址值
	fmt.Println("\ta value:", *a) // *a 取出 a 存储的地址上的数据并打印

	*a = 2                       // *a 取出a存储的地址上并给他赋上新值 2
	fmt.Println("\ta addr:", a)  // 打印 a 存储的地址值
	fmt.Println("\tb addr:", &b) // 打印 a 存储的地址值
	fmt.Println("\tb value:", b) // *a 取出 a 存储的地址上的数据 并打印
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
}
