package main

import "fmt"

/*
	1.基本 for 循环
	2.类似 while 循环
	3.range循环，range循环和for循环区别
	4.continue和break
	5.goto
*/

// Steps1 通过 for 循环累加 0-9
func Steps1() {
	sum := 0
	// for 关键字
	// i := 0 初始化语句：在第一次迭代前执行
	// i < 10 条件表达式：在每次迭代前求值
	// i++    后置语句：在每次迭代的结尾执行
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("\tsum: %d\n", sum)
}

// Steps2 for循环初始化语句和后置语句不是必须的
func Steps2() {
	sum := 0
	// 初始化语句和后置语句是可选的
	for sum < 5 {
		sum++
	}

	// 死循环
	//for {
	//
	//}

	fmt.Printf("\tsum: %d\n", sum)
}

// Steps3 range形式的循环遍历
func Steps3() {
	str := "Golang Tutorial"
	// 遍历字符串打印每个字符
	// i: 字符串中每个字符的索引,从0开始
	// v: 字符串中每个字符
	for i, v := range str {
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

// Steps5 for 循环中的 break 和 continue
func Steps5() {
	for i := 0; i < 10; i++ {
		if i == 5 { // 下一小节介绍
			fmt.Printf("\ti:%d, continue\n", i)
			continue
		}

		if i == 6 {
			fmt.Printf("\ti:%d, break\n", i)
			break
		}
	}
}

// Steps6 goto 实现循环
func Steps6() {
	i := 0

Next: // 跳转标签声明
	fmt.Printf("\ti:%d\n", i)
	i++
	if i < 5 {
		goto Next // 跳转
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
	fmt.Println("Steps5():")
	Steps5()
	fmt.Println("Steps6():")
	Steps6()
}
