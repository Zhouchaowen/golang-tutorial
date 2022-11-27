package main

import "fmt"

func createPerson() {
	p := Person{}
	p.Sports()
	fmt.Println("------------------")
	p.Eat()
}

func createDog() {
	d := Dog{}
	d.Run()
}

func main() {
	//createPerson()
	createDog()
}
