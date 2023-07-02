package main

import (
	"fmt"
	"time"
)

/*
	1.通过make定义channel
	2.读和写channel
*/

// 定义 channel, channel 是带有类型的管道，可以通过信道操作符 <- 来发送或者接收值
func main() {
	// 信道在使用前必须通过内建函数 make 来创建

	// make(chan T,size)  表示用内建函数 make 来创建 1 个 T 类型的缓冲大小为 size 的 channel
	// 如下: make(chan int) 用内建函数 make 来创建 1 个 int 类型的缓冲大小为 0 的 channel
	c := make(chan int)

	go func() {
		// 通过信道操作符 <- 从 c 接收值并赋予 num
		num := <-c
		fmt.Printf("recover:%d\n", num)
	}()

	// 通过信道操作符 <- 将 1 发送至信道 c
	c <- 1

	// 等待 3s 让 goroutine 得到执行
	<-time.After(time.Second * 3)

	fmt.Println("return")
}
