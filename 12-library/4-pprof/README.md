# PProf

在 `Go` 语言中，可以使用内置的测试框架来进行单元测试, 基准测试。帮助开发者编写、运行和管理测试代码，确保代码的质量和稳定性。

上一小节提到，可以通过`-memprofile`和`-cpuprofile`生成内存分配文件和`cpu`分析文件。我们可以使用 `Go` 自带的 `pprof` 工具来查看 `mem.out` 和 `cpu.out` 文件中的分析结果。

具体的步骤如下：

1. 生成 `mem.out` 和 `cpu.out` 文件：

   ```go
   go test -bench=BenchmarkAdd -memprofile=mem.out -cpuprofile=cpu.out
   ```

2. 使用 `go tool pprof` 命令打开 `cpu.out` 文件：

   ```go
   go tool pprof cpu.out
   ```

3. 在 `pprof` 命令行中输入 `top` 命令，可以查看 CPU 使用率最高的函数：

   ```go
   (pprof) top
   ```

4. 输入 `list` 命令，可以查看当前函数的源代码：

   ```go
   (pprof) list funxxx
   ```

5. 使用 `web` 命令生成火焰图，并在浏览器中查看：

   ```go
   (pprof) svg
   ```

   执行完该命令后，会在当前目录下生成一个 `profile00x.svg` 文件，可以用浏览器打开该文件来查看火焰图。

## 目录

- 排查`Heap`占用过高
- 排查`CPU`占用过高
- 排查频繁内存回收
- 排查`Goroutine`泄露
- 排查`Mutex`的竞争
- 排查阻塞操作

## Heap占用过高

```go
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

// go tool pprof http://127.0.0.1:6060/debug/pprof/heap
func main() {
	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	heap()
}
```

返回结果：

```bash
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/heap
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/heap
Type: inuse_space
Time: Apr 21, 2023 at 6:34pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top # 通过top列出占用排行
Showing nodes accounting for 1.31GB, 99.85% of 1.31GB total
Dropped 15 nodes (cum <= 0.01GB)
      flat  flat%   sum%        cum   cum%
    1.31GB 99.85% 99.85%     1.31GB 99.85%  main.heap
         0     0% 99.85%     1.31GB 99.85%  main.main
         0     0% 99.85%     1.31GB 99.89%  runtime.main
(pprof) list main.heap	# 通过list查看main.heap函数
Total: 1.31GB
ROUTINE ======================== main.heap in /Users/zdns/Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_1/main.go
    1.31GB     1.31GB (flat, cum) 99.85% of Total
         .          .     18:func heap() {
         .          .     19:	for {
         .          .     20:		max := Gi
         .          .     21:		var buffer [][Mi]byte
         .          .     22:		for len(buffer)*Mi < max {
    1.31GB     1.31GB     23:			buffer = append(buffer, [Mi]byte{}) # 排查到具体占用位置
         .          .     24:		}
         .          .     25:		time.Sleep(time.Second)
         .          .     26:	}
         .          .     27:}
         .          .     28:
```

## 排查CPU占用过高

```go
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

// pprof 分析 cpu
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
```

```bash
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/profile
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/profile
Type: cpu
Time: Apr 22, 2023 at 5:18pm (CST)
Duration: 30s, Total samples = 8.45s (28.17%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top # 通过top列出占用排行
Showing nodes accounting for 8420ms, 99.64% of 8450ms total
Dropped 9 nodes (cum <= 42.25ms)
      flat  flat%   sum%        cum   cum%
    8180ms 96.80% 96.80%     8430ms 99.76%  main.cpu (inline)
     240ms  2.84% 99.64%      240ms  2.84%  runtime.asyncPreempt
         0     0% 99.64%     8430ms 99.76%  main.main
         0     0% 99.64%     8430ms 99.76%  runtime.main
(pprof) list main.cpu  # 通过list查看main.cpu函数
Total: 8.45s
ROUTINE ======================== main.cpu in /Users/zdns/Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_2/main.go
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
         .          .     26:
```

## 排查频繁内存回收

```go
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
```

```bash
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/allocs
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/allocs
Type: alloc_space
Time: Apr 22, 2023 at 5:44pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top # 通过top列出占用排行
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
(pprof)
```

## 排查协程泄露

```go
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
```

```bash
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/goroutine
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/goroutine
Type: goroutine
Time: Apr 22, 2023 at 5:28pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top # 通过top列出占用排行
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
```

## 排查mutex的竞争

```go
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
```

```bash
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/mutex
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/mutex
Type: delay
Time: Apr 23, 2023 at 9:30am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top # 通过top列出占用排行
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
```

## 排查阻塞操作

```go
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
```

```bash
harris-3% go tool pprof http://127.0.0.1:6060/debug/pprof/block
Fetching profile over HTTP from http://127.0.0.1:6060/debug/pprof/block
Type: delay
Time: Apr 23, 2023 at 9:31am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top # 通过top列出占用排行
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
```

## 思考题

## 自检

- `pprof`分析程序？
- 示例函数的编写和展示

## 参考
https://blog.wolfogre.com/posts/go-ppof-practice/

https://colobu.com/2019/08/20/use-pprof-to-compare-go-memory-usage/

https://www.sofastack.tech/blog/is-pprof-enough-for-go-memory-leak/

https://blog.xizhibei.me/2021/06/27/golang-heap-profiling/

