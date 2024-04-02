package n_pool

import (
	"fmt"
	"time"
)

var ErrScheduleTimeout = fmt.Errorf("schedule error: timed out")

type Pool struct {
	sem  chan struct{}
	work chan func()
}

// NewPool 创建具有给定大小的新goroutine池。
// 它还创建了给定大小的工作队列。
// 最后，它立即生成给定数量的goroutine。
// size 最大goroutine数量;
// queue 队列等待长度
// spawn 初始创建的 goroutine 数量
func NewPool(size, queue, spawn int) *Pool {
	if spawn <= 0 && queue > 0 {
		panic("dead queue configuration detected")
	}
	if spawn > size {
		panic("spawn > workers")
	}
	p := &Pool{
		sem:  make(chan struct{}, size),
		work: make(chan func(), queue),
	}
	for i := 0; i < spawn; i++ {
		p.sem <- struct{}{}
		go p.worker(func() {})
	}

	return p
}

// Schedule 调度要在池内执行的任务
func (p *Pool) Schedule(task func()) {
	p.schedule(task, nil)
}

// ScheduleTimeout schedules task to be executed over pool's workers.
// It returns ErrScheduleTimeout when no free workers met during given timeout.
func (p *Pool) ScheduleTimeout(timeout time.Duration, task func()) error {
	return p.schedule(task, time.After(timeout))
}

func (p *Pool) schedule(task func(), timeout <-chan time.Time) error {
	select {
	case <-timeout:
		return ErrScheduleTimeout
	case p.work <- task:
		return nil
	case p.sem <- struct{}{}:
		go p.worker(task)
		return nil
	}
}

func (p *Pool) worker(task func()) {
	defer func() { <-p.sem }()

	task()

	for task := range p.work {
		task()
	}
}
