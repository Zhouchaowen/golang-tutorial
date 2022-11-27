package main

import (
	"fmt"
	"time"
)

// 并发与并行：https://gfw.go101.org/article/control-flows-more.html
// goroutine是轻量级的用户态线程，可以在代码里创建成千上万个goroutine来并发工作

// 使用 goroutine 打印数据
func main() {
	language := []string{"golang", "java", "c++", "python", "rust", "js"}
	tutorial := []string{"入门", "初级", "中级", "高级", "专家"}

	// Go 程（goroutine）是由 Go 运行时管理的轻量级线程
	// 在函数调⽤语句前添加 go 关键字，就可创建一个 goroutine
	go listLanguage(language) // 通过goroutine启动该函数
	go listTutorial(tutorial)

	<-time.After(time.Second * 10) // 10s后执行下一行
	fmt.Println("return")
}

func listLanguage(items []string) {
	for i := range items {
		fmt.Printf("language: %s\n", items[i])
		time.Sleep(time.Second)
	}
}

func listTutorial(items []string) {
	for i := range items {
		fmt.Printf("tutorial: %s\n", items[i])
		time.Sleep(time.Second)
	}
}
