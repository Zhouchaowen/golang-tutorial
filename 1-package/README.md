# 包

## 目录

- 入门
- 导入标准库
- 导出函数与属性


## 入门 
每个 `Golang` 文件都必须以 `package namexxx` 语句开头, 如下代码中`package main`表示当前代码隶属于`main`包下。

`main`包下定义`func main()` 函数可以独立运行，并且所有`Golang`函数的启动入口都是`main`包下的`func main() `函数。

`import()`表示要导入的标准库或第三方包, 在实际开发者会引用许多标准库和第三方包来加简化业务开发。

`func main()`函数是程序入口。

```go
package main // 定义了包名,只有定义了main的包才能独立运行

// 导入: 标准库和第三方库
import (
	"fmt" // 引入一个标准库包
)

// main 函数,程序的入口
func main() {
	// 调用标准库 fmt 在控制台打印 hello world 字符串
	fmt.Println("hello world")
}

// 通过命令行运行
// go run main.go
// go build main.go && ./main
```

## 导入标准库

通过`import`到入了`math/rand`标准库,在`func main()`函数中调用标准库的函数`rand.Intn(n)`获取一个`[0,n)`的伪随机`int`值.

```go
package main

// 导入: 标准库和第三方库
import (
	"fmt"
	"math/rand" // 导入rand库
)

// main 函数,程序的入口。
func main() {
	// 调用标准库 fmt 在控制台打印 hello world 字符串
	// rand.Intn(10) 函数返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
	fmt.Println("hello world", rand.Intn(10))
}
```

## 导出函数与属性

在`golang`语言中想调用其它包的函数或变量**需要被调用函数或变量是导出的**, 导出一个函数或变量非常简单, **只需首字母大写就代表该函数或变量导出**.

```go
package main

// 导入: 标准库和第三方库
import (
	"fmt"
	"math"
)

// main 函数,程序的入口。
func main() {
	// 注意：在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问

	// 什么是导出？首字母大写代表导出,小写代表不导出
	//fmt.Println("intSize",math.intSize) // 引用未导出变量将报错
	fmt.Println("MaxInt", math.MaxInt)
}
```

`math`包如下(省略一些信息和注释方便查看)

```go
package math

.........

// Integer limit values.
const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64 // 未导出

	MaxInt    = 1<<(intSize-1) - 1 // 导出
	MinInt    = -1 << (intSize - 1)
  
  .........
)
```



## 思考题

1. 导入`math`包, 通过`Sqrt`函数求 9 的平方根并打印

## 参考