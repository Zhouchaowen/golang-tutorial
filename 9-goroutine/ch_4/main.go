package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	1.并发操作同一变量的安全问题
	2.sync.Mutex解决并发安全问题
*/

// NoConcurrence 并发操作一个变量是不安全的，需要加锁
func NoConcurrence() {
	sum := 0

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ { // sum做累加
			sum++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ { // sum做累加
			sum++
		}
	}()

	wg.Wait()

	fmt.Println(sum) // 结果应该等于20000000
}

func Concurrence() {
	var sum int64 = 0

	var wg sync.WaitGroup

	wg.Add(2) // 设置需要等待 goroutine 的数量,目前为2

	go func() {
		defer wg.Done() // 程序运行完毕, 将 goroutine 等待数量减1
		for i := 0; i < 10000000; i++ {
			atomic.AddInt64(&sum, 1) // 原子操作 +1
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ {
			atomic.AddInt64(&sum, 1) // 原子操作 +1
		}
	}()

	wg.Wait()

	fmt.Println(sum) // 结果应该等于20000000
}

// goroutine 的并发安全问题
func main() {
	NoConcurrence()
	Concurrence()
}
