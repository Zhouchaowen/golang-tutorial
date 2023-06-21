package main

import (
	"fmt"
	"unsafe"
)

/*
	1.切片参数和指针切片参数的区别
*/

func appendValueSlice(slice []int) {
	fmt.Printf("\tslice value addr: %p\n", slice) // 参数 slice 存储的值等于 sliceInt
	fmt.Printf("\tslice       addr: %p\n", &slice)
	// slice 的addr是新的，对应的len,cap也是和sliceInt隔离的,但底层指向 array 的地址是一样的,
	// 当添加100后实际底层 array 上其实是有值的,只是len限制,我们无法在sliceInt中看到。
	slice = append(slice, 100)
	fmt.Printf("\tslice:%+v len:%d cap:%d\n",
		slice,
		len(slice),
		cap(slice))
	fmt.Printf("\tslice value addr: %p\n", slice)
	fmt.Printf("\tslice       addr: %p\n", &slice)
}

// Steps5 切片作为函数参数时传递的是指针类型的全拷贝(array的 uintptr 指针,len,cap)
func Steps5() {
	//var sliceInt = make([]int,2,2)
	var sliceInt = make([]int, 2, 10)

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	fmt.Printf("\tsliceInt value addr: %p\n", sliceInt)
	fmt.Printf("\tsliceInt       addr: %p\n", &sliceInt)

	fmt.Printf("\t------------------------------------\n")
	appendValueSlice(sliceInt)
	fmt.Printf("\t------------------------------------\n")

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	// 通过不安全的方式强行获取 sliceInt[2]
	fmt.Printf("\tsliceInt[2]:%d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sliceInt[0])) + uintptr(16))))
}

func appendPointerSlice(slice *[]int) {
	fmt.Printf("\tslice       value addr: %p\n", slice)  // slice 表示这个变量存储的值 == &sliceInt
	fmt.Printf("\tslice value addr value: %p\n", *slice) // *slice 表示这个变量存储的地址上的值 == sliceInt
	fmt.Printf("\tslice             addr: %p\n", &slice) // &slice 表示这个变量本身的地址
	// slice 的addr是新的，对应的len,cap也是和sliceInt隔离的，
	// 所以添加100后实际地址上其实是有值，只是通过len限制，我们无法看到。
	*slice = append(*slice, 100)
	fmt.Printf("\tslice:%+v len:%d cap:%d\n",
		*slice,
		len(*slice),
		cap(*slice))
	fmt.Printf("\tslice value addr: %p\n", slice)
	fmt.Printf("\tslice       addr: %p\n", &slice)
}

func Steps6() {
	//var sliceInt = make([]int,2,2)
	var sliceInt = make([]int, 2, 10)

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	fmt.Printf("\tsliceInt value[0] addr: %p\n", &sliceInt[0])
	fmt.Printf("\tsliceInt    value addr: %p\n", sliceInt)
	fmt.Printf("\tsliceInt          addr: %p\n", &sliceInt)

	fmt.Printf("\t------------------------------------\n")
	appendPointerSlice(&sliceInt)
	fmt.Printf("\t------------------------------------\n")

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
}

func main() {
	fmt.Println("Steps5():")
	Steps5()
	fmt.Println("Steps6():")
	Steps6()
}
