package main

import "fmt"

// ResponseStatus 自定义类型的方法
type ResponseStatus int

const (
	QuerySuccess ResponseStatus = iota
	QueryError
)

func (r ResponseStatus) ToCN() string {
	switch r {
	case 0:
		return "query success"
	case 1:
		return "query error"
	default:
		return "non"
	}
}

type Person struct {
	Name string
	sex  string
	age  int
}

func (p Person) Eat() {
	fmt.Println(p.Name, "eat")
}

func (p Person) Sleep() {
	fmt.Println(p.Name, "sleep")
}

func (p Person) Work() {
	fmt.Println(p.Name, "work")
}

func (p Person) Entertainment() {
	fmt.Println(p.Name, "entertainment")
}

func (p Person) Age() {
	fmt.Println(p.age, "age")
}

func main() {
	fmt.Println(QuerySuccess.ToCN())
	fmt.Println(QueryError.ToCN())

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

	fmt.Println("name", p.Name)
	fmt.Println("sex", p.sex)
	fmt.Println("age", p.age)
}
