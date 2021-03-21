package basic

import (
	"fmt"
	"time"
)

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

// channel的超时机制
// 使用空struct来申明一个用来关闭的channel
func chanWait() {
	ch := make(chan int)
	quit := make(chan struct{})

	go func() {
		var num int
		for {
			select {
			case num = <-ch:
				fmt.Println("num is ", num)
			case <-time.After(time.Second * 3): // 超时时间3秒
				fmt.Println("send quit")
				close(quit)
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}

	c := <-quit
	fmt.Println("quit", c)
}
