package main

import (
	"fmt"
)

/*
	1.定义值结构体方法
	2.值结构体方法中调用属性值
	3.值结构体方法改变属性值
*/

// 方法就是一类带特殊的 接收者 参数的函数
// 接收者(可以是struct或自定义类型) 分为：
//  	1.值接收者
// 		2.指针接收者

// Demo 值接收者
type Demo struct {
	a bool
	// 大写表示导出，包外能引用
	B byte
	C int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	D float32 // float64
	E string
	F []int
	G map[string]int
}

func (d Demo) print() {
	fmt.Printf("%+v\n", d)
}

func (d Demo) printB() {
	fmt.Printf("%+v\n", d.B)
}

func (d Demo) ModifyE() {
	d.E = "Hello World"
}

func (d Demo) printAddr1() {
	fmt.Printf("%p\n", &d.a)
}

func (d Demo) printAddr2() {
	fmt.Printf("%p\n", &d.a)
}

func main() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}}
	v.print()
	v.printB()

	// 值接收者 无法通过方法改变接收者内部值
	v.ModifyE()
	v.print()

	// 值接收者
	v.printAddr1()
	v.printAddr1()
	v.printAddr2()
}
