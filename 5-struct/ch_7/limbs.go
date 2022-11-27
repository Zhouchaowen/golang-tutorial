package main

import "fmt"

type Limbs struct {
	Hands string
	Feet  string
}

func (l Limbs) HandsWork() {
	fmt.Println("hands is working")
}

func (l Limbs) FeetWork() {
	fmt.Println("feet is working")
}
