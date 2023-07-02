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
	// 类型 [n]T 表示拥有 n 个 T 类型值的数组
	// 类型 [3]int 表示拥有 3 个 int 类型值的数组, 默认值为 0
	var arrayInt [3]int // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	arrayInt[0] = 1     // 通过 [] 获取对应索引数据并修改
	arrayInt[1] = 2
	fmt.Printf("\tarrayInt: %+v\n", arrayInt)

	// 定义并初始化数组
	arrayBool := [3]bool{false, true}
	fmt.Printf("\tarrayBool: %+v\n", arrayBool)

	arrayFloat32 := [3]float32{1.0, 2.0} // float64
	fmt.Printf("\tarrayFloat32: %+v\n", arrayFloat32)

	arrayString := [3]string{"Golang", "Tutorial"}
	fmt.Printf("\tarrayString: %+v\n", arrayString)

	// 数组结构体
	arrayStruct := [3]dome{{a: 1, b: 2.0}, {a: 11, b: 22.0}}
	fmt.Printf("\tarrayStruct: %+v\n", arrayStruct)

	fmt.Printf("\t------------------------------\n")

	// 数组可以直接通过下标访问 T[x]
	fmt.Printf("\tarrayInt[0]: %d\n", arrayInt[0])

	// 数组可以直接通过下标修改 T[x] = y
	arrayInt[0] = 11
	fmt.Printf("\tarrayInt[0]: %d\n", arrayInt[0])

	fmt.Printf("\t------------------------------\n")

	// 不同大小的数组被认为是不同的类,不能直接赋值
	var arrayInt1 [5]int
	var arrayInt2 = [4]int{1, 2, 3}
	//arrayInt1 = arrayInt2 // panic: cannot use arrayInt2 (variable of type [3]int) as [4]int value in assignment
	fmt.Printf("\tarrayInt1: %+v\n", arrayInt1)
	fmt.Printf("\tarrayInt2: %+v\n", arrayInt2)

	fmt.Printf("\t------------------------------\n")
	// 数组遍历方式一
	for i := 0; i < len(arrayInt); i++ {
		fmt.Printf("\tarrayInt[%d]:%d\n", i, arrayInt[i])
	}

	// 数组地址
	fmt.Printf("\tarrayInt: %p\n", arrayInt) // arrayInt: %!p([3]int=[11 2 0]),arrayInt存储的不是地址值
	fmt.Printf("\t&arrayInt: %p\n", &arrayInt)
	// 数组遍历方式二
	for i, v := range arrayInt { // 数组的地址等于数组第一个元素的地址
		fmt.Printf("\t&arrayInt[%d]:%p value:%d\n", i, &arrayInt[i], v)
	}

	fmt.Printf("\tarrayInt len: %d\n", len(arrayInt))
	fmt.Printf("\tarrayInt cap: %d\n", cap(arrayInt))
}

// Steps2 二维数组
func Steps2() {
	arrayArrayString := [5][10]string{} // 初始化一个 5x10 的二维数组
	for i := 0; i < len(arrayArrayString); i++ {
		for ii := 0; ii < len(arrayArrayString[i]); ii++ {
			arrayArrayString[i][ii] = "-"
		}
	}

	// 遍历方式一
	for i := 0; i < len(arrayArrayString); i++ {
		fmt.Printf("\t")
		for ii := 0; ii < len(arrayArrayString[i]); ii++ {
			fmt.Printf(arrayArrayString[i][ii])
		}
		fmt.Println()
	}
	fmt.Printf("\t*************\n")
	// 遍历方式二
	for _, v := range arrayArrayString {
		fmt.Printf("\t")
		for _, vv := range v {
			fmt.Printf(vv)
		}
		fmt.Println()
	}
}

// Steps3 数组的内存占用
func Steps3() {
	arrayBool := [3]bool{true} // 3*1 byte
	fmt.Printf("\t   arrayBool size: %+v\n", unsafe.Sizeof(arrayBool))

	arrayFloat32 := [3]float32{1.0, 2.0} // 3*4 byte
	fmt.Printf("\tarrayFloat32 size: %+v\n", unsafe.Sizeof(arrayFloat32))

	arrayString := [3]string{"Golang", "Tutorial", "Hello"} // 3*16 byte
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
