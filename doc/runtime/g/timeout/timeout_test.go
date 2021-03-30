package timeout

// go routine的超时问题与处理方法

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// 例子1： 函数主题timeout设定了1ms的超时时间，而协程doBadThing需要1秒的运行时间；
// 通过NumGoroutine可以看到此时有1000多个协程，代表子协程都没有退出；
func TestTimeout(t *testing.T) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		_ = timeout(doBadThing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func doBadThing(done chan struct{}) {
	time.Sleep(time.Second)
	done <- struct{}{} // 会在这里阻塞住； 其实这里日常开发用的更多的是 close(done)
}

func timeout(fn func(chan struct{})) error {
	done := make(chan struct{})
	go fn(done)
	select {
	case <-done:
		fmt.Println("done!")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

// 例子2： 对上面的解决办法是,
// 1. 使用buffer channel来防止阻塞;
// 2. select增加default防止阻塞;
// 这里select适应的场景相对更复杂，见timeout2函数：
// chan回复done之后可以继续下面的打印代码，如果因为超时则走default直接返回了；
// 如果使用buffer chan，协程往chan发送数据后逻辑会直接往下继续走；

// 使用buffer channel
func TestTimeoutByBuffer(t *testing.T) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		_ = timeoutByBuffer(doBadThing)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func timeoutByBuffer(fn func(chan struct{})) error {
	done := make(chan struct{}, 1)
	go fn(done)
	select {
	case <-done:
		fmt.Println("done!")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

// 使用select default防止阻塞;
func TestTimeoutBySelect(t *testing.T) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		_ = timeoutBySelect(doBadThing2)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine())
}

func doBadThing2(done chan struct{}) {
	time.Sleep(time.Second)
	select {
	case done <- struct{}{}:
	default:
		return
	}
	fmt.Println("aa")
}

func timeoutBySelect(fn func(chan struct{})) error {
	done := make(chan struct{})
	go fn(done)
	select {
	case <-done:
		fmt.Println("done!")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}
