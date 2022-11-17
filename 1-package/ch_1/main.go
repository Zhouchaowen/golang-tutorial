package main // 定义了包名,只有定义了main的包才能独立运行

// 导入: 标准库和第三方库
import (
	"fmt" // 引入一个标准库包
)

// main 函数,程序的入口
func main() {
	// 调用标准库 fmt 在控制台打印 hello world 字符串
	fmt.Println("hello world")
}

// 通过命令行运行
// go run main.go
// go build main.go && ./main
