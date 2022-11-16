package main

import (
	"fmt"
	"sync"
)

// NoConcurrence 并发操作一个变量是不安全的，需要加锁
func NoConcurrence() {
	sum := 0

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ {
			sum++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ {
			sum++
		}
	}()

	wg.Wait()

	fmt.Println(sum)
}

func Concurrence() {
	sum := 0

	var wg sync.WaitGroup
	var mu sync.Mutex // 互斥锁（保护临界区，同一时刻只能有一个 goroutine 可以操作临界区）

	wg.Add(2) // 设置需要等待 goroutine 的数量,目前为2

	go func() {
		defer wg.Done() // 程序运行完毕, 将 goroutine 等待数量减1
		for i := 0; i < 10000000; i++ {
			mu.Lock() // 加锁保护临界区
			sum++
			mu.Unlock() // 操作完成解锁,临界区
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000000; i++ {
			mu.Lock() // 加锁保护临界区
			sum++
			mu.Unlock() // 操作完成解锁,临界区
		}
	}()

	wg.Wait()

	fmt.Println(sum)
}

// goroutine 的并发安全问题
func main() {
	NoConcurrence()
	Concurrence()
}
