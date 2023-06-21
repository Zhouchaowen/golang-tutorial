package main

import (
	"fmt"
	"unsafe"
)

/*
	1.定义局部切片
	2.通过append向切片添加值
	3.定义局部切片并赋值
	4.通过 make 创建切片
	5.通过下标对切片赋值
	6.定义二维切片并遍历
	7.截取切片
	8.遍历切片
*/

// 切片也可以定义在全局
var sliceByte []byte

// Steps1 定义切片
func Steps1() {
	// Steps 1-1: 类型 []T 表示一个元素类型为 T 的切片
	// 切片拥有长度和容量, 切片在添加数据时会自动扩容, 可以通过len(),cap()获取切片长度和容量

	// 类型为 int 的切片, 初始化后长度容量都为 0, 不指向任何底层数组
	var sliceInt []int // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	fmt.Printf("\t&sliceInt:%p sliceInt:%p sliceInt:%+v len:%d cap:%d\n",
		&sliceInt,
		sliceInt,
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
	//_ = sliceInt[0] // 在未初始化长度前直接通过下标读取或赋值数据将会报错, 只能通过 append 添加元素

	// Steps 1-2: append 向切片中添加元素（可能会导致内存重新分配）
	for i := 1; i < 11; i++ {
		sliceInt = append(sliceInt, i)
		fmt.Printf("\t&sliceInt:%p sliceInt:%p sliceInt:%+v len:%d cap:%d\n",
			&sliceInt,
			sliceInt,
			sliceInt,
			len(sliceInt),
			cap(sliceInt))
	}

	// Steps 1-3: 获取切片长度
	fmt.Println("\tsliceInt len:", len(sliceInt))

	// Steps 1-4: 获取切片的容量
	fmt.Println("\tsliceInt cap:", cap(sliceInt))

	// Steps 1-5: 类型为 bool 的切片, 初始化后长度和容量为 0 且没有底层数组
	var sliceBool []bool
	fmt.Printf("\tsliceBool:%+v len:%d cap:%d\n",
		sliceBool,
		len(sliceBool),
		cap(sliceBool))
}

// Steps2 定义并初始化切片
func Steps2() {
	// Steps 2-1: 定义并初始化切片
	sliceString := []string{"Golang", "Tutorial"}
	fmt.Printf("\tsliceString:%+v len:%d cap:%d\n",
		sliceString,
		len(sliceString),
		cap(sliceString))

	// 数组地址
	fmt.Printf("\t         &sliceString addr: %p\n", &sliceString)
	fmt.Printf("\t    sliceString value addr: %p\n", sliceString)
	fmt.Printf("\t&sliceString[0] value addr: %p\n", &sliceString[0])
	fmt.Printf("\t&sliceString[1] value addr: %p\n", &sliceString[1])

	// Steps 2-2: 初始化切片
	sliceInt := []int{1, 2, 3} // uint8,int8,uint16,int16,uint32,int32,uint64,int64,uintptr
	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	// 数组地址
	fmt.Printf("\t         &sliceInt addr: %p\n", &sliceInt)
	fmt.Printf("\t    sliceInt value addr: %p\n", sliceInt)
	fmt.Printf("\t&sliceInt[0] value addr: %p\n", &sliceInt[0])
	fmt.Printf("\t&sliceInt[1] value addr: %p\n", &sliceInt[1])
	fmt.Printf("\t&sliceInt[2] value addr: %p\n", &sliceInt[2])
}

