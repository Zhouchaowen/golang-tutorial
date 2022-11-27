package main

import "fmt"

type Body struct {
	Heart   string // 心脏
	Lung    string // 肺
	Stomach string // 胃
}

func (b Body) HeartWork() {
	fmt.Println("heart is working")
}

func (b Body) LungWork() {
	fmt.Println("lung is working")
}

func (b Body) StomachWork() {
	fmt.Println("stomach is working")
}
