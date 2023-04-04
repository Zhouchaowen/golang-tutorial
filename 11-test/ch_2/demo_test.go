package ch_1

import "testing"

// 匹配后缀为Add的函数
// go test -bench='Add$'

// 达到指定的时间 t,默认值为 1 秒
// go test -bench='Add$' -benchtime=5s

// 运行基准测试 N 次 针对(b.N)
// go test -bench='Add$' -benchtime=1000x

// 运行基准测试 N 次 针对整个测试
// go test -bench='Add$' -benchtime=5s -count=3

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}

// 打印内存分配的指标
// go test -bench='MA$' -benchmem .
func BenchmarkMA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		memoryAllocation(1024)
	}
}
