# Goroutine

`Goroutine` 是 `Go` 语言中轻量级的并发处理方式之一。它可以看作是一个轻量级线程，一个程序可以包含成百上千个 `Goroutine`。`Goroutine` 的启动非常快，只需要几纳秒的时间，而且 `Goroutine` 的调度是由 `Go` 运行时系统自动完成的，开发者不需要手动进行线程调度。

## 目录

- `Goroutine` 基础用法
- `Goroutine` 中 `sync.WaitGroup` 的使用
- `Goroutine` 小实验: 并发的下载图片
- `Goroutine` 的并发安全问题

## Goroutine基础

`golang`中想要并发的执行一短逻辑可以通过`go func() `实现。

```go
go func() {
    // goroutine 执行的代码
}()
```

一个`go func()`会启动一个后台并发任务, 大概流程是通过`go`关键字将这个`func()`打包成一个任务，然后提交给`golang`的并发调度器，并发调度器会根据一定策略来执行这些任务。

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

再上一小节中通过`<-time.After(time.Second * 10)`来等待`goroutine`执行完成, 这是非常难以控制的。

在真实的场景中我们并不那么容易知道一个`Goroutine`什么时候执行完成, 我们需要一种更简单的方式来等待`Goroutine`的结束。

`sync.WaitGroup` 是 `Go` 语言中用于并发控制的一个结构体，它可以用于等待一组 `Goroutine` 的完成。

`WaitGroup` 包含三个方法：

1. `Add(delta int)`：向 `WaitGroup` 中添加 `delta` 个等待的 `Goroutine`。
2. `Done()`：表示一个等待的 `Goroutine` 已经完成了，向 `WaitGroup` 中减少一个等待的 `Goroutine`。
3. `Wait()`：等待所有添加到 `WaitGroup` 中的 `Goroutine` 都完成。

使用 `WaitGroup` 进行并发控制的基本流程如下：

1. 创建 `WaitGroup` 对象 `wg`。
2. 启动多个 `Goroutine`，在每个 `Goroutine` 的开始处调用 `wg.Add(1)` 将等待的 `Goroutine` 数量加 1。
3. 在每个 `Goroutine` 中进行任务处理，当任务处理完毕后，在 `Goroutine` 的结束处调用 `wg.Done()` 将已完成的 `Goroutine` 数量减 1。
4. 在主 `Goroutine` 中调用 `wg.Wait()` 等待所有的 `Goroutine` 完成任务。

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

## 并发下载图片小练习

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

`Goroutine` 的出现使得 `Go` 语言可以更加方便地进行并发编程。但是在使用 `Goroutine` 时需要注意避免资源竞争和死锁等问题。

当多个`goroutine`并发修改同一个变量有可能会产生并发安全问题导致结果错误，因为修改可能是非原子的。这种情况可以将修改变成原子操作(`atomic`)或通过加锁保护(`sync.Mutex`, `sync.RWMutex`)，让修改的步骤串行防止并发安全问题。

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
  // var rmu sync.RWMutex

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

`Mutex` 和 `RWMutex` 都是 `Go` 语言中的并发控制机制，它们都可以用于保护共享资源，避免并发访问导致的数据竞争和不一致性。

`Mutex` 是最简单的并发控制机制，它提供了两个方法：

1. `Lock()`：获取互斥锁，如果互斥锁已经被其他 `Goroutine` 获取，则当前 `Goroutine` 会阻塞等待。
2. `Unlock()`：释放互斥锁，如果当前 `Goroutine` 没有获取互斥锁，则会引发运行时 `panic`。(必须先`Lock`, 在`Unlock`)

`Mutex` 适用于对共享资源的互斥访问，即同一时间只能有一个 `Goroutine` 访问共享资源的情况。

`RWMutex` 是在 `Mutex` 的基础上进行了扩展，它允许多个 `Goroutine` 同时读取共享资源，但只允许一个 `Goroutine` 写共享资源。`RWMutex` 提供了三个方法：

1. `RLock()`：获取读锁，允许多个 `Goroutine` 同时获取读锁。
2. `RUnlock()`：释放读锁。
3. `Lock()`：获取写锁，只允许一个 `Goroutine` 获取写锁。
4. `Unlock()`：释放互斥锁。

`RWMutex` 适用于读写分离的场景，可以提高共享资源的并发读取性能。

## 思考题



## 参考

https://blog.boot.dev/golang/gos-waitgroup-javascripts-promiseall/

https://gfw.go101.org/article/control-flows-more.html

https://larrylu.blog/race-condition-in-golang-c49a6e242259