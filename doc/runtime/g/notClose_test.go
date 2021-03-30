package g

import (
	"runtime"
	"testing"
	"time"
)

// 未关闭channel

// 情况1
// 梳理：
// 1. end g比start g多1个; 因为sendTask没有阻塞行为，只可能是do func未退出;两个函数增加defer输出可证明；
// 2. do因为没有退出的逻辑，始终等待task chan的数据（如果在发送端close掉task chan，会导致do的task chan无限返回空值）；
// 3. 解决办法：1. sendTask 在退出前close掉chan，do程在读chan前判断chan是否关闭； 2. 增加done chan通知子协程退出;

func TestSendTask(t *testing.T) {
	t.Log("task start goroutine num:", runtime.NumGoroutine())
	sendTask()
	time.Sleep(time.Second)
	t.Log("task end goroutine num:", runtime.NumGoroutine())
}

func do(task chan int) {
	//defer fmt.Println("do func is quit")
	for {
		select {
		case _, notClosed := <-task:
			if !notClosed {
				return
			}
			time.Sleep(time.Millisecond)
		}
	}
}

func sendTask() {
	task := make(chan int, 10)
	//defer fmt.Println("sendTask func is quit")
	defer close(task)
	go do(task)
	for i := 1; i < 1000; i++ {
		task <- i
	}
}
