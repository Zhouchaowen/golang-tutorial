package main

import (
	"fmt"
	"unsafe"
)

/*
	1.切片浅拷贝
	2.切片深拷贝
*/

// 指针持有者类型的拷贝问题

// Steps1 浅拷贝
func Steps1() {
	var sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var sliceIntTmp []int

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	sliceIntTmp = sliceInt
	fmt.Printf("\tsliceIntTmp:%+v len:%d cap:%d\n",
		sliceIntTmp,
		len(sliceIntTmp),
		cap(sliceIntTmp))

	sliceIntTmp[0] = 111

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
	fmt.Printf("\tsliceIntTmp:%+v len:%d cap:%d\n",
		sliceIntTmp,
		len(sliceIntTmp),
		cap(sliceIntTmp))
}

// Steps2 深拷贝
func Steps2() {
	var sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var sliceIntTmp []int

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	sliceIntTmp = make([]int, len(sliceInt))

	copy(sliceIntTmp, sliceInt) // 深拷贝

	fmt.Printf("\tsliceIntTmp:%+v len:%d cap:%d\n",
		sliceIntTmp,
		len(sliceIntTmp),
		cap(sliceIntTmp))

	sliceIntTmp[0] = 111

	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))
	fmt.Printf("\tsliceIntTmp:%+v len:%d cap:%d\n",
		sliceIntTmp,
		len(sliceIntTmp),
		cap(sliceIntTmp))
}

// Steps3 解释浅拷贝产生原因
func Steps3() {
	var sliceInt = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("\t            sliceInt address:%p\n", &sliceInt) // &sliceInt 获取的是sliceInt变量本身的地址
	fmt.Printf("\t      sliceInt value address:%p\n", sliceInt)  //  sliceInt 获取切片底层数据的地址

	fmt.Printf("\t   sliceInt[0] value address:%p\n", &sliceInt[0]) // &sliceInt[0] 获取切片第一个值的地址
	fmt.Printf("\t   sliceInt[1] value address:%p\n", &sliceInt[1]) // &sliceInt[1] 获取切片第二个值的地址
	fmt.Printf("\t   sliceInt[2] value address:%p\n", &sliceInt[2]) // &sliceInt[2] 获取切片第三个值的地址
	fmt.Printf("\t   sliceInt[3] value address:%p\n", &sliceInt[3]) // &sliceInt[3] 获取切片第四个值的地址

	tmpSliceInt := sliceInt
	fmt.Printf("\t-----------------------------\n")
	fmt.Printf("\t         tmpSliceInt address:%p\n", &tmpSliceInt) // &tmpSliceInt 获取的是tmpSliceInt变量本身的地址
	fmt.Printf("\t   tmpSliceInt value address:%p\n", tmpSliceInt)  //  tmpSliceInt 获取切片底层数据的地址

	fmt.Printf("\ttmpSliceInt[0] value address:%p\n", &tmpSliceInt[0]) // &tmpSliceInt[0] 获取切片第一个值的地址
	fmt.Printf("\ttmpSliceInt[1] value address:%p\n", &tmpSliceInt[1]) // &tmpSliceInt[1] 获取切片第二个值的地址
	fmt.Printf("\ttmpSliceInt[2] value address:%p\n", &tmpSliceInt[2]) // &tmpSliceInt[2] 获取切片第三个值的地址
	fmt.Printf("\ttmpSliceInt[3] value address:%p\n", &tmpSliceInt[3]) // &tmpSliceInt[3] 获取切片第四个值的地址
}

// Steps4 证明切片底层由三部分组成 array ptr, len, cap
func Steps4() {
	var sliceInt = make([]int, 5, 10)
	sliceInt[0] = 1
	sliceInt[1] = 2
	sliceInt[2] = 33
	sliceInt[3] = 4
	sliceInt[4] = 5
	fmt.Printf("\tsliceInt:%+v len:%d cap:%d\n",
		sliceInt,
		len(sliceInt),
		cap(sliceInt))

	fmt.Printf("\tsliceInt memery size:%d\n", unsafe.Sizeof(sliceInt)) // &sliceInt 获取的是sliceInt变量本身的地址

	fmt.Printf("\tsliceInt variable address:%p\n", &sliceInt) // &sliceInt 获取的是sliceInt变量本身的地址
	fmt.Printf("\tsliceInt value    address:%p\n", sliceInt)  //  sliceInt 获取切片底层数据的地址

	fmt.Printf("\t-----------------------------\n")
	/*
		切片是一个指针持有者类型，底层由三部分组成  array ptr, len, cap
		type slice struct {
			array unsafe.Pointer
			len   int
			cap   int
		}
	*/

	fmt.Printf("\tsliceInt data array pointer  value 0x%x\n", (*uintptr)(unsafe.Pointer(&sliceInt)))
	fmt.Printf("\tsliceInt data array pointer *value 0x%x\n", *(*uintptr)(unsafe.Pointer(&sliceInt)))
	fmt.Printf("\tsliceInt data len  value 0x%x\n", (*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&sliceInt))+uintptr(8))))
	fmt.Printf("\tsliceInt data len *value %d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sliceInt)) + uintptr(8))))
	fmt.Printf("\tsliceInt data cap  value 0x%x\n", (*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&sliceInt))+uintptr(16))))
	fmt.Printf("\tsliceInt data cap *value %d\n", *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sliceInt)) + uintptr(16))))

	fmt.Printf("\t-----------------------------\n")
	// 等同于上面的结果
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}
	s := *(*slice)(unsafe.Pointer(&sliceInt))
	fmt.Printf("\tsliceInt %+v\n", s)

	fmt.Printf("\t-----------------------------\n")

	// 通过偏移量获取切片上的值
	fmt.Printf("\tsliceInt data array pointer value[0] %d\n", *(*int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sliceInt)))))
	fmt.Printf("\tsliceInt data array pointer value[1] %d\n", *(*int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sliceInt)) + uintptr(8))))
	fmt.Printf("\tsliceInt data array pointer value[2] %d\n", *(*int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sliceInt)) + uintptr(16))))
	fmt.Printf("\tsliceInt data array pointer value[3] %d\n", *(*int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sliceInt)) + uintptr(24))))
	fmt.Printf("\tsliceInt data array pointer value[4] %d\n", *(*int)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sliceInt)) + uintptr(32))))

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
