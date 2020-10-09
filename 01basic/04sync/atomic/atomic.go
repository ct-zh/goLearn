package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

// 原子操作
func (a *atomicInt) incr() {
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value = a.value + 1
	}()
}

//
func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

// 使用 sync.Mutex 的lock功能对数据进行原子操作
func main() {
	// go run -race 查看数据冲突 去掉该例子内的锁操作,即可查看出对a的同时操作导致冲突的情况

	var a atomicInt

	a.incr()
	fmt.Printf("main value: %d \n", a.get())
	go func() {
		for i := 1; i <= 10; i++ {
			a.incr()
			fmt.Printf("goroutine value: %d \n", a.get())
		}
	}()
	//time.Sleep(time.Millisecond * 5)
	for i := 1; i <= 10; i++ {
		a.incr()
		fmt.Printf("main value: %d \n", a.get())
	}

	time.Sleep(time.Second * 3)
}
