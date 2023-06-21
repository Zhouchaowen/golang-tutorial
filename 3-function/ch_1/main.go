package main

import "fmt"

/*
	1.定义无参函数
	2.定义有参函数
	3.定义多返回参数函数
	4.定义不定长参数函数
	5.定义指针参数函数
*/

/* 定义函数格式
func name([parameter list]) [return_type list] {
  do some things
}
*/

// 定义一个 print 函数, 它接受 0 个参数，返回 0 个值
func print() {
	fmt.Println("hello world")
}

// 定义一个 add 函数, 它接受 2 个 int 类型的参数,返回 1 个 int 类型的值
func add(x int, y int) int {
	return x + y
}

// 定义一个 swap 函数, 它接受 2 个 int 类型的参数,返回 2 个 int 类型的值
func swap(x, y int) (int, int) {
	return y, x
}

// 返回值可被命名(a,b)，它们会被视作定义在函数顶部的变量
func swap2(x, y int) (a int, b int) {
	a = y
	b = x
	return
}

// 传递不定长参数 ...T
func variableCut(x int, y ...int) int {
	for _, v := range y {
		x += v
	}
	return x
}

// 值类型参数
func modifyValue(x int) {
	x = x * 10
}

// 指针类型参数
func modifyPointer(x *int) {
	*x = (*x) * 10
}

// 值类型参数
func printValue(j int) {
	fmt.Printf("value addr:%p\n", &j)
}

// 指针类型参数
func printPointer(j *int) {
	fmt.Printf("value addr:%p\n", j)
	fmt.Printf("pointer addr:%p\n", &j)
}

func main() {
	// 调用 print() 函数
	print()

	// 调用 add() 函数并传递 2 个 int 类型数据 1 和 2
	fmt.Printf("add() return: %d\n", add(1, 2))

	// 调用 swap() 函数并传递 2 个 int 类型数据 1 和 2 并接收两个返回 x 和 y
	x, y := swap(1, 2)
	fmt.Printf("swap() x: %d, y: %d \n", x, y)

	// 调用 variableCut() 函数并传递 1 个 int 类型数据 1 和 1个不定长的 int 类型数据 2,3,4,5
	fmt.Printf("advariableCutd() return: %d\n", variableCut(1, 2, 3, 4, 5))

	// 值类型参数与指针类型参数的区别
	//		值类型参数是指在调用函数时将实际数据复制一份传递给函数中对应的接受参数,这样在函数中如果对参数进行修改，将不会影响到外面的数据
	//		指针类型参数(引用传递)是指在调用函数时将实际数据的地址复制一份传递给函数中对应的指针接受参数，那么在函数中对参数进行的修改，将影响到外面的数据
	// 不管是值类型参数还是指针类型参数都是【值传递】，只是传递的是值数据还在指针地址数据
	x = 1
	modifyValue(x)
	fmt.Printf("modifyValue() return: %d\n", x)

	x = 1
	modifyPointer(&x) // 和 modifyValue 区别是使用了 &
	fmt.Printf("modifyPointer() return: %d\n", x)

	// 参数传递的数据地址
	j := 10
	fmt.Printf("value addr:%p\n", &j)
	printValue(j)
	printPointer(&j)
}
