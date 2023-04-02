# 变量与常量

变量和常量简单来说就是给内存中某一个地址起一个名字, 然后用这个地址存储某个特定类型的值。

## 目录

- 数据类型
- 定义变量

- 变量赋值
- 类型转换
- 定义常量
- 定义函数变量
- 定义指针变量
- 占位符
- 运算符

## 数据类型

**数据类型分类：**

- 布尔类型：`bool`。
- 整数类型：`int8`、`uint8`、`int16`、`uint16`、`int32`、`uint32`、`int64`、`uint64`、`int`、`uint`、 `uintptr`。
- 浮点数类型: `float32`、`float64`。
- 复数类型：`complex64`、`complex128`。
- 字符串类型：`string`。
- 字符类型：`byte`。
- 指针持有者类型：`[size]T`(数组)、`[]T`(切片)、`map[T]T`、`struct`、`func`。

**数据类型占用大小：**

```go
bool
string

uint        either 32 or 64 bits
int         same size as uint
uintptr     an unsigned integer large enough to store the uninterpreted bits of
            a pointer value
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers
            (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32 (represents a Unicode code point)
```

## 定义变量

在`Golang`通过`var`关键字定义变量, 格式有多种最常用的两种为`var variableName T`和`variableName := Value`。

变量可以定义在函数外, 当做该`package`下的全局变量, 也可以定义在函数内当, 做该函数内的局部变量。

```go
package main

import "fmt"

/*
	通过 var 定义变量
	var variableName T			// 初始化为零值
	var variableName T = Value	// 初始化为Value
	var variableName = Value	// 初始化为Value
	variableName := Value		// 初始化为Value
*/

// var 语句可以声明全局变量
// 全局变量: 函数外声明的变量，全局变量作用域可以在当前的整个包甚至外部包(被导出后)使用
var aa int64

// 可以在 var 中定义多个全局变量
var (
	bb int8
	cc int16
	dd bool
	ee string
)

func main() {
	// var 语句用于声明一个变量列表,默认值为对应零值，并且声明变量后不使用该变量的话将会抛出错误
	// 如下 var a int 定义了一个 int 类型的局部变量 a, 局部变量：函数内声明的变量，作用域只在函数体内。
	// 这意味着 a 只能在 main 函数内使用（函数的参数和返回值也是局部变量）
	var a int         // 整型 uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	var b float32     // float64
	var c bool        // 布尔型
	var d string      // 字符串
	var e byte        // 等同于 uint8
	var f rune        // 等同于 int32,表示一个 Unicode 码点
	var g interface{} // 接口型

	// 多变量声明
	var h, i string

	// 没有明确初始值的变量声明会被赋予它们的 零值
	// 零值是:
	//    数值类型为 0
	//    布尔类型为 false
	//    字符串为 ""（空字符串）
	// 	  接口为 nil

	// 打印对应零值
	fmt.Println("int zero value: ", a)
	fmt.Println("int64 zero value: ", aa)
	fmt.Println("float32 zero value: ", b)
	fmt.Println("bool zero value: ", c)
	fmt.Println("string zero value: ", d)
	fmt.Println("byte zero value: ", e)
	fmt.Println("rune zero value: ", f)
	fmt.Println("interface zero value: ", g)

	fmt.Println("string zero value: ", h)
	fmt.Println("string zero value: ", i)
}
```

## 变量赋值

`Golang`中通过 `=` 对变量进行赋值, `=` 可以在变量初始化时赋值也可以在变量定义后赋值。可以通过上文提到的简洁赋值 `:=`, `:=` 表示定义变量并赋值, 可以替代`Var`。

需要注意的是，如果您在定义变量时不指定类型，Go 会自动推断其类型, 如下`var b = 2`。

```go
package main

import (
	"fmt"
	"reflect"
)

// var 语句可以声明全局变量并赋值
var aa int64 = 3

func main() {
	// 声明变量并赋初始值
	var a int = 1
	// 如果初始化值已存在，则可以省略类型，go编辑器会自动推导类型
	var b = 2
	var c = 3.2
	var d = "hello"

	// 在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明
	// 但 := 结构不能在函数外使用
	e := true

	var f byte = 'a'
	var g rune = '\u72d7' // \u72d7 = 狗

	var h interface{} = "golang"
	var i interface{} = true

	fmt.Printf("a value:%d  a type:%s\n", a, reflect.TypeOf(a))
	fmt.Printf("aa value:%d  aa type:%s\n", aa, reflect.TypeOf(aa))
	fmt.Printf("b value:%d  b type:%s\n", b, reflect.TypeOf(b))
	fmt.Printf("c value:%f  c type:%s\n", c, reflect.TypeOf(c))
	fmt.Printf("d value:%s  d type:%s\n", d, reflect.TypeOf(d))
	fmt.Printf("e value:%t  e type:%s\n", e, reflect.TypeOf(e))
	fmt.Printf("f value:%c  f type:%s\n", f, reflect.TypeOf(f))
	fmt.Printf("g value:%c  g type:%s\n", g, reflect.TypeOf(g))
	fmt.Printf("h value:%s  h type:%s\n", h, reflect.TypeOf(h))
	fmt.Printf("i value:%t  i type:%s\n", i, reflect.TypeOf(i))
}
```

