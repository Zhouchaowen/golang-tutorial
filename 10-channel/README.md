# Channel



## 目录

- channel 定义
- channel 阻塞发送 

- 非缓冲和缓冲 channel 的对比
- 关闭 channel 
- 遍历 channel
- channel+select 控制 goroutine 退出

## 定义Channel
channel 是解决 goroutine 的同步问题以及 goroutine 之间数据共享（数据传递）的问题。

```go
package main

import (
	"fmt"
	"time"
)

// 定义 channel, channel 是带有类型的管道，可以通过信道操作符 <- 来发送或者接收值
func main() {
	// 信道在使用前必须通过内建函数 make 来创建

	// make(chan T,size)  标识用内建函数 make 来创建 一个T类型的缓冲大小为 size 的 channel
	// 如下: make(chan int) 用内建函数 make 来创建 一个 int 类型的缓冲大小为 0 的 channel
	c := make(chan int)

	go func() {
		// 从 c 接收值并赋予 num
		num := <-c
		fmt.Printf("recover:%d\n", num)
	}()

	// 将 1 发送至信道 c
	c <- 1

	<-time.After(time.Second * 3)

	fmt.Println("return")
}
```

## 阻塞Channel

golang中无缓存channel的收发都需要对端(发送端，接收端)准备好，如果有一短未准备好接收，则另一端将会阻塞。

```go
package main

import (
	"fmt"
	"time"
)

// 发送端和接收端的阻塞问题
// 发送端在没有准备好之前会阻塞,同样接收端在发送端没有准备好之前会阻塞
func main() {
	c := make(chan string)

	go func() {
		<-time.After(time.Second * 10)
		fmt.Println("发送端准备好了 send: ping")
		c <- "ping" // 发送
	}()

	// 发送端10s后才准备好，所以阻塞在当前位置
	fmt.Println("阻塞在当前位置，发送端发送数据后才继续执行")
	num := <-c
	fmt.Printf("recover: %s\n", num)
}
```

小练习：通过goroutine+channel计算数组之和。

```go
package main

import "fmt"

// 对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。

// sum 求和函数
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
```

## 缓冲Channel

缓冲channel通过make(chan T, size)定义。缓冲channel可以储存一定的数据，当数据未达到上限时写入端可以非阻塞的一直写；当到达上限后，写入端会阻塞。当缓冲区有数据时，读取端可以一直读；没有数据时，读取端将阻塞。

```go
package main

import "fmt"

// 不带缓冲的 channel
func NoBufferChan() {
	ch := make(chan int)
	ch <- 1 // 被阻塞,执行报错 fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
}

// 带缓冲的 channel
func BufferChan() {
	// channel 有缓冲、是非阻塞的，直到写满 cap 个元素后才阻塞
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)
}

func main() {
	//NoBufferChan()
	BufferChan()
}
```

## 关闭channel

channle可以通过close()函数关闭，关闭一个channel后，可以从中读取数据不过读取的数据全是当前channel类型的零值，但不能向这个channel写入数据会发送panic。

```go
package main

func main() {
	ch := make(chan bool)
	close(ch)
	fmt.Println(<- ch)
	//ch <- true // panic: send on closed channel
}
```

| 操作     | 一个零值nil通道 | 一个非零值但已关闭的通道 | 一个非零值且尚未关闭的通道 |
| -------- | --------------- | ------------------------ | -------------------------- |
| 关闭     | 产生恐慌        | 产生恐慌                 | 成功关闭                   |
| 发送数据 | 永久阻塞        | 产生恐慌                 | 阻塞或者成功发送           |
| 接收数据 | 永久阻塞        | 永不阻塞                 | 阻塞或者成功接收           |

## 持续读取Channel

可以通过range持续读取channel，直到channel关闭。

```go
package main

import (
	"fmt"
	"time"
)

// 通过 range 遍历 channel, 并通过关闭 channel 来退出循环

// 复制一个 channel 或用于函数参数传递时, 只是拷贝了一个 channel 的引用, 因此调用者和被调用者将引用同一个channel对象
func genNum(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	// 发送者可通过 close 关闭一个信道来表示没有需要发送的值了
	close(c)
}

func main() {
	c := make(chan int, 10)
	go genNum(c)

	// 循环 for v := range c 会不断从信道接收值，直到它被关闭
	// 并且只有发送者才能关闭信道，而接收者不能, 向一个已经关闭的信道发送数据会引发程序恐慌（panic）
	for v := range c {
		fmt.Println("receive:", v)
	}

	// 接收者可以通过 v,ok := <- c 表达式接收第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭,那么 v 为对应类型零值,ok 为 false
	v, ok := <-c
	fmt.Printf("value:%d, ok:%t\n", v, ok)

	fmt.Println("close")
}
```

## 通过select操作channel

通过`select-case`可以选择一个准备好数据`channel`执行，会从这个`channel`中读取或写入数据。

```go
package main

import (
	"fmt"
	"time"
)

// 通过 channel+select 控制 goroutine 退出
func genNum(c, quit chan int) {
	for i := 0; ; i++ {
		// select 可以等待多个通信操作
		// select 会阻塞等待可执行分支。当多个分支都准备好时会随机选择一个执行。
		select {
		case <-quit:
			// 发送者可通过 close 关闭一个信道来表示没有需要发送的值了。
			close(c)
			return
		default: // 等同于 switch 的 default。当所以case都阻塞时如果有default则，执行default
			c <- i
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go genNum(c, quit)

	// 循环 for v := range c 会不断从信道接收值，直到它被关闭
	// 并且只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。
	for i := 0; i < 10; i++ {
		fmt.Println("receive:", <-c)
	}

	// 通知 genNum() 退出
	quit <- 1

	// 接收者可以通过 v,ok := <- c 表达式第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完
	v, ok := <-c
	fmt.Printf("value:%d, ok:%t\n", v, ok)

	fmt.Println("close")
}
```

## 思考题



## 参考



