package main

import (
	"fmt"
	"time"
)

func genNum(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	// 发送者可通过 close 关闭一个信道来表示没有需要发送的值了。
	close(c)
}

// 通过 range 变量 channel
func main() {
	c := make(chan int, 10)
	go genNum(c)

	// 循环 for v := range c 会不断从信道接收值，直到它被关闭
	// 并且只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。
	for v := range c {
		fmt.Println("receive:", v)
	}

	// 接收者可以通过 v,ok := <- c 表达式第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完
	v, ok := <-c
	fmt.Printf("value:%d, ok:%t\n", v, ok)

	fmt.Println("close")
}
