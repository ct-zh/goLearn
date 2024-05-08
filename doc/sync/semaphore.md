## semaphore

是Go语言锁实现的基础, 所有同步原语的基础设施. 源码处于`runtime/sema.go`

### 基本结构

顶层结构如下:
```go
type semTable [semTabSize]struct { // 固定大小的结构体数组
	root semaRoot
	pad  [cpu.CacheLinePadSize - unsafe.Sizeof(semaRoot{})]byte
}
```

- semTable是一个结构体数组, 其固定大小semTabSize写死为251, 也就是该数组有0-250、一共251个元素. Go语言runtime中存在一个semTable类型, 名字也是semTable的全局变量; 每次对锁相关内容进行操作时, 对应的g会通过一种hash算法计算出自身对应的心好像应该落在semTable数组中的哪一个子元素上来进行锁操作; 

- root字段对应的semaRoot类型代表semTable数组中单个子元素的结构;
- pad字段是一个固定大小的字节数组, 用于填充semTable, 其大小为cpu缓存行大小减去一个semaRoot的大小, 也就是说单个`semTable`结构体的大小刚好是一个cpu缓存行的大小. 

> 为什么不直接使用semaRoot结构作为数组的元素, 而是额外套一层并且增加一个固定大小的pad字段?这里涉及到一个cpu缓存行的问题:
>
> ### cpu缓存行
>
> 现代 CPU 使用CPU缓存来提高内存访问速度。CPU 优先从CPU缓存中读取数据，如果数据不在CPU缓存中，则需要从主内存中读取，这会带来明显的延迟。这些CPU缓存通常以缓存行（Cache Line）为单位进行操作，一个缓存行通常包含连续的多个字节的数据。当 CPU 需要读取或写入数据时，它会将整个缓存行加载到缓存中，即使只需要其中的一部分。
>
> **伪共享问题：**
>
> 如果多个变量共享同一缓存行，但并非所有线程都同时访问这些变量，则会出现伪共享问题。当两个或多个线程在同一缓存行的不同位置读写数据时，即使他们访问的是不同的数据，由于缓存系统以缓存行为单位进行数据同步，这就导致了不必要的数据同步操作。当一个线程修改共享变量时，整个缓存行都会被标记为脏，并回写到主内存。其他线程即使没有修改共享变量，也需要重新加载整个缓存行，这会带来不必要的开销。
>
> #### 举例
>
> 假设有两个线程，线程 A 和线程 B，它们分别在同一缓存行的不同位置写入数据。即使 A 和 B 写入的是不同的数据，但由于它们位于同一缓存行，所以每次 A 写入数据时，缓存行就会被标记为 “脏”，需要同步到主内存。然后，当 B 需要写入数据时，它首先需要从主内存中获取最新的缓存行，然后再写入数据。这个过程会反复进行，导致了大量的不必要的数据同步操作，从而降低了程序的性能。
>
> 因此，为了避免伪共享，我们通常需要确保不同线程操作的数据位于不同的缓存行中。这就是`semTable`结构体定义中 `pad` 字段的作用，它通过增加额外的填充，确保每个 `semaRoot` 实例都能独占一个 CPU 缓存行，从而避免了伪共享的问题。

对于`semTable.semaRoot`字段,  其结构为:

```go
type semaRoot struct {
	lock  mutex		// 互斥锁, 提供线程安全
	treap *sudog        // 主要结构, 是一个tree+heap 树堆的根节点, 这个平衡树中的每个节点都是一个唯一的等待者（waiter）。
	nwait atomic.Uint32 // 原子操作的无符号32位整数，表示waiter的数量。由于它是原子操作的，所以可以在不加锁的情况下读取它。
}
```

`semaRoot`的主要目的是管理在特定地址上等待的 `sudog`。每个 `sudog` 可能会通过 `s.waitlink` 指向在同一地址上等待的其他 `sudog` 的列表。对于这些 `sudog` 的内部列表操作都是 O(1) 的时间复杂度。而对于顶级的 `semaRoot` 列表的扫描则是 O(log n) 的时间复杂度，其中 n 是阻塞在给定 `semaRoot` 哈希值上的不同地址的 `goroutine` 的数量。

