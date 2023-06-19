package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
	1.声明局部常量
	2.声明全局变量
	3.iota赋值
*/

/*
	const NameOfVariable [type] = value  type 可以省略让编译器推导
	const constantName = value
	const constantName T = value
*/

// 常量定义的时候必须赋值，定义后值不能被修改

const PI = 3.14
const NAME = "Golang-tutorial"
const OK bool = true

// 可以通过 const() 定义多个常量
const (
	MaxUint8  = math.MaxUint8
	MaxUint16 = math.MaxUint16
)

// iota 定义常量, iota的值是const语句块里的行索引,行索引从0开始
const (
	One   = iota // 第一行 One值等于 0
	Two          // 第二行 Two值等于 1
	Three        // 第三行 Three值等于 2
)

func main() {
	// 函数内也可以定义常量(局部常量)
	const World = "World"
	fmt.Println("Hello", World)

	fmt.Printf("MaxUint8 value:%d  MaxUint8 type:%s\n", MaxUint8, reflect.TypeOf(MaxUint8))
	fmt.Printf("MaxUint16 value:%d  MaxUint16 type:%s\n", MaxUint16, reflect.TypeOf(MaxUint16))

	fmt.Printf("One value:%d  One type:%s\n", One, reflect.TypeOf(One))
	fmt.Printf("Two value:%d  Two type:%s\n", Two, reflect.TypeOf(Two))
	fmt.Printf("Three value:%d  Three type:%s\n", Three, reflect.TypeOf(Three))

	//MaxUint8 = math.MaxUint32 // 修改常量值将会报错
}
