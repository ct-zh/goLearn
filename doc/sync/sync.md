# 并发同步与锁


## sync.Once


## sync.Mutex
首先举出Mutex中最重要的几个常量, Mutex的基本结构, 代码见`sync/mutex.go`

```go
const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

// Mutex 是一个互斥锁, 其零值是未上锁的
// Mutex 不支持copy
type Mutex struct {
	state int32
	sema  uint32
}
```

- 
- 
- sema即是semaphore结构. 所有和锁相关的最终都会有这么一个成员.代表信号量地址


// todo

state则是状态, 为了优化又做了一些比较恶心的就是，他把最后三位拿出来当一个标记位:

- 其中最后一位其实就是表示我当前的所是不是已经上锁了, 是不是已经被被人抢到这个锁了？
- 倒数第二位的话是说我当前 note是不是已经有唤醒的等所的G在等待这个锁
- 倒数第三位的话是指的是我当前是我的这个呼吸量是不是也要进入饥饿模式. 如果是饥饿模式的话后来的goroutine不能主动抢这个锁

state是32位的，除了这三位应该还是29位，这29位其实就是当前所有等待这个锁的goroutine的一个简单的基数。

但是因为要做优化，把最后三位拿来做标志位，所以看代码的时候会发现，做一些操作的时候, 每次都要向右移三位。


### Lock 、TryLock

先看两个方法的声明:

```go
func (m *Mutex) TryLock() bool
func (m *Mutex) Lock()
```

可以看到Lock是没有返回值的, 如果锁已经在使用中，则调度的g会阻塞，直到Mutex可用。而TryLock则是尝试获取一次锁并返回获取结果;

先看TryLock源码:

```go
func (m *Mutex) TryLock() bool {
	old := m.state
	if old&(mutexLocked|mutexStarving) != 0 {
		return false
	}
  // ...
}
```

old代表Mutex的当前状态, 如果`old`的状态是`mutexLocked`或`mutexStarving`，则表示互斥锁已被锁定或正在饥饿状态，此时返回`false`;

> 这里使用位运算检查 `old` 的值是否包含 `mutexLocked` 或 `mutexStarving` 的标志位。todo 疑问:如何进行位运算的? 

```go
func (m *Mutex) TryLock() bool {
  // ......
	if !atomic.CompareAndSwapInt32(&m.state, old, old|mutexLocked) {
		return false
	}

	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
	return true
}
```

如果互斥锁未被锁定，那么就会尝试使用`atomic.CompareAndSwapInt32`原子操作来改变互斥锁的状态。这个函数会比较`Mutex`的当前状态和`old`，如果它们相同，那么就将`Mutex`的状态设置为`old|mutexLocked`，并返回`true`。如果它们不同，那么就返回`false`。

如果`race`条件启用，就调用 `race.Acquire` 方法来标记当前 goroutine 已经获得了锁.

再看看Lock方法:

```go
func (m *Mutex) Lock() {
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		return
	}
	m.lockSlow()
}
```

Lock的快速路径与TryLock类似, 直接使用`atomic.CompareAndSwapInt32`原子操作来进行修改;

如果失败了则开始慢速操作`lockSlow()`

```go
func (m *Mutex) lockSlow() {
	var waitStartTime int64	// 等待开始时间
	starving := false  // 饥饿状态标志
	awoke := false		// 唤醒状态标志
	iter := 0			// 迭代次数
	old := m.state		// 互斥锁的当前状态
	for {
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}
			runtime_doSpin()
			iter++
			old = m.state
			continue
		}
    // ......
  }
}
```

- 先定义了一些变量，包括`waitStartTime`（等待开始时间）、`starving`（饥饿状态标志）、`awoke`（唤醒状态标志）和`iter`（迭代次数）。然后获取了互斥锁的当前状态`old`
- 在一个无限循环中，首先检查互斥锁是否处于锁定状态或饥饿状态。如果已被锁，并且可以进行自旋操作，那么就尝试设置`mutexWoken`标志，以通知`Unlock`方法不唤醒其他阻塞的goroutine。
  - `old&(mutexLocked|mutexStarving) == mutexLocked`  
- 然后进行自旋操作，迭代次数+1, 并更新互斥锁的状态`old`。

