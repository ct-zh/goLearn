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
	fmt.Printf("GOMAXPROCS = %d\n", runtime.GOMAXPROCS(0))

	c := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("send:", i)
			c <- i
		}
	}()

	time.Sleep(time.Millisecond) // go park等到c buf被塞满

	for i := 0; i < 10; i++ {
		fmt.Printf("got: %d = %d\n", i, <-c)
	}
}
