package main

import "fmt"

/*
	1.If 分支
	2.If-Else 分支
*/

// if 分支打印不同字符
func main() {
	flag := 10
	if flag > 5 { // 判断表达式
		fmt.Println("flag:", flag)
	}

	flag = 14
	//flag = 16
	//flag = 21

	if flag > 20 {
		fmt.Println("flag:", flag)
	} else if flag < 15 {
		fmt.Println("flag:", flag)
	} else {
		fmt.Println("flag:", flag)
	}
}
