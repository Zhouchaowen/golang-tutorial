# 工程实践

[TOC]

## 依赖注入

```go
type Server struct {
  Conf *Config // 依赖 Config
}

func NewServer() *Server {
  return &Server{
    Conf: config.New()
  }
}
```

上面的例子中，`Server` 依赖 `Config`，在Server的初始化中直接使用了config.New方法；如果后面config.New增加了参数，那我们还需要到Server的初始化方法中修改Conf。 具体说来，上述这段的设计中存在几个明显的问题：

- 耦合性高：Server 承担了 Config 的初始化，如果 Config的初始化过程发生改变，我们需要一并修改Server；
- 扩展性差：一个Server的实例化过程只能构造自己的Config，它不能使用其它已经实例化好的Config；
- 不利于单元测试：在完成单元测试的时候，我们无法测试不同类型的Config对于Server的影响。

为此我们可以修改成下面这样：

```go
type Server struct {
  Conf *Config // 依赖 Config
}

func NewServer(c *Config) *Server {
  return &Server{
    Conf: c
  }
}
```

这样，我们可以在外部对Config进行初始化然后注入到Server中。上述修改完成后，耦合性和扩展性的问题便解决掉了，对于Server来说，接收从外部传入的Config，其不关心Config的初始化的过程；另外，多个Server之间可共用一份Config，并且可做到随意替换；此外，测试的时候我们可以较为轻松对关键部件进行替换。

## 结构体初始化

使用结构体文字初始化以避免无效的中间状态。尽可能内联结构声明。

```go
// 不推荐的方式
foo, err := newFoo(
    *fooKey,
    bar,
    100 * time.Millisecond,
    nil,
)
if err != nil {
    log.Fatal(err)
}
defer foo.close()


// 不推荐的方式
cfg := fooConfig{}
cfg.Bar = bar
cfg.Period = 100 * time.Millisecond
cfg.Output = nil

foo, err := newFoo(*fooKey, cfg)
if err != nil {
    log.Fatal(err)
}
defer foo.close()
```

```go
// 推荐的方式
cfg := fooConfig{
    Bar:    bar,
    Period: 100 * time.Millisecond,
    Output: nil,
}

foo, err := newFoo(*fooKey, cfg)
if err != nil {
    log.Fatal(err)
}
defer foo.close()

// 推荐的方式
foo, err := newFoo(*fooKey, fooConfig{
    Bar:    bar,
    Period: 100 * time.Millisecond,
    Output: nil,
})
if err != nil {
    log.Fatal(err)
}
defer foo.close()
```

## 默认值设置

通过默认的无操作实现避免 nil 检查，让零值变得有用，尤其是在配置对象中。

```go
func (f *foo) process() {
    if f.Output != nil {  // 不要在使用的时候判断
        fmt.Fprintf(f.Output, "start\n")
    }
    // ...
}
```

```go
func newFoo(..., cfg fooConfig) *foo {
    if cfg.Output == nil { // 在初始化的时候填上默认值
        cfg.Output = ioutil.Discard
    }
    // ...
}

func (f *foo) process() {
     fmt.Fprintf(f.Output, "start\n")
     // ...
}
```

## 明确依赖关系

```go
func (f *foo) process() {
    fmt.Fprintf(f.Output, "start\n")
    result := f.Bar.compute()
    log.Printf("bar: %v", result) // Whoops! 隐式依赖全局日志操作
    // ...
}
```

```go
func (f *foo) process() {
    fmt.Fprintf(f.Output, "start\n")
    result := f.Bar.compute()
    f.Logger.Printf("bar: %v", result) // Better. 变成显示依赖
    // ...
}
```

## 依赖接口不依赖实现

```go
// 版本一
func foo() {
    resp, err := http.Get("http://zombo.com")
    // ...
}
```

```go
// 版本二
func foo(client *http.Client) {
    resp, err := client.Get("http://zombo.com")
    // ...
}
```

```go
// 版本三
type Doer interface {
    Do(*http.Request) (*http.Response, error)
}

func foo(d Doer) {
    req, _ := http.NewRequest("GET", "http://zombo.com", nil)
    resp, err := d.Do(req)
    // ...
}
```

