package test

import "fmt"

func init() {
	fmt.Println("test.....")
	fmt.Println(name)
	fmt.Println(tep)
}

var tep int = 101

const name = "golang"

func Print() {
	fmt.Println("this is test package")
}
