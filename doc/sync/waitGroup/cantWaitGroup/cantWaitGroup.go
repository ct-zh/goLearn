package cantWaitGroup

import (
	"errors"
	"sync"
	"time"
)

var CantWaitTimeoutErr = errors.New("cantWaitGroup timeout")

type CantWaitGroup struct {
	wg      sync.WaitGroup
	timeout time.Duration
}

func NewCantWaitGroup(timeout time.Duration) *CantWaitGroup {
	return &CantWaitGroup{
		wg:      sync.WaitGroup{},
		timeout: timeout,
	}
}

func (c *CantWaitGroup) Add(delta int) {
	c.wg.Add(delta)
}

func (c *CantWaitGroup) Done() {
	c.wg.Done()
}

func (c *CantWaitGroup) Wait() error {
	done := make(chan struct{})
	go func() {
		c.wg.Wait()
		close(done)
	}()

	timer := time.NewTimer(c.timeout)
	select {
	case <-done:
		return nil
	case <-timer.C:
		return CantWaitTimeoutErr
	}
}
