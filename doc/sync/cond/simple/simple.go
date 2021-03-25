package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	mail := 1

	// 数15秒
	go func() {
		for count := 0; count <= 10; count++ {
			time.Sleep(time.Second)
			mail = count
			cond.Broadcast()
		}
	}()

	// worker1
	go func() {
		// 触发的条件，如果不等于5，就会进入cond.Wait()等待，
		// 此时cond.Broadcast()通知进来的时候，wait阻塞解除，进入下一个循环，此时发现mail != 5，跳出循环，开始工作。
		cond.L.Lock()
		for mail != 5 {
			cond.Wait()
		}
		cond.L.Unlock()
		fmt.Println("worker1 started to work")
		time.Sleep(3 * time.Second)
		fmt.Println("worker1 work end")
	}()

	// worker2
	go func() {
		cond.L.Lock()
		for mail != 10 {
			cond.Wait()
		}
		cond.L.Unlock()
		fmt.Println("worker2 started to work")
		time.Sleep(3 * time.Second)
		fmt.Println("worker2 work end")
	}()

	// worker3
	go func() {
		cond.L.Lock()
		for mail != 10 {
			cond.Wait()
		}
		cond.L.Unlock()
		fmt.Println("worker3 started to work")
		time.Sleep(3 * time.Second)
		fmt.Println("worker3 work end")
	}()

	select {}

}
