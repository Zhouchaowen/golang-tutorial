package main

import (
	"fmt"
	"time"
)

// 定义 channel, channel 是带有类型的管道，可以通过信道操作符 <- 来发送或者接收值
func main() {
	// 信道在使用前必须通过内建函数 make 来创建

	// make(chan T,size)  标识用内建函数 make 来创建 一个T类型的缓冲大小为 size 的 channel
	// 如下: make(chan int) 用内建函数 make 来创建 一个 int 类型的缓冲大小为 0 的 channel
	c := make(chan int)

	go func() {
		// 从 c 接收值并赋予 num
		num := <-c
		fmt.Printf("recover:%d\n", num)
	}()

	// 将 1 发送至信道 c
	c <- 1

	<-time.After(time.Second * 3)

	fmt.Println("return")
}
