package with

// context的四个with函数 demo
import (
	"context"
	"fmt"
	"time"
)

// withCancel函数,返回一个cancel函数，调用该函数会给子进程的ctx.Done发送信号
func withCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, "with cancel func")
	time.Sleep(time.Second * 2)
	cancel() // 调用cancel函数，
	time.Sleep(time.Second)
}

func withValue() {
	key := "ky name"
	val := "context set value"

	// ctx的特点，可以叠加传递上下文
	ctx, cancel := context.WithCancel(context.Background())
	ctx2 := context.WithValue(ctx, key, val)

	go workWithValue(ctx2, "work with value 1", key)
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second)
}

// 设置过期时间
func withTimeOut() {
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*2)
	go work(ctx, "work with timeout")
	time.Sleep(time.Second * 4)
	cancelFn()
	time.Sleep(time.Second)
}

// 设置过期时间点，类似上面的timeout
func withDeadline() {
	ctx, cancelFn := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	go work(ctx, "work with deadline")
	time.Sleep(time.Second * 4)
	cancelFn()
	time.Sleep(time.Second)
}

func workWithValue(ctx context.Context, name string, key string) {
	for {
		select {
		case <-ctx.Done():
			val := ctx.Value(key)
			fmt.Printf("name %s get done, key:%s value: %s\n",
				name, key, val)
			return
		default:
			fmt.Printf("[%s] name(%s) waiting for signal, key is %s\n",
				time.Now().Format("2006-01-02 15:04:05"), name, key)
			time.Sleep(time.Second)
		}
	}
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("name ", name, " get done")
			return
		default:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), ": name ", name, " waiting for signal")
			time.Sleep(time.Second)
		}
	}
}
