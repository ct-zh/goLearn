package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	send()
}

func send() {
	runtime.GOMAXPROCS(1)
	fmt.Printf("GOMAXPROCS = %d\n", runtime.GOMAXPROCS(0))
	const count = 12

	c := make(chan int, 2)
	go func() {
		for i := 0; i < count; i++ {
			fmt.Println("send:", i)
			c <- i
		}
	}()

	time.Sleep(time.Millisecond) // go park等到c buf被塞满

	for i := 0; i < count; i++ {
		fmt.Printf("got: %d = %d\n", i, <-c)
	}
}
