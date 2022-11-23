# Goroutine



## 目录

- ch_1 goroutine 基础用法
- ch_2 goroutine 中 sync.WaitGroup 的使用
- ch_3 goroutine 小实验: 并发的下载图片
- ch_4 goroutine 的并发安全问题
- ch_5 goroutine 小实验: 实现一个工作池



## Goroutine基础

`golang`中想要并发的执行一短逻辑可以通过`go func() `实现，一个`go func()`会启动一个后台并发任务。大概流程是通过`go`关键字将这个`func()`打包成一个任务，然后提交给`golang`的并发调度器，并发调度器会根据一定策略来执行这些任务。

```go
package main

import (
	"fmt"
	"time"
)

// 并发与并行：https://gfw.go101.org/article/control-flows-more.html

// 使用 goroutine 打印数据
func main() {
	language := []string{"golang", "java", "c++", "python", "rust", "js"}
	tutorial := []string{"入门", "初级", "中级", "高级", "专家"}

	// Go 程（goroutine）是由 Go 运行时管理的轻量级线程
	// 在函数调⽤语句前添加 go 关键字，就可创建一个 goroutine
	go listLanguage(language) // 通过goroutine启动该函数
	go listTutorial(tutorial)

	<-time.After(time.Second * 10) // 10s后执行下一行
	fmt.Println("return")
}

func listLanguage(items []string) {
	for i := range items {
		fmt.Printf("language: %s\n", items[i])
		time.Sleep(time.Second)
	}
}

func listTutorial(items []string) {
	for i := range items {
		fmt.Printf("tutorial: %s\n", items[i])
		time.Sleep(time.Second)
	}
}
```

## WaitGroup使用

再上一小节中通过`<-time.After(time.Second * 10)`来等待`goroutine`执行完成，这是非常难以控制的；在真实的场景中我们并不那么容易知道一个`goroutine`什么时候执行完成，我们需要一种更简单的方式来等待`goroutine`的结束。`sync.WaitGroup`可以等待一组`goroutine`执行完成

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func listLanguage(items []string, wg *sync.WaitGroup) { // 一般不建议这样使用
	defer wg.Done()

	for i := range items {
		fmt.Printf("language: %s\n", items[i])
		time.Sleep(time.Second)
	}
}

func listTutorial(items []string) {
	for i := range items {
		fmt.Printf("tutorial: %s\n", items[i])
		time.Sleep(time.Second)
	}
}

// 使用 WaitGroup等待goroutine执行完成
func main() {
	language := []string{"golang", "java", "c++", "python", "rust", "js"}
	tutorial := []string{"入门", "初级", "中级", "高级", "专家"}

	var wg sync.WaitGroup

	wg.Add(2) // 设置需要等待 goroutine 的数量,目前为2

	go listLanguage(language, &wg) // 通过 goroutine 启动该函数

	go func() { // 建议使用方式
		defer wg.Done() // 程序运行完毕, 将等待数量减1
		listTutorial(tutorial)
	}()

	wg.Wait() // 当等待数量为0后执行下一行
	//<-time.After(time.Second * 10) // 10s后执行下一行。 通过 wg.Wait() 代替
	fmt.Println("return")
}
```

并发下载图片小练习

```go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
)

func getImageData(url, name string) {
	resp, _ := http.Get(url) // 通过 http.get 请求读取 url 的数据

	// 创建一个缓存读取返回的 response 数据
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	dir, _ := os.Getwd()             // 获取当前执行程序目录
	fileName := path.Join(dir, name) // 拼接保存图片的文件地址

	// 将数据写到指定文件地址，权限为0666
	err := ioutil.WriteFile(fileName, buf.Bytes(), 0666)
	if err != nil {
		fmt.Printf("Save to file failed! %v", err)
	}
}

// 并发下载图片
func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(3)

	go func() {
		defer wg.Done()
		getImageData("https://img2.baidu.com/it/u=3125736368,3712453346&fm=253&fmt=auto&app=138&f=JPEG?w=800&h=500", "1.jpg")
	}()

	go func() {
		defer wg.Done()
		getImageData("https://img2.baidu.com/it/u=4284966505,4095784909&fm=253&fmt=auto&app=138&f=JPEG?w=640&h=400", "2.jpg")
	}()

	go func() {
		defer wg.Done()
		getImageData("https://img1.baidu.com/it/u=3580024761,2271795904&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=667", "3.jpg")
	}()
}
```

## Goroutine并发安全

`golang`中多个`goroutine`并发修改同一个变量有可能会产生并发安全问题导致结果错误，因为修改可能是非原子的。这种情况可以将修改变成原子操作(`atomic`)或通过加锁保护(`sync.Mutex`)，让修改的步骤串行防止并发安全问题。

```go
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
```

## 思考题



## 参考

https://blog.boot.dev/golang/gos-waitgroup-javascripts-promiseall/

https://gfw.go101.org/article/control-flows-more.html