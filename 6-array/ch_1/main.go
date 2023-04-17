package main

import "fmt"

/*
	1.定义局部数组
	2.定义局部数组并赋值
	3.通过下标赋值
	4.定义全局数组
*/

type dome struct {
	a int
	b float32
}

var arrayUInt [3]uint

// 定义数组, 数组必须指定大小
func main() {
	// 类型 [n]T 表示拥有 n 个 T 类型的值的数组
	// 类型 [3]int 表示拥有 3 个 int 类型的值的数组, 默认值为0
	var arrayInt = [3]int{} // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	arrayInt[0] = 1
	arrayInt[1] = 2
	fmt.Printf("arrayInt: %+v\n", arrayInt)

	arrayBool := [3]bool{false, true}
	fmt.Printf("arrayBool: %+v\n", arrayBool)

	arrayFloat32 := [3]float32{1.0, 2.0} // float64
	fmt.Printf("arrayFloat32: %+v\n", arrayFloat32)

	arrayString := [3]string{"Golang", "Tutorial"}
	fmt.Printf("arrayString: %+v\n", arrayString)

	arrayStruct := [3]dome{{a: 1, b: 2.0}, {a: 11, b: 22.0}}
	fmt.Printf("arrayStruct: %+v\n", arrayStruct)

	// 数组可以直接通过下标访问 T[x]
	fmt.Printf("arrayInt[0]: %d\n", arrayInt[0])

	// 数组可以直接通过下标修改 T[x] = y
	arrayInt[0] = 11
	fmt.Printf("arrayInt[0]: %d\n", arrayInt[0])

	// 数组地址
	fmt.Printf("arrayInt: %p\n", &arrayInt)
	fmt.Printf("arrayInt[0]: %p\n", &arrayInt[0])

	fmt.Printf("arrayInt len: %d\n", len(arrayInt))
	fmt.Printf("arrayInt cap: %d\n", cap(arrayInt))
}
