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