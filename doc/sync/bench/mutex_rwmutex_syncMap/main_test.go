package mutex_rwmutex_syncMap

import (
	"sync"
	"testing"
)

// 性能对比：
// 1. sync.Mutex
// 2. sync.RwMutex
// 3. sync.Map
// 单次结论：
// 只写：sync.Map > sync.RwMutex > sync.Mutex
// 只读：sync.Mutex > sync.RwMutex > sync.Map
// 读写各一半：sync.RwMutex > sync.Mutex > sync.Map

// 2023.05.09 询问chatGPT: 在读多写少的环境下，sync.RWMutex 性能可能更好。因为它支持多个读锁同时获取，而不需要排他锁，这可以提高并发读的性能。

type myMutex struct {
	sync.Mutex
	data map[int]struct{}
}

type myRwMutex struct {
	sync.RWMutex
	data map[int]struct{}
}

const GenNum int = 100
const Length int = 50

func BenchmarkMutexWriteOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := sync.WaitGroup{}
		tmp := &myMutex{
			data: make(map[int]struct{}, Length),
		}
		w.Add(GenNum)
		for i := 0; i < GenNum; i++ {
			go func() {
				defer w.Done()
				for i := 0; i < Length; i++ {
					tmp.Lock()
					tmp.data[i] = struct{}{}
					tmp.Unlock()
				}
			}()
		}
		w.Wait()
	}
}

func BenchmarkMutexReadOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := sync.WaitGroup{}
		tmp := &myMutex{
			data: make(map[int]struct{}, Length),
		}
		w.Add(GenNum)
		for i := 0; i < GenNum; i++ {
			go func() {
				defer w.Done()
				for i := 0; i < Length; i++ {
					tmp.Lock()
					_ = tmp.data[0]
					tmp.Unlock()
				}
			}()
		}
		w.Wait()
	}
}

func BenchmarkMutexReadAndWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := sync.WaitGroup{}
		tmp := &myMutex{
			data: make(map[int]struct{}, Length),
		}
		w.Add(GenNum)
		for i := 0; i < GenNum; i++ {
			if i%2 == 0 {
				go func() {
					defer w.Done()
					for i := 0; i < Length; i++ {
						tmp.Lock()
						_ = tmp.data[0]
						tmp.Unlock()
					}
				}()
			} else {
				go func() {
					defer w.Done()
					for i := 0; i < Length; i++ {
						tmp.Lock()
						tmp.data[i] = struct{}{}
						tmp.Unlock()
					}
				}()
			}
		}
		w.Wait()
	}
}

func BenchmarkRwMutexWriteOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := sync.WaitGroup{}
		tmp := &myRwMutex{
			data: make(map[int]struct{}, Length),
		}
		w.Add(GenNum)
		for i := 0; i < GenNum; i++ {
			go func() {
				defer w.Done()
				for i := 0; i < Length; i++ {
					tmp.Lock()
					tmp.data[i] = struct{}{}
					tmp.Unlock()
				}
			}()
		}
		w.Wait()
	}
}

func BenchmarkRwMutexReadOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := sync.WaitGroup{}
		tmp := &myRwMutex{
			data: make(map[int]struct{}, Length),
		}
		w.Add(GenNum)
		for i := 0; i < GenNum; i++ {
			go func() {
				defer w.Done()
				for i := 0; i < Length; i++ {
					tmp.RLock()
					_ = tmp.data[0]
					tmp.RUnlock()
				}
			}()
		}
		w.Wait()
	}
}

func BenchmarkRwMutexReadAndWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := sync.WaitGroup{}
		tmp := &myRwMutex{
			data: make(map[int]struct{}, Length),
		}
		w.Add(GenNum)
		for i := 0; i < GenNum; i++ {
			if i%2 == 0 {
				go func() {
					defer w.Done()
					for i := 0; i < Length; i++ {
						tmp.RLock()
						_ = tmp.data[0]
						tmp.RUnlock()
					}
				}()
			} else {
				go func() {
					defer w.Done()
					for i := 0; i < Length; i++ {
						tmp.Lock()
						tmp.data[i] = struct{}{}
						tmp.Unlock()
					}
				}()
			}
		}
		w.Wait()
	}
}

func BenchmarkSyncMapWriteOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := sync.WaitGroup{}
		tmp := &sync.Map{}
		w.Add(GenNum)
		for i := 0; i < GenNum; i++ {
			go func() {
				defer w.Done()
				for i := 0; i < Length; i++ {
					tmp.Store(i, struct{}{})
				}
			}()
		}
		w.Wait()
	}
}

func BenchmarkSyncMapReadOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := sync.WaitGroup{}
		tmp := &sync.Map{}
		w.Add(GenNum)
		for i := 0; i < GenNum; i++ {
			go func() {
				defer w.Done()
				for i := 0; i < Length; i++ {
					tmp.Load(i)
				}
			}()
		}
		w.Wait()
	}
}

func BenchmarkSyncMapReadAndWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := sync.WaitGroup{}
		tmp := &sync.Map{}
		w.Add(GenNum)
		for i := 0; i < GenNum; i++ {
			if i%2 == 0 {
				go func() {
					defer w.Done()
					for i := 0; i < Length; i++ {
						tmp.Load(i)
					}
				}()
			} else {
				go func() {
					defer w.Done()
					for i := 0; i < Length; i++ {
						tmp.Store(i, struct{}{})
					}
				}()
			}
		}
		w.Wait()
	}
}
