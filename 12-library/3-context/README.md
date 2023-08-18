# Context

在`Go`语言中上下文 `Context`是用来实现多函数中传递相关值、设置请求截止日期、进行超时控制等的重要接口，它与` Goroutine` 密切协作实现对函数调用链路的控制。

`Context`接口定义了四个需要实现的方法，其中包括：

- `Deadline` 返回被取消的时间，也就是完成工作的截止日期；

- `Done` 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 `Done` 方法会返回同一个 Channel；

- `Err` 返回Context结束的原因，它只会在 `Done` 方法对应的 Channel 关闭时返回非空的值；

  - 如果Context被取消，会返回 `context canceled` 错误；

  - 如果Context超时，会返回 `context deadline exceeded` 错误；

-  `Value` 从Context中获取键对应的值，对于同一个上下文来说，多次调用 `Value` 并传入相同的 `Key` 会返回相同的结果，该方法可以用来传递请求特定的数据；

```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

包中提供的函数会返回实现该接口的私有结构体，我们会在后面详细介绍它们的工作原理。

## WithCancel 



##WithDeadline



##WithTimeout



##WithValue




## 参考
https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/

https://www.lixueduan.com/posts/go/context/

https://zhuanlan.zhihu.com/p/34417106

https://learnku.com/articles/61838

https://taoshu.in/go/go-http-server-timeout.html

https://xiaorui.cc/archives/7131

https://colobu.com/2016/07/01/the-complete-guide-to-golang-net-http-timeouts/

https://zhuanlan.zhihu.com/p/612406500

