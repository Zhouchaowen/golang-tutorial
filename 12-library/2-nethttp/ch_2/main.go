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

	// [/hello] 访问 http://ip:port/hello 将匹配到hello函数
	http.HandleFunc("/hello", hello)

	// [/hi] 访问 http://ip:port/hi 将匹配到SayHi实例的ServeHTTP()方法上
	http.Handle("/hi", SayHi{})

	// 启动服务端, 监听的地址 IP:Port
	http.ListenAndServe(fmt.Sprintf("%s:%d", Ip, Port), nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// SayHi 实现 Handler 接口
type SayHi struct {
}

func (s SayHi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!")
}
