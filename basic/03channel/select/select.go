package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生产者c1、c2, 给chan传送int数据
	c1, c2 := generator(), generator()

	// 消费者worker
	worker := createWorker(1)

	var values []int                   // 从c1、c2拿的数据统一存入values
	tm := time.After(10 * time.Second) // 10秒后停止
	tick := time.Tick(time.Second)     // 每秒触发一次tick
	for {
		// activeValue从values取数据，写到activeWorker里面
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:] // 注意activeWorker拿到数据了，values才减一

		case <-time.After(800 * time.Millisecond): // 超时操作
			fmt.Println("timeout")
		case <-tick: // 每秒检测一次queue的长度
			fmt.Println("queue len = ", len(values))
		case <-tm: // 接收到停止的信号
			fmt.Println("end")
			return
		}
	}

}

// 生成器，生成一个chan，死循环给chan传递递增的整数
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			// 随机sleep[0,1500)毫秒
			time.Sleep(
				// 伪随机从[0,n)中间随机一个值
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

// 给id读取chan
func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d \n",
			id, n)
	}
}

// 创建一个worker付给id
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
