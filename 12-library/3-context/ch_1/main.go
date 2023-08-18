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
				return // 不返回可能发生泄露
			default:
				fmt.Println("generate a number:", n)
				time.Sleep(time.Second * 3)
			}
			n++
		}
	}()
}

func main() {
	// 取消上下文
	ctx, cancel := context.WithCancel(context.Background())
	go doSomeTask(ctx)

	// 10s后调用取消函数
	<-time.After(10 * time.Second)
	cancel()

	<-time.After(3 * time.Second)
	fmt.Println("exit main")
}