## 类型转换

`Golang`中可以将相近类型的数据进行强转, 格式为`variableName2 := T(variableName1)`。 需要注意的是，不是所有类型之间都可以进行转换。例如，您不能将一个字符串转换为一个整数，除非该字符串表示一个整数。

当进行类型转换时，必须确保类型转换是安全的，如果转换失败，则会引发一个运行时错误。另外，由于类型转换可能导致精度损失或溢出，因此在进行类型转换时应特别小心。

```go
package main

import (
	"fmt"
	"reflect"
)

/*
	类型转换
	variableName2 := T(variableName1)

*/

func main() {
	var a int = 1

	// 表达式 T(v) 将值 v 转换为类型 T
	// 如下将 int 类型的值转换为 float64 类型的值
  // float64(a) 将值 a 转换为类型 float64
	var b float64 = float64(a)

	// 简洁形式
	c := uint(b)

	fmt.Printf("a value:%d  a type:%s\n", a, reflect.TypeOf(a))
	fmt.Printf("b value:%f  b type:%s\n", b, reflect.TypeOf(b))
	fmt.Printf("c value:%d  c type:%s\n", c, reflect.TypeOf(c))

	fmt.Printf("*************************************************************\n")

	d := uint8(255)

	fmt.Printf("d value:%d  d type:%s\n", d, reflect.TypeOf(d))
	// 类型转换 不能超过转换类型的范围
	//e := uint8(256) // 编译错误, 常量256溢出了uint8

	// 超过转换类型的范围溢出
	var f int = 256
	g := uint8(f)
	h := uint8(f + 1)

	fmt.Printf("g value:%d  g type:%s\n", g, reflect.TypeOf(g))
	fmt.Printf("h value:%d  h type:%s\n", h, reflect.TypeOf(h))

	j := 10
	l := 100.1
	// 不同类型在golang中不能计算，需进行类型转换
	p := float64(j) * l

	fmt.Printf("p value:%f  p type:%s\n", p, reflect.TypeOf(p))
}
```

注意：如果要在类型之间进行转换，但不确定是否能够安全地进行转换，可以使用类型断言。类型断言可以判断**一个接口类型的值是否属于指定的类型**，如果是，则返回转换后的值，否则返回一个错误。（该知识点将在`Interface`知识点讲解）

## 定义常量

`Golang`中通过`const`定义常量, 格式为`const constantName = value`或`const constantName T = value`, 其中 `constantName` 表示常量的名称，`T` 表示常量的类型（可以省略，此时类型会根据值自动推导），`value` 表示常量的值。

常量可以定义在函数外当做该包下的全局常量, 也可以定义在函数内当做该函数内的局部常量。注意：常量定义的时候必须赋值，定义后值不能被修改。

```go
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

// const NameOfVariable [type] = value  type 可以省略让编译器推导

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
```

在上面的示例中，我们声明了三个常量：`PI`, `NAME` 和 `OK`。`PI` 的类型为 `float64`，其值为 `3.14`,  `NAME` 的类型为 `string`，其值为 `"Golang-tutorial"`, `OK`的类型为 `bool`, 其值为 `true`。

## 定义函数变量

在`Golang`中，函数可以像普通变量一样被声明和使用。这意味着可以将一个函数赋值给一个变量, 也可以将一个函数作为另一个函数的参数或返回值。

```go
package main

import (
	"fmt"
)

/*
	定义函数变量
	var variableName = func
*/

func compute(x, y int, handler func(x, y int) int) int {
	x = x * 10
	y = y * 10
	return handler(x, y)
}

// 函数也可以当做类型,可以像其它值一样传递
func main() {
  // 1
	var add = func(x, y int) int {
		return x + y
	}
	fmt.Println("add",add(1, 2))

  // 2
	Multi := func(x, y int) int {
		return x * y
	}
	fmt.Println("Multi",compute(1, 2, Multi))
}
```

在上面的示例中第1部分，我们在 `main` 函数中声明了一个名为 `add` 的函数变量，它接受两个整数作为参数并返回它们的和。然后调用 `add` 变量，将其作为一个普通的函数来计算 1 和 2 的和，最后将结果输出到控制台上。

