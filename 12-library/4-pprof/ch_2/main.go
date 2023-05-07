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

// pprof 分析 cpu 占用过高
func cpu() {
	for {
		for i := 0; i < Gi; i++ {
			// do nothing
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

	cpu()
}

/*
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/profile
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/profile
Type: cpu
Time: Apr 22, 2023 at 5:18pm (CST)
Duration: 30s, Total samples = 8.45s (28.17%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 8420ms, 99.64% of 8450ms total
Dropped 9 nodes (cum <= 42.25ms)
      flat  flat%   sum%        cum   cum%
    8180ms 96.80% 96.80%     8430ms 99.76%  main.cpu (inline)
     240ms  2.84% 99.64%      240ms  2.84%  runtime.asyncPreempt
         0     0% 99.64%     8430ms 99.76%  main.main
         0     0% 99.64%     8430ms 99.76%  runtime.main
(pprof) list main.cpu
Total: 8.45s
ROUTINE ======================== main.cpu in /Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_2/main.go
     8.18s      8.43s (flat, cum) 99.76% of Total
         .          .     15:	Pi = Ki * Ti
         .          .     16:)
         .          .     17:
         .          .     18:func cpu() {
         .          .     19:	for {
     8.18s      8.42s     20:		for i := 0; i < Gi; i++ {
         .          .     21:			// do nothing
         .          .     22:		}
         .       10ms     23:		time.Sleep(time.Second)
         .          .     24:	}
         .          .     25:}
*/
