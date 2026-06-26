package taskpool

import (
	"context"
	"errors"
	"sync"
)

// Task 表示一个由执行池调度的任务。
//
// 后续实现时，你需要决定：
// 1. worker 执行任务时传入什么上下文
// 2. pool 关闭时，是否以及如何取消正在执行的任务
type Task func(ctx context.Context) error

var (
	ErrInvalidWorkerCount = errors.New("taskpool: worker count must be > 0")
	ErrInvalidQueueSize   = errors.New("taskpool: queue size must be >= 0")
	ErrNilTask            = errors.New("taskpool: nil task")
	ErrPoolStopped        = errors.New("taskpool: pool stopped")
)

// Stats 预留给后续阶段。
// 第 1 阶段不要求真正维护这些字段，但建议在结构设计时预留空间。
type Stats struct {
	Submitted int64
	Started   int64
	Succeeded int64
	Failed    int64
	Rejected  int64
}

// Pool 是本题的核心结构。
//
// 第 1 阶段建议你先只关注：
// - workers 如何启动
// - tasks 如何入队
// - Submit 在队列满时如何正确响应 ctx.Done()
//
// 第 2 阶段再补充关闭语义。
type Pool struct {
	workerN   int
	queueSize int

	tasks chan Task

	done chan struct{}
	once sync.Once
	wg   sync.WaitGroup
	lock sync.RWMutex
}

// New 创建一个新的任务执行池。
//
// 阶段 1 最低要求：
// - 校验参数
// - 初始化有界队列
// - 启动固定数量 worker
func New(workerN, queueSize int) (*Pool, error) {
	if workerN <= 0 {
		return nil, ErrInvalidWorkerCount
	}
	if queueSize < 0 {
		return nil, ErrInvalidQueueSize
	}

	p := &Pool{
		workerN:   workerN,
		queueSize: queueSize,
		tasks:     make(chan Task, queueSize),
		done:      make(chan struct{}),
	}

	// 启动固定数量的 worker，并从chan 中持续读取task
	for i := 0; i < workerN; i++ {
		p.wg.Add(1)
		go func() {
			for t := range p.tasks {
				ctx := context.Background() // todo 后续优化为 pool 级 context
				err := t(ctx)
				if err != nil {
					// todo 后续做失败计数等操作
				}
			}
			p.wg.Done()
		}()
	}

	return p, nil
}

// Submit 提交一个任务。
//
// 阶段 1 需要实现的关键语义：
// - task 为 nil 时返回 ErrNilTask
// - 队列未满时成功入队
// - 队列已满时阻塞等待
// - 等待期间若 ctx.Done()，返回 ctx.Err() 或包装后的错误
//
// 当前骨架没有实现任何并发语义，你需要自行补齐。
func (p *Pool) Submit(ctx context.Context, task Task) error {
	if task == nil {
		return ErrNilTask
	}
	p.lock.RLock()
	defer p.lock.RUnlock()

	select {
	case <-p.done:
		return ErrPoolStopped
	default:
	}

	select {
	case p.tasks <- task:
		return nil
	case <-ctx.Done(): // 阻塞超时直接报错
		return ctx.Err()
	}
}

// Stop 仅拒绝接受新任务，保证已提交任务执行完毕
func (p *Pool) Stop(ctx context.Context) error {
	var err error
	p.once.Do(func() {
		done := make(chan struct{})
		go func() {
			p.lock.Lock()
			close(p.done) // 停止新任务接收
			close(p.tasks)
			p.lock.Unlock()
			p.wg.Wait() // 确保worker全部退出
			close(done)
		}()
		select {
		case <-done:
			return
		case <-ctx.Done():
			err = ctx.Err()
			return
		}
	})
	return err
}

// Stats 是阶段 3 任务。
func (p *Pool) Stats() Stats {
	return Stats{}
}
