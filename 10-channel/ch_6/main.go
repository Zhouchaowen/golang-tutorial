package main

import (
	"fmt"
	"time"
)

// 通过 channel+select 控制 goroutine 退出
func genNum(c, quit chan int) {
	for i := 0; ; i++ {
		// select 可以等待多个通信操作
		// select 会阻塞等待可执行分支。当多个分支都准备好时会随机选择一个执行。
		select {
		case <-quit:
			// 发送者可通过 close 关闭一个信道来表示没有需要发送的值了。
			close(c)
			return
		default: // 等同于 switch 的 default。当所以case都阻塞时如果有default则，执行default
			c <- i
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go genNum(c, quit)

	// 循环 for v := range c 会不断从信道接收值，直到它被关闭
	// 并且只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。
	for i := 0; i < 10; i++ {
		fmt.Println("receive:", <-c)
	}

	// 通知 genNum() 退出
	quit <- 1

	// 接收者可以通过 v,ok := <- c 表达式第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完
	v, ok := <-c
	fmt.Printf("value:%d, ok:%t\n", v, ok)

	fmt.Println("close")
}
