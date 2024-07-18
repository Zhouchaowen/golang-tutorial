package main

import "fmt"

/*
	1.声明局部指针变量(指针地址与指针类型)
	2.指针变量的赋值与取值
	3.使用指针修改值
	4.通过new()创建指针变量
*/

/*
	通过 var 定义指针变量
	var variableName *T
	var variableName *T = &Value
	var variableName = &Value
	variableName := &Value
*/

// Steps1 定义指针变量
func Steps1() {
	// 定义一个 int 类型的指针变量, 默认值零值为nil
	var a *int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	var b *float32 // float64
	var c *bool
	var d *string // 定义一个 string 的指针类型
	var e *byte   // 定义一个 byte 的指针类型
	var f *rune
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

// Steps2 指针变量赋值与取值
func Steps2() {
	// 定义了一个指针变量 a, 指针变量只能存储地址
	var a *int

	fmt.Println("\ta value:", a) // 打印 a 变量存储的地址值
	fmt.Println("\ta addr:", &a) // 打印 a 变量自己的地址值

	// 取空指针变量存储地址上的值会导致 panic: runtime error: invalid memory address or nil pointer dereference
	//fmt.Println("a value:", *a) // *a 代表取出 a 变量存储的地址, 并获取该地址上存储的值

	fmt.Println("\tb value:", b)  // 打印 b 的值
	fmt.Println("\tb addr :", &b) // 打印 b 的地址

	// & 表示取 b 变量的地址并赋值给 a, 改动 a 就相当于改动 b
	a = &b
	fmt.Println("\ta value:", a)             // 打印 a 变量存储的地址值
	fmt.Println("\ta value over value:", *a) // *a 代表取出 a 变量存储的地址, 并获取该地址上存储的值

	*a = 2                       // *a 取出 a 变量存储的地址，并修该地址上存储的值(赋值为 2)
	fmt.Println("\ta value:", a) // 打印 a 存储的地址值
	fmt.Println("\ta value over value:", *a)
	fmt.Println("\tb value:", b) // 打印 b 的值

	c := &a
	// Go指针不支持算术运算, 下面这两行编译不通过。
	// c++
	// c = (&a) + 8
	_ = c

	// Go指针不支持算术运算, 可以通过 unsafe.Pointer 打破这个限制
}

// Steps3 内置函数 new 创建指针
func Steps3() {
	// 通过内置函数 new 创建一个 int 的指针类型
	a := new(int)
	var b *int
	fmt.Println("\tnew(int) value: ", a)
	fmt.Println("\tnew(int) value over value: ", *a)
	fmt.Println("\t*int value: ", b)
	// 取空指针变量存储地址上的值会导致 panic,这就是a := new(int)和var b *int创建指针的区别,new分配空间并赋值零值，定义指针并不会创建内存空间
	//fmt.Println("\t*int value over value: ", *b)

	c := new(string)
	var d *string
	fmt.Println("\tnew(string) value: ", c)
	fmt.Println("\tnew(string) value over value: ", *c)
	fmt.Println("\t*string value: ", d)

	e := new([]int)
	fmt.Println("\tnew(string) value: ", e)
	fmt.Println("\tnew(string) value over value: ", *e)
}

// Steps4 判断指针是否为nil
func Steps4() {
	var ptr *int // 指向nil
	if ptr == nil {
		fmt.Println("\tptr is nil")
	}

	var a int = 42
	var ptr2 *int = &a // 指向a的地址内存空间
	if ptr2 != nil {
		fmt.Println("\tptr2 is not nil")
	}
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
}