// Steps3 通过 make 创建切片
func Steps3() {
	// Steps 3-1: 用内建函数 make 来创建切片
	// make([]T,len,cap) 如下：创建一个 float32 类型, 长度为 5 的数组
	// 和 var sliceFloat32 []float32 的区别是 make 创建的切片会分配底层数组并赋零值
	sliceFloat32 := make([]float32, 5)
	fmt.Printf("\t&sliceFloat32:%p sliceFloat32:%p sliceFloat32:%+v len:%d cap:%d\n",
		&sliceFloat32,
		sliceFloat32,
		sliceFloat32,
		len(sliceFloat32),
		cap(sliceFloat32))
	for i := 0; i < len(sliceFloat32); i++ {
		sliceFloat32[i] = float32(i)
	}

	fmt.Printf("\t&sliceFloat32:%p sliceFloat32:%p sliceFloat32:%+v len:%d cap:%d\n",
		&sliceFloat32,
		sliceFloat32,
		sliceFloat32,
		len(sliceFloat32),
		cap(sliceFloat32))

	// 创建一个 float64 类型, 长度为 5, 容量为 10 的数组
	sliceFloat64 := make([]float64, 5, 10)
	fmt.Printf("\t&sliceFloat64:%p sliceFloat64:%p sliceFloat64:%+v len:%d cap:%d\n",
		&sliceFloat64,
		sliceFloat64,
		sliceFloat64,
		len(sliceFloat64),
		cap(sliceFloat64))
	//for i := 0 ;i < cap(sliceFloat64);i++ { // cap-len的部分并没有分配，不能直接赋值
	//	sliceFloat64[i] = float64(i) // panic: runtime error: index out of range [5] with length 5
	//}
	for i := 0; i < len(sliceFloat64); i++ {
		sliceFloat64[i] = float64(i)
	}
	fmt.Printf("\t&sliceFloat64:%p sliceFloat64:%p sliceFloat64:%+v len:%d cap:%d\n",
		&sliceFloat64,
		sliceFloat64,
		sliceFloat64,
		len(sliceFloat64),
		cap(sliceFloat64))
}

// Steps4 二维切片
func Steps4() {
	// Steps 4-1: 定义二维切片，并赋值
	sliceStringString := [][]string{
		[]string{"0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0"},
		[]string{"0", "0", "0", "0", "0"},
	}
	fmt.Printf("\tsliceStringString:%+v len:%d cap:%d\n",
		sliceStringString,
		len(sliceStringString),
		cap(sliceStringString))

	// Steps 4-3: 添加一行
	sliceStringString = append(sliceStringString, []string{"1", "1", "1", "1", "1"})
	fmt.Printf("\tsliceStringString:%+v len:%d cap:%d\n",
		sliceStringString,
		len(sliceStringString),
		cap(sliceStringString))

	// Steps 4-3: 打印二维数组
	for i := 0; i < len(sliceStringString); i++ { // len(sliceStringString) y轴数组长度
		fmt.Print("\t")
		for j := 0; j < len(sliceStringString[i]); j++ { // len(sliceStringString[i]) 第i行 x轴数组长度
			fmt.Printf("%s ", sliceStringString[i][j])
		}
		fmt.Println()
	}
}

