package main

import (
	"fmt"
	"time"
)

// channel 是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值
// <- “箭头”就是数据流的方向
func main() {
	// 信道在使用前必须创建, 内建函数 make 来创建
	// make(chan T,size)  用内建函数 make 来创建 一个T类型的缓冲大小为 size 的 channel
	// make(chan int) 用内建函数 make 来创建 一个 int 类型的缓冲大小为 0 的 channel
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
