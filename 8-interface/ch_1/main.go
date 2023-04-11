package main

import "fmt"

/*
	1.定义interface并赋值
	2.断言interface
*/

// 类型断言
// 断言 interface
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// 类型断言, 断言失败一般会导致panic的发生, 所以为了防止panic的发生, 我们需要在断言前进行一定的判断。
	// 如果断言失败, 那么ok的值将会是false
	// 如果断言成功, 那么ok的值将会是true, 同时s将会得到正确类型的值。
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 如果断言失败 报错(panic)
	fmt.Println(f)
}
