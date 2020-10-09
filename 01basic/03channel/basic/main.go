package main

import (
	"fmt"
	"time"
)

func createWork(i int) chan int {
	c := make(chan int)
	go func(i int, c chan int) {
		for n := range c {
			fmt.Printf("Channel Id: %d Value: %d\n", i, n)
		}
	}(i, c)
	return c
}

// 一个最简单的channel
func main() {
	work := createWork(1)
	for i := 0; i < 10; i++ {
		work <- i
	}
	time.Sleep(2 * time.Second)
}
