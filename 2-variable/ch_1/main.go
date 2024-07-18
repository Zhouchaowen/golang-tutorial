package main

import (
	"fmt"
	"unsafe"
)

/*
	1.声明局部变量
	2.声明全局变量
	3.声明多变量
	4.各类型变量内存占比
*/

/*
	通过 var 定义变量
	var variableName T			// 初始化为零值
	var variableName T = Value	// 初始化为Value
	var variableName = Value	// 初始化为Value
	variableName := Value		// 初始化为Value
*/

// var 语句可以声明全局变量
// 全局变量: 函数外声明的变量,全局变量作用域可以在当前的整个包甚至外部包(被导出后)使用
var aa int64

// 可以在 var() 中定义多个全局变量
var (
	bb int8
	cc int16
	dd bool
	ee string
)

func Steps1() {
	// var 语句用于声明一个变量列表,默认值为对应类型零值,
	// 并且声明变量后不使用该变量的话将会抛出错误(该规则适用于函数内定义的变量)

	// 如下 var a int 定义了一个 int 类型的局部变量 a (局部变量：函数内声明的变量,作用域只在函数体内),
	// 这意味着 a 只能在 main 函数内使用（函数的参数和返回值也是局部变量）

	var a int         // 整型 uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	var b float32     // 浮点型 float64
	var c bool        // 布尔型
	var d string      // 字符串
	var e byte        // 等同于 uint8
	var f rune        // 等同于 int32,表示一个 Unicode 码点
	var g interface{} // 接口型 类似于java中的Object 可以存储任意类型的值

	// 多变量声明 通过 , 隔开
	var h, i string

	// 没有明确初始值的变量声明编译时会被赋予 零值
	// 不同类型的零值:
	//    数值类型为   0
	//    布尔类型为   false
	//    字符串类型为 ""（空字符串）
	// 	  接口类型为   nil

	// 打印对应零值
	fmt.Println("\tint zero value: ", a)
	fmt.Println("\tint64 zero value: ", aa)
	fmt.Println("\tfloat32 zero value: ", b)
	fmt.Println("\tbool zero value: ", c)
	fmt.Println("\tstring zero value: ", d)
	fmt.Println("\tbyte zero value: ", e)
	fmt.Println("\trune zero value: ", f)
	fmt.Println("\tinterface zero value: ", g)

	fmt.Println("\tstring zero value: ", h)
	fmt.Println("\tstring zero value: ", i)
}

// Steps2 不同数据类型在内存中占用的字节
func Steps2() {
	var i int // 整型 uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64
	var ptr uintptr

	var f32 float32 // 浮点型 float64
	var f64 float64

	var c bool           // 布尔型
	var d string         // 字符串
	var e byte           // 等同于 uint8
	var g interface{}    // 接口型 类似于java中的Object 可以存储任意类型的值
	var arr [10]bool     // 数组类型
	var slice []bool     // 切片类型
	var mp map[bool]bool // map(映射)类型

	var ptrUi8 *uint8
	var ptrSlice *[]bool

	fmt.Printf("\t     int memory size:%d\n", unsafe.Sizeof(i))
	fmt.Printf("\t   uint8 memory size:%d\n", unsafe.Sizeof(ui8))
	fmt.Printf("\t  uint16 memory size:%d\n", unsafe.Sizeof(ui16))
	fmt.Printf("\t  uint32 memory size:%d\n", unsafe.Sizeof(ui32))
	fmt.Printf("\t  uint64 memory size:%d\n", unsafe.Sizeof(ui64))
	fmt.Printf("\t uintptr memory size:%d\n", unsafe.Sizeof(ptr))
	fmt.Printf("\t     f32 memory size:%d\n", unsafe.Sizeof(f32))
	fmt.Printf("\t     f64 memory size:%d\n", unsafe.Sizeof(f64))
	fmt.Printf("\t    bool memory size:%d\n", unsafe.Sizeof(c))
	fmt.Printf("\t  string memory size:%d\n", unsafe.Sizeof(d))
	fmt.Printf("\t    byte memory size:%d\n", unsafe.Sizeof(e))
	fmt.Printf("\tinterfac memory size:%d\n", unsafe.Sizeof(g))
	fmt.Printf("\t     arr memory size:%d\n", unsafe.Sizeof(arr))
	fmt.Printf("\t   slice memory size:%d\n", unsafe.Sizeof(slice))
	fmt.Printf("\t     map memory size:%d\n", unsafe.Sizeof(mp))

	fmt.Printf("\t  ptrUi8 memory size:%d\n", unsafe.Sizeof(ptrUi8))
	fmt.Printf("\tptrSlice memory size:%d\n", unsafe.Sizeof(ptrSlice))
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
}