// Steps5 切片上截取切片
func Steps5() {
	// Steps 5-1: 定义切片并初始化
	sliceInt := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\t&sliceInt:%p sliceInt:%p sliceInt:%+v len:%d cap:%d\n",
		&sliceInt,
		sliceInt,
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
	for i := 0; i < len(sliceInt); i++ {
		fmt.Printf("\t&sliceInt[%d]:%p\n", i, &sliceInt[i])
	}

	fmt.Printf("\t--------------------------------\n")
	// Steps 5-2: 可以用 array[low : high] or slice[low : high] 来截取数组或切片的一个片段长度为 high-low
	// 注意: sliceInt[0:3] 等同于 sliceInt[:3]
	interceptionSliceInt := sliceInt[1:3] // 获取 sliceInt 下标 1-2 的元素:[1,2] 长度为2 容量为9
	fmt.Printf("\t&interceptionSliceInt:%p interceptionSliceInt:%p interceptionSliceInt:%+v len:%d cap:%d\n",
		&interceptionSliceInt,
		interceptionSliceInt,
		interceptionSliceInt,
		len(interceptionSliceInt),
		cap(interceptionSliceInt))
	for i := 0; i < len(interceptionSliceInt); i++ {
		fmt.Printf("\t&interceptionSliceInt[%d]:%p\n", i, &interceptionSliceInt[i])
	}
	/*
		对比sliceInt[1],sliceInt[2]的地址和interceptionSliceInt[0],interceptionSliceInt[1]的地址, 会发现他们是相等滴
		证明他们底层共用一片地址空间
		&sliceInt[1]:0xc0000200f8
		&sliceInt[2]:0xc000020100

		&interceptionSliceInt[0]:0xc0000200f8
		&interceptionSliceInt[1]:0xc000020100
	*/

	fmt.Printf("\t--------------------------------\n")
	// Steps 5-3: 可以用 slice[low : high: cap] 来截取切片或数组的一个片段长度为 high-low,容量为cap
	interceptionSliceIntCap := sliceInt[1:3:5] // 获取 sliceInt 下标 1-2 的元素:[1,2,3] 长度为2, 容量为4
	fmt.Printf("\t&interceptionSliceIntCap:%p interceptionSliceIntCap:%p interceptionSliceIntCap:%+v len:%d cap:%d\n",
		&interceptionSliceIntCap,
		interceptionSliceIntCap,
		interceptionSliceIntCap,
		len(interceptionSliceIntCap),
		cap(interceptionSliceIntCap))
	for i := 0; i < len(interceptionSliceInt); i++ {
		fmt.Printf("\t&interceptionSliceIntCap[%d]:%p\n", i, &interceptionSliceIntCap[i])
	}

	fmt.Printf("\t--------------------------------\n")
	// Steps 5-4: 切片并不存储任何数据，它只是描述了底层数组中的一段
	// 更改切片的元素会修改其底层数组中对应的元素,与它共享底层数组的其它切片都会观测到这些修改
	fmt.Printf("\t[modify before] sliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
	fmt.Printf("\t[modify before] interceptionSliceInt:%+v len:%d cap:%d\n",
		interceptionSliceInt,
		len(interceptionSliceInt),
		cap(interceptionSliceInt))
	interceptionSliceIntCap[0] = 111
	fmt.Printf("\t[modify after ] sliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
	fmt.Printf("\t[modify after ] interceptionSliceInt:%+v len:%d cap:%d\n",
		interceptionSliceInt,
		len(interceptionSliceInt),
		cap(interceptionSliceInt))

	fmt.Printf("\t--------------------------------\n")
	// Steps 5-5: 通过unsafe.Pointer函数强行获取截取切片之外的数据
	// interceptionSliceIntCap[2] 超出当前len, 打印报错 panic: runtime error: index out of range [2] with length 2
	//fmt.Printf("interceptionSliceIntCap[2]:%d",interceptionSliceIntCap[2])

	// 通过指针偏移强行获取interceptionSliceIntCap[2]底层元素（这种方式是不安全的）
	fmt.Printf("\tinterceptionSliceCap[2]:%d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&interceptionSliceIntCap[0])) + uintptr(16))))

	fmt.Printf("\t[modify before] sliceInt:%+v\n", sliceInt)
	// Steps 5-6: 修改interceptionSliceCap[2]的值为33,底层切片sliceInt对应[3]位置改变33
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&interceptionSliceIntCap[0])) + uintptr(16))) = 33
	fmt.Printf("\t[modify after ] sliceInt:%+v\n", sliceInt)
}

// Steps6 range 遍历切片
func Steps6() {
	// Steps 6-1: 初始化切片
	sliceString := []string{"Golang", "Tutorial"}

	for idx := range sliceString {
		fmt.Printf("\tindex: %d, value:%s\n", idx, sliceString[idx])
	}

	for idx, v := range sliceString {
		fmt.Printf("\tindex: %d, value:%s\n", idx, v)
	}
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
	fmt.Println("Steps6():")
	Steps6()
}