在上面的示例中第2部分，我们定义了一个函数变量`Multi`，它接受两个整数作为参数并返回它们的乘积。然后，我们定义了一个函数 `compute`，它接受两个整数，以及一个函数作为参数。`compute` 函数会调用传入的函数`handler`，并将两个整数+10后作为参数传递给它。最后，我们在 `main` 函数调用 `compute` 函数，将两个整数和 `Multi`函数作为参数传递给它，然后将结果输出到控制台上。

需要注意的是，函数变量不能像普通变量那样赋值为 `nil`。如果要声明一个为空的函数变量，可以使用函数类型的零值 `func() {}`。

## 定义指针变量

`Golang`中也是通过`var`定义指针变量, 格式有多种选择常用的三种`var variableName *T`, `variableName := &Value`, `variableName := new(T)`，指针变量还可以与 `nil` 进行比较，以检查它们是否指向了某个地址。如果一个指针变量为 `nil`，则表示它没有指向地址。

```go
package main

import "fmt"

/*
	通过 var 定义指针变量
	var variableName *T
	var variableName *T = Value
	var variableName = &Value
	variableName := &Value
*/

// Steps1 定义指针变量
func Steps1() {
	// 定义一个 int 的指针类型, 并且未指定任何地址
	var a *int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	var b *float32 // float64
	var c *bool
	var d *string // 定义一个 string 的指针类型
	var e *byte   // 定义一个 byte 的指针类型
	var f *rune
	var g *interface{}

	fmt.Println("\t*int zero value: ", a)
	fmt.Println("\t*float32 zero value: ", b)
	fmt.Println("\t*bool zero value: ", c)
	fmt.Println("\t*string zero value: ", d)
	fmt.Println("\t*byte zero value: ", e)
	fmt.Println("\t*rune zero value: ", f)
	fmt.Println("\t*rune zero value: ", g)
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
}

```

- 指针变量的赋值

在 Go 中，还可以通过取地址符 `&` 来获取一个变量的内存地址，并将其赋值给指针变量

```go
package main

import "fmt"

// 注释 ①
var b = 1

// Steps2 指针变量赋值与取值
func Steps2() {
	// 定义了一个指针变量 a, 指针变量只能存储地址
	var a *int

	fmt.Println("\ta addr:", a) // 打印 a 存储的地址值
	// 取空指针变量存储地址上的值会导致 
  // panic: runtime error: invalid memory address or nil pointer dereference
  
	//fmt.Println("a value:", *a) // *a 取出 a 存储的地址上的数据并打印

	fmt.Println("\tb value:", b) // 打印 b 的值

	// 注释 ② & 表示取 b 变量的地址并赋值给 a, 改动 a 就相当于改动 b
	a = &b
	fmt.Println("\ta addr:", a)   // 打印 a 存储的地址值
	fmt.Println("\ta value:", *a) // *a 取出 a 存储的地址上的数据并打印

	*a = 2                       // *a 取出a存储的地址上并给他赋上新值 2
	fmt.Println("\ta addr:", a)  // 打印 a 存储的地址值
	fmt.Println("\tb addr:", &b) // 打印 a 存储的地址值
	fmt.Println("\tb value:", b) // *a 取出 a 存储的地址上的数据 并打印

	c := &a
	// Go指针不支持算术运算, 下面这两行编译不通过。
	// c++
	// c = (&a) + 8
	_ = c

	// Go指针不支持算术运算, 可以通过 unsafe.Pointer 打破这个限制
}

func main() {
	fmt.Println("Steps2():")
	Steps2()
}
```

在上面的代码中，(注释①) 我们定义了一个整数变量 `b`，并将其初始化为 1。 (注释②) 然后，我们使用取地址符 `&` 来获取变量 `b` 的内存地址，并将其赋值给指针变量 `a`。现在，指针变量 `a` 指向了变量 `b` 的内存地址，打印a存储的地址和地址上的具体值。

需要注意的是，指针变量只能指向相同类型的变量。例如，如果我们定义了一个整数类型的指针变量，那么它只能指向整数类型的变量，而不能指向字符串、浮点数或其他类型的变量。

- 通过内置函数new创建指针

使用 `new` 函数创建的变量是零值，也就是说，它们的值为 0、false 或者空指针，具体取决于变量类型。

```go
package main

import "fmt"

// Steps3 内置函数 new 创建指针
func Steps3() {
	// 通过内置函数 new 创建一个 int 的指针类型
	a := new(int)
	var b *int
	fmt.Println("\tnew(int) value: ", a) // 有默认值
	fmt.Println("\t*int value: ", b) // nil
}

func main() {
	fmt.Println("Steps3():")
	Steps3()
}
```

- 空指针nil

在 Go 中，空指针非常有用，可以用于判断一个指针变量是否已经分配了内存空间。

