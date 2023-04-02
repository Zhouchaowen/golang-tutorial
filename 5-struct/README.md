# 结构体

在 `Go` 语言中，`struct` 是一种用户自定义的复合类型，可以将多个字段组合在一起，形成一个新的结构体类型。通常情况下，结构体类型用于封装多个相关的数据字段，以便更方便地进行操作和管理。

## 目录

- 结构体定义
- 定义结构体值方法
- 定义结构体指针方法
- 自定义类型

## 结构体定义

结构体类型的定义可以通过 `type` 关键字和 `struct` 关键字来完成, 语法如下：

```go
type StructName struct {
    Field1 FieldType1
    Field2 FieldType2
    ...
    FieldN FieldTypeN
}
```

其中，`StructName` 表示结构体类型的名称，`Field1`、`Field2` 等表示结构体的数据字段，`FieldType1`、`FieldType2` 等表示字段的数据类型。

- 如下展示结构体中定义常用字段并初始化结构体：

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

	fmt.Printf("%+v\n", d) // 打印整个结构体

	// 结构体字段使用点号来访问
	d.a = false // 修改a字段的值

	fmt.Printf("%+v\n", d)

	fmt.Printf("dome.B: %c\n", d.B)
}

func main() {
	Steps1()
}
```

以上代码，我们定义了一个Demo结构体，包含了一些常见字段。在定义结构体类型之后，我们可以通过结构体字面量的方式来创建结构体变量, 并初始化一些数据。

```go
d := Demo{ // 创建一个 Demo 类型的结构体
  a: true,
  B: 'b',
  C: 1,
  D: 1.0,
  E: "E",
  F: []int{1},
  G: map[string]int{"GOLANG": 1},
}
```

在创建结构体变量之后，我们可以通过`.`运算符来修改或访问结构体的数据字段。

```go
// 结构体字段使用点号来访问
d.a = false // 修改a字段的值

fmt.Printf("%+v\n", d)

fmt.Printf("dome.B: %c\n", d.B)
```

- 函数内定义结构体：

```go
package main

import "fmt"

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
	Steps2()
}
```

## 结构体方法

除了定义数据字段之外，结构体类型还可以定义相关的方法。方法是一种与特定类型相关联的函数，可以对该类型的值进行操作。在 Go 语言中，可以通过 `func` 关键字和结构体类型的名称来定义方法，语法如下：

```go
func (p StructName) MethodName(parameter1 Type1, parameter2 Type2, ...) ReturnType {
    // 方法的实现代码
}
```

其中，`StructName` 表示当前方法属于这结构体。方法名和参数列表后面的部分与普通函数的定义类似，用于指定方法的输入和输出。

### 值方法

```go
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
	fmt.Printf("%p\n", &d)
}

func (d Demo) printAddr2() {
	fmt.Printf("%p\n", &d)
}

func main() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}}
	v.print()
	v.printB()

	// 值接收者 无法通过方法改变接收者内部值
	v.ModifyE()
	fmt.Printf("%+v\n", v)

	// 值接收者
	v.printAddr1()
	v.printAddr1()
	v.printAddr2()
}
```

在上面的代码中， `print(),printB(),ModifyE(),printAddr1 (),printAddr2()` 方法绑定到 `Demo` 结构体上，并且只能通过`Demo`结构体的实例才能调用。

每个方法中都使用 `d` 作为接收者名称 (当然接收者d可以任意取名)，表示当 `Demo` 类型的实例调用该方法时，实例本身的数据会被赋值给接收者 `d` ，从而可以通过接收者`d`在结构体方法中访问该实例的数据字段,  例如：

```go
v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}}
v.printB() // 打印 G
```

`v`是`Demo`结构体的一个实例，当调用`v.printB()`结构体方法时，`v`实例中的数据会拷贝一份给`d`, 这样在`printB()`中调用`d.B`时就可以获取到`G`这个数据了。

需要注意的是，上面定义的这些方法都是值方法。`v`实例赋值到接收者`d`也是通过拷贝一份数据的方式，所以在方法中修改接收者`d`的数据并不会影响到`v`实例的数据。

```go
v.ModifyE()
fmt.Printf("%+v\n", v)

// 执行结果
{a:true B:71 C:1 D:1 E:Golang Tutorial F:[1 2] G:map[Golang:0 Tutorial:1]}
```

以上两个方法调用证明了这一点，`ModifyE()`方法中修改了`E`字段，并不会影响到v实例。

### 指针方法

```go
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
	fmt.Printf("%+v\n", d)
}

func (d *Demo) printB() {
	fmt.Printf("%+v\n", d.B)
}

func (d *Demo) ModifyE() {
	d.E = "Hello World"
}

func (d *Demo) printAddr1() {
	fmt.Printf("%p\n", d)
}

func (d *Demo) printAddr2() {
	fmt.Printf("%p\n", d)
}

func main() {
	v := Demo{true, 'G', 1, 1.0, "Golang Tutorial", []int{1, 2}, map[string]int{"Golang": 0, "Tutorial": 1}}
	v.print()
	v.printB()

	// 指针接收者 可以通过方法改变接收者内部值
	v.ModifyE()
	fmt.Printf("%+v\n", v)

	v.printAddr1()
	v.printAddr1()
	v.printAddr2()
}

```

在上面的代码中， `print(),printB(),ModifyE(),printAddr1 (),printAddr2()` 这些方法都是定义的指针方法，与值方法不同的是定义方法时结构体使用指针类型：`func (p *StructName) MethodName(parameter1 Type1, parameter2 Type2, ...) ReturnType {}`

并且 `v`实例赋值到接收者`d`是通过传递指针的方式，所以通过接收者`d`修改数据会影响`v`实例的数据。

```go
v.ModifyE()
fmt.Printf("%+v\n", v)

// 执行结果
{a:true B:71 C:1 D:1 E:Hello World F:[1 2] G:map[Golang:0 Tutorial:1]}
```

所以如果方法需要修改接收者的值，那么必须使用指针类型的接收者。如果使用值类型的接收者，则只能访问接收者的数据字段，而不能修改接收者的值。

## 自定义类型定义方法

自定义类型方法和结构体方法使用方式基本一致。

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
```go
type numb struct {
	a,b int
}

func (n numb) add() int {
	return n.a+n.b
}
```

2. 定义一个圆结构体,并定义求圆面积,周长和输入角度求弧长等方法。

```go
type circle struct{
  radius float64
}
```

## 参考

