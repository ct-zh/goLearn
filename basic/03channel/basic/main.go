package main

import (
	"fmt"
	"time"
)

// 一个最简单的channel
func main() {
	work := createWork(1)

	// 往channel里面写数
	for i := 0; i < 10; i++ {
		work <- i
	}
	time.Sleep(2 * time.Second)
	close(work) // 两秒后关闭channel
}

// 创建channel
func createWork(i int) chan int {
	c := make(chan int)
	go func(a int, c2 chan int) {
		for n := range c2 {
			fmt.Printf("Channel Id: %d Value: %d\n", a, n)
		}
	}(i, c)
	return c
}
