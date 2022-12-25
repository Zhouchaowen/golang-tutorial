package main

import (
	"fmt"
	"time"
)

/*
	1.发送端和接收端的阻塞问题
*/

// 发送端和接收端的阻塞问题
// 发送端在没有准备好之前会阻塞,同样接收端在发送端没有准备好之前会阻塞
func main() {
	c := make(chan string)

	go func() {
		<-time.After(time.Second * 10)
		fmt.Println("发送端准备好了 send: ping")
		c <- "ping" // 发送
	}()

	// 发送端10s后才准备好，所以阻塞在当前位置
	fmt.Println("阻塞在当前位置，发送端发送数据后才继续执行")
	num := <-c
	fmt.Printf("recover: %s\n", num)
}
