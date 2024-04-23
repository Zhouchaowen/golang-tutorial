package main

import "fmt"

/*
	1.定义interface并赋值
	2.断言interface
*/

func Steps1() {
	var i interface{} = "hello"

	s := i.(string) // 将变量 i 断言为 string 类型并赋值给 s
	fmt.Println(s)

	// 类型断言, 断言失败一般会导致panic的发生, 所以为了防止panic的发生, 我们需要在断言时使用双值检查。
	// 如果断言失败, 那么ok的值将会是false
	// 如果断言成功, 那么ok的值将会是true, 同时s将会得到正确类型的值。
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 如果断言失败 报错(panic)
	fmt.Println(f)
}

// question ?
/*
	type Print interface {
		print()
	}

	type Test struct {
		Name string
	}

	func (t Test) print() {
		fmt.Println(t.Name)
	}

	Test 结构体实现了 Print 接口, Test能断言成Print接口: Test.(Print), 那[]Test能断言成[]Print吗？[]Test.([]Print) x or √ ？
*/

type Print interface {
	print()
}

type Test struct {
	Name string
}

func (t Test) print() {
	fmt.Println(t.Name)
}

func Steps2() {
	var tmp interface{} = Test{Name: "golang"}
	value, ok := tmp.(Test)
	value2, ok2 := tmp.(Print)
	fmt.Printf("tmp.(Test)     : %+v,%+v\n", value, ok)
	fmt.Printf("tmp.(Print)    : %+v,%+v\n", value2, ok2)

	var tmps interface{} = []Print{Test{Name: "golang"}, Test{Name: "tutorial"}}
	value3, ok3 := tmps.([]Print) // []Print和[]Print 是同类型所以可以断言成功
	value4, ok4 := tmps.([]Test)  // []Print和[]Test 不是同类型所以不可以断言成功
	fmt.Printf("tmps.([]Print) : %+v,%+v\n", value3, ok3)
	fmt.Printf("tmps.([]Test)  : %+v,%+v\n", value4, ok4)
	// 回答上面的问题：[]Test能断言成[]Print吗？[]Test.([]Print)；
	//	不能，应为切片也是一种类型，所以[]Print和[]Test不是同一种类型(即使Test实现了Print)
}

// 类型断言
// 断言 interface
func main() {
	//Steps1()
	Steps2()
}
