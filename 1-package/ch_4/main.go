package main

// 导入: 标准库和第三方库
import (
	"fmt"
	m "math" // 给math包起一个别名 m
)

// main 函数,程序的入口。
func main() {
	fmt.Println("MaxInt", m.MaxInt) // 引用导出变量的时候也要用别名
}
