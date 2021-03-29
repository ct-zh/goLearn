package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	log.Println("start")

	test()

	log.Println("force gc")
	runtime.GC()

	log.Println("end")
	time.Sleep(time.Second * 360)
}

func test() {
	container := make([]int, 8)

	log.Println("loop start")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}

	log.Println("loop end")
}
