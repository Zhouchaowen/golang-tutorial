package person

import "fmt"

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

func (p Person) answerSex() {
	fmt.Println(p.sex, "sex")
}
