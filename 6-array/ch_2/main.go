package main

import (
	"fmt"
	"unsafe"
)

// Steps1 定义切片
func Steps1() {
	// Steps 1-1: 类型 []T 表示一个元素类型为 T 的切片
	// 切片拥有长度和容量, 切片在添加数据时会自动扩容, 可以通过len(),cap()获取切片长度和容量
	var arrayInt []int // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr

	// Steps 1-2: append 向切片中添加元素（可能会导致内存重新分配）
	for i := 0; i < 10; i++ {
		arrayInt = append(arrayInt, i)
	}
	fmt.Printf("\tarrayInt:%+v len:%d cap:%d\n",
		arrayInt,
		len(arrayInt),
		cap(arrayInt))

	// Steps 1-3: 获取切片长度
	fmt.Println("\tarrayInt len:", len(arrayInt))

	// Steps 1-4: 获取切片的容量
	fmt.Println("\tarrayInt cap:", cap(arrayInt))

	// Steps 1-5: nil 切片的长度和容量为 0 且没有底层数组
	var arrayBool []bool
	fmt.Printf("\tarrayBool:%+v len:%d cap:%d\n",
		arrayBool,
		len(arrayBool),
		cap(arrayBool))
}

// Steps2 定义并初始化切片
func Steps2() {
	// Steps 2-1: 初始化切片
	arrayString := []string{"Golang", "Tutorial"}
	fmt.Printf("\tarrayString:%+v len:%d cap:%d\n",
		arrayString,
		len(arrayString),
		cap(arrayString))
}

// Steps3 通过 make 创建切片
func Steps3() {
	// Steps 3-1: 用内建函数 make 来创建切片
	// make([]T,len,cap)
	arrayFloat32 := make([]float32, 5)
	fmt.Printf("\tarrayFloat32:%+v len:%d cap:%d\n",
		arrayFloat32,
		len(arrayFloat32),
		cap(arrayFloat32))
	arrayFloat64 := make([]float64, 5, 10)
	fmt.Printf("\tarrayFloat64:%+v len:%d cap:%d\n",
		arrayFloat64,
		len(arrayFloat64),
		cap(arrayFloat64))
}

// Steps4 二维切片
func Steps4() {
	// Steps 4-1: 定义二维切片，并赋值
	arrayStringString := [][]string{
		[]string{"0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0"},
	}
	fmt.Printf("\tarrayStringString:%+v len:%d cap:%d\n",
		arrayStringString,
		len(arrayStringString),
		cap(arrayStringString))
	// Steps 4-3: 添加一行
	arrayStringString = append(arrayStringString, []string{"1", "1", "1", "1", "1"})
	fmt.Printf("\tarrayStringString:%+v len:%d cap:%d\n",
		arrayStringString,
		len(arrayStringString),
		cap(arrayStringString))

	// Steps 4-3: 打印二维数组
	for i := 0; i < len(arrayStringString); i++ { // len(arrayStringString) y轴数组长度
		fmt.Print("\t")
		for j := 0; j < len(arrayStringString[i]); j++ { // len(arrayStringString[i]) 第i行 x轴数组长度
			fmt.Printf("%s ", arrayStringString[i][j])
		}
		fmt.Println()
	}
}

// Steps5 切片上截取切片
func Steps5() {
	// Steps 5-1: 定义切片并初始化
	arrayInt := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\tarrayInt:%+v len:%d cap:%d\n",
		arrayInt,
		len(arrayInt),
		cap(arrayInt))

	// Steps 5-2: 可以用 array[low : high] or slice[low : high] 来截取数组或切片的一个片段长度为 high-low
	// 注意: arrayInt[0:3] 等同于 arrayInt[:3]
	interceptionArrayInt := arrayInt[1:3] // 获取 arrayInt 下标 1-2 的元素:[1,2,3] 长度为2
	fmt.Printf("\tinterceptionArrayInt:%+v len:%d cap:%d\n",
		interceptionArrayInt,
		len(interceptionArrayInt),
		cap(interceptionArrayInt))

	// Steps 5-3: 可以用 array[low : high: cap] 来截取切片或数组的一个片段长度为 high-low,容量为cap
	interceptionArrayIntCap := arrayInt[1:3:5] // 获取 arrayInt 下标 1-2 的元素:[1,2,3] 长度为2, 容量为4
	fmt.Printf("\tinterceptionArrayIntCap:%+v len:%d cap:%d\n",
		interceptionArrayIntCap,
		len(interceptionArrayIntCap),
		cap(interceptionArrayIntCap))

	// Steps 5-4: 切片并不存储任何数据，它只是描述了底层数组中的一段
	// 更改切片的元素会修改其底层数组中对应的元素,与它共享底层数组的切片都会观测到这些修改

	// interceptionArrayIntCap[2] 超出当前len, 打印报错 panic: runtime error: index out of range [2] with length 2
	//fmt.Printf("interceptionArrayIntCap[2]:%d",interceptionArrayIntCap[2])

	// 通过指针偏移强行获取底层元素（这种方式时不安全的）
	fmt.Printf("\tinterceptionSliceCap[2]:%d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&interceptionArrayIntCap[0])) + uintptr(16))))

	// Steps 5-6: 修改interceptionSliceCap[2]的值为33,底层切片arrayInt对应[3]位置改变33
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&interceptionArrayIntCap[0])) + uintptr(16))) = 33
	fmt.Printf("\tarrayInt[3]:%d\n", arrayInt[3])

	interceptionArrayIntCap[0] = 11
	fmt.Printf("\tarrayInt[1]:%d\n", arrayInt[1])
}

// 每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角
func main() {
	fmt.Println("Steps1():")
	Steps1()
	fmt.Println("Steps2():")
	Steps2()
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
	fmt.Println("Steps5():")
	Steps5()
}

// 参考:
//   https://tour.go-zh.org/moretypes/7
//   https://blog.go-zh.org/go-slices-usage-and-internals

// 小实验:
//    1.定义一个切片并初始化11以内的偶数,然后打印这些数字和切片的长度和容量
//    2.
