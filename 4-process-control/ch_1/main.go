package main

import "fmt"

// 通过 for 循环打印 0-9
func main() {
	sum := 0
	// for 循环
	// i := 0 初始化语句：在第一次迭代前执行
	// i < 10 条件表达式：在每次迭代前求值
	// i++    后置语句：在每次迭代的结尾执行
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 0
	// 初始化语句和后置语句是可选的
	for sum < 5 {
		sum++
	}
	fmt.Println(sum)
}
