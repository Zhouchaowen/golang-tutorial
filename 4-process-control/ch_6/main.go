package main

import (
	"fmt"
	"os"
)

/*
	1.defer+exit()
*/

// 注意：如果函数因为执行了os.Exit而退出，而不是正常return退出或者panic退出，那程序会立即停止，被defer的函数调用不会执行。

func Steps1() {
	fmt.Println("Steps1")
}

func main() {
	fmt.Println("start")
	defer Steps1()
	fmt.Println("stop")
	os.Exit(0)
}
