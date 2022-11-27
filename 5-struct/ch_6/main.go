package main

import (
	"fmt"
	"golang-tutorial/5-struct/person"
)

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
