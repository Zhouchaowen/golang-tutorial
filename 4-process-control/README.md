# 流程控制



## 目录

- for/range 循环
- if 分支判断
- switch 分支选择
- defer 函数延时调用(栈)
- defer+recover() 函数捕获 panic 错误

## For循环

`Golang`中通过`For`关键字来定义一个循环并且只有`For`关键字(`Golang`中没有`while`关键字)

```go
package main

import "fmt"

// Steps1 通过 for 循环累加 0-9
func Steps1() {
	sum := 0
	// for 循环
	// i := 0 初始化语句：在第一次迭代前执行
	// i < 10 条件表达式：在每次迭代前求值
	// i++    后置语句：在每次迭代的结尾执行
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("\tsum: %d\n", sum)
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
}
```

通过`For`实现类似`while`的语义

```go
package main

import "fmt"

// Steps2 for循环初始化语句和后置语句不是必须的
func Steps2() {
	sum := 0
	// 初始化语句和后置语句是可选的
	for sum < 5 {
		sum++
	}
	fmt.Printf("\tsum: %d\n", sum)
}

func main() {
	fmt.Println("Steps2():")
	Steps2()
}
```

通过`Range`关键字来遍历字符串,数组,切片或映射

```go
package main

import "fmt"

// Steps3 range形式的循环遍历
func Steps3() {
	str := "Golang Tutorial"
	for i, v := range str { // 遍历字符串
		fmt.Printf("\ti:%d,v:%c\n", i, v)
	}
}

func main() {
	fmt.Println("Steps3():")
	Steps3()
}
```

`Range`和`for`遍历的区别

```go
package main

import "fmt"

// Steps4 range和for遍历的区别
func Steps4() {
	str := "Golang 教程"
	for i := 0; i < len(str); i++ {
		fmt.Printf("\ti:%d,v:%c\n", i, str[i])
	}

	for i, v := range str { // 遍历字符串
		fmt.Printf("\ti:%d,v:%c\n", i, v)
	}
}

func main() {
	fmt.Println("Steps4():")
	Steps4()
}
```

`For`循环中的`break`和`continue`

```go
package main

import "fmt"

// Steps5 for 循环中的 break 和 continue
func Steps5() {
	for i := 0; i < 10; i++ {
		if i == 5 { // 下一小节介绍
			fmt.Printf("\ti:%d, continue\n", i)
			continue
		}

		if i == 6 {
			fmt.Printf("\ti:%d, break\n", i)
			break
		}
	}
}

func main() {
	fmt.Println("Steps5():")
	Steps5()
}
```

`Goto`实现循环

```go
package main

import "fmt"

// Steps6 goto 实现循环
func Steps6() {
	i := 0

Next: // 跳转标签声明
	fmt.Printf("\ti:%d\n", i)
	i++
	if i < 5 {
		goto Next // 跳转
	}
}

func main() {
	fmt.Println("Steps6():")
	Steps6()
}
```

## If判断

`Golang`中`If`语句和其它语言语义相同

```go
package main

import "fmt"

// if 分支打印不同字符
func main() {
	flag := 10
	if flag > 5 { // 判断表达式
		fmt.Println("flag:", flag)
	}

	flag = 14
	//flag = 16
	//flag = 21

	if flag > 20 {
		fmt.Println("flag:", flag)
	} else if flag < 15 {
		fmt.Println("flag:", flag)
	} else {
		fmt.Println("flag:", flag)
	}
}
```

## Switch选择

`Golang`中可以通过`switch-case`来实现分支选择, 每一个`case`分支都是唯一的，从上往下逐一判断，直到匹配为止，如果某些`case`分支条件重复了，编译会报错。

每个`case`分支最后自带`break`效果，匹配成功就不会执行其它`case`; 如果所有分支都没有匹配成功并且又定义了`default`分支, 那最终会走`default`分支

