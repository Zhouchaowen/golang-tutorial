package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
	1.声明局部变量并赋值
	2.多变量声明并赋值
	3.声明全局变量并赋值
	4.简短方式声明变量并赋值
	5.声明接口类型并赋值
	6.声明字符串类型并调用内置len()函数
	7.获取字符串的某一段字符
	8.原样字符串输出
*/

// var 语句声明全局变量并赋值
var aa int64 = 3

// Steps1 声明变量并赋值
func Steps1() {
	// 声明局部变量并赋初始值
	var a int = 1

	// 如果初始化值已存在,则可以省略类型,Go编译器会自动推导类型
	var b = 2       // 自动推导为int型
	var c = 3.2     // 自动推导为float64型
	var d = "hello" // 自动推导为string型

	// 在函数中,简洁赋值语句 := 可在类型明确的地方代替 var 声明,但 := 结构不能在函数外使用,也就是说不能用于声明全局变量
	e := true

	var f byte = 'a'
	var g rune = '\u72d7' // \u72d7 = 狗

	var h interface{} = "golang"
	var i interface{} = true
	var j, k = 1, "Golang Tutorial" // 多变量声明并赋值

	// 打印变量的 value, 通过 reflect.TypeOf 函数获取变量对应类型
	fmt.Printf("\ta value:%d  a type:%s\n", a, reflect.TypeOf(a))
	fmt.Printf("\taa value:%d  aa type:%s\n", aa, reflect.TypeOf(aa))
	fmt.Printf("\tb value:%d  b type:%s\n", b, reflect.TypeOf(b))
	fmt.Printf("\tc value:%f  c type:%s\n", c, reflect.TypeOf(c))
	fmt.Printf("\td value:%s  d type:%s\n", d, reflect.TypeOf(d))
	fmt.Printf("\te value:%t  e type:%s\n", e, reflect.TypeOf(e))
	fmt.Printf("\tf value:%c  f type:%s\n", f, reflect.TypeOf(f))
	fmt.Printf("\tg value:%c  g type:%s\n", g, reflect.TypeOf(g))
	fmt.Printf("\th value:%s  h type:%s\n", h, reflect.TypeOf(h))
	fmt.Printf("\ti value:%t  i type:%s\n", i, reflect.TypeOf(i))
	fmt.Printf("\tj value:%d  j type:%s\n", j, reflect.TypeOf(j))
	fmt.Printf("\tk value:%s  k type:%s\n", k, reflect.TypeOf(k))
}

// Steps2 字符串
func Steps2() {
	// 定义字符串变量并初始化为 Golang Tutorial
	str := "Golang Tutorial"
	strLength := len(str) // len() 函数可以获取字符串的长度
	fmt.Printf("\tstr value:%s str length:%d str type:%s\n", str, strLength, reflect.TypeOf(str))
}

// Steps3 interface被动态赋值
func Steps3() {
	var i interface{} = true
	fmt.Printf("\ti value:%t  i type:%s\n", i, reflect.TypeOf(i))

	// i 被重新赋值, 类型转换为string
	i = "tutorial"
	fmt.Printf("\ti value:%s  i type:%s\n", i, reflect.TypeOf(i))
}

// Steps4 证明interface底层是由type和data组成
func Steps4() {
	/*
		interface底层由两部分组成  _type ptr, data ptr
		type eface struct {
			_type *_type
			data  unsafe.Pointer
		}
	*/

	var i interface{} = true
	fmt.Printf("\ti value:%t  i type:%s\n", i, reflect.TypeOf(i))

	fmt.Printf("\ti size:%d\n", unsafe.Sizeof(i))
	fmt.Printf("\ti addr:%p\n", &i)
	fmt.Printf("\ti type pointer value 0x%x\n", *(*uintptr)(unsafe.Pointer(&i)))
	fmt.Printf("\ti data pointer value 0x%x\n", *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + uintptr(8))))
	fmt.Printf("\ti data pointer *value %t\n", *(*bool)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + uintptr(8))))))

	fmt.Printf("\t---------------------------\n")
	// i 被重新赋值, 类型转换为string
	i = "tutorial"
	fmt.Printf("\ti value:%s  i type:%s\n", i, reflect.TypeOf(i))

	fmt.Printf("\ti size:%d\n", unsafe.Sizeof(i))
	fmt.Printf("\ti addr:%p\n", &i)
	fmt.Printf("\ti type pointer value 0x%x\n", *(*uintptr)(unsafe.Pointer(&i)))
	fmt.Printf("\ti data pointer value 0x%x\n", *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + uintptr(8))))
	fmt.Printf("\ti data pointer *value %s\n", *(*string)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + uintptr(8))))))

	// 等同如上
	type eface struct {
		_type *struct{} // *_type
		data  unsafe.Pointer
	}
	s := *(*eface)(unsafe.Pointer(&i))
	fmt.Printf("\ti %+v\n", s)
	fmt.Printf("\ti._type: %+v\n", s._type)
	fmt.Printf("\ti.data:  %+v\n", *(*string)(s.data))
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
