package main

import "fmt"

// dlv debug main.go
// b main.main:2
// c
// disass
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int, 3)

	fmt.Printf("ch1 len=%d, cap=%d\n", len(ch1), cap(ch1))
	fmt.Printf("ch2 len=%d, cap=%d\n", len(ch2), cap(ch2))

	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	ch2 <- 4
}
