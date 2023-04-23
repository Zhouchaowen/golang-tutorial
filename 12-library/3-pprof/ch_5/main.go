package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

const (
	Ki = 1024
	Mi = Ki * Ki
	Gi = Ki * Mi
	Ti = Ki * Gi
	Pi = Ki * Ti
)

func mutex() {
	for {
		m := &sync.Mutex{}
		m.Lock()
		go func() {
			time.Sleep(time.Second)
			m.Unlock()
		}()
		m.Lock()
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

	mutex()
}

/*
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/mutex
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/mutex
Type: delay
Time: Apr 23, 2023 at 9:30am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 32.05s, 100% of 32.05s total
      flat  flat%   sum%        cum   cum%
    32.05s   100%   100%     32.05s   100%  sync.(*Mutex).Unlock (inline)
         0     0%   100%     32.05s   100%  main.mutex.func1
(pprof) list main.mutex.func1
Total: 32.05s
ROUTINE ======================== main.mutex.func1 in /Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_5/main.go
         0     32.05s (flat, cum)   100% of Total
         .          .     20:	for {
         .          .     21:		m := &sync.Mutex{}
         .          .     22:		m.Lock()
         .          .     23:		go func() {
         .          .     24:			time.Sleep(time.Second)
         .     32.05s     25:			m.Unlock()
         .          .     26:		}()
         .          .     27:		m.Lock()
         .          .     28:		time.Sleep(time.Second)
         .          .     29:	}
         .          .     30:}
(pprof)

*/
