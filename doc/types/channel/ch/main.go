package main

import (
	"fmt"
)

// dlv debug main.go
// b main.main:2
// c
// disass
func main() {
	// channel初始化
	ch1 := make(chan int)
	ch2 := make(chan int, 3)

	l1 := len(ch1)
	p1 := cap(ch1)

	l2 := len(ch2)
	p2 := cap(ch2)
	fmt.Printf("ch1 len=%d, cap=%d\n", l1, p1)
	fmt.Printf("ch2 len=%d, cap=%d\n", l2, p2)

	// 写入unbuffer ch
	go func() {
		ch1 <- 1
	}()

	// 读取unbuffer ch
	<-ch1

	// 读取buffer ch
	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	for {
	// 		<-ch2
	// 	}
	// }()

	// 写入buffer ch
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	// ch2 <- 4 // buffer ch 写满并挂起

	// 写入已经关闭的ch  panic
	// close(ch1)
	// ch1 <- 1

	// 读取已经关闭的ch
	close(ch1)
	<-ch1

	// var ch3 chan int
	// ch3 <- 1 // 写入 nil chan
	// <-ch3 // 读取nil chan

}
