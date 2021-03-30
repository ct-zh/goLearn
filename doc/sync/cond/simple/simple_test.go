package simple

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Cond条件变量用来协调想要访问共享资源的那些goroutine，当共享资源的状态发生变化的时候，它可以用来通知被互斥锁阻塞的goroutine。

// 也就是 N个协程需要阻塞等待1个协程的工作做到某个程度，当这个协程达到某个条件时，可以用Cond来通知其他协程可以开始了；

// 例子1

func TestSimple(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})

	mail := 1
	fn := func(signal int, id int) {
		cond.L.Lock()
		for mail != signal {
			cond.Wait()
		}
		cond.L.Unlock()
		fmt.Printf("worker:%d get signal: %d started to work\n", id, signal)
		time.Sleep(time.Second)
		fmt.Printf("worker%d work end\n", id)
	}

	go fn(4, 1)
	go fn(5, 2)
	go fn(10, 3)

	for count := 0; count <= 10; count++ {
		time.Sleep(time.Second)
		mail = count
		cond.Broadcast()
	}

	time.Sleep(time.Second * 2)
}
