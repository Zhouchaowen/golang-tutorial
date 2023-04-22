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

- 排查Heap占用过高
- 排查CPU占用过高

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
(pprof) top 
Showing nodes accounting for 1.31GB, 99.85% of 1.31GB total
Dropped 15 nodes (cum <= 0.01GB)
      flat  flat%   sum%        cum   cum%
    1.31GB 99.85% 99.85%     1.31GB 99.85%  main.heap
         0     0% 99.85%     1.31GB 99.85%  main.main
         0     0% 99.85%     1.31GB 99.89%  runtime.main
(pprof) list main.heap
Total: 1.31GB
ROUTINE ======================== main.heap in /Users/zdns/Desktop/code_study/golang-tutorial/12-library/3-pprof/ch_1/main.go
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

```
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



## 排查协程泄露



## 排查mutex的竞争



## 排查阻塞操作





## 思考题

## 自检

- `pprof`分析程序？
- 示例函数的编写和展示

## 参考
https://bingdoal.github.io/backend/2022/05/unit-test-on-golang/

https://blog.wolfogre.com/posts/go-ppof-practice/

https://colobu.com/2019/08/20/use-pprof-to-compare-go-memory-usage/

https://www.sofastack.tech/blog/is-pprof-enough-for-go-memory-leak/

https://blog.xizhibei.me/2021/06/27/golang-heap-profiling/

