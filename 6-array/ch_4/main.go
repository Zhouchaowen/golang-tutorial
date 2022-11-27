package main

import (
	"fmt"
)

// 数组参数和切片参数的区别

func modifySlice0(slice []int) {
	fmt.Printf("\tslice value addr: %p\n", slice)
	fmt.Printf("\tslice addr: %p\n", &slice)
	slice[0] = 1000
}

// 切片作为函数参数时传递的是指针类型的全拷贝(array的uintptr指针，len，cap)
func Steps3() {
	var sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	fmt.Printf("\tsliceInt value addr: %p\n", sliceInt)
	fmt.Printf("\tsliceInt addr: %p\n", &sliceInt)
	modifySlice0(sliceInt)

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
}

func modifyArr0(arr [10]int) {
	fmt.Printf("\tarr value addr: %p\n", &arr[0])
	fmt.Printf("\tarr addr: %p\n", &arr)
	arr[0] = 1000
}

// 数组作为函数参数时传递的是值类型的全拷贝([10]int的全部数据)
func Steps4() {
	var arrInt = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("\tarrInt:%+v len:%d cap:%d\n",
		arrInt,
		len(arrInt),
		cap(arrInt))

	fmt.Printf("\tarrInt value addr: %p\n", &arrInt[0])
	fmt.Printf("\tarrInt addr: %p\n", &arrInt)
	modifyArr0(arrInt) // 数组地址很特别,数组地址就等于第一个元素地址

	fmt.Printf("\tarrInt:%+v len:%d cap:%d\n",
		arrInt,
		len(arrInt),
		cap(arrInt))
}

func main() {
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
}
