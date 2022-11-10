package main

import (
	"fmt"
	"time"
)

// 发送和接收操作在另一端准备好之前都会阻塞
func main() {
	c := make(chan int)
	go func() {
		fmt.Println("阻塞在当前位置，发送端发送数据后才继续执行")
		// 发送端3s后才准备好，所以阻塞在当前位置
		num := <-c
		fmt.Printf("recover:%d\n", num)
	}()

	<-time.After(time.Second * 10)

	fmt.Println("send....")
	c <- 1

	<-time.After(time.Second * 3)

	fmt.Println("return")
}