`treap`是按照lock addr 排列的一棵二叉搜索树. 如果现在某个g要加锁, 锁存在一个对应的信号成员，Go语言会通过这个信号成员的地址在这个二叉搜索树上找到它处于的节点; 如果程序有任何锁发生阻塞，最终都是挂在semaRoot上面。sudog根据信号成员地址找到属于它的节点以后，挂在其等待列表后面.

treap作为一棵二叉搜索树, 为了保证treap的平衡，在g挂上树时, Go语言会给对应节点赋值了一个ticket, 这个ticket是在sudog初始化的时候，用fastrand的函数来生成的。

### semaphore相关函数

```go
//go:linkname sync_runtime_Semacquire sync.runtime_Semacquire
func sync_runtime_Semacquire(addr *uint32) {
	semacquire1(addr, false, semaBlockProfile, 0, waitReasonSemacquire)
}

//go:linkname sync_runtime_Semrelease sync.runtime_Semrelease
func sync_runtime_Semrelease(addr *uint32, handoff bool, skipframes int) {
	semrelease1(addr, handoff, skipframes)
}
```

在`sema.go`中可以看到很多格式类似上面的函数,这些函数是 Go 语言中信号量实现的一部分。它们用于获取和释放信号量，信号量是一种同步原语，用于控制对共享资源的访问。

- `go:linkname` 表明这些函数导出到其他包，允许它们从 `sync` 包外部调用;
- `sync_runtime_Semacquire`用于获取信号量。接受一个 `uint32` 值的指针作为参数，该指针表示信号量的地址; 
- `sync_runtime_Semrelease`用于释放信号量。它也接受一个 `uint32` 值的指针作为参数，表示信号量的地址; `handoff`: 是否将信号量移交给另一个 goroutine; `skipframes`: 计数时跳过的堆栈帧数(用于分析);

从这些函数可以得知sema的所有信号量操作都会指向两个函数: `semacquire1`与 `semrelease1`, 下面我们着重分析这两个函数. 


### semacquire1 

`semacquire1` 函数是 Go 语言中用于获取信号量的核心函数之一。它尝试获取由 `addr` 参数指定的信号量。该函数会执行一系列操作，包括检查信号量状态、加入等待队列以及在获取不到信号量时进入睡眠等待。

**函数参数：**

- `addr`: 指向信号量地址的指针
- `lifo`: 布尔值，指示是否使用后进先出 (LIFO) 策略管理等待队列
- `profile`: 用于标识需要进行阻塞或互斥锁的性能分析
- `skipframes`: 跳过的堆栈帧数(用于分析)
- `reason`: 表示进行信号量获取操作的原因

**代码解析:**

```go
func semacquire1(addr *uint32, lifo bool, profile semaProfileFlags, skipframes int, reason waitReason) {
	gp := getg()
	if gp != gp.m.curg {	// gp.m 代表当前g绑定的m, m的curg字段代表当前m正在运行的g
		throw("semacquire not on the G stack")
	}
	if cansemacquire(addr) {   // Easy case.
		return
	}
  // ....
}

func cansemacquire(addr *uint32) bool {
	for {
		v := atomic.Load(addr)
		if v == 0 {
			return false
		}
		if atomic.Cas(addr, v, v-1) {
			return true
		}
	}
}
```

- 获取当前 goroutine, 确保当前的操作是在当前 goroutine 的栈上进行的(gp.m代表当前g对应的m, 如果m当前的g(curg字段)不等于该g, 说明不是在当前g上运行的); 如果不是,则抛出异常, **因为信号量获取操作必须在调用它的线程上进行. **

- easy case 简单的情况: 如果可以通过CAS操作直接操作信号量，那么就直接返回;

