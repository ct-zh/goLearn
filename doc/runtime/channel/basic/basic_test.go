package basic

import (
	"testing"
	"time"
)

func TestCreateWork(t *testing.T) {
	work := createWork(1)

	// 往channel里面写数
	for i := 0; i < 10; i++ {
		work <- i
	}
	time.Sleep(2 * time.Second)
	close(work) // 两秒后关闭channel
}

func TestChanWait(t *testing.T) {
	chanWait()
}
