package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
通过 MAP 如何实现 key 不存在 get 操作等待直到ky存在或者超时，保证并发安全,且需要实现以下接口
type sp interface {
	Out(key string, val interface{})   // 存储 key/val，如果读取 key 的 goroutine 挂起，则唤醒它。此方法不应阻塞，应立即执行。
	Rd(key string, timeout time.Duration) interface{} // 读取一个 key，如果 key 不存在则阻塞，等待 key 存在或者超时。
}

知识点：
	1.Cond.Wait()会解锁当持有的锁，Wait()返回后会重新加锁，所以需要循环检查静态条件
		 c.L.Lock()
		 for !condition() {
			 c.Wait()
		 }
		 ... make use of condition ...
		 c.L.Unlock()
*/

type sp interface {
	Out(key string, val interface{})                  // 存储 key/val，如果读取 key 的 goroutine 挂起，则唤醒它。此方法不应阻塞，应立即执行。
	Rd(key string, timeout time.Duration) interface{} // 读取一个 key，如果 key 不存在则阻塞，等待 key 存在或者超时。
}

type ConcurrentMap struct {
	data map[string]interface{}
	mu   sync.Mutex
	cond *sync.Cond
}

func NewConcurrentMap() *ConcurrentMap {
	cm := &ConcurrentMap{
		data: make(map[string]interface{}),
	}
	cm.cond = sync.NewCond(&cm.mu)
	return cm
}

func (c *ConcurrentMap) Out(key string, val interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = val

	c.cond.Broadcast()
}

func (c *ConcurrentMap) Rd(key string, timeout time.Duration) interface{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		if val, ok := c.data[key]; ok {
			return val
		}

		ch := make(chan struct{})

		go func() {
			c.cond.Wait()
			close(ch)
		}()

		select {
		case <-ch:
			fmt.Println("wake")
		case <-timer.C:
			fmt.Println("timeout!")
			return nil
		}
	}
}

func main() {
	cm := NewConcurrentMap()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Int63n(10)))
			cm.Out(fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
			fmt.Println(fmt.Sprintf("key-%d", i), fmt.Sprintf("value-%d", i))
		}(i)
	}

	value := cm.Rd(fmt.Sprintf("key-%d", 55), time.Second*10)
	fmt.Println("value:", value)
}