```go
func semacquire1(addr *uint32, lifo bool, profile semaProfileFlags, skipframes int, reason waitReason) {
	// ......
  // hard case
  s := acquireSudog()
	root := semtable.rootFor(addr)	
	t0 := int64(0)
	s.releasetime = 0
	s.acquiretime = 0
	s.ticket = 0
	if profile&semaBlockProfile != 0 && blockprofilerate > 0 {
		t0 = cputicks()
		s.releasetime = -1
	}
	if profile&semaMutexProfile != 0 && mutexprofilerate > 0 {
		if t0 == 0 {
			t0 = cputicks()
		}
		s.acquiretime = t0
	}
  // ......
} 
```

- `acquireSudog`   获取一个sudog对象; 具体解析见下面的关联函数;
- `semtable.rootFor` 找到信号量的根节点;
- 根据 `profile` 标志判断是否需要进行阻塞态分析和互斥锁分析，并设置相应的时间戳

下面这段是quire的核心代码:

```go
func semacquire1(addr *uint32, lifo bool, profile semaProfileFlags, skipframes int, reason waitReason) {
	// ......
	for {
		lockWithRank(&root.lock, lockRankRoot)	// 获取锁
		// 添加到nwait以禁用semrelease中的“easy case”。
		root.nwait.Add(1)
		// 检查cansemaccessed以避免错过唤醒。
		if cansemacquire(addr) {
			root.nwait.Add(-1)
			unlock(&root.lock)
			break
		}
		root.queue(addr, s, lifo)
		goparkunlock(&root.lock, reason, traceEvGoBlockSync, 4+skipframes)
		if s.ticket != 0 || cansemacquire(addr) {
			break
		}
	}
  // ......
} 
```

- for循环不断尝试获取信号量
- `lockWithRank ` 获取root节点的锁
- 将等待者计数 (`root.nwait`) 加 1，以防止在 `semrelease` 函数中误判没有等待者
- 再次尝试获取信号量 (`cansemacquire(addr)` easy case的操作 )。如果成功，则释放锁, 等待计数-1, 并跳出循环
- 如果未成功获取信号量，进入阻塞sleep流程: 将等待描述符 (`s`) 加入到信号量的等待队列 (`root.queue`) 中，并使用 `goparkunlock` 函数使当前 Goroutine 进入睡眠等待

```go
func semacquire1(addr *uint32, lifo bool, profile semaProfileFlags, skipframes int, reason waitReason) {
	// ......
	if s.releasetime > 0 {
		blockevent(s.releasetime-t0, 3+skipframes)
	}
	releaseSudog(s)
}
```

如果启用了阻塞分析，那么就记录阻塞事件; 然后`releaseSudog`释放 `sudog` ，表示当前的等待者已经完成了信号量的获取操作;

> 该函数的流程是, 传入信号量地址并试图通过CAS对其进行操作( Easy Case ), 如果不能操作说明存在其他线程占用; 
>
> 于是该函数去申请一个空闲的sudog做专门的无限循环, 循环内容仍然是试图通过CAS对信号量进行操作;
>
> 每次循环都要给root加锁,并且给root的waiter数量+1; 
>
> 如果CAS操作失败了, 则会将当前sudog挂载到root后,并且将当前sudog gopark暂停;
>
> 如果CAS操作成功, 退出无限循环, 释放sudog;


### semrelease1 

`semrelease1`函数是 Go 语言中用于释放信号量的核心函数之一。它用于释放由 addr 参数指定的信号量，并唤醒等待该信号量的 Goroutine。

**函数参数：**

- addr: 指向信号量地址的指针
- handoff: 布尔值，指示是否将 Goroutine 所有权转移给等待该信号量的 Goroutine
- skipframes: 跳过的堆栈帧数，用于分析目的

**代码解析:**

```go
func semrelease1(addr *uint32, handoff bool, skipframes int) {
	root := semtable.rootFor(addr)
	atomic.Xadd(addr, 1)
	if root.nwait.Load() == 0 { 	// Easy case: no waiters? 此检查必须在xadd之后进行，以避免错过唤醒
		return
	}
  // .....
}
```

easy case操作:

- 获取根节点
- 使用原子操作 (`atomic.Xadd(addr, 1)`) 将信号量值加 1
- 由于 `semacquire` 函数在循环中可能会错过唤醒，因此这里需要**在**增加信号量值**之后** 再次检查等待者数量 (`root.nwait.Load()`)
- 如果没有等待者 (`root.nwait` 为 0)说明没有阻塞的sudog，则直接返回，释放操作完成