## 尽早 return

这是一个来自 `bytes` 包的例子:

```go
func (b *Buffer) UnreadRune() error {
	if b.lastRead <= opInvalid {
		return errors.New("bytes.Buffer: UnreadRune: previous operation was not a successful ReadRune")
	}
	if b.off >= int(b.lastRead) {
		b.off -= int(b.lastRead)
	}
	b.lastRead = opInvalid
	return nil
}
```

进入 `UnreadRune` 后，将检查 `b.lastRead` 的状态，如果之前的操作不是 `ReadRune`，则会立即返回错误。 之后，函数的其余部分继续进行 `b.lastRead` 大于 `opInvalid` 的断言。

```go
func (b *Buffer) UnreadRune() error {
	if b.lastRead > opInvalid {
		if b.off >= int(b.lastRead) {
			b.off -= int(b.lastRead)
		}
		b.lastRead = opInvalid
		return nil
	}
	return errors.New("bytes.Buffer: UnreadRune: previous operation was not a successful ReadRune")
}
```

最常见的执行成功的情况是嵌套在第一个if条件内，成功的退出条件是 `return nil`，而且必须通过仔细匹配大括号来发现。 函数的最后一行是返回一个错误，并且被调用者必须追溯到匹配的左括号，以了解何时执行到此点。

对于读者和维护程序员来说，这更容易出错，因此 Go 语言更喜欢使用 `guard clauses` 并尽早返回错误。

## 设计难以被误用的 API

警惕采用几个相同类型参数的函数

简单, 但难以正确使用的 API 是采用两个或更多相同类型参数的 API。 让我们比较两个函数签名：

```go
func Max(a, b int) int
func CopyFile(to, from string) error
```

这两个函数有什么区别？ 显然，一个返回两个数字最大的那个，另一个是复制文件，但这不重要。

```go
Max(8, 10) // 10
Max(10, 8) // 10
```

`Max` 是可交换的; 参数的顺序无关紧要。 无论是 8 比 10 还是 10 比 8，最大的都是 10。

但是，却不适用于 `CopyFile`。

```go
CopyFile("/tmp/backup", "presentation.md")
CopyFile("presentation.md", "/tmp/backup")
```

这些声明中哪一个备份了 `presentation.md`，哪一个用上周的版本覆盖了 `presentation.md`？ 没有文档，你无法分辨。 如果没有查阅文档，代码审查员也无法知道你写对了顺序。

一种可能的解决方案是引入一个 `helper` 类型，它会负责如何正确地调用 `CopyFile`。

```go
type Source string

func (src Source) CopyTo(dest string) error {
	return CopyFile(dest, string(src))
}

func main() {
	var from Source = "presentation.md"
	from.CopyTo("/tmp/backup")
}
```

通过这种方式，`CopyFile` 总是能被正确调用 - 还可以通过单元测试 - 并且可以被设置为私有，进一步降低了误用的可能性。

## 让函数定义它们所需的行为

假设我需要编写一个将 `Document` 结构保存到磁盘的函数的任务。

```go
// Save writes the contents of doc to the file f.
func Save(f *os.File, doc *Document) error
```

我可以指定这个函数 `Save`，它将 `*os.File` 作为写入 `Document` 的目标。但这样做会有一些问题

`Save` 的签名排除了将数据写入网络位置的选项。假设网络存储可能在以后成为需求，则此功能的签名必须改变，从而影响其所有调用者。

`Save` 测试起来也很麻烦，因为它直接操作磁盘上的文件。因此，为了验证其操作，测试时必须在写入文件后再读取该文件的内容。

而且我必须确保 `f` 被写入临时位置并且随后要将其删除。

`*os.File` 还定义了许多与 `Save` 无关的方法，比如读取目录并检查路径是否是符号链接。 如果 `Save` 函数的签名只用 `*os.File` 的相关内容，那将会很有用。

我们能做什么 ？

```go
// Save writes the contents of doc to the supplied
// ReadWriterCloser.
func Save(rwc io.ReadWriteCloser, doc *Document) error
```

