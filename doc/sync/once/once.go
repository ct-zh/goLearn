package once

import (
	"fmt"
	"sync"
	"time"
)

func fn() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func doOnce() {
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		once.Do(fn)
		fmt.Println(i) // i会被循环打印出来，但是once对应的函数
	}
}
