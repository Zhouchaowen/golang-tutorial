# 包

Package是golang中用来划分模块的，编写代码时我们通常将相同模块的代码划分到一个package中。我们首先来看第一个示例。

## 目录

- 入门
- 导入标准库
- 导出函数与属性


## 入门 
在`golang`中每个 `go` 文件都必须以 `package xxxname` 语句开头, 如代码中`package main`表示当前代码隶属于`main`包。`main`包是一个特殊的包，我们的启动函数`func main()`必须要定义在`package main`下才能执行。

`import()`表示要导入标准库(也可以叫做导入一个包)或第三方包, 在实际开发者会引用许多标准库和第三方包来简化业务开发。这里我们导入一个`fmt`的标准库，这个标准库定义了一些函数在控制台输出一些信息，如`Println()`函数。

`func main()`函数是程序入口。我们可以通过这个函数来运行我们编写的代码，我们在`func main()`函数中调用`fmt`标准库的`Println()`函数在控制台打印一个字串`hello world`。

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

上一节简单提了一下`import()`,这一节我们详细介绍`import()`。`import()`函数用来导入标准库和第三方库，只有通过`import()`导入的标准库和第三方库我们才能在当前`package`下调用并且只能调用改`package`下已经导出的属性和函数。

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
	intSize = 32 << (^uint(0) >> 63) // 32 or 64 // 未导出, 其余package不能调用

	MaxInt    = 1<<(intSize-1) - 1 // 导出, 其余package能调用
	MinInt    = -1 << (intSize - 1)
  
  .........
)
```



## 思考题

1. 导入`math`包, 通过`Sqrt`函数求 9 的平方根并打印

## 参考