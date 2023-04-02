package main

import (
	"fmt"
	"reflect"
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

// var 语句可以声明全局变量并赋值
var aa int64 = 3

func main() {
	// 声明变量并赋初始值
	var a int = 1
	// 如果初始化值已存在，则可以省略类型，go编辑器会自动推导类型
	var b = 2
	var c = 3.2
	var d = "hello"

	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
	// 但 := 结构不能在函数外使用,也就是说不能用于声明全局变量
	e := true

	var f byte = 'a'
	var g rune = '\u72d7' // \u72d7 = 狗

	var h interface{} = "golang"
	var i interface{} = true
	//var j,k = 1,"Golang Tutrial" // 多变量声明并赋值

	fmt.Printf("a value:%d  a type:%s\n", a, reflect.TypeOf(a))
	fmt.Printf("aa value:%d  aa type:%s\n", aa, reflect.TypeOf(aa))
	fmt.Printf("b value:%d  b type:%s\n", b, reflect.TypeOf(b))
	fmt.Printf("c value:%f  c type:%s\n", c, reflect.TypeOf(c))
	fmt.Printf("d value:%s  d type:%s\n", d, reflect.TypeOf(d))
	fmt.Printf("e value:%t  e type:%s\n", e, reflect.TypeOf(e))
	fmt.Printf("f value:%c  f type:%s\n", f, reflect.TypeOf(f))
	fmt.Printf("g value:%c  g type:%s\n", g, reflect.TypeOf(g))
	fmt.Printf("h value:%s  h type:%s\n", h, reflect.TypeOf(h))
	fmt.Printf("i value:%t  i type:%s\n", i, reflect.TypeOf(i))

	// 定义字符串变量并初始化为 Golang Tutorial
	str := "Golang Tutorial"
	strLength := len(str) // len() 函数可以获取字符串的长度
	fmt.Printf("str value:%s  str length:%d  str type:%s\n", str, strLength, reflect.TypeOf(i))
}