```go
package main

import "fmt"

// Steps1 基础用法
func Steps1() {
	flag := 1
	//flag = 2
	//flag = 3
	//flag = 4
	//flag = 5

	switch flag { // flag 待判断条件
	case 1: // 条件 flag 是否等于 1。是：执行该case下的流程，否：选择其它满足条件的 case
		fmt.Println("\tcase:", flag)
		// Golang 中每个 case 后面不需要 break 语句。当然 return 是可选的
	case 2:
		fmt.Println("\tcase:", flag)
	case 3, 4: // case 可以设置多个条件。只要 flag 等于3或4都能执行当前case流程
		fmt.Println("\tcase:", flag)
	case 5:
		fmt.Println("\tcase:", flag)
		return
	default: // 当所有case都无法满足, 会执行 default 的流程。如果没有 default 那当前 switch 执行完成
		fmt.Println("\tdefault:", flag)
	}
}

// Steps2 switch 条件可以是任何支持判断的类型
func Steps2() {
	flag := "Hello"
	flag = "World"
	flag = "Golang"
	flag = "Tutorial"
	flag = "Process"

	switch flag { // flag 待判断条件
	case "Hello": // 条件 flag 是否等于 "Hello"。是：执行该case下的流程，否：选择其它满足条件的 case
		fmt.Println("\tcase:", flag)
	case "World":
		fmt.Println("\tcase:", flag)
	case "Golang", "Tutorial": // case 可以设置多个条件。只要 flag 等于"Golang"或"tutorial"都能执行当前case流程
		fmt.Println("\tcase:", flag)
	default: // 当所有case都无法满足, 会执行 default 的流程。如果没有 default 那当前 switch 执行完成
		fmt.Println("\tdefault:", flag)
	}
}

// Steps3 switch true 可以将一长串 if-then-else 写得更加清晰
func Steps3() {
	flag := 1
	//flag = 2
	//flag = 3
	//flag = 4
	//flag = 5
	//flag = 7

	switch { // flag 待判断条件
	case flag < 2: // 条件 flag 是否小于 2。是：执行该case下的流程，否：选择其它满足条件的 case
		fmt.Println("\tcase flag < 2 flag:", flag)
	case flag < 4:
		fmt.Println("\tcase flag < 4, flag:", flag)
	case flag > 6, flag < 10: // case 可以设置多个条件。flag 大于6或小于10都能执行当前case流程
		fmt.Println("\tcase flag > 6 || flag < 10 flag:", flag)
	case flag > 6 && flag < 10: // case 可以设置组合条件。flag 大于6并且小于10都才能执行当前case流程
		fmt.Println("\tcase flag > 6 || flag < 10 flag:", flag)
	}
}

// Steps4 for + switch 的使用
func Steps4() {
	for flag := 0; flag < 11; flag++ {
		switch { // flag 待判断条件
		case flag < 2: // 条件 flag 是否小于 2。是：执行该case下的流程，否：选择其它满足条件的 case
			fmt.Println("\tcase flag < 2 flag:", flag)
		case flag < 4:
			fmt.Println("\tcase flag < 4, flag:", flag)
		case flag > 6, flag < 8: // case 可以设置多个条件。flag 大于6或小于10都能执行当前case流程
			fmt.Println("\tcase flag > 6 || flag < 8 flag:", flag)
		case flag > 6 && flag < 10: // case 可以设置组合条件。flag 大于6并且小于10都才能执行当前case流程
			fmt.Println("\tcase flag > 6 && flag < 10 flag:", flag)
		}
	}
}

// switch 是编写一连串 if - else 语句的简便方法
func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
}
```

## Defer

` Golang`中通过`defer`来实现延时调用, 常用来做一些收尾工作: **关闭连接,清理资源**

```go
package main

import "fmt"

// defer作用:
//    释放占用的资源
//    捕捉处理异常 recover

// Steps1 defer 语句会将函数推迟到外层函数返回之后执行。
// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
func Steps1() {
	defer fmt.Printf(" world\n")

	fmt.Printf("\thello")
}

// Steps2 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
func Steps2() {
	fmt.Println("\tbegin")
	for i := 0; i < 3; i++ {
		defer fmt.Println("\t\ti:", i)
		fmt.Printf("\t\ti:%d\n", i)
	}
	fmt.Println("\tend")
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
}

/*
   ----- -----
  |   | |    |
  | | V | |  |
	|	|     |  V
		| ... |
		|  3  |
		|  2  |
		|  1  |
		|  0  |
		 —————
*/
```

## recover

通过`defer+recover`来拦截程序捕获`panic`

```go
package main

import "fmt"

// 捕捉处理异常 recover
func main() {
	defer func() {
		if err := recover(); err != nil {
			// 捕捉错误 run err: runtime error: integer divide by zero
			fmt.Println("run err:", err)
		}
	}()

	a := 10
	b := 0
	_ = a / b
	fmt.Println("return")
}
```

## 思考题

1. 计算 100000 以内偶数,并且不是 4 的倍数外的所有数值和
2. 定义函数`Calculation`通过`Switch`实现加减乘除

```go
// 参考
func Calculation(option byte, a float64,b float64) float64{
	switch option {
	case '-':
	.......	
	}
}
```

3. 通过`For`循环打印如下图形

```bigquery
*
**
***
****
*****
******
```
## 参考
https://gfw.go101.org/article/control-flows.html
