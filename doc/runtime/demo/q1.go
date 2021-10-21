package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go fmt.Println(i) // 会输出什么？
	}
	time.Sleep(time.Minute)
}
