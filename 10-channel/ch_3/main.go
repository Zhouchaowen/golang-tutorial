package main

import "fmt"

// 对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。
func sum(s []int, c chan int) {
	ans := 0
	for _, v := range s {
		ans += v
	}
	c <- ans // 将和送入 c
}

func main() {
	s := []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
