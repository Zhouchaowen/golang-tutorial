1. 依赖注入

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

2. 使用结构体文字初始化以避免无效的中间状态。尽可能内联结构声明。

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

3. 通过默认的无操作实现避免 nil 检查，让零值变得有用，尤其是在配置对象中。

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

4. 明确依赖关系

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

5. 依赖测试

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

6. [包名的定义(准确，尽量不要更改)](https://dave.cheney.net/2016/08/20/solid-go-design)

遵循的经验法则不是“我应该在这个包中放入什么类型的？”。相反，我要问是“该包提供的服务是什么？”通常这个问题的答案不是“这个包提供 `X` 类型”，而是“这个包提供 `HTTP`”。

在项目中，每个包名称应该是唯一的。包的名称应该描述其目的的建议很容易理解 - 如果你发现有两个包需要用相同名称，它可能是:

1. 包的名称太通用了。
2. 该包与另一个类似名称的包重叠了。在这种情况下，您应该检查你的设计，或考虑合并包。

- 正确的包名

在 Go 中，所有代码都位于包内，设计良好的包以其名称开头。包的名称既是其用途的描述，也是名称空间前缀。Go 标准库中的一些优秀包的示例：

- `net/http`，它提供http客户端和服务器。
- `os/exec`，它运行外部命令。
- `encoding/json`，它实现了JSON文档的编码和解码。

当您在自己的包中使用另一个包的符号时，这是通过“import”声明来完成的，它在两个包之间建立了源级耦合。他们现在互相了解。

- 错误的包名

这种对名字的关注不仅仅是迂腐。一个命名不当的包会错过枚举其用途的机会，即使它确实有过用途。

提供什么`package server`？……好吧，希望是一个服务器，但是哪个协议呢？

提供什么`package private`？我不应该看到的东西？它应该有任何公共符号吗？

**而且`package common`，就像它的犯罪同伙一样，`package utils`经常被发现与其他犯罪者很接近。**

捕获所有像这样的包成为杂项的垃圾场，并且因为它们有很多责任，所以它们经常无缘无故地更改。

7. 尽早 `return` 而不是深度嵌套

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

8. 设计难以被误用的 API（API 应该易于使用且难以被误用）

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

9.让函数定义它们所需的行为

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

10.当遇到难以忍受的错误处理时，请尝试将某些操作提取到辅助程序类型中

如：bufio.Scanner

```go
func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}
```

```go
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0

	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}
```

在处理底层网络协议时，有必要使用 `I/O` 原始的错误处理来直接构建响应，这样就可能会变得重复。看一下构建 `HTTP` 响应的 `HTTP` 服务器的这个片段。

```go
type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}

	for _, h := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
		if err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, "\r\n"); err != nil {
		return err
	}

	_, err = io.Copy(w, body)
	return err
}
```

首先，我们使用 `fmt.Fprintf` 构造状态码并检查错误。 然后对于每个标题，我们写入键值对，每次都检查错误。 最后，我们使用额外的 `\r\n` 终止标题部分，检查错误之后将响应主体复制到客户端。 最后，虽然我们不需要检查 `io.Copy` 中的错误，但我们需要将 `io.Copy` 返回的两个返回值形式转换为 `WriteResponse` 的单个返回值。

这里很多重复性的工作。 我们可以通过引入一个包装器类型 `errWriter` 来使其更容易。

`errWriter` 实现 `io.Writer` 接口，因此可用于包装现有的 `io.Writer`。 `errWriter` 写入传递给其底层 `writer`，直到检测到错误。 从此时起，它会丢弃任何写入并返回先前的错误。

```go
type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, nil
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprint(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}
```

将 `errWriter` 应用于 `WriteResponse` 可以显着提高代码的清晰度。 每个操作不再需要自己做错误检查。 通过检查 `ew.err` 字段，将错误报告移动到函数末尾，从而避免转换从 `io.Copy` 的两个返回值。

11.错误只处理一次

最后，我想提一下你应该只处理错误一次。 处理错误意味着检查错误值并做出单一决定。

```
// WriteAll writes the contents of buf to the supplied writer.
func WriteAll(w io.Writer, buf []byte) {
        w.Write(buf)
}
```

如果你做出的决定少于一个，则忽略该错误。 正如我们在这里看到的那样， `w.WriteAll` 的错误被丢弃。

但是，针对单个错误做出多个决策也是有问题的。 以下是我经常遇到的代码。

```
func WriteAll(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		log.Println("unable to write:", err) // annotated error goes to log file
		return err                           // unannotated error returned to caller
	}
	return nil
}
```

在此示例中，如果在 `w.Write` 期间发生错误，则会写入日志文件，注明错误发生的文件与行数，并且错误也会返回给调用者，调用者可能会记录该错误并将其返回到上一级，一直回到程序的顶部。

调用者可能正在做同样的事情

```
func WriteConfig(w io.Writer, conf *Config) error {
	buf, err := json.Marshal(conf)
	if err != nil {
		log.Printf("could not marshal config: %v", err)
		return err
	}
	if err := WriteAll(w, buf); err != nil {
		log.Println("could not write config: %v", err)
		return err
	}
	return nil
}
```

因此你在日志文件中得到一堆重复的内容，

```
unable to write: io.EOF
could not write config: io.EOF
```

但在程序的顶部，虽然得到了原始错误，但没有相关内容。

```
err := WriteConfig(f, &conf)
fmt.Println(err) // io.EOF
```

我想深入研究这一点，因为作为个人偏好, 我并没有看到 `logging` 和返回的问题。

```
func WriteConfig(w io.Writer, conf *Config) error {
	buf, err := json.Marshal(conf)
	if err != nil {
		log.Printf("could not marshal config: %v", err)
		// oops, forgot to return
	}
	if err := WriteAll(w, buf); err != nil {
		log.Println("could not write config: %v", err)
		return err
	}
	return nil
}
```

很多问题是程序员忘记从错误中返回。正如我们之前谈到的那样，Go 语言风格是使用 `guard clauses` 以及检查前提条件作为函数进展并提前返回。

在这个例子中，作者检查了错误，记录了它，但忘了返回。这就引起了一个微妙的错误。

Go 语言中的错误处理规定，如果出现错误，你不能对其他返回值的内容做出任何假设。由于 `JSON` 解析失败，`buf` 的内容未知，可能它什么都没有，但更糟的是它可能包含解析的 `JSON` 片段部分。

由于程序员在检查并记录错误后忘记返回，因此损坏的缓冲区将传递给 `WriteAll`，这可能会成功，因此配置文件将被错误地写入。但是，该函数会正常返回，并且发生问题的唯一日志行是有关 `JSON` 解析错误，而与写入配置失败有关。

**为错误添加相关内容**

发生错误的原因是作者试图在错误消息中添加 `context` 。 他们试图给自己留下一些线索，指出错误的根源。

让我们看看使用 `fmt.Errorf` 的另一种方式。

```
func WriteConfig(w io.Writer, conf *Config) error {
	buf, err := json.Marshal(conf)
	if err != nil {
		return fmt.Errorf("could not marshal config: %v", err)
	}
	if err := WriteAll(w, buf); err != nil {
		return fmt.Errorf("could not write config: %v", err)
	}
	return nil
}

func WriteAll(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		return fmt.Errorf("write failed: %v", err)
	}
	return nil
}
```

通过将注释与返回的错误组合起来，就更难以忘记错误的返回来避免意外继续。

如果写入文件时发生 `I/O` 错误，则 `error` 的 `Error()` 方法会报告以下类似的内容;

```
could not write config: write failed: input/output error
```

**使用 `github.com/pkg/errors` 包装 `errors**`

`fmt.Errorf` 模式适用于注释错误 `message`，但这样做的代价是模糊了原始错误的类型。 我认为将错误视为不透明值对于松散耦合的软件非常重要，因此如果你使用错误值做的唯一事情是原始错误的类型应该无关紧要的面孔

1. 检查它是否为 `nil`。
2. 输出或记录它。

但是在某些情况下，我认为它们并不常见，您需要恢复原始错误。 在这种情况下，使用类似我的 `errors` 包来注释这样的错误, 如下

```
func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.WithMessage(err, "could not read config")
}

func main() {
	_, err := ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

现在报告的错误就是 `K＆D` [[11\]](http://www.gopl.io/)样式错误，

```
could not read config: open failed: open /Users/dfc/.settings.xml: no such file or directory
```

并且错误值保留对原始原因的引用。

```
func main() {
	_, err := ReadConfig()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		os.Exit(1)
	}
}
```

因此，你可以恢复原始错误并打印堆栈跟踪;

```
original error: *os.PathError open /Users/dfc/.settings.xml: no such file or directory
stack trace:
open /Users/dfc/.settings.xml: no such file or directory
open failed
main.ReadFile
        /Users/dfc/devel/practical-go/src/errors/readfile2.go:16
main.ReadConfig
        /Users/dfc/devel/practical-go/src/errors/readfile2.go:29
main.main
        /Users/dfc/devel/practical-go/src/errors/readfile2.go:35
runtime.main
        /Users/dfc/go/src/runtime/proc.go:201
runtime.goexit
        /Users/dfc/go/src/runtime/asm_amd64.s:1333
could not read config
```

使用 `errors` 包，你可以以人和机器都可检查的方式向错误值添加上下文。 如果昨天你来听我的演讲，你会知道这个库在被移植到即将发布的 Go 语言版本的标准库中。

12.如果你的 `goroutine` 在得到另一个结果之前无法取得进展，那么让自己完成此工作而不是委托给其他 `goroutine` 会更简单。这通常会消除将结果从 `goroutine` 返回到其启动程序所需的大量状态跟踪和通道操作。

13.将并发性留给调用者

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

14.永远不要启动一个停止不了的 goroutine。

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





https://github.com/llitfkitfk/go-best-practice