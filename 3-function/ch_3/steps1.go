package main

import "fmt"

func counter() func() int {

	total := 0 // 闭包打包的环境

	return func() int { // 闭包函数
		total++
		return total
	}
}

func Steps1() {
	f := counter()
	fmt.Printf("\tnum: %d\n", f()) // 输出 1
	fmt.Printf("\tnum: %d\n", f()) // 输出 2
	fmt.Printf("\tnum: %d\n", f()) // 输出 3
	fmt.Printf("\t-------\n")
	ff := counter()
	fmt.Printf("\tnum: %d\n", ff()) // 输出 1
	fmt.Printf("\tnum: %d\n", ff()) // 输出 2
	fmt.Printf("\tnum: %d\n", ff()) // 输出 3
}
