package main

import "fmt"

func NoBufferChan() {
	ch := make(chan int)
	ch <- 1 //被阻塞,执行报错 fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
}

func BufferChan() {
	// channel 有缓冲、是非阻塞的，直到写满 cap 个元素后才阻塞写入
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)
}

// 带缓冲 channel
func main() {
	//NoBufferChan()
	BufferChan()
}
