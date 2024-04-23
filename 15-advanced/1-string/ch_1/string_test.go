// 来源：https://www.yuque.com/qyuhen/go/ab78rx
package ch_1

import (
	"strings"
	"testing"
	"unsafe"
)

var S = strings.Repeat("a", 100)

// 普通转换 normalConv 调用 runtime.stringtoslicebyte、runtime.slicebytetostring， 引发 mallocgc、memmove 等操作。
func normalConv() bool {
	b := []byte(S)
	s := string(b)
	return s == S
}

func unsafeConv() bool {
	b := unsafe.Slice(unsafe.StringData(S), len(S))

	s := unsafe.String(unsafe.SliceData(b), len(b))
	return s == S
}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !normalConv() {
			b.Fail()
		}
	}
}

func BenchmarkUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !unsafeConv() {
			b.Fail()
		}
	}
}
