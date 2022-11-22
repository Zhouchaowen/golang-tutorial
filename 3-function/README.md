# 函数



## 目录

- ch_1 函数定义
- ch_2 init() 函数的使用
- ch_3 跨包的 init() 函数调用
- ch_4 函数导出和跨包调用



## 函数定义

`Golang`中通过关键字`func`定义一个函数, 格式`func functionName(parameterName T,....) resultName T ...{}`

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

// 传递必定长参数
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

`Golang`中想在`main()`函数启动前初始化一些配置或调用需要通过`init()`函数。在运行程序时会先调用`init()`中的语句执行完成后再继续执行`main()`函数

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



## 思考题

1. 通过传参实现加减乘除这四个函数,如加函数
```golang
func add(a,b int) int {
	return a+b
}
```

## 参考

