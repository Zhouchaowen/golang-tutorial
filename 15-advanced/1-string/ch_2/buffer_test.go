package ch_2

import (
	"bytes"
	"strings"
	"testing"
)

var S = strings.Repeat("a", 1000)

func concat() bool {
	var s2 string
	for i := 0; i < 1000; i++ {
		s2 += "a"
	}
	return s2 == S
}

func join() bool {
	b := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		b[i] = "a"
	}
	return strings.Join(b, "") == S
}

func buffer() bool {
	var b bytes.Buffer
	b.Grow(1000)
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String() == S
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !concat() {
			b.Fatal()
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !join() {
			b.Fatal()
		}
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !buffer() {
			b.Fatal()
		}
	}
}
