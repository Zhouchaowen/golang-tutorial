package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	1.sync.WaitGroup的使用
*/

func listLanguage(items []string, wg *sync.WaitGroup) { // 一般不建议这样使用
	defer wg.Done()

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

// 使用 WaitGroup等待goroutine执行完成
func main() {
	language := []string{"golang", "java", "c++", "python", "rust", "js"}
	tutorial := []string{"入门", "初级", "中级", "高级", "专家"}

	var wg sync.WaitGroup // 改结构用于等待 goroutine 执行完成

	wg.Add(2) // 设置需要等待 goroutine 的数量,目前为2

	go listLanguage(language, &wg) // 通过 goroutine 启动该函数, 将wg传递到函数中去(注意是传递的指针到函数中)

	go func() { // 建议使用方式
		defer wg.Done() // 程序运行完毕, 将等待数量减1
		listTutorial(tutorial)
	}()

	wg.Wait() // 当等待数量为0后执行下一行
	//<-time.After(time.Second * 10) // 10s后执行下一行。 通过 wg.Wait() 代替
	fmt.Println("return")
}
