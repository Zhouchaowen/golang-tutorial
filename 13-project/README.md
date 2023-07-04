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



