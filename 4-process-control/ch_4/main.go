package main

import "fmt"

/*
	1.defer
*/

// defer作用:
//    释放占用的资源
//    捕捉处理异常 recover

// Steps1 defer 语句会将函数推迟到外层函数返回之后执行。
// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
func Steps1() {
	defer fmt.Printf(" world\n")

	fmt.Printf("\thello")
}

// Steps2 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
func Steps2() {
	fmt.Println("\tbegin")
	for i := 0; i < 3; i++ {
		defer fmt.Println("\t\ti:", i)
		fmt.Printf("\t\ti:%d\n", i)
	}
	fmt.Println("\tend")
}

func Steps3() {
	fmt.Println("\tbegin")
	x := 2
	defer func() {
		x = x * x
		fmt.Println("\tx =", x) // x = 9
	}()
	fmt.Println("\tend")
	x = 3
}

func Steps4() {
	fmt.Println("\tbegin")
	x := 2
	defer func(x int) {
		x = x * x
		fmt.Println("\tx =", x) // x = 4
	}(x)
	fmt.Println("\tend")
	x = 3
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
