package main

import (
	"fmt"
	"golang-tutorial/5-struct/person"
)

/*
	1.跨包内引用结构体和结构体方法
	2.跨包内引用导出属性和非导出属性
	3.跨包内引用导出方法和非导出方法
*/

func main() {
	p := person.Person{
		Name: "golang",
	}

	p.Eat()
	p.Sleep()
	p.Work()
	p.Entertainment()
	p.Age()
	//p.answerSex() //  无法引用结构体未导出方法

	fmt.Println(p.Name)
	//fmt.Println(p.sex) // 无法引用结构体未导出变量
	//fmt.Println(p.age)
}
