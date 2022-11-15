package main

import (
	"fmt"
	"sync"
)

// goroutine 的并发安全问题
func main() {
	sum := 0

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			sum++
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			sum++
		}
	}()

	wg.Wait()

	fmt.Println(sum)
}
