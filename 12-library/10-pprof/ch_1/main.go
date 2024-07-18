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

// pprof 分析 heap 占用过高
func heap() {
	for {
		max := Gi
		var buffer [][Mi]byte
		for len(buffer)*Mi < max {
			buffer = append(buffer, [Mi]byte{})
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

	heap()
}

/*
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/heap
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/heap
Type: inuse_space
Time: Apr 21, 2023 at 6:34pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1.31GB, 99.85% of 1.31GB total
Dropped 15 nodes (cum <= 0.01GB)
      flat  flat%   sum%        cum   cum%
    1.31GB 99.85% 99.85%     1.31GB 99.85%  main.heap
         0     0% 99.85%     1.31GB 99.85%  main.main
         0     0% 99.85%     1.31GB 99.89%  runtime.main
(pprof) list main.heap
Total: 1.31GB
ROUTINE ======================== main.heap in /Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_1/main.go
    1.31GB     1.31GB (flat, cum) 99.85% of Total
         .          .     18:func heap() {
         .          .     19:	for {
         .          .     20:		max := Gi
         .          .     21:		var buffer [][Mi]byte
         .          .     22:		for len(buffer)*Mi < max {
    1.31GB     1.31GB     23:			buffer = append(buffer, [Mi]byte{})
         .          .     24:		}
         .          .     25:		time.Sleep(time.Second)
         .          .     26:	}
         .          .     27:}
         .          .     28:
*/
