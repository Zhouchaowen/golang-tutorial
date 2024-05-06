package ch_1

import "testing"

const (
	MAX   = 50000000
	BATCH = 500
	CAP   = 100
)

func normal() {
	done := make(chan struct{})
	c := make(chan int, CAP)
	go func() {
		defer close(done)

		count := 0
		for x := range c {
			count += x
		}
	}()

	for i := 0; i < MAX; i++ {
		c <- i
	}
	close(c)
	<-done
}

func batch() {
	done := make(chan struct{})
	c := make(chan [BATCH]int, CAP)
	go func() {
		defer close(done)

		count := 0
		for a := range c {
			for _, x := range a {
				count += x
			}
		}
	}()

	for i := 0; i < MAX; i += BATCH {
		var b [BATCH]int
		for n := 0; n < BATCH; n++ {
			b[n] = i + n
			if i+n == MAX-1 {
				break
			}
		}

		c <- b
	}

	close(c)
	<-done
}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normal()
	}
}

func BenchmarkBatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		batch()
	}
}

// go test -bench . -benchmem
