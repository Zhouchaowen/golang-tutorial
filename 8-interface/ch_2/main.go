package main

import "fmt"

/*
	1.通过接口定义方法
	2.实现接口定义方法
*/

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
