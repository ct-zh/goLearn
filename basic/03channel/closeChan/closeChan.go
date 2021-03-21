package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// The Channel Closing Principle 通道关闭原则
// 不要从接收端关闭channel，也不要关闭有多个并发发送者的channel

// 如何优雅地关闭channel
// 1. 一个生产者时由生产者控制关闭channel；
// 2. 多个生产者时由第三方主动关闭channel；
// 3. 下游协程用for-range等待channel关闭;
// 4. worker需要select多个chan，等待done信号；

func main() {
	// 情况1：一个sender与N个worker
	//close1()

	// 情况2：N个生产者1个消费者
	//closeCh2()

	// 情况3： N个生产者N个消费者,需要第三方程序来控制开关
	//closeCh3()
}

func close1() {
	rand.Seed(time.Now().Unix()) // 设置随机数种子
	const MAXRANDINT = 10000     // 最大随机数10000
	const WORKERNUMS = 100       // worker数

	ch := make(chan int, 10)  // buffer 为10的int channel
	clock := sync.WaitGroup{} // 同步
	clock.Add(WORKERNUMS)     // 根据worker的数量添加同步标记

	// sender生产者
	go func() {
		for {
			// 随机数在 MAXRANDINT 以内，当val随机到0时退出
			if val := rand.Intn(MAXRANDINT); val == 0 {
				fmt.Println("sender end val=", val)
				close(ch)
				return
			} else {
				ch <- val
			}
		}
	}()

	// workers
	for i := 0; i < WORKERNUMS; i++ {
		go func(id int) {
			defer func() {
				clock.Done() // done必须在worker调用，使用defer在结束时返回
				fmt.Printf("worker[%d] done\n", id)
			}()
			for val := range ch {
				fmt.Printf("worker[%d] get num: %d \n", id, val)
			}
		}(i)
	}

	clock.Wait()
}

func closeCh2() {
	rand.Seed(time.Now().Unix()) // 设置随机数种子
	log.SetFlags(0)

	const MAXRANDINT = 10000 // 最大随机数10000
	const SENDERNUM = 10     // 10个生产者

	wg := sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan int, 100)   // 数据传输channel
	quit := make(chan struct{}) // 关闭信号传输channel, 使用空struct类型,节约空间

	// 生产者
	for i := 0; i < SENDERNUM; i++ {
		go func(id int) {
			for {
				select {
				case <-quit:
					return
				case ch <- rand.Intn(MAXRANDINT):
				}
			}
		}(i)
	}

	// 消费者
	go func() {
		defer wg.Done()
		for val := range ch {
			// 触发某个逻辑，关闭ch
			// 这个逻辑用(val == MAXRANDINT - 1)替代
			if val == MAXRANDINT-1 {
				fmt.Println("worker get end")
				close(quit)
				return
			} else {
				fmt.Printf("worker received %d \n", val)
			}
		}
	}()

	wg.Wait()
	close(ch)
}

const (
	ManyProducers = 100   // 生产者数量
	ManyWorkers   = 10    // 消费者数量
	MaxSeed       = 10000 // 随机数的最大值
)

func closeCh3() {
	rand.Seed(time.Now().Unix()) // 设置随机数种子
	log.SetFlags(0)

	wg := sync.WaitGroup{}
	wg.Add(ManyWorkers) // 等待消费者结束

	ch := make(chan int, 100)   // 数据传递channel
	quit := make(chan struct{}) // 停止信号直接使用空struct类型,节约空间

	// 用于给moderator传递停止信号的channel，
	// buffer=1防止在moderator还未创建时就已经有写数据进去(没有buffer的channel是同步的)
	toStop := make(chan string, 1)

	var stopText string

	// moderator 主持人协程,
	// 接收某个生产者或者消费者发出的停止信号(toStop)，并向所有生产者消费者发送退出信号(quit)
	go func() {
		select { // 没有default的select将会一直阻塞在这里，直到case成立；
		case stopText = <-toStop:
			close(quit)
			return
		}
	}()

	// 生产者
	for i := 0; i < ManyProducers; i++ {
		go func(id int) {
			for {
				val := rand.Intn(MaxSeed)

				// 触发某个逻辑从而return,这里用(val == MaxSeed-1)替代
				// 向moderator发送停止信号
				if val == MaxSeed-1 {
					toStop <- fmt.Sprintf("#produccer: %d stopped", id)
					return
				}

				// 这样做的目的是能尽早退出
				select {
				case <-quit:
					return
				default:
				}

				select {
				case <-quit: // 接收停止的信号，收到后立马return
					return
				case ch <- val:
				}
			}
		}(i)
	}

	// 消费者
	for i := 0; i < ManyWorkers; i++ {
		go func(id int) {
			defer wg.Done()

			for {
				select { // 这样做的目的是能尽早退出
				case <-quit:
					return
				default:
				}

				select {
				case <-quit:
					return
				case val := <-ch:
					if val == MaxSeed-1 {
						toStop <- fmt.Sprintf("worker#%d stopped", id)
						return
					}

					fmt.Printf("worker:%d get num=%d\n", id, val)
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("moderator sent close signal: ", stopText)
}
