package main

import "fmt"

type Brain struct {
	Eyes  string
	Ears  string
	Nose  string
	Mouth string
}

func (b Brain) EyesWork() {
	fmt.Println("eyes is working")
}

func (b Brain) EarsWork() {
	fmt.Println("ears is working")
}

func (b Brain) NoseWork() {
	fmt.Println("nose is working")
}

func (b Brain) MouthWork() {
	fmt.Println("mouth is working")
}
