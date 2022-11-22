# 结构体

## 目录

- ch_1 结构体基础用法
- ch_2 定义结构体值方法
- ch_3 定义结构体指针方法
- ch_4 定义自定义类型的方法

## 结构体定义

golang可以通过struct关键字定义结构体来抽象一个数据集或对象

```go
package main

import "fmt"

// Demo 定义结构体
type Demo struct {
	// 小写表示不导出,包外不能引用
	a bool
	// 大写表示导出，包外能引用
	B byte
	C int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	D float32 // float64
	E string
	F []int
	G map[string]int
	H *int64
}

func Steps1() {
	d := Demo{ // 创建一个 Demo 类型的结构体
		a: true,
		B: 'b',
		C: 1,
		D: 1.0,
		E: "E",
		F: []int{1},
		G: map[string]int{"GOLANG": 1},
	}

	fmt.Printf("%+v\n", d)

	// 结构体字段使用点号来访问
	d.a = false // 修改a字段的值

	fmt.Printf("%+v\n", d)

	fmt.Printf("dome.B: %c\n", d.B)
}

func Steps2() {
	// 结构体也可以定义在函数内
	type Demo struct {
		a int
		B string
	}

	d := Demo{ // 创建一个 Demo 类型的结构体
		a: 1,
	}

	fmt.Printf("%+v\n", d)

	// 结构体字段使用点号来访问
	d.a = 2 // 修改a字段的值

	fmt.Printf("%+v\n", d)
}

func main() {
	Steps1()
	Steps2()
}
```

## 结构体方法

### 值方法

```go
package main

import (
	"fmt"
)

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
```

### 指针方法

```go
package main

import (
	"fmt"
)

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
```

## 自定义类型定义方法

```go
package main

import "fmt"

// ResponseStatus 自定义类型的方法
type ResponseStatus int

const (
	QuerySuccess ResponseStatus = iota
	QueryError
)

func (r ResponseStatus) ToCN() string {
	switch r {
	case 0:
		return "query success"
	case 1:
		return "query error"
	default:
		return "non"
	}
}

func main() {
	fmt.Println(QuerySuccess.ToCN())
	fmt.Println(QueryError.ToCN())
}
```




## 思考题
1. 通过结构体方法的形式实现加减乘除
```bigquery
type numb struct {
	a,b int
}

func (n numb) add() int {
	return n.a+n.b
}
```

2. 定义一个圆结构体,并定义求圆面积,周长和输入角度求弧长等方法。

## 参考

