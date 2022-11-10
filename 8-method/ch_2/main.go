package main

import (
	"fmt"
)

// 使用指针接收者的原因：
// 		首先，方法能够修改其接收者指向的值。
// 		其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

// 指针接收者
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
	fmt.Printf("%+v\n", d)
}

func (d *Demo) printB() {
	fmt.Printf("%+v\n", d.B)
}

func (d *Demo) ModifyE() {
	d.E = "Hello World"
}

func (d *Demo) printAddr1() {
	fmt.Printf("%p\n", &d.a)
}

func (d *Demo) printAddr2() {
	fmt.Printf("%p\n", &d.a)
}

func main() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}}
	v.print()
	v.printB()

	// 指针接收者 可以通过方法改变接收者内部值
	v.ModifyE()
	v.print()

	v.printAddr1()
	v.printAddr1()
	v.printAddr2()
}
