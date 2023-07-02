package main

import "fmt"

/*
	1.阻塞channel与非阻塞channel
*/

// NoBufferChan 不带缓冲的 channel
func NoBufferChan() {
	ch := make(chan int)
	ch <- 1 // 被阻塞,执行报错 fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
}

// BufferChan 带缓冲的 channel
func BufferChan() {
	// channel 有缓冲、是非阻塞的，直到写满 cap 个元素后才阻塞
	ch := make(chan int, 1)
	ch <- 1 // 未写满 cap 不会被阻塞
	fmt.Println(<-ch)
}

func main() {
	//NoBufferChan()
	BufferChan()
}
