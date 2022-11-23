# Interface

## 目录

- ch_1 interface 基本用法
- ch_2 通 interface 定义方法约定

## 定义接口

golang中interface是非常重要的一个特性，接口定义了一组方法签名的集合

```go
package main

import "fmt"

// Print 接口类型 是由一组方法签名定义的集合
// 定义接口约定
type Print interface {
	print(name string)
	// ....
}

type Float float64

func (f Float) print(name string) {
	fmt.Printf("%s, score:%f", name, f)
}

func main() {
	f := Float(80.5)
	f.print("Golang")
}
```

## 接口断言

golang中interface不仅可以定义一组方法签名，还可以当作Object用

```go
package main

import "fmt"

// 类型断言
// 断言 interface
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// 类型断言, 断言失败一般会导致panic的发生, 所以为了防止panic的发生, 我们需要在断言前进行一定的判断。
	// 如果断言失败, 那么ok的值将会是false
	// 如果断言成功, 那么ok的值将会是true, 同时s将会得到正确类型的值。
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)
}
```

## 思考题



## 参考

