package main

import (
	"fmt"
	"unsafe"
)

/*
	1.用内建函数 make 来创建map
	2.赋值
	3.获取元素
	4.删除元素
	5.通过双赋值检测某个key是否存在
	6.通过range遍历map
*/

// Steps2 map的基础使用
func Steps2() {
	// Steps 2-1: 用内建函数 make 来创建map
	mpIntString := make(map[int]string)
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// Steps 2-2: 映射 mpIntString 中插入或修改元素
	// 映射添加元素容量不够时会自动扩容
	mpIntString[0] = "Golang"
	mpIntString[1] = "World"
	mpIntString[1] = "Tutorial" // 覆盖mpIntString[1]的value
	mpIntString[2] = "Study"
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	fmt.Printf("\t---------------------\n")

	// Steps 2-3: 获取元素
	elem := mpIntString[0]
	fmt.Printf("\telem:%+v\n", elem)

	fmt.Printf("\t---------------------\n")

	// Steps 2-4: 删除元素
	// 通过内置函数 delete(map,key)
	delete(mpIntString, 0)
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	fmt.Printf("\t---------------------\n")

	// Steps 2-5: 通过双赋值检测某个key是否存在
	// 若 key 在 mpIntString 中，ok 为 true ; 否则 ok 为 false
	elem, ok := mpIntString[0]
	fmt.Printf("\telem:%+v ok:%t\n", elem, ok)

	fmt.Printf("\t---------------------\n")

	// Steps 2-6: 通过range遍历map
	// 方法1，拿到key，再根据key获取value
	for k := range mpIntString {
		fmt.Printf("\tKey:%d, Value:%s\n", k, mpIntString[k])
	}

	fmt.Printf("\t---------------------\n")

	// 方法2，同时拿到key和value
	for k, v := range mpIntString {
		fmt.Printf("\tKey:%d, Value:%s\n", k, v)
	}
}

// Steps3 不同方式创建map的区别
func Steps3() {
	var mpIntBool map[int]bool
	fmt.Printf("\tmpIntBool:%+v len:%d\n",
		mpIntBool,
		len(mpIntBool))
	fmt.Printf("\tmpIntBool       size:%d\n", unsafe.Sizeof(mpIntBool))
	fmt.Printf("\tmpIntBool       addr:%p\n", &mpIntBool)
	fmt.Printf("\tmpIntBool value addr:%p\n", mpIntBool)

	fmt.Printf("\t---------------------\n")

	var mpIntBool1 = map[int]bool{} // 与 var mpIntBool map[int]bool 的区别; 会开辟内存空间
	fmt.Printf("\tmpIntBool1:%+v len:%d\n",
		mpIntBool1,
		len(mpIntBool1))
	fmt.Printf("\tmpIntBool1       size:%d\n", unsafe.Sizeof(mpIntBool1))
	fmt.Printf("\tmpIntBool1       addr:%p\n", &mpIntBool1)
	fmt.Printf("\tmpIntBool1 value addr:%p\n", mpIntBool1)

	fmt.Printf("\t---------------------\n")

	mpIntBool2 := map[int]bool{} // 与 var mpIntBool map[int]bool 的区别; 会开辟内存空间
	fmt.Printf("\tmpIntBool2:%+v len:%d\n",
		mpIntBool2,
		len(mpIntBool2))
	fmt.Printf("\tmpIntBool2       size:%d\n", unsafe.Sizeof(mpIntBool2))
	fmt.Printf("\tmpIntBool2       addr:%p\n", &mpIntBool2)
	fmt.Printf("\tmpIntBool2 value addr:%p\n", mpIntBool2)

	fmt.Printf("\t---------------------\n")

	var mpIntBool3 = make(map[int]bool, 10) // 与 var mpIntBool map[int]bool 的区别; 会开辟内存空间
	fmt.Printf("\tmpIntBool3:%+v len:%d\n",
		mpIntBool3,
		len(mpIntBool3))
	fmt.Printf("\tmpIntBool3       size:%d\n", unsafe.Sizeof(mpIntBool3))
	fmt.Printf("\tmpIntBool3       addr:%p\n", &mpIntBool3)
	fmt.Printf("\tmpIntBool3 value addr:%p\n", mpIntBool3)
}

func main() {
	fmt.Println("Steps2():")
	Steps2()
	fmt.Println("Steps3():")
	Steps3()
}
