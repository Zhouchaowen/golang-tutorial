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
	fmt.Printf("\td %+v\n", d)
}

func (d *Demo) printB() {
	fmt.Printf("\td.B %+v\n", d.B)
}

func print(d *Demo) {
	fmt.Printf("\td %+v\n", d)
}

func printB(d *Demo) {
	fmt.Printf("\td.B %+v\n", d.B)
}

func Steps1() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, [3]int{3, 4, 5}, map[string]int{"Golang": 0, "Tutorial": 1}}
	v.print()
	v.printB()

	print(&v)  // 等同于 v.print()
	printB(&v) // 等同于 v.printB()
}

func (d *Demo) modifyE() {
	d.E = "Hello World"
}

func (d *Demo) printE1() {
	fmt.Printf("\te1  value:%s\n", d.E)
}

func (d Demo) printE2() {
	fmt.Printf("\te2  value:%s\n", d.E)
}

func Steps2() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, [3]int{3, 4, 5}, map[string]int{"Golang": 0, "Tutorial": 1}}
	// 指针接收者 可以通过方法改变接收者内部值
	fmt.Printf("\t%+v\n", v)
	v.modifyE()
	fmt.Printf("\t%+v\n", v)
	v.printE1()
	v.printE2()
}

func (d *Demo) modifyF() {
	d.F = []int{11, 22}
}

func (d *Demo) printF1() {
	fmt.Printf("\tf1   value:%v\n", d.F)
}

func (d Demo) printF2() {
	fmt.Printf("\tf2   value:%v\n", d.F)
}

func Steps3() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, [3]int{3, 4, 5}, map[string]int{"Golang": 0, "Tutorial": 1}}
	fmt.Printf("\t%+v\n", v)
	v.modifyF()
	fmt.Printf("\t%+v\n", v)
	v.printF1()
	v.printF2()
}

func (d *Demo) modifyH() {
	d.H = map[string]int{"Hello": 0, "World": 1}
}

func (d *Demo) printH1() {
	fmt.Printf("\th1   value:%v\n", (*d).H) // (*d).H
}

func (d Demo) printH2() {
	fmt.Printf("\th2   value:%v\n", d.H)
}

func Steps4() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, [3]int{3, 4, 5}, map[string]int{"Golang": 0, "Tutorial": 1}}
	fmt.Printf("\t%+v\n", v)
	v.modifyH()
	fmt.Printf("\t%+v\n", v)
	v.printH1()
	v.printH2()
}

func (d *Demo) modifyStruct() {
	*d = Demo{a: false}
	fmt.Printf("\t%+v\n", d)
}

func Steps5() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, [3]int{3, 4, 5}, map[string]int{"Golang": 0, "Tutorial": 1}}
	// 有趣的值方法对比实验
	fmt.Printf("\t%+v\n", v)
	v.modifyStruct()
	fmt.Printf("\t%+v\n", v)
}

func (d *Demo) printAddr1() {
	fmt.Printf("\td address:%p\n", &d) // &d 去d本身的内存地址
	fmt.Printf("\td   value:%p\n", d)  // d 是一个指针变量，所以存储的就是一个指针值，使用 d.X 时其实等于 (*d).X 这是Golang做了语法糖
}

func (d *Demo) printAddr2() {
	fmt.Printf("\td address:%p\n", &d)
	fmt.Printf("\td   value:%p\n", d)
}

func Steps6() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, [3]int{3, 4, 5}, map[string]int{"Golang": 0, "Tutorial": 1}}
	fmt.Printf("\tv address:%p\n", &v)
	fmt.Println("\t--------------")
	v.printAddr1()
	fmt.Println("\t--------------")
	v.printAddr1()
	fmt.Println("\t--------------")
	v.printAddr2()
}

// 永远记住 Golang 是值传递，基础类型(int,bool...)是传递具体值，指针类型传递一个地址(也相当于一个int类型的值)
func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
	fmt.Println("Steps5():")
	Steps5()
	fmt.Println("Steps6():")
	Steps6()
}
