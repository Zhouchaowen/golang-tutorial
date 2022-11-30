# 变量与常量

## 目录

- ch_1 定义变量
- ch_2 变量赋值
- ch_3 类型转换
- ch_4 常量定义
- ch_5 定义函数变量
- ch_6 定义指针变量
- ch_7 占位符

## 定义变量

变量简单来说就是给内存中某一个地址起一个名字, 然后用这个地址存储某个特定类型的值。

在`golang`通过`var`关键字定义变量, 格式有多种最常用的两种为`var variableName T`和`variableName := Value`。

变量可以定义在函数外当做该包下的全局变量, 也可以定义在函数内当做该函数内的局部变量。

```go
package main

import "fmt"

/*
	通过 var 定义变量
	var variableName T
	var variableName T = Value
	var variableName = Value
	variableName := Value
*/

// var 语句可以声明全局变量
var aa int64

// 可以在 var 中定义多个全局变量
var (
	bb int8
	cc int16
	dd bool
	ee string
)

func main() {
	// var 语句用于声明一个变量列表,默认值为对应零值
	var a int     // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	var b float32 // float64
	var c bool
	var d string
	var e byte // 等同于 uint8
	var f rune // 等同于 int32,表示一个 Unicode 码点
	var g interface{}

	// 多变量声明
	var h, i string

	// 没有明确初始值的变量声明会被赋予它们的 零值
	// 零值是:
	//    数值类型为 0
	//    布尔类型为 false
	//    字符串为 ""（空字符串）

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

## 数据类型

**数据类型分类：**

- 布尔类型：`bool`。
- 整数类型：`int8`、`uint8`、`int16`、`uint16`、`int32`、`uint32`、`int64`、`uint64`、`int`、`uint`、 `uintptr`。
- 浮点数类型: `float32`、`float64`。
- 复数类型：`complex64`、`complex128`。
- 字符串类型：`string`。
- 指针持有者类型：`[size]T`、`[]T`、`map[T]T`、`struct`、`func`。

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
## 变量赋值

`Golang`中通过 `=` 对变量进行赋值, `=` 可以在变量初始化时赋值也可以在变量定义时赋值。还一种上文提到的简洁赋值 `:=`, `:=` 表示定义变量并赋值, 可以替代`Var`。

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

`Golang`中可以将相近类型的数据进行强转格式为`variableName2 := T(variableName1)`, 注意强转可能导致数据溢出或精度丢失.

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
	// 如下 float64(a) 将值 a 转换为类型 float64
	var b float64 = float64(a)

	// 简洁形式
	c := uint(b)

	fmt.Printf("a value:%d  a type:%s\n", a, reflect.TypeOf(a))
	fmt.Printf("b value:%f  b type:%s\n", b, reflect.TypeOf(b))
	fmt.Printf("c value:%d  c type:%s\n", c, reflect.TypeOf(c))

	fmt.Printf("*************************************************************\n")

	d := uint8(255)

	// 类型转换 不能超过转换类型的范围
	//e := uint8(256) // 编译错误, 常量256溢出了uint8

	// 超过转换类型的范围溢出
	var f int = 256
	g := uint8(f)
	h := uint8(f + 1)

	fmt.Printf("d value:%d  d type:%s\n", d, reflect.TypeOf(d))

	fmt.Printf("g value:%d  g type:%s\n", g, reflect.TypeOf(g))
	fmt.Printf("h value:%d  h type:%s\n", h, reflect.TypeOf(h))
}
```

## 定义常量

`Golang`中通过const定义常量, 格式为`const constantName = value`和`const constantName T = value`。

常量可以定义在函数外当做该包下的全局变量, 也可以定义在函数内当做该函数内的局部常量。

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

// 常量是在程序运行时，不会被修改的量
// 常量的声明与变量类似，使用 const 关键字, 常量中的数据类型只可以是字符、字符串、布尔值或数值

// const NameOfVariable [type] = value  type 可以省略然编译器推导

const PI = 3.14
const NAME = "Golang-tutorial"
const OK bool = true

// 可以在 const 中定义多个常量
const (
	MaxUint8  = math.MaxUint8
	MaxUint16 = math.MaxUint16
)

func main() {
	// 函数内也可以定义常量
	const World = "World"
	fmt.Println("Hello", World)

	fmt.Printf("MaxUint8 value:%d  MaxUint8 type:%s\n", MaxUint8, reflect.TypeOf(MaxUint8))
	fmt.Printf("MaxUint16 value:%d  MaxUint16 type:%s\n", MaxUint16, reflect.TypeOf(MaxUint16))
}
```

## 定义函数变量

`Golang`中函数也是可以定义为变量,并且可以当做参数传递

```go
package main

import (
	"fmt"
)

/*
	定义函数变量
	var variableName = func

*/

// 定义一个自定义类型的函数, 用 Handler 表示这个自定义类型
type Handler func(x, y int) int

func compute(x, y int, handler Handler) int {
	x = x * 10
	y = y * 10
	return handler(x, y)
}

// 函数也可以当做类型,可以像其它值一样传递
func main() {
	var add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	Multi := func(x, y int) int {
		return x * y
	}
	fmt.Println(compute(1, 2, Multi))
}

```

## 定义指针变量

`Golang`中也是通过`var`定义指针变量, 格式有多种选择常用的三种`var variableName *T`, `variableName := &Value`, `variableName := new(T)`

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
	// 定义一个 int 的指针类型
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

指针变量的赋值

```go
package main

import "fmt"

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

	// & 表示取 b 变量的地址并赋值给 a, 改动 a 就相当于改动 b
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

通过内置函数new创建指针变赋默认值

```go
package main

import "fmt"

// Steps3 内置函数 new 创建指针
func Steps3() {
	// 通过内置函数 new 创建一个 int 的指针类型
	a := new(int)
	var b *int
	fmt.Println("\tnew(int) value: ", a)
	fmt.Println("\t*int value: ", b)
}

func main() {
	fmt.Println("Steps3():")
	Steps3()
}
```

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
1. 定义一个值为 1024 的`int`变量`a`, 再定义一个值为 10.1 的`float64`的变量`b`,将这两个变量加减乘除并打印结果。

## 参考
https://gfw.go101.org/article/operators.html

https://gfw.go101.org/article/type-system-overview.html

https://github.com/jincheng9/go-tutorial/blob/main/workspace/lesson2/readme.md

https://github.com/jincheng9/go-tutorial/blob/main/workspace/lesson9/readme.md