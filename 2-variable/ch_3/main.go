package main

import (
	"fmt"
	"reflect"
)

/*
	1.相近类型强转
	2.数据溢出与精度丢失
	3.不同类型的数据计算
*/

/*
	类型转换
	variableName2 := T(variableName1)
*/

func main() {
	var a int = 1

	// 表达式 T(v) 将值 v 转换为类型 T

	// 如下将 int 类型的值转换为 float64 类型的值, float64(a) 将值 a 转换为 float64 类型并赋值给 b
	var b float64 = float64(a)

	// 简洁形式 uint(b) 将值 b 转换为 uint 类型并赋值给 c
	c := uint(b)

	fmt.Printf("a value:%d  a type:%s\n", a, reflect.TypeOf(a))
	fmt.Printf("b value:%f  b type:%s\n", b, reflect.TypeOf(b))
	fmt.Printf("c value:%d  c type:%s\n", c, reflect.TypeOf(c))

	fmt.Printf("*************************************************************\n")

	d := uint8(255) // 将常量值 255 转换为 uint8 类型并赋值给 d
	fmt.Printf("d value:%d  d type:%s\n", d, reflect.TypeOf(d))

	// 注意：类型转换 不能超过转换类型的范围
	//e := uint8(256) // 编译错误, 常量256超出了uint8最大存储限制, 不能转换

	// 转换时, 超过转换的类型范围时将导致数据溢出
	var f uint16 = 256
	g := uint8(f) // uint8最大为255, 溢出后从0开始, 所以g等于0
	ff := f + 1
	h := uint8(ff) // 如上可知h等于1

	fmt.Printf("f  binary value:%016b f  value:%d   f type:%s\n", f, f, reflect.TypeOf(f))
	fmt.Printf("g  binary value:%016b g  value:%d     g type:%s\n", g, g, reflect.TypeOf(g))
	fmt.Printf("ff binary value:%016b ff value:%d  ff type:%s\n", ff, ff, reflect.TypeOf(ff))
	fmt.Printf("h  binary value:%016b h  value:%d     h type:%s\n", h, h, reflect.TypeOf(h))

	j := 10    // 自动推导为int型
	l := 100.1 // 自动推导为float64型

	// 在Go中不同类型的数据不能直接计算，需进行类型转换
	p := float64(j) * l

	fmt.Printf("p value:%f  p type:%s\n", p, reflect.TypeOf(p))
}
