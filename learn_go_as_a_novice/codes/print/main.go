package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	B()
}

func A() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
	ch := make(chan int)
	<-ch
}

func B() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
