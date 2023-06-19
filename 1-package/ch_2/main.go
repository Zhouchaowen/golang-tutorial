package main

/*
	1.导入标准库
	2.导入第三方库
*/

// 导入: 标准库和第三方库
import (
	"fmt"
	"math/rand" // 导入 rand 库
	"time"
)

// main 函数,程序的入口
func main() {
	// 设置随机数种子（暂不用了解）
	rand.Seed(time.Now().UnixNano())
	// rand.Intn(10) 函数返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
	fmt.Println("hello world", rand.Intn(10))
}
