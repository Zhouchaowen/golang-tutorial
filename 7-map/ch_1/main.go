package main

import "fmt"

/*
	1.定义map
	2.定义map并赋值
	3.通过key赋值
	4.通过make创建map
*/

// Steps1 定义映射
func Steps1() {
	// Steps 1-1: map[T]X 表示定义了一个 Key 类型为 T，Value 类型为 X 的映射
	// 注意key的T, key 的 T 必须支持 == 和 != 比较，才能用作 map 的 key
	// 定义一个 int->int 的map
	var mpIntInt map[int]int // 零值map
	fmt.Printf("\tmpIntInt:%+v len:%d\n",
		mpIntInt,
		len(mpIntInt)) // len 可以获取当前 map 存储的映射数量
	// mpIntInt[1] =1 // 定义后没有初始化的 map 为 nil, nil 映射不能添加键,添加报错 panic: assignment to entry in nil map

	// nil 映射可以读, 值为对应类型零值, ok 为 false
	v, ok := mpIntInt[1]
	fmt.Printf("\tmpIntInt[1]:%+v ok:%t\n", v, ok)

	// Steps 1-2: 定义一个 int->string 的map并初始化
	mpIntString := map[int]string{1: "Golang", 2: "Tutorial"}
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// Steps 1-3: 用内建函数 make 来创建map
	// make(map[T][T],size) 表示定义一个key为T类型value为T类型,容量为size的映射
	mpIntBool := make(map[int]bool)
	fmt.Printf("\tmpIntBool:%+v len:%d\n",
		mpIntBool,
		len(mpIntBool))
	mpIntBool[0] = true
	fmt.Printf("\tmpIntBool:%+v len:%d\n",
		mpIntBool,
		len(mpIntBool))

	// 创建了一个key为int类型value为float32类型,容量为10的映射
	mpIntFloat32 := make(map[int]float32, 10)
	fmt.Printf("\tmpIntFloat32:%+v len:%d\n",
		mpIntFloat32,
		len(mpIntFloat32))

	mpStringSliceInt := make(map[string][]int, 10)
	fmt.Printf("\tmpStringSliceInt:%+v len:%d\n",
		mpStringSliceInt,
		len(mpStringSliceInt))
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
}
