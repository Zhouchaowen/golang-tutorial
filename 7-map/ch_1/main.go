package main

import "fmt"

// Steps1 定义映射
func Steps1() {
	// Steps 1-1: map[T]X 表示定义了一个 Key 类型为 T，Value 类型为 X 的映射
	// 定义一个 int->int 的map
	var mpIntInt map[int]int // 零值map
	fmt.Printf("\tmpIntInt:%+v len:%d\n",
		mpIntInt,
		len(mpIntInt)) // len 可以获取当前 map 存储的映射数量
	// mpIntInt[1] =1 // nil 映射不能添加键,添加报错 panic: assignment to entry in nil map

	// Steps 1-2: 定义一个 int->string 的map并初始化
	mpIntString := map[int]string{1: "Golang", 2: "Tutorial"}
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// Steps 1-3: 用内建函数 make 来创建map
	mpIntBool := make(map[int]bool)
	fmt.Printf("\tmpIntBool:%+v len:%d\n",
		mpIntBool,
		len(mpIntBool))
	mpIntBool[0] = true
	fmt.Printf("\tmpIntBool:%+v len:%d\n",
		mpIntBool,
		len(mpIntBool))

	mpIntFloat32 := make(map[int]float32, 10)
	fmt.Printf("\tmpIntFloat32:%+v len:%d\n",
		mpIntFloat32,
		len(mpIntFloat32))
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
}