```go
func semrelease1(addr *uint32, handoff bool, skipframes int) {
  // ......
  lockWithRank(&root.lock, lockRankRoot) // harder case: 寻找waiter并唤醒
	if root.nwait.Load() == 0 {
		unlock(&root.lock) 		// 计数已经被另一个goroutine消耗，所以不需要唤醒另一个goroutine
		return
	}
	s, t0 := root.dequeue(addr)
	if s != nil {
		root.nwait.Add(-1)
	}
	unlock(&root.lock)
  // ......
} 
```

- 获取 `root` 节点的锁 (`lockWithRank`)
- 再次检查等待者数量 (`root.nwait.Load()`)。如果此时没有等待者，可能是因为其他 Goroutine 已经消费了该信号量，不需要再唤醒其他等待者。释放锁并返回
- 如果还有等待者，则从队列中取出一个sudog并记录唤醒时间 (`t0`)
- 如果取到了等待描述符 (`s` 不为 nil)，则将等待者数量减 1 (`root.nwait.Add(-1)`)。
- 释放锁 (`unlock(&root.lock)`)

下面是处理从dequeue中取到sudog waiter后的操作

```go
func semrelease1(addr *uint32, handoff bool, skipframes int) {
  // .......
	if s != nil {
		acquiretime := s.acquiretime
		if acquiretime != 0 {
			mutexevent(t0-acquiretime, 3+skipframes) // 性能分析
		}
		if s.ticket != 0 {
			throw("corrupted semaphore ticket")
		}
		if handoff && cansemacquire(addr) {	// handoff=true饥饿模式 
			s.ticket = 1
		}
		readyWithTime(s, 5+skipframes)
		if s.ticket == 1 && getg().m.locks == 0 {
      // 直接进行G的切换
      // readyWithTime已将waiter G添加为当前P中的runNext;现在调用调度程序，直接运行waiter G
			goyield()
		}
	}
}

func readyWithTime(s *sudog, traceskip int) {
	if s.releasetime != 0 {
		s.releasetime = cputicks()
	}
	goready(s.g, traceskip)
}
```

- 获取等待者开始等待的时间 (`acquiretime`)，用于性能分析; 如果进行了互斥锁分析，则记录唤醒等待者消耗的时间
- 检查sudog的 (`s.ticket`) 是否为 0。如果不是，则抛出异常，表示信号量ticket损坏
- 检查 `handoff` 标志。如果是 `true` (饥饿模式)，并且可以直接再次获取信号量 (`cansemacquire(addr)`)，则将等待描述符的(`s.ticket`) 设置为 1
- 将等待描述符 (`s`) 放入就绪队列 (`readyWithTime`,  goReady的包装)，等待 Goroutine 调度
- 如果等待描述符的ticket为 1 并且当前 Goroutine 没有持有的互斥锁 (`getg().m.locks == 0`)，则进行 Goroutine 所有权转移 (仅限饥饿模式)

  - 将被唤醒的 Goroutine (`s`) 标记为当前处理器 (P) 的下一个可运行 Goroutine (`runnext`)

  - 调用 `goyield` 函数进行调度，使被唤醒的 Goroutine 立即开始运行

  - 在饥饿模式下，通过所有权转移可以避免信号量长期占用处理器

### 关联方法 rootFor

`semaTable`存在唯一方法`rootFor`, 用于根据地址获取信号量根节点:

```go
const semTabSize = 251

type semTable [semTabSize]struct {
	root semaRoot
	pad  [cpu.CacheLinePadSize - unsafe.Sizeof(semaRoot{})]byte
}

func (t *semTable) rootFor(addr *uint32) *semaRoot {
	return &t[(uintptr(unsafe.Pointer(addr))>>3)%semTabSize].root
}
```

这是一种hash算法, 将addr的地址转换成`uintptr`类型, 再对其进行位运算`(>>>3)`, 将获得的结果余`semTabSize`, 即可获得该地址对应的信号量根节点.

