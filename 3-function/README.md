# 函数

在 Go 语言中，函数是一种基本的代码块，用于执行某些操作并返回结果。函数用来划分不同功能， 我们可以将不同的功能封装成不同的函数，减少主干逻辑的复杂度。

## 目录

- 定义函数
- 函数的使用
- init() 函数调用

## 函数定义

`Golang`中通过关键字`func`定义一个函数, 并指定函数名称、参数列表和返回类型。 格式`func functionName(parameterName T,....) resultName T ...{}`

```go
package main

import "fmt"

// 定义一个 print 函数, 它接受 0 个参数，返回 0 个值
func print() {
	fmt.Println("hello world")
}

// 定义一个 add 函数, 它接受两个 int 类型的参数,返回 一个 int 类型的值 （参数的声明请看下一章）
func add(x int, y int) int {
	return x + y
}

// 定义一个 swap 函数, 它接受两个 int 类型的参数,返回 两个 int 类型的值
func swap(x, y int) (int, int) {
	return y, x
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量
func swap2(x, y int) (a int, b int) {
	a = y
	b = x
	return
}

// 传递不定长参数
func variableCut(x int, y ...int) int {
	for _, v := range y {
		x += v
	}
	return x
}

// 值类型参数
func modifyValue(x int) {
	x = x * 10
}

// 指针类型参数
func modifyPointer(x *int) {
	*x = (*x) * 10
}

func main() {
	print()
	fmt.Printf("add() return: %d\n", add(1, 2))
	x, y := swap(1, 2)
	fmt.Printf("swap() x: %d, y: %d \n", x, y)
	fmt.Printf("advariableCutd() return: %d\n", variableCut(1, 2, 3, 4, 5))

	// 值类型传递与指针类型传递的区别
	//		值传递是指在调用函数时将实际参数复制一份传递到函数中，
  //       这样在函数中如果对参数进行修改，将不会影响到实际参数
	//		指针传递(引用传递)是指在调用函数时将实际参数的地址传递到函数中，
  //       那么在函数中对参数所进行的修改，将影响到实际参数
	x = 1
	modifyValue(x)
	fmt.Printf("modifyValue() return: %d\n", x)

	x = 1
	modifyPointer(&x)
	fmt.Printf("modifyPointer() return: %d\n", x)
}
```

## Init函数

在 `Go` 语言中，`init` 函数是一个特殊的函数，用于在程序执行前自动执行一些初始化操作，例如初始化全局变量或加载配置文件等。`init` 函数没有任何参数和返回值，并且不能被显式调用。当程序启动时，Go 会自动在主函数执行之前调用所有包中的 `init` 函数，包括导入的所有包的 `init` 函数, `init()`函数执行完成后再执行`main()`函数

```go
package main

import (
	"fmt"
)

// init 函数会在 main 函数之前执行，而且无需调用就会执行
func init() {
	fmt.Println("Golang Tutorial")
}

func main() {
	fmt.Println("Hello World")
}
```

跨包的`init()`函数调用

```go
package main

import (
	"fmt"
	_ "golang-tutorial/3-function/test" // 引用 test 包, 会先执行 test 包的init函数
)

// init 函数会在 main 函数之前执行，而且无需调用就会执行
func init() {
	fmt.Println("Golang Tutorial")
}

func main() {
	fmt.Println("Hello World")
}
```

注意，Go 语言的初始化顺序是先导入所有包，然后按照包导入的顺序依次初始化每个包中的 `const`、`var` 和 `init`，最后调用 `main` 函数。

## 思考题

1. 通过传参实现加减乘除这四个函数,如: 加法函数
```go
func add(a,b int) int {
	return a+b
}
```

2. 通过指针的方式实现加减乘除这四个函数，要求：加减乘除不能通过返回参数的形式收集结果
3. 通过init函数给初始化全局变量str的值为"Hello Golang tutorial"

## 参考

