package main

import "fmt"

/*
	1.同包内引用结构体和结构体方法
	2.同包内引用导出属性和非导出属性
	3.同包内引用导出方法和非导出方法
*/

// 引用包内结构体的导出变量与导出方法
func main() {
	p := Person{
		Name: "golang",
		sex:  "man",
		age:  14,
	}
	p.Eat()
	p.Sleep()
	p.Work()
	p.Entertainment()
	p.Age()
	p.answerSex()

	fmt.Println("name", p.Name)
	fmt.Println("sex", p.sex)
	fmt.Println("age", p.age)
}
