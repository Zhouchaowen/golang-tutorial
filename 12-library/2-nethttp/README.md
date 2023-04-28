# Net/Http

在`Go`语言中,  `net/http` 包是用于构建HTTP服务器和客户端的标准库。它提供了一组功能强大的函数，使开发者可以轻松地编写可靠、高效的Web应用程序。

```go
package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	Ip   string
	Port int
)

func init() {
	flag.StringVar(&Ip, "i", "0.0.0.0", "server ip")
	flag.IntVar(&Port, "p", 8080, "server port")
}

func main() {
	flag.Parse()

	// HandleFunc 函数用于处理匹配成功的每个请求
	// [/] 匹配所有路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Golang Tutorial!")
	})

	// 启动服务端, 监听的地址 Ip:Port
	http.ListenAndServe(fmt.Sprintf("%s:%d", Ip, Port), nil)
}
```

在上面的例子中，我们创建一个HTTP服务器, 使用 `http.HandleFunc` 函数来处理HTTP请求。这个函数接受一个字符串和匿名函数作为参数，这个函数将在每个HTTP请求到达服务器时被调用(如果第一个参数和访问的路由匹配的话)。

```go
func(w http.ResponseWriter, r *http.Request) {
    // 处理HTTP请求
}
```

这个匿名函数接受两个参数，`http.ResponseWriter` 和 `*http.Request`。`http.ResponseWriter` 用于向客户端发送HTTP响应，而 `*http.Request` 用于接收来自客户端的HTTP请求的所有信息。

除了传递匿名函数外还可以单独定义有名函数来传递(匿名函数不适合写太多逻辑，而单独的函数则不存在这个问题), 如下代码中`http.HandleFunc("/hello", hello)`,路由`/hello`匹配到`func hello(w http.ResponseWriter, r *http.Request)`函数。

```go
package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	Ip   string
	Port int
)

func init() {
	flag.StringVar(&Ip, "i", "0.0.0.0", "server ip")
	flag.IntVar(&Port, "p", 8080, "server port")
}

func main() {
	flag.Parse()

	// [/hello] 访问 http://ip:port/hello 将匹配到hello函数
	http.HandleFunc("/hello", hello)

	// 启动服务端, 监听的地址 IP:Port
	http.ListenAndServe(fmt.Sprintf("%s:%d", Ip, Port), nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
```