使用 `io.ReadWriteCloser`，我们可以应用[接口隔离原则](https://zh.wikipedia.org/wiki/接口隔离原则)来重新定义 `Save` 以获取更通用文件形式。

通过此更改，任何实现 `io.ReadWriteCloser` 接口的类型都可以替换以前的 `*os.File`。

这使 `Save` 在其应用程序中更广泛，并向 `Save` 的调用者阐明 `*os.File` 类型的哪些方法与其操作有关。

而且，`Save` 的作者也不可以在 `*os.File` 上调用那些不相关的方法，因为它隐藏在 `io.ReadWriteCloser` 接口后面。

但我们可以进一步采用[接口隔离原则](https://zh.wikipedia.org/wiki/接口隔离原则)。

首先，如果 `Save` 遵循[单一功能原则](https://zh.wikipedia.org/wiki/单一功能原则)，它不可能读取它刚刚写入的文件来验证其内容 - 这应该是另一段代码的功能。

```go
// Save writes the contents of doc to the supplied
// WriteCloser.
func Save(wc io.WriteCloser, doc *Document) error
```

因此，我们可以将我们传递给 `Save` 的接口的规范缩小到只写和关闭。

其次，通过向 `Save` 提供一个关闭其流的机制，使其看起来仍然像一个文件，这就提出了在什么情况下关闭 `wc` 的问题。

可能 `Save` 会无条件地调用 `Close`，或者在成功的情况下调用 `Close`。

这给 `Save` 的调用者带来了问题，因为它可能希望在写入文档后将其他数据写入流。

```go
// Save writes the contents of doc to the supplied
// Writer.
func Save(w io.Writer, doc *Document) error
```

一个更好的解决方案是重新定义 `Save` 仅使用 `io.Writer`，它只负责将数据写入流。

将[接口隔离原则](https://zh.wikipedia.org/wiki/接口隔离原则)应用于我们的 `Save` 功能，同时, 就需求而言, 得出了最具体的一个函数 - 它只需要一个可写的东西 - 并且它的功能最通用，现在我们可以使用 `Save` 将我们的数据保存到实现 `io.Writer` 的任何事物中。

## 将并发性留给调用者

```go
// ListDirectory returns the contents of dir.
func ListDirectory(dir string) ([]string, error)
```

```go
// ListDirectory returns a channel over which
// directory entries will be published. When the list
// of entries is exhausted, the channel will be closed.
func ListDirectory(dir string) chan string
```

```go
// ListDirectory returns a channel over which
// directory entries will be published. When the list
// of entries is exhausted, the channel will be closed.
func ListDirectory(dir string) chan string
```

## 永远不要启动一个停止不了的 goroutine。

```go
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	go http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux) // debug
	http.ListenAndServe("0.0.0.0:8080", mux)                       // app traffic
}
```

```go
func serveApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}

func serveDebug() {
	if err := http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func main() {
	go serveDebug()
	go serveApp()
	select {}
}
```

```go
func serveApp() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	return http.ListenAndServe("0.0.0.0:8080", mux)
}

func serveDebug() error {
	return http.ListenAndServe("127.0.0.1:8001", http.DefaultServeMux)
}

func main() {
	done := make(chan error, 2)
	go func() {
		done <- serveDebug()
	}()
	go func() {
		done <- serveApp()
	}()

	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error: %v", err)
		}
	}
}
```

```go
func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop // wait for stop signal
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func serveApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello, QCon!")
	})
	return serve("0.0.0.0:8080", mux, stop)
}

func serveDebug(stop <-chan struct{}) error {
	return serve("127.0.0.1:8001", http.DefaultServeMux, stop)
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		done <- serveDebug(stop)
	}()
	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error: %v", err)
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
```

## 切片大小预分配

过去，我曾使用`make(a, 10)`预分配数组空间，但是，

我经常习惯性地误用`append()`方法，导致数组中出现了许多前导零

为了避免这种情况，我改用了一种更有效率的预分配方法：`make(a, 0, 10)`

```go
func main() {
	a := make([]int, 5)
	a = append(a, 1) // [0,0,0,0,0,1]

	b := make([]int, 0, 5)
	b = append(b, 1) // [1]
}
```



https://github.com/llitfkitfk/go-best-practice