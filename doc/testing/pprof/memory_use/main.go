package main

import (
	"fmt"
	"time"
)

func main() {
	var o []int
	o = make([]int, 10240000000)
	fmt.Printf("%d", len(o))

	//bigMemoryOccupy()
	time.Sleep(time.Minute * 20)
}

func bigMemoryOccupy() {
	var o []int
	o = make([]int, 10240000000)
	fmt.Printf("%d", len(o))
}
