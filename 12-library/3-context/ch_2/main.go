// 演示传递了一个带有任意截止日期的上下文来告诉一个阻塞函数它应该在它到达它时立即放弃它的工作。
package main

import (
	"context"
	"fmt"
	"time"
)

func doSomeTask(ctx context.Context) {
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exit task, reason:", ctx.Err())
				return
			default:
				fmt.Println("generate a number:", n)
				time.Sleep(time.Second * 3)
			}
			n++
		}
	}()
}

func Steps1() {
	// 截止时间上下文,到达指定时间后ctx.Done()返回的chanel会接收到空结构体值
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	go doSomeTask(ctx)

	<-time.After(15 * time.Second)
	fmt.Println("exit main")
}

func Steps2() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	go doSomeTask(ctx)
	<-time.After(5 * time.Second)
	cancel()

	<-time.After(3 * time.Second)
	fmt.Println("exit main")
}

func main() {
	Steps1()
	Steps2()
}
