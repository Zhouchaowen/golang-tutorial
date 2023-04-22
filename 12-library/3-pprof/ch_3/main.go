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

// pprof 分析频繁分配内存
func gc() {
	for {
		_ = make([]byte, 16*Mi)
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

	gc()
}

/* GODEBUG=gctrace=1 go run main.go | grep gc
gc 1 @0.004s 2%: 0.017+0.47+0.002 ms clock, 0.017+0.27/0.13/0+0.002 ms cpu, 16->16->0 MB, 17 MB goal, 1 P
gc 2 @1.025s 0%: 0.020+0.41+0.003 ms clock, 0.020+0/0.061/0.32+0.003 ms cpu, 16->16->0 MB, 17 MB goal, 1 P
gc 3 @2.031s 0%: 0.040+0.58+0.003 ms clock, 0.040+0/0.11/0.42+0.003 ms cpu, 16->16->0 MB, 17 MB goal, 1 P
gc 4 @3.036s 0%: 0.041+0.57+0.004 ms clock, 0.041+0/0.12/0.41+0.004 ms cpu, 16->16->0 MB, 17 MB goal, 1 P

gc 1        					第1次执行GC
@0.004s     					程序已经执行了0.004秒(我们可以看到这列数据一直在递增)
2%          					gc时间占程序总执行时间的2%
0.017+0.47+0.002 ms clock 		垃圾回收各阶段占用的时间(wall-clock，现实意义消耗的时间): STW(stop-the-world)清扫终止+并发标记和扫描的时间+STW标记终止的时间。
0.017+0.27/0.13/0+0.002 ms cpu  也是gc各阶段占用的时间，但是程序在cpu上消耗的时间。 STW(stop-the-world)清扫的时间+并发标记和扫描的时间(辅助时间/后台gc时间/闲置gc时间)+STW标记的时间。
16->16->0 MB	        		堆在gc开始时的大小、gc结束时的大小、当前活跃的大小
17 MB goal						全局堆的大小
1 P								P(process)的数量是1

*/

/*
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/allocs
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/allocs
Type: alloc_space
Time: Apr 22, 2023 at 5:44pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 17921.33kB, 100% of 17921.33kB total
Showing top 10 nodes out of 16
      flat  flat%   sum%        cum   cum%
   16384kB 91.42% 91.42%    16384kB 91.42%  main.gc (inline)
 1025.12kB  5.72% 97.14%  1025.12kB  5.72%  runtime.allocm
  512.20kB  2.86%   100%   512.20kB  2.86%  runtime.malg
         0     0%   100%    16384kB 91.42%  main.main
         0     0%   100%    16384kB 91.42%  runtime.main
         0     0%   100%  1025.12kB  5.72%  runtime.mstart
         0     0%   100%  1025.12kB  5.72%  runtime.mstart0
         0     0%   100%  1025.12kB  5.72%  runtime.mstart1
         0     0%   100%  1025.12kB  5.72%  runtime.newm
         0     0%   100%   512.20kB  2.86%  runtime.newproc.func1
(pprof) list main.gc
Total: 17.50MB
ROUTINE ======================== main.gc in /Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_3/main.go
      16MB       16MB (flat, cum) 91.42% of Total
         .          .     15:	Pi = Ki * Ti
         .          .     16:)
         .          .     17:
         .          .     18:func gc() {
         .          .     19:	for {
      16MB       16MB     20:		_ = make([]byte, 16*Mi)
         .          .     21:		time.Sleep(time.Second)
         .          .     22:	}
         .          .     23:}
         .          .     24:
         .          .     25:func main() {
(pprof)
*/