```go
package main

import "fmt"

// Steps4 判断指针是否为nil
func Steps4() {
	var ptr *int
	if ptr == nil {
		fmt.Println("\tptr is nil")
	}

	var a int = 42
	var ptr2 *int = &a
	if ptr2 != nil {
		fmt.Println("\tptr2 is not nil")
	}
}

func main() {
	fmt.Println("Steps4():")
	Steps4()
}
```

在上面的代码中，我们首先定义了一个指针变量 `ptr`，并检查它是否为 `nil`。由于它没有指向任何变量，因此应该输出 "ptr is nil"。然后，我们定义了一个整数变量 `a`，并使用 `&a` 获取它的内存地址，并将其赋值给指针变量 `ptr2`。由于 `ptr2` 指向了变量 `a`，因此应该输出 "ptr2 is not nil"。

需要注意的是，指针变量可能会产生空指针异常。如果我们在访问指针变量所指向的变量之前没有检查指针变量是否为 `nil`，则可能会出现空指针异常，导致程序崩溃。因此，在使用指针变量之前，应该先检查它是否为 `nil`，以避免出现空指针异常。

## 占位符

`Golang`中常用占位符为`%d,%f,%s,%T,%+v`

```go
package main

import (
   "fmt"
)

// 占位符
func main() {
   var a byte = 255            // byte = uint8 rune = int32
   fmt.Printf("%v:%T\n", a, a) // 255:uint8

   var b int = 380
   // 不足位数前面补0
   fmt.Printf("%05d:%T\n", b, b)  // 00380:int
   fmt.Printf("%010d:%T\n", b, b) // 0000000380:int

   var c int = 88
   // 十进制 -> 二进制
   fmt.Printf("%b:%T\n", c, c) // 1011000:int
   // 十进制 -> 十六进制
   fmt.Printf("%x:%T\n", c, c) // 58:int

   var d string = "Golang"
   fmt.Printf("%s:%T\n", d, d) // Golang:string

   var e float64 = 3.14
   fmt.Printf("%f:%T\n", e, e) // 3.140000:float64
}

/*
[常用]
   %d    十进制表示
   %s    字符串或切片的无解译字节
   %f    有小数点而无指数，例如 123.456

   %v    相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
   %#v    相应值的 Go 语法表示
   %T    相应值的类型的 Go 语法表示
   %%    字面上的百分号，并非值的占位符

[布尔]

　　%t    单词 true 或 false。

[整数]

　　%b    二进制表示
　　%c    相应 Unicode 码点所表示的字符
　　%d    十进制表示
　　%o    八进制表示
　　%q    单引号围绕的字符字面值，由 Go 语法安全地转义
　　%x    十六进制表示，字母形式为小写 a-f
　　%X    十六进制表示，字母形式为大写 A-F
　　%U    Unicode 格式：U+1234，等同于 "U+%04X"

[浮点数及其复合构成]

　　%b    无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78
　　%e    科学计数法，例如 -1234.456e+78
　　%E    科学计数法，例如 -1234.456E+78
　　%f    有小数点而无指数，例如 123.456
　　%g    根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的 0）输出
　　%G    根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的 0）输出

[字符串与字节切片]

　　%s    字符串或切片的无解译字节
　　%q    双引号围绕的字符串，由 Go 语法安全地转义
　　%x    十六进制，小写字母，每字节两个字符
　　%X    十六进制，大写字母，每字节两个字符

[指针]

　　%p    十六进制表示，前缀 0x

[其它标记]

　　+    总打印数值的正负号；对于 %q（%+q）保证只输出 ASCII 编码的字符。
　　-    在右侧而非左侧填充空格（左对齐该区域）
　　#    备用格式：为八进制添加前导 0（%#o），为十六进制添加前导 0x（%#x）或
　　0X（%#X），为 %p（%#p）去掉前导 0x；如果可能的话，%q（%#q）会打印原始（即反引号围绕的）字符串；如果是可打印字符，%U（%#U）会写出该字符的 Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
　　' '    （空格）为数值中省略的正负号留出空白（% d）；
                以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开：fmt.Printf("% x\n", "Hello")
// 参考 http://www.manongjc.com/detail/25-pmahqixdhaombky.html
*/
```

## 运算符

Go支持五个基本二元算术运算符：+、-、*、/、%
Go支持六种位运算符：&、|、^、&^、<<、>>

## 思考题
1. 定义一个值为 1024 的`int`变量`a`, 再定义一个值为 0.1 的`float64`的变量`b`,将这两个变量加减乘除并打印结果。

## 参考
https://gfw.go101.org/article/operators.html

https://gfw.go101.org/article/type-system-overview.html

https://github.com/jincheng9/go-tutorial/blob/main/workspace/lesson2/readme.md

https://github.com/jincheng9/go-tutorial/blob/main/workspace/lesson9/readme.md