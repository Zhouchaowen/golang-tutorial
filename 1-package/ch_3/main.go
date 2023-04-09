package main

/*
	1.导出变量
	2.导出函数
*/

// 导入: 标准库和第三方库
import (
	"fmt"
	"math"
)

// main 函数,程序的入口。
func main() {
	// 注意：在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问

	// 什么是导出？首字母大写代表导出,小写代表不导出
	//fmt.Println("intSize",math.intSize) // 引用未导出变量将报错
	fmt.Println("MaxInt", math.MaxInt)
}
