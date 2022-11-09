package main

import "fmt"

type dome struct {
	a int
	b float32
}

// 定义切片
// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。
func main() {
	// 类型 []T 表示一个元素类型为 T 的切片
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
}
