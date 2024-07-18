package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

const (
	Ki = 1024
	Mi = Ki * Ki
	Gi = Ki * Mi
	Ti = Ki * Gi
	Pi = Ki * Ti
)

func block() {
	for {
		<-time.After(time.Second)
	}
}

// go tool pprof http://127.0.0.1:6060/debug/pprof/block
func main() {
	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	block()
}

/*
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/block
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/block
Type: delay
Time: Apr 23, 2023 at 9:31am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 4s, 100% of 4s total
      flat  flat%   sum%        cum   cum%
        4s   100%   100%         4s   100%  runtime.chanrecv1
         0     0%   100%         4s   100%  main.block (inline)
         0     0%   100%         4s   100%  main.main
         0     0%   100%         4s   100%  runtime.main
(pprof) list main.block
Total: 4s
ROUTINE ======================== main.block in /Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_6/main.go
         0         4s (flat, cum)   100% of Total
         .          .     15:	Pi = Ki * Ti
         .          .     16:)
         .          .     17:
         .          .     18:func block() {
         .          .     19:	for {
         .         4s     20:		<-time.After(time.Second)
         .          .     21:	}
         .          .     22:}
         .          .     23:
(pprof)
*/
