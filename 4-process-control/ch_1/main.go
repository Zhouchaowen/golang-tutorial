package main

import "fmt"

// Steps1 通过 for 循环打印 0-9
func Steps1() {
	sum := 0
	// for 循环
	// i := 0 初始化语句：在第一次迭代前执行
	// i < 10 条件表达式：在每次迭代前求值
	// i++    后置语句：在每次迭代的结尾执行
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("\tsum: %d\n", sum)
}

// Steps2 初始化语句和后置语句不是必须的
func Steps2() {
	sum := 0
	// 初始化语句和后置语句是可选的
	for sum < 5 {
		sum++
	}
	fmt.Printf("\tsum: %d\n", sum)
}

// Steps3 range形式的循环遍历
func Steps3() {
	str := "Golang Tutorial"
	for i, v := range str { // 遍历字符串
		fmt.Printf("\ti:%d,v:%c\n", i, v)
	}
}

// Steps4 range和for遍历的区别
func Steps4() {
	str := "Golang 教程"
	for i := 0; i < len(str); i++ {
		fmt.Printf("\ti:%d,v:%c\n", i, str[i])
	}

	for i, v := range str { // 遍历字符串
		fmt.Printf("\ti:%d,v:%c\n", i, v)
	}
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