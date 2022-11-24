package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
	定义常量
	const constantName = value
	const constantName T = value

*/

// 常量定义的时候必须赋值，定义后值不能被修改
// 常量的声明与变量类似，使用 const 关键字, 常量中的数据类型只可以是字符、字符串、布尔值或数值

// const NameOfVariable [type] = value  type 可以省略然编译器推导

// 全局常量
const PI = 3.14
const NAME = "Golang-tutorial"
const OK bool = true

// 可以在 const 中定义多个常量
const (
	MaxUint8  = math.MaxUint8
	MaxUint16 = math.MaxUint16
)

// iota 定义常量
// iota的值是const语句块里的行索引，行索引从0开始
const (
	One = iota
	Two
	Three
)

func main() {
	// 函数内也可以定义常量（局部常量）
	const World = "World"
	fmt.Println("Hello", World)

	fmt.Printf("MaxUint8 value:%d  MaxUint8 type:%s\n", MaxUint8, reflect.TypeOf(MaxUint8))
	fmt.Printf("MaxUint16 value:%d  MaxUint16 type:%s\n", MaxUint16, reflect.TypeOf(MaxUint16))

	fmt.Printf("One value:%d  One type:%s\n", One, reflect.TypeOf(One))
	fmt.Printf("Two value:%d  Two type:%s\n", Two, reflect.TypeOf(Two))
	fmt.Printf("Three value:%d  Three type:%s\n", Three, reflect.TypeOf(Three))

}