### treap的相关操作

`seamRoot.treap` 是一个树堆, 也就是说它既有树的结构, 也有堆的结构; 

- **树（Tree）特性**: Treap 是一种二叉搜索树，这意味着它的每个节点都有两个子节点（可能为空），并且对于每个节点，其左子节点的值小于节点值，右子节点的值大于节点值。这使得我们可以快速地在 Treap 中查找、插入和删除元素。
- **堆（Heap）特性**: Treap 的每个节点都有一个优先级值，这个优先级值是随机生成的。在 Treap 中，父节点的优先级总是大于或等于其子节点的优先级。这个特性使得 Treap 的形状总是接近于平衡，从而保证了操作的效率。

semaphore使用了四个与树堆相关的方法:

```go
func (root *semaRoot) queue(addr *uint32, s *sudog, lifo bool)
func (root *semaRoot) dequeue(addr *uint32) (found *sudog, now int64)
func (root *semaRoot) rotateLeft(x *sudog)
func (root *semaRoot) rotateRight(y *sudog)
```

在queue时进行左旋, 在dequeue时进行右旋;

### 关联函数 acquireSudog

proc函数, 用来获取一个`sudog`对象，如果缓存中没有，则创建一个新的`sudog`对象。源码处于`runtime/proc.go`

- 首先调用 `acquirem()` 函数，获取当前M（machine，即操作系统线程）的上下文; 
- 获取当前P（processor，处理器）的指针 `pp`;
- 检查当前P的 `sudogcache` 缓存中是否有可用的`sudog`对象。
- 如果缓存中没有`sudog`对象，则尝试从全局的`sched.sudogcache`中获取一批。这里的 `sched` 是调度器的全局变量，`sudogcache` 是一个链表，存储了可用的`sudog`对象。
- 如果全局缓存也为空，则分配一个新的`sudog`对象。
- 从缓存中取出一个`sudog`对象，并进行一些检查。如果`elem`字段不为`nil`，则抛出异常。
- 释放之前获取的M的上下文。
- 返回获取到的`sudog`对象。

整个函数的目的是为了尽量避免频繁地创建和销毁`sudog`对象，而是通过缓存的方式复用已有的对象，从而提高性能。

### 关联函数 releaseSudog

这个函数是用来释放一个`sudog`类型的对象，并将其放回到缓存中以便后续重用。源码处于`runtime/proc.go`

1. 对`sudog`对象的各个字段进行检查，确保它们都是`nil`，如果有任何一个字段不是`nil`，则会抛出异常。

2. 获取当前goroutine的`g`结构体，`getg()`函数用于获取当前goroutine的`g`结构体。

3. 调用`acquirem()`函数获取当前M的上下文，这是为了避免在函数执行期间将该goroutine调度到另一个P上。

4. 获取当前P的指针 `pp`。

5. 检查当前P的`sudogcache`缓存是否已满，如果已满，则将一半的本地缓存转移到全局缓存`sched.sudogcache`中。

6. 将待释放的`sudog`对象添加到当前P的`sudogcache`缓存中。

7. 最后释放之前获取的M的上下文。

这个函数的主要目的是尽量避免频繁地创建和销毁`sudog`对象，而是通过缓存的方式复用已有的对象，从而提高性能。

### 关联函数 goparkunlock

将当前goroutine置于等待状态并解开锁锁。通过调用goready可以使goroutine再次运行。

其实就是包了一下`gopark`  

### 关联函数 lockWithRank

这是一个实现互斥锁的函数，用于获取锁。让我们逐步解析：

1. 首先，获取当前goroutine的`g`结构体。

2. 检查当前goroutine持有的锁的数量是否小于0，如果小于0，则抛出异常。这个检查是为了确保锁的数量没有出现异常情况。

3. 将当前goroutine持有的锁的数量加1。

4. 尝试通过原子操作 `atomic.Casuintptr` 获取锁。如果成功获取锁，则直接返回。

5. 如果无法立即获取到锁，那么需要创建一个信号量 `semacreate(gp.m)`，用于在无法获取锁时进行等待。

