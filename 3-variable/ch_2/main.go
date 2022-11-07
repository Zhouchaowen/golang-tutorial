package main

import (
	"fmt"
	"reflect"
)

// var 语句可以出现在包或函数级别
var aa int64 = 3

func main() {
	// 变量声明可以包含初始值
	var a int = 1
	// 如果初始化值已存在，则可以省略类型，go编辑器会自动推导类型
	var b = 2
	var c = 3.2
	var d = "hello"

	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
	// 但 := 结构不能在函数外使用
	e := true

	var f byte = 'a'
	var g rune = '\u72d7' // \u72d7 = 狗

	var h interface{} = "golang"
	var i interface{} = true

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
}
