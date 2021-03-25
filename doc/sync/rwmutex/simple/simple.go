package simple

import "sync"

type process struct {
	mu   sync.RWMutex
	data int
}

func (p *process) w(i int, do func()) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.data = i
	do()
}

func (p *process) r() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.data
}
