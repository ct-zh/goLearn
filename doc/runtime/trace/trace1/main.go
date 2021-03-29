package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

// 使用go tool trace分析调度流程

// go run main.go
// go tool trace trace.out

func main() {
	// 创建分析文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 写入trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 业务流程
	do()

	// 结束
	trace.Stop()
}

func do() {
	ch := make(chan int)
	quit := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case <-quit:
				return
			default:
			}
			ch <- i
			i++
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		timer := time.NewTimer(time.Second * 5)
		for {
			select {
			case <-timer.C:
				close(quit)
				return
			case val := <-ch:
				fmt.Println(val)
			}
		}
	}()

	select {
	case <-quit:
	}
}
