package atomic

import (
	"fmt"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var a atomicInt

	a.incr()
	fmt.Printf("main value: %d \n", a.get())
	go func() {
		for i := 1; i <= 10; i++ {
			a.incr()
			fmt.Printf("goroutine value: %d \n", a.get())
		}
	}()

	time.Sleep(time.Millisecond * 2)
	for i := 1; i <= 10; i++ {
		a.incr()
		fmt.Printf("main value: %d \n", a.get())
	}
}
