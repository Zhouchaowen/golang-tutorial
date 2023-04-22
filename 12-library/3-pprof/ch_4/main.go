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

// pprof 分析goroutine泄露
func coroutinesLeaked() {
	for {
		for i := 0; i < 10; i++ {
			go func() {
				time.Sleep(30 * time.Second)
			}()
		}
		time.Sleep(time.Second)
	}
}

func main() {
	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	coroutinesLeaked()
}

/*
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/goroutine
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/goroutine
Type: goroutine
Time: Apr 22, 2023 at 5:28pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 293, 99.66% of 294 total
Showing top 10 nodes out of 28
      flat  flat%   sum%        cum   cum%
       292 99.32% 99.32%        292 99.32%  runtime.gopark
         1  0.34% 99.66%          1  0.34%  runtime/pprof.runtime_goroutineProfileWithLabels
         0     0% 99.66%          1  0.34%  internal/poll.(*FD).Accept
         0     0% 99.66%          1  0.34%  internal/poll.(*pollDesc).wait
         0     0% 99.66%          1  0.34%  internal/poll.(*pollDesc).waitRead (inline)
         0     0% 99.66%          1  0.34%  internal/poll.runtime_pollWait
         0     0% 99.66%          1  0.34%  main.coroutinesLeaked
         0     0% 99.66%        290 98.64%  main.coroutinesLeaked.func1
         0     0% 99.66%          1  0.34%  main.main
         0     0% 99.66%          1  0.34%  main.main.func1
(pprof) list main.coroutinesLeaked
Total: 294
ROUTINE ======================== main.coroutinesLeaked in /Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_3/main.go
         0          1 (flat, cum)  0.34% of Total
         .          .     20:		for i := 0; i < 10; i++ {
         .          .     21:			go func() {
         .          .     22:				time.Sleep(30 * time.Second)
         .          .     23:			}()
         .          .     24:		}
         .          1     25:		time.Sleep(time.Second)
         .          .     26:	}
         .          .     27:}
ROUTINE ======================== main.coroutinesLeaked.func1 in /Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_3/main.go
         0        290 (flat, cum) 98.64% of Total
         .          .     17:
         .          .     18:func coroutinesLeaked() {
         .          .     19:	for {
         .          .     20:		for i := 0; i < 10; i++ {
         .          .     21:			go func() {
         .        290     22:				time.Sleep(30 * time.Second)
         .          .     23:			}()
         .          .     24:		}
         .          .     25:		time.Sleep(time.Second)
         .          .     26:	}
         .          .     27:}
*/
