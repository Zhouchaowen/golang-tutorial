package main

import "fmt"

// defer作用:
//    释放占用的资源
//    捕捉处理异常 recover

// Steps1 defer 语句会将函数推迟到外层函数返回之后执行。
// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
func Steps1() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

// Steps2 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
func Steps2() {
	fmt.Println("begin")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
		fmt.Println(i)
	}
	fmt.Println("end")
}

func main() {
	//Steps1()
	Steps2()
}

/*
     -----   -----
    |     | |     |
    |   | V | |   |
	|	|     |   V
		| ... |
		|  3  |
		|  2  |
		|  1  |
		|  0  |
		 —————
*/
