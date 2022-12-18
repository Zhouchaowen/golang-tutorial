# Interface



## 目录

- 定义接口
- 接口断言
- 

## 定义接口

`golang`中`interface{}`是一个非常重要的特性，`interface{}`定义了一组方法签名的集合, 用于抽象同一类事物的行为。如下：`type Duck interface{}`定义了一个名字为`GaGaga()`的方法, `DonaldDuck`实现了`GaGaga()`方法, 可以说`DonaldDuck就是Duck`, 所有可以将`DonaldDuck`赋值给`Duck`;  `Dog`也实现了GaGaga()方法，所有也可以将`Dog`赋值给`Duck`。

```go
package main

import "fmt"

// Duck 接口类型 定义一组方法签名的集合
// 定义接口约定
type Duck interface {
	GaGaga()
	// ....
}

type DonaldDuck string

func (d DonaldDuck) GaGaga() {
	fmt.Printf("%s, ga ga ga\n", d)
}

type RubberDuck string

func (d RubberDuck) GaGaga() {
	fmt.Printf("%s, ga ga ga\n", d)
}

type Dog struct {
	Name string
	age  int
}

func (d Dog) GaGaga() {
	fmt.Printf("%s, ga ga ga\n", d.Name)
}

func main() {
	var d Duck

	d = DonaldDuck("🦆 唐老鸭")
	d.GaGaga()

	d = RubberDuck("🦆 小黄鸭")
	d.GaGaga()

	d = Dog{
		Name: "小狗",
		age:  5,
	}
	d.GaGaga()
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



## 为什么需要接口

- 接口允许 Go 具有多态性, 在需要多态性的 Go 中使用接口。
- 在可以传递多种类型的函数中，可以使用接口。

- 接口还用于帮助减少重复/样板代码。

在需要动态类型参数的函数和方法的情况下，接口非常有用，例如接受任何类型值的 Println 函数。



## 思考题



## 参考

https://blog.knoldus.com/how-to-use-interfaces-in-golang/

https://stackoverflow.com/questions/39092925/why-are-interfaces-needed-in-golang

https://stackoverflow.com/questions/23148812/whats-the-meaning-of-interface

https://blog.boot.dev/golang/golang-interfaces/