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
	G [3]int
	H map[string]int
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

func (d *Demo) modifyE() {
	d.E = "Hello World"
}

func (d *Demo) printE1() {
	fmt.Printf("e1  value:%s\n", d.E)
}

func (d Demo) printE2() {
	fmt.Printf("e2  value:%s\n", d.E)
}

func (d *Demo) modifyF() {
	d.F = []int{11, 22}
}

func (d *Demo) printF1() {
	fmt.Printf("f1   value:%v\n", d.F)
}

func (d Demo) printF2() {
	fmt.Printf("f2   value:%v\n", d.F)
}

func (d *Demo) modifyH() {
	d.H = map[string]int{"Hello": 0, "World": 1}
}

func (d *Demo) printH1() {
	fmt.Printf("h1   value:%v\n", (*d).H) // (*d).H
}

func (d Demo) printH2() {
	fmt.Printf("h2   value:%v\n", d.H)
}

func (d *Demo) modifyStruct() {
	d = &Demo{a: false}
	fmt.Printf("d   value:%v\n", d)
}

func (d *Demo) printAddr1() {
	fmt.Printf("d address:%p\n", &d) // &d 去d本身的内存地址
	fmt.Printf("d   value:%p\n", d)  // d 是一个指针变量，所以存储的就是一个指针值，使用 d.X 时其实等于 (*d).X 这是Golang做了语法糖
}

func (d *Demo) printAddr2() {
	fmt.Printf("d address:%p\n", &d)
	fmt.Printf("d   value:%p\n", d)
}

func main() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, [3]int{3, 4, 5}, map[string]int{"Golang": 0, "Tutorial": 1}}
	v.print()
	v.printB()

	print(&v)  // 等同于 v.print()
	printB(&v) // 等同于 v.printB()

	// 指针接收者 可以通过方法改变接收者内部值
	fmt.Println("--------------")
	fmt.Printf("%+v\n", v)
	v.modifyE()
	fmt.Printf("%+v\n", v)
	v.printE1()
	v.printE2()

	fmt.Println("--------------")
	fmt.Printf("%+v\n", v)
	v.modifyF()
	fmt.Printf("%+v\n", v)
	v.printF1()
	v.printF2()

	fmt.Println("--------------")
	fmt.Printf("%+v\n", v)
	v.modifyH()
	fmt.Printf("%+v\n", v)
	v.printH1()
	v.printH2()

	// 有趣的值方法对比实验
	fmt.Println("--------------")
	fmt.Printf("%+v\n", v)
	v.modifyStruct()
	fmt.Printf("%+v\n", v)

	// 永远记住 Golang 是值传递，基础类型(int,bool...)是传递具体值，指针类型传递一个地址(也相当于一个int类型的值)
	fmt.Println("--------------")
	fmt.Printf("v address:%p\n", &v)
	fmt.Println("--------------")
	v.printAddr1()
	fmt.Println("--------------")
	v.printAddr1()
	fmt.Println("--------------")
	v.printAddr2()
}
