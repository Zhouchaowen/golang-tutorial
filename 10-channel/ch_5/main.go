package main

import (
	"fmt"
	"time"
)

// 通过 range 遍历 channel, 并通过关闭 channel 来退出循环

// 复制一个 channel 或用于函数参数传递时, 只是拷贝了一个 channel 的引用, 因此调用者和被调用者将引用同一个channel对象
func genNum(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	// 发送者可通过 close 关闭一个信道来表示没有需要发送的值了
	close(c)
}

func main() {
	c := make(chan int, 10)
	go genNum(c)

	// 循环 for v := range c 会不断从信道接收值，直到它被关闭
	// 并且只有发送者才能关闭信道，而接收者不能, 向一个已经关闭的信道发送数据会引发程序恐慌（panic）
	for v := range c {
		fmt.Println("receive:", v)
	}

	// 接收者可以通过 v,ok := <- c 表达式接收第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭,那么 v 为对应类型零值,ok 为 false
	v, ok := <-c
	fmt.Printf("value:%d, ok:%t\n", v, ok)

	fmt.Println("close")
}
