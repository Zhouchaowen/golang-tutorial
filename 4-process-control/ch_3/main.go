package main

import "fmt"

// switch 是编写一连串 if - else 语句的简便方法
func main() {
	Demo1()
	Demo2()
	Demo3()
}

// 基础用法
func Demo1() {
	flag := 1
	//flag = 2
	//flag = 3
	//flag = 4
	//flag = 5

	switch flag { // flag 待判断条件
	case 1: // 条件 flag 是否等于 1。是：执行该case下的流程，否：选择其它满足条件的 case
		fmt.Println("case:", flag)
		// Golang 中每个 case 后面不需要 break 语句。当然 return 是可选的
	case 2:
		fmt.Println("case:", flag)
	case 3, 4: // case 可以设置多个条件。只要 flag 等于3或4都能执行当前case流程
		fmt.Println("case:", flag)
	case 5:
		fmt.Println("case:", flag)
		return
	default: // 当所有case都无法满足, 会执行 default 的流程。如果没有 default 那当前 switch 执行完成
		fmt.Println("default:", flag)
	}
}

// switch 条件可以是任何支持判断的类型
func Demo2() {
	flag := "Hello"
	flag = "World"
	flag = "Golang"
	flag = "Tutorial"
	flag = "Process"

	switch flag { // flag 待判断条件
	case "Hello": // 条件 flag 是否等于 "Hello"。是：执行该case下的流程，否：选择其它满足条件的 case
		fmt.Println("case:", flag)
	case "World":
		fmt.Println("case:", flag)
	case "Golang", "Tutorial": // case 可以设置多个条件。只要 flag 等于"Golang"或"tutorial"都能执行当前case流程
		fmt.Println("case:", flag)
	default: // 当所有case都无法满足, 会执行 default 的流程。如果没有 default 那当前 switch 执行完成
		fmt.Println("default:", flag)
	}
}

// switch true 可以将一长串 if-then-else 写得更加清晰
func Demo3() {
	flag := 1
	//flag = 2
	//flag = 3
	//flag = 4
	//flag = 5
	//flag = 7

	switch { // flag 待判断条件
	case flag < 2: // 条件 flag 是否小于 2。是：执行该case下的流程，否：选择其它满足条件的 case
		fmt.Println("case flag < 2 flag:", flag)
	case flag < 4:
		fmt.Println("case flag < 4, flag:", flag)
	case flag > 6, flag < 10: // case 可以设置多个条件。flag 大于6或小于10都能执行当前case流程
		fmt.Println("case flag > 6 || flag < 10 flag:", flag)
	case flag > 6 && flag < 10: // case 可以设置组合条件。flag 大于6并且小于10都才能执行当前case流程
		fmt.Println("case flag > 6 || flag < 10 flag:", flag)
	}
}
