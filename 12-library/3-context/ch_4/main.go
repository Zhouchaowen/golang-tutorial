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
				fmt.Println("ctx value", ctx.Value("golang"), ",generate a number:", n)
				time.Sleep(time.Second * 3)
			}
			n++
		}
	}()
}

func main() {
	base, cancel := context.WithCancel(context.Background())
	ctx := context.WithValue(base, "golang", "tutorial")
	go doSomeTask(ctx)

	<-time.After(7 * time.Second)
	cancel()

	<-time.After(10 * time.Second)
	fmt.Println("exit main")
}
