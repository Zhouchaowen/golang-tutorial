// 来源：https://www.yuque.com/qyuhen/go/sfagor
package main

import "fmt"

// 工厂方法将goroutine和channel绑定

func newRecv[T any](cap int, fn func(T)) (data chan T, done chan struct{}) {
	data = make(chan T, cap)
	done = make(chan struct{})

	go func() {
		defer close(done)
		for v := range data {
			fn(v)
		}
	}()
	return
}

func main() {
	fn := func(i int) {
		fmt.Println(i)
	}
	data, done := newRecv[int](3, fn)
	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)

	<-done
}
