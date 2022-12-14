package main

import "fmt"

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
	mpIntString[1] = "Tutorial" // 修改mpIntString[1]元素
	mpIntString[2] = "Study"
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// Steps 2-3: 获取元素
	elem := mpIntString[0]
	fmt.Printf("\telem:%+v\n", elem)

	// Steps 2-4: 删除元素
	// 通过内置函数 delete(map,key)
	delete(mpIntString, 0)
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// Steps 2-5: 通过双赋值检测某个key是否存在
	// 若 key 在 mpIntString 中，ok 为 true ; 否则, ok 为 false
	elem, ok := mpIntString[0]
	fmt.Printf("\telem:%+v ok:%t\n", elem, ok)

	// Steps 2-6: 通过range遍历map
	fmt.Printf("\tmpIntString:%+v len:%d\n",
		mpIntString,
		len(mpIntString))

	// 方法1，拿到key，再根据key获取value
	for k := range mpIntString {
		fmt.Printf("\tKey:%d, Value:%s\n", k, mpIntString[k])
	}
	// 方法2，同时拿到key和value
	for k, v := range mpIntString {
		fmt.Printf("\tKey:%d, Value:%s\n", k, v)
	}
}

func main() {
	fmt.Println("Steps2():")
	Steps2()
}
