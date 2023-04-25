// https://medium.com/@j.d.livni/write-a-go-worker-pool-in-15-minutes-c9b42f640923
package main

import (
	"golang-tutorial/9-goroutine/ch_6/pool"
	"golang-tutorial/9-goroutine/ch_6/work"
	"log"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range work.CreateJobs(JOB_COUNT) {
		collector.Work <- pool.Work{Job: job, ID: i}
	}
}
