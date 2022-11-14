package main

import (
	"fmt"
	"reflect"
)

// 类型转换
func main() {
	var a int = 1

	// 表达式 T(v) 将值 v 转换为类型 T
	// 如下 float64(a) 将值 a 转换为类型 float64
	var b float64 = float64(a)

	// 简洁形式
	c := uint(b)

	fmt.Printf("a value:%d  a type:%s\n", a, reflect.TypeOf(a))
	fmt.Printf("b value:%f  b type:%s\n", b, reflect.TypeOf(b))
	fmt.Printf("c value:%d  c type:%s\n", c, reflect.TypeOf(c))

	fmt.Printf("*************************************************************\n")

	d := uint8(255)

	// 类型转换 不能超过转换类型的范围
	//e := uint8(256)

	// 超过转换类型的范围溢出
	var f int = 256
	g := uint8(f)
	h := uint8(f + 1)

	fmt.Printf("d value:%d  d type:%s\n", d, reflect.TypeOf(d))

	fmt.Printf("g value:%d  g type:%s\n", g, reflect.TypeOf(g))
	fmt.Printf("h value:%d  h type:%s\n", h, reflect.TypeOf(h))
}
