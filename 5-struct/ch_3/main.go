package main

import (
	"fmt"
)

/*
	1.定义指针结构体方法
	2.指针结构体方法中调用属性值
	3.指针结构体方法改变属性值
*/

// 使用指针接收者的原因：
// 		首先，方法能够修改其接收者指向的值。
// 		其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

// Demo 指针接收者
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

func (d *Demo) print() {
	fmt.Printf("d %+v\n", d)
}

func (d *Demo) printB() {
	fmt.Printf("d.B %+v\n", d.B)
}

func print(d *Demo) {
	fmt.Printf("d %+v\n", d)
}

func printB(d *Demo) {
	fmt.Printf("d.B %+v\n", d.B)
}

func (d *Demo) ModifyE() {
	d.E = "Hello World"
}

func (d *Demo) printAddr1() {
	fmt.Printf("d address:%p\n", &d)
	fmt.Printf("d   value:%p\n", d)
}

func (d *Demo) printAddr2() {
	fmt.Printf("d address:%p\n", &d)
	fmt.Printf("d   value:%p\n", d)
}

func main() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}}
	v.print()
	v.printB()

	print(&v)  // 等同于 v.print()
	printB(&v) // 等同于 v.printB()

	// 指针接收者 可以通过方法改变接收者内部值
	v.ModifyE()
	fmt.Printf("%+v\n", v)

	fmt.Println("--------------")
	fmt.Printf("v address:%p\n", &v)
	fmt.Println("--------------")
	v.printAddr1()
	fmt.Println("--------------")
	v.printAddr1()
	fmt.Println("--------------")
	v.printAddr2()
}
