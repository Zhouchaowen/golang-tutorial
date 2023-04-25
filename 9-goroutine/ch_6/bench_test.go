package main

import (
	"golang-tutorial/9-goroutine/ch_6/pool"
	"golang-tutorial/9-goroutine/ch_6/work"
	"testing"
)

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range work.CreateJobs(20) {
			collector.Work <- pool.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, job := range work.CreateJobs(20) {
			work.DoWork(job, 1)
		}
	}
}
