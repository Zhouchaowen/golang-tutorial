// 来源：https://www.yuque.com/qyuhen/go/sfagor
package main

import (
	"fmt"
	"sync"
	"time"
)

// 通道实现信号量，在同一时刻指定goroutine工作的数量

type Sema struct {
	c chan struct{}
}

func NewSema(n int) *Sema {
	return &Sema{
		c: make(chan struct{}, n),
	}
}

func (m *Sema) Acquire() {
	m.c <- struct{}{}
}

func (m *Sema) Release() {
	<-m.c
}

func main() {
	var wg sync.WaitGroup
	sem := NewSema(2)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem.Acquire()
			defer sem.Release()
			for n := 0; n < 3; n++ {
				time.Sleep(2 * time.Second)
				fmt.Println(id, time.Now())
			}
		}(i)
	}

	wg.Wait()
}
