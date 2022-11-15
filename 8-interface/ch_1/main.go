package main

import "fmt"

// Print 接口类型 是由一组方法签名定义的集合
// 定义接口约定
type Print interface {
	print(name string)
	// ....
}

type Float float64

func (f Float) print(name string) {
	fmt.Printf("%s, score:%f", name, f)
}

func main() {
	f := Float(80.5)
	f.print("Golang")
}
