package main

import (
	"fmt"
	"unsafe"
)

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

// Steps1 定义数组, 数组必须指定大小
func Steps1() {
	// 类型 [n]T 表示拥有 n 个 T 类型的值的数组
	// 类型 [3]int 表示拥有 3 个 int 类型的值的数组, 默认值为0
	var arrayInt [3]int // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	arrayInt[0] = 1
	arrayInt[1] = 2
	fmt.Printf("\tarrayInt: %+v\n", arrayInt)

	arrayBool := [3]bool{false, true}
	fmt.Printf("\tarrayBool: %+v\n", arrayBool)

	arrayFloat32 := [3]float32{1.0, 2.0} // float64
	fmt.Printf("\tarrayFloat32: %+v\n", arrayFloat32)

	arrayString := [3]string{"Golang", "Tutorial"}
	fmt.Printf("\tarrayString: %+v\n", arrayString)

	arrayStruct := [3]dome{{a: 1, b: 2.0}, {a: 11, b: 22.0}}
	fmt.Printf("\tarrayStruct: %+v\n", arrayStruct)

	// 数组可以直接通过下标访问 T[x]
	fmt.Printf("\tarrayInt[0]: %d\n", arrayInt[0])

	// 数组可以直接通过下标修改 T[x] = y
	arrayInt[0] = 11
	fmt.Printf("\tarrayInt[0]: %d\n", arrayInt[0])

	// 数组地址
	fmt.Printf("\tarrayInt: %p\n", arrayInt) // arrayInt: %!p([3]int=[11 2 0])，arrayInt没有地址
	fmt.Printf("\t&arrayInt: %p\n", &arrayInt)
	for i, v := range arrayInt {
		fmt.Printf("\t&arrayInt[%d]:%p value:%d\n", i, &arrayInt[i], v)
	}

	fmt.Printf("\tarrayInt len: %d\n", len(arrayInt))
	fmt.Printf("\tarrayInt cap: %d\n", cap(arrayInt))
}

// Steps2 二维数组
func Steps2() {
	arrayArrayString := [5][10]string{}
	for i := 0; i < len(arrayArrayString); i++ {
		for ii := 0; ii < len(arrayArrayString[i]); ii++ {
			arrayArrayString[i][ii] = "-"
		}
	}

	for i := 0; i < len(arrayArrayString); i++ {
		fmt.Printf("\t")
		for ii := 0; ii < len(arrayArrayString[i]); ii++ {
			fmt.Printf(arrayArrayString[i][ii])
		}
		fmt.Println()
	}
}

// Steps3 数组的内存占用
func Steps3() {
	arrayBool := [3]bool{true}
	fmt.Printf("\t   arrayBool size: %+v\n", unsafe.Sizeof(arrayBool))

	arrayFloat32 := [3]float32{1.0, 2.0}
	fmt.Printf("\tarrayFloat32 size: %+v\n", unsafe.Sizeof(arrayFloat32))

	arrayString := [3]string{"Golang", "Tutorial", "Hello"}
	fmt.Printf("\t arrayString size: %+v\n", unsafe.Sizeof(arrayString))
}

// Steps4 数组的内存占用
func Steps4() {
	arrayArrayBool := [3][1]bool{}
	fmt.Printf("\tarrayArrayBool[0]:%p\n", &arrayArrayBool[0])
	fmt.Printf("\tarrayArrayBool[1]:%p\n", &arrayArrayBool[1])
	fmt.Printf("\tarrayArrayBool[2]:%p\n", &arrayArrayBool[2])
	fmt.Printf("\t------------------\n")

	arrayArrayBool2 := [3][3]bool{}
	fmt.Printf("\tarrayArrayBool2[0]:%p\n", &arrayArrayBool2[0])
	fmt.Printf("\tarrayArrayBool2[1]:%p\n", &arrayArrayBool2[1])
	fmt.Printf("\tarrayArrayBool2[2]:%p\n", &arrayArrayBool2[2])
	fmt.Printf("\t------------------\n")

	arrayArrayInt := [3][1]int{}
	fmt.Printf("\tarrayArrayInt[0]:%p\n", &arrayArrayInt[0])
	fmt.Printf("\tarrayArrayInt[1]:%p\n", &arrayArrayInt[1])
	fmt.Printf("\tarrayArrayInt[2]:%p\n", &arrayArrayInt[2])
	fmt.Printf("\t------------------\n")

	arrayArrayInt2 := [3][2]int{}
	fmt.Printf("\tarrayArrayInt2[0]:%p\n", &arrayArrayInt2[0])
	fmt.Printf("\tarrayArrayInt2[1]:%p\n", &arrayArrayInt2[1])
	fmt.Printf("\tarrayArrayInt2[2]:%p\n", &arrayArrayInt2[2])
	fmt.Printf("\t------------------\n")

	arrayArrayString := [3][1]string{}
	fmt.Printf("\tarrayArrayString[0]:%p\n", &arrayArrayString[0])
	fmt.Printf("\tarrayArrayString[1]:%p\n", &arrayArrayString[1])
	fmt.Printf("\tarrayArrayString[2]:%p\n", &arrayArrayString[2])
}

func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
}
