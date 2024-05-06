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
	// 大写表示导出,包外能引用
	B byte
	C int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	D float32 // float64
	E string
	F []int
	G map[string]int
	H *string
}

func (d Demo) print() {
	fmt.Printf("\td %+v\n", d)
}

func (d Demo) printB() {
	fmt.Printf("\td.B %+v\n", d.B)
}

func print(d Demo) {
	fmt.Printf("\td %+v\n", d)
}

func printB(d Demo) {
	fmt.Printf("\td.B %+v\n", d.B)
}

// Steps1 值方法调用等同于函数调用
func Steps1() {
	H := "Golang"
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}, &H}
	v.print()
	v.printB()

	print(v)  // 等同于 v.print()
	printB(v) // 等同于 v.printB()
}

func (d Demo) ModifyE() {
	d.E = "Hello World"
}

// Steps2 值方法操作的是结构体副本,在值方法修改结构体属性不会影响传入的值
func Steps2() {
	H := "Golang"
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}, &H}

	// 值接收者 无法通过方法改变接收者内部值
	v.ModifyE()
	fmt.Printf("\t%+v, H:%s\n", v, *v.H)
}

func (d Demo) ModifyFGH() {
	d.F[1] = 200
	d.G["Hello"] = 2
	*(d.H) = "Tutorial"
}

// Steps3 注意, 如果结构体内部有指针持有者字段，那值方法修改结构体中指针持有者字段会影响传入的值
func Steps3() {
	H := "Golang"
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}, &H}

	v.ModifyFGH()
	fmt.Printf("\t%+v, H:%s\n", v, *v.H)
}

func (d Demo) printAddr1() {
	fmt.Printf("\td address:%p\n", &d)
	fmt.Printf("\td.E address:%p\n", &d.E)
	fmt.Printf("\td.F address %p value address:%p\n", &d.F, d.F)
	fmt.Printf("\td.G address %p value address:%p\n", &d.G, d.G)
	fmt.Printf("\td.H address %p value address:%p\n", &d.H, d.H)
}

func (d Demo) printAddr2() {
	fmt.Printf("\td address:%p\n", &d)
	fmt.Printf("\td.E address:%p\n", &d.E)
	fmt.Printf("\td.F address %p value address:%p\n", &d.F, d.F)
	fmt.Printf("\td.G address %p value address:%p\n", &d.G, d.G)
	fmt.Printf("\td.H address %p value address:%p\n", &d.H, d.H)
}

// Steps4 解释为什么？值方法修改结构体中指针持有者字段会影响传入的值
func Steps4() {
	H := "Golang"
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}, &H}

	v.printAddr1()
	fmt.Println("\t----------------------")
	v.printAddr1()
	fmt.Println("\t----------------------")
	v.printAddr2()
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
