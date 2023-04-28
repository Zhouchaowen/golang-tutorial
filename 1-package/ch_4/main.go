package main

/*
	1.包别名
	2.匿名包
*/

// 导入: 标准库和第三方库
import (
	"fmt"
	m "math"      // 给math包起一个别名 m
	_ "math/rand" // 匿名包，主要用于引入一些驱动的init函数的初始化
)

// main 函数,程序的入口。
func main() {
	fmt.Println("MaxInt", m.MaxInt) // 引用导出变量的时候也要用别名
}
