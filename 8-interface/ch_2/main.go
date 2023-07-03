package main

import "fmt"

/*
	1.通过接口定义方法
	2.实现接口定义方法
*/

// Duck 接口类型 定义一组方法签名的集合
type Duck interface {
	GaGaga() // 定义接口约定
	// ....
}

// 定义一个新类型(相当于给string起了个别名)
type DonaldDuck string

// DonaldDuck 实现了GaGaga()函数
func (d DonaldDuck) GaGaga() {
	fmt.Printf("%s, ga ga ga\n", d)
}

type RubberDuck string

// RubberDuck 实现了GaGaga()函数
func (d RubberDuck) GaGaga() {
	fmt.Printf("%s, ga ga ga\n", d)
}

type BlackSwan struct {
	Name  string
	Color string
}

// BlackSwan 实现了GaGaga()函数
func (d BlackSwan) GaGaga() {
	fmt.Printf("%s, ga ga ga\n", d.Name)
}

type Dog struct {
	Name string
}

// Dog 实现了GaGaga()函数
func (d Dog) GaGaga() {
	fmt.Printf("%s, ga ga ga\n", d.Name)
}

func (d Dog) WangWangWang() {
	fmt.Printf("%s, wang wang wang\n", d.Name)
}

func main() {
	var d Duck

	// 可以将DonaldDuck具体类型赋值给接口Duck类型，因为实现了接口类型的方法集合
	d = DonaldDuck("🦆 唐老鸭")
	d.GaGaga()

	d = RubberDuck("🦆 小黄鸭")
	d.GaGaga()

	d = BlackSwan{
		Name:  "黑天鹅",
		Color: "黑色",
	}
	d.GaGaga()

	// 接口与具体实现类调用时的对比
	d = Dog{
		Name: "小狗",
	}
	d.GaGaga()

	dog := Dog{
		Name: "哈士奇",
	}
	dog.GaGaga()
	dog.WangWangWang()
	fmt.Println(dog.Name)
}
