package main

import (
	"fmt"
)

/*
	1.切片传参
	2.数组传参
*/

// 数组参数和切片参数的区别

func modifySlice0(slice []int) {
	fmt.Printf("\t[modifySlice0] slice value    addr: %p\n", slice) // 改值等于 sliceInt 的值
	fmt.Printf("\t[modifySlice0] slice variable addr: %p\n", &slice)
	slice[0] = 1000
}

// Steps3 切片作为函数参数时传递的是指针类型的全拷贝(array的uintptr指针，len，cap)
func Steps3() {
	var sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("\t[Steps3] sliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	fmt.Printf("\t[Steps3] sliceInt value    addr: %p\n", sliceInt)
	fmt.Printf("\t[Steps3] sliceInt variable addr: %p\n", &sliceInt)
	modifySlice0(sliceInt)

	fmt.Printf("\t[Steps3] sliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
}

func modifyArr0(arr [10000000]int) {
	fmt.Printf("\t[modifyArr0] arr value    addr: %p\n", &arr[0])
	fmt.Printf("\t[modifyArr0] arr variable addr: %p\n", &arr)
	arr[0] = 1000
}

// Steps4 数组作为函数参数时传递的是值类型的全拷贝([10]int的全部数据)
func Steps4() {
	var arrInt = [10000000]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("\t[Steps4] len:%d cap:%d\n",
		len(arrInt),
		cap(arrInt))

	fmt.Printf("\t[Steps4] arrInt value    addr: %p\n", &arrInt[0])
	fmt.Printf("\t[Steps4] arrInt variable addr: %p\n", &arrInt)
	modifyArr0(arrInt) // 数组地址很特别,数组地址就等于第一个元素地址

	fmt.Printf("\t[Steps4] arrInt[0]:%+v len:%d cap:%d\n",
		arrInt[0],
		len(arrInt),
		cap(arrInt))
}

func main() {
	fmt.Println("Steps3():")
	Steps3()
	fmt.Println("Steps4():")
	Steps4()
}
