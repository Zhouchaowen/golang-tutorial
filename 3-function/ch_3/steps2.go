package main

import "fmt"

func Steps2() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Printf("\tx: %d\n", x)

	fmt.Printf("\tnum: %d\n", increment()) // 输出 1
	fmt.Printf("\tnum: %d\n", increment()) // 输出 2
	fmt.Printf("\tnum: %d\n", increment()) // 输出 3

	fmt.Printf("\tx: %d\n", x)

	// 改变x的值
	x = 10
	fmt.Printf("\tx: %d\n", x)
	fmt.Printf("\tnum: %d\n", increment())
	fmt.Printf("\tnum: %d\n", increment())
}