```go
func (m *Mutex) lockSlow() {
  // ......
  for {
    // .....
		new := old
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		if old&(mutexLocked|mutexStarving) != 0 {
			new += 1 << mutexWaiterShift
		}

		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}
		if awoke {
			if new&mutexWoken == 0 {
				throw("sync: inconsistent mutex state")
			}
			new &^= mutexWoken
		}
    // ....
  }
}
```

如果互斥锁不处于饥饿状态，那么就尝试获取锁。如果互斥锁处于锁定状态或饥饿状态，那么就增加等待者的数量。如果当前goroutine处于饥饿状态，并且互斥锁处于锁定状态，那么就将互斥锁切换到饥饿模式。





```go
func (m *Mutex) lockSlow() {
  // ......
  for {
    // .....
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&(mutexLocked|mutexStarving) == 0 {
				break
			}
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 {
				waitStartTime = runtime_nanotime()
			}
			runtime_SemacquireMutex(&m.sema, queueLifo, 1)
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			old = m.state
			if old&mutexStarving != 0 {
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					throw("sync: inconsistent mutex state")
				}
				delta := int32(mutexLocked - 1<<mutexWaiterShift)
				if !starving || old>>mutexWaiterShift == 1 {
					delta -= mutexStarving
				}
				atomic.AddInt32(&m.state, delta)
				break
			}
			awoke = true
			iter = 0
		} else {
			old = m.state
		}
	}
	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
}
```

使用`atomic.CompareAndSwapInt32`原子操作尝试更新锁的状态。如果更新成功且原来未被持有,则成功获取锁,跳出循环。





### 饥饿模式

饥饿模式主要是为了实现锁的公平性。因为在设计锁的时候给了gorouine一个优先级，就是最新来讲锁的goroutine优先级是最高的，因为在代码里面有一个for循环是一个自旋的逻辑，这个逻辑一直在检查我当前能不能抢到这把锁，然后如果抢得到的话，那就直接去把lock函数返回，然后用户去执行lock后面的代码。

如果当前的mutex已经饥饿了，那说明在队列里面已经有go rou ti ne等待了很长时间了，我记得那个值是一毫秒，也就说有goroutine在等锁已经等一毫秒了，那这时候你新来的goroutine还有很高的优先级去拿到这把锁，那就不太合理了，所以它要进入饥饿模式。然后进入饥饿模式以后，所有的going都是老老实实的去排队，然后排队之后是由那个排队之后有一个把饥饿模式重新置回来的。



## rwmutex

```go

type rwmutex struct {
	rLock      mutex
	readers    muintptr
	readerPass uint32

	wLock  mutex
	writer muintptr

	readerCount atomic.Int32
	readerWait  atomic.Int32
}

```


## sync.Pool 

结构
```go
runtime.allPools    // runtime全局对象, 整个程序创建的pool都放在同一个切片里面, 操作该对象需要加锁

type Pool struct {
	noCopy noCopy   // no copy 特性, 不能进行拷贝

	local     unsafe.Pointer    // poolLocal的数组对象
	localSize uintptr           // 与gomaxprocs相等

	victim     unsafe.Pointer   // 
	victimSize uintptr

	New func() any // 创建新对象的方法
}

type poolLocal struct {
    poolLocalInternal   // 主字段
    pad [128 - unsafe.Sizeof(poolLocalInternal{})%128]byte // 对齐用
}

type poolLocalInternal struct {
	private any       // 只能由相应的P使用, 相当于private cache
	shared  poolChain // 如果private为空, 就需要在shared 无锁链表中找.
}
```

不考虑victim的情况下:
- 先从当前P的poolLocalInternal.private找(相当于L1 cache)
- 再从当前P对应的poolLocalInternal.shared中找(相当于L2 cache)
- 再没找到, 则去其他P的poolLocalInternal.shared中找 相当于L3 cache

shared后续优化成了一个无锁队列, 通过`pushHead`写入obj, 通过`popTail`弹出obj. 早期的shared是带锁的, 瞬时请求太大会导致延迟波动.

gc时将local与localSize直接平移到victim与victimSize, 如果之前有值,则直接丢弃掉.如果local为空,则会去victim中找.

