package taskpool

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_InvalidArgs(t *testing.T) {
	// invalid worker count
	_, err1 := New(0, 10)
	assert.Equal(t, err1, ErrInvalidWorkerCount)

	// invalid queue sieze
	_, err2 := New(5, -1)
	assert.Equal(t, err2, ErrInvalidQueueSize)
}

func TestPool_Submit_TaskExecuted(t *testing.T) {
	p, err := New(1, 3)
	assert.Nil(t, err)

	ctx := context.Background()

	var nilTask Task
	err1 := p.Submit(ctx, nilTask)
	assert.Equal(t, err1, ErrNilTask)

	testChan := make(chan int8)
	var task Task = func(ctx context.Context) error {
		testChan <- 1
		return nil
	}

	err2 := p.Submit(ctx, task)
	assert.Nil(t, err2)
	data, ok := <-testChan
	assert.True(t, ok)
	assert.Equal(t, data, int8(1))
}

func TestPool_Submit_WorkersConsumeConcurrently(t *testing.T) {
	ctx := context.Background()
	var (
		workerN   = 3
		queueSize = 10
		taskNum   = 20
		errChan   = make(chan error, taskNum)
	)

	// 并发阻塞task，确保同时有 workerN 个worker在运行
	startCh := make(chan struct{}, workerN)
	release := make(chan struct{})
	doneCh := make(chan struct{}, taskNum)
	// 该task仅用于测试是否同时有workerN个 worker在运行
	var syncTask Task = func(ctx context.Context) error {
		startCh <- struct{}{}
		<-release
		doneCh <- struct{}{}
		return nil
	}

	p, err := New(workerN, queueSize)
	assert.Nil(t, err)

	for i := 0; i < taskNum; i++ {
		// t.Logf("submit idx=%d", i)
		if i < workerN { // 只有前workerN个task执行syncTask，用于验证同时有workerN个worker在执行
			go func() {
				errChan <- p.Submit(ctx, syncTask)
			}()
		} else {
			go func() {
				errChan <- p.Submit(ctx, func(ctx context.Context) error {
					doneCh <- struct{}{}
					return nil
				})
			}()
		}
	}

	// 读满workerN次startCh才close，此时会有 workerN 个 work阻塞等待close
	// 确保了同时有workerN 个 worker在执行
	for i := 0; i < workerN; i++ {
		<-startCh
	}
	close(release)

	// 验证有 taskNum条任务已提交 （但是不代表有这么多条任务已完成）
	for i := 0; i < taskNum; i++ {
		e := <-errChan
		assert.Nil(t, e)
	}

	// 验证有 taskNum 条任务已完成
	for i := 0; i < taskNum; i++ {
		<-doneCh
	}
}

// 队列先被占满
// 再次 Submit 时不能立刻成功
// 这次阻塞中的 Submit 会因为传入的 ctx 超时或取消而返回错误
func TestPool_Submit_ContextCanceledWhenQueueFull(t *testing.T) {

	var (
		workerN   = 3
		queueSize = 5
	)
	p, err := New(workerN, queueSize)
	assert.Nil(t, err)

	// 先占满队列
	ctxOccupy := context.Background()
	occupyChan := make(chan struct{})
	var occupyTask Task = func(ctx context.Context) error {
		<-occupyChan
		return nil
	}

	for i := 0; i < queueSize+workerN; i++ {
		err := p.Submit(ctxOccupy, occupyTask)
		assert.Nil(t, err)
	}

	// 再次submit
	doneCtx, cancel := context.WithCancel(context.Background())
	cancelChan := make(chan struct{})
	go func() {
		doneErr := p.Submit(doneCtx, func(ctx context.Context) error {
			return nil
		})
		assert.Equal(t, doneErr, context.Canceled)
		close(cancelChan)
	}()
	cancel()
	<-cancelChan
}

// 建议你在阶段 1 至少补上以下测试：
//
// 1. TestNew_InvalidArgs
// 2. TestPool_Submit_TaskExecuted
// 3. TestPool_Submit_WorkersConsumeConcurrently
// 4. TestPool_Submit_ContextCanceledWhenQueueFull
//
// 写测试时尽量避免直接依赖 time.Sleep 判断结果，
// 优先使用 channel、WaitGroup、barrier 等可控同步手段。