6. 在多处理器系统中，会进行自旋尝试获取锁，循环进行一定次数的自旋。如果是单处理器系统，则没有自旋的意义，直接跳过自旋。

7. 进入一个无限循环，尝试获取锁。循环中首先尝试读取锁的状态。

8. 如果锁是未锁定状态，则尝试通过原子操作 `atomic.Casuintptr` 获取锁。如果获取成功，则直接返回。

9. 如果自旋次数小于设定的自旋次数，则通过 `procyield` 进行主动让出CPU时间片，以便让其他goroutine有机会执行。

10. 如果自旋次数超过了设定的自旋次数，但还未到达完全放弃CPU的被动自旋次数，则通过 `osyield` 函数暂时放弃CPU。

11. 如果以上两种自旋都没有获取到锁，说明有其他goroutine持有锁，需要将当前goroutine加入到等待锁的链表中。这里通过原子操作将当前goroutine的指针加入到锁的等待链表中。

12. 最后，如果还是无法获取到锁，则通过 `semasleep` 进行等待。

### 关联函数 blockevent

该函数实现了阻塞事件的跟踪和记录.

- `blockevent` 函数是一个入口函数，用于触发阻塞事件的跟踪和记录。它接受两个参数：`cycles`表示事件发生的时间（以CPU周期为单位），`skip`表示跳过的调用栈帧数。如果 `cycles` 小于等于 0，则将其设置为 1。然后，通过调用 `blocksampled` 函数来判断当前事件是否应该被采样，如果需要采样，则调用 `saveblockevent` 函数保存事件信息。
- `blocksampled` 函数用于判断当前事件是否应该被采样。它接受两个参数：`cycles`表示事件发生的时间，`rate`表示采样率。如果 `rate` 小于等于 0，或者 `(rate > cycles && int64(fastrand()) % rate > cycles)` 条件成立，则返回 false；否则返回 true。这个函数的逻辑是：如果 `rate` 大于 `cycles`，则按照 `cycles / rate` 的概率进行采样；如果 `rate` 小于等于 `cycles`，则始终采样。
- `saveblockevent` 函数用于保存阻塞事件的信息。它接受四个参数：`cycles`表示事件发生的时间，`rate`表示采样率，`skip`表示跳过的调用栈帧数，`which`表示事件的类型（blockProfile 或 mutexProfile）。首先，获取当前goroutine的调用栈信息，并将其保存到一个数组中。然后根据事件的类型和采样率计算事件的统计信息，并将其更新到相应的桶（bucket）中。需要注意的是，在计算事件统计信息时，根据事件是否被采样，对事件的数量和时间进行了相应的缩放，以反映实际采样情况。

### 关联函数 goyield

> Create by chatGPT

这个函数是一个用于让出CPU执行权的函数，类似于 `Gosched` 函数，但有一些不同之处。

1. `goyield` 函数用于让出CPU执行权。在 Go 语言中，goroutine 之间的调度是由 Go 运行时系统负责的。当一个 goroutine 需要等待或者暂时让出 CPU 时，可以调用 `goyield` 函数。

2. 这个函数调用了 `checkTimeouts` 函数，用于检查是否有超时的 goroutine 需要被唤醒。

3. `mcall(goyield_m)` 是一个函数调用，用于执行特定的处理函数 `goyield_m`。`mcall` 函数是一个内部函数，用于在当前的 M（machine，表示一个线程）上执行一个函数。在这里，它用于在当前的 M 上执行 `goyield_m` 函数。

4. `goyield_m` 函数的作用是将当前的 goroutine 放回到当前 P（Processor，表示处理器）的运行队列中，而不是全局的运行队列中。这个操作是通过将当前 goroutine 放回当前 P 的运行队列中来实现的，而不是放回全局的运行队列中。这一点与 `Gosched` 函数的行为不同。

总的来说，`goyield` 函数是用于让出 CPU 执行权的函数，与 `Gosched` 函数相比，它的行为略有不同：它发出 `GoPreempt` 跟踪事件，将当前 goroutine 放回当前 P 的运行队列中。

