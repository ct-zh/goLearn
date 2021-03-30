# go的调度模型

## 简单讲完GMP模型
> [可视化分析GMP流程代码](./trace/.)

### GMP是什么
- G: goroutine
- M: 线程, 在进程内核空间中运行
- P: 处理器processor

有关名词:
1. global run queue, 全局队列,存放等待运行的G;
2. P的本地队列(local run queue),存放当前P等待运行的G,最大256;
3. P列表: `GOMAXPROCS`确定值后将创建所有P,这些P就是P list;
4. 时间片: 每个G都有执行时间上限,一个G最多占用CPU 10ms;

### GMP的调度流程
1. 执行`go func`,也就是新建G;
2. G加入当前P的本地队列,如果本地队列满了,则需要加入全局队列;
3. P从队列里面取出G,给M执行;
4. 如果执行完毕,则直接返回:
5. 如果G在时间片内未执行完毕,则放回P的本地列表;
6. 如果M产生阻塞:
  1. 此时runtime会新开一个线程M2/或者从空闲线程池里取一个线程M2,将M对应的P分配给M2; 
  2. 当阻塞解决, G尝试获取一个空闲P.如果没有,则G进入全局队列,M进入空闲线程池;

### 调度策略:
1. work stealing: 当P从队列里取G时,如果队列为空,则去全局队列里取G; 如果都为空, 则P从其他P的队列偷取G;
2. hand off: 当M因G产生阻塞时,P会去绑定其他空闲的M;
3. 抢占: Golang不存在抢占, 而是时间片主动让出, 一个G最多连续运行10ms就要重回队列;

关于lrq为空时,先去grq里取还是先去偷的问题,网上部分文章含糊不清,可以直接看源码: 版本1.14 `src/runtime/proc.go:findrunnable()`
```go
// local runq
if gp, inheritTime := runqget(_p_); gp != nil {
  return gp, inheritTime
}

// global runq
if sched.runqsize != 0 {
  lock(&sched.lock)
  gp := globrunqget(_p_, 0)
  unlock(&sched.lock)
  if gp != nil {
    return gp, false
  }
}

// ...

// Steal work from other P's.
if gp := runqsteal(_p_, p2, stealRunNextG); gp != nil {
  return gp, false
}
```



### GMP的生命周期
1. 启动程序,系统创建进程和线程,这个线程称为M0;
2. M0启动第一个G也就是G0,G0只做初始化,根据`GOMAXPROCS`创建P列表以及对应的队列;
3. 创建main函数对应的G,加入P的本地队列:
4. P把本地队列里的G交给M执行,M开始调度-执行-执行完毕会销毁,未执行完毕则保存上下文-返回 四个步骤;

ps. 一个Go程序只有一个M0,用于初始化runtime. 每个M都会有一个G0,用于初始化和调度. 上面的G0指的是M0的G0.


### 何时创建GMP?对应的数量是多少?
程序开始时会创建M0; 如果没有足够的M执行P的任务，则会创建M;

在M0G0初始化的时候根据`GOMAXPROCS`的值就能确定P的值,并执行对应P的初始化,之后P的数量保持不变;

P和M的数量没有关联, M的数量应该大于等于P,当M出现阻塞时会新建M去绑定P.


## 源码分析
> [源码位置](https://github.com/golang/go/blob/go1.14.15/src/runtime/runtime2.go#L395)

- G(go routine)代表协程,是go运行的最小单元;值得注意的是main函数本身也是一个G;

- M(machine)在当前版本的golang里等同于*系统线程*;
  M可以运行两种代码:
    - 原生代码(例如阻塞的syscall),不需要P;
    - go代码,即G,需要一个P;

- P(precess)代表M运行G所需要的资源;数量通过环境变量`GOMAXPROC`修改,默认等于核心数(实际无任何关联);


### G
应用程序的内存会分成堆区(Heap)和栈区(Stack)两个部分，程序在运行期间可以主动从堆区申请内存空间，这些内存由内存分配器分配并由垃圾收集器负责回收。栈区的内存由编译器自动进行分配和释放，栈区中存储着函数的参数以及局部变量，它们会随着函数的创建而创建，函数的返回而销毁。

> 堆和栈都是编程语言里的虚拟概念，并不是说在物理内存上有堆和栈之分，两者的主要区别是栈是每个线程或者协程独立拥有的，从栈上分配内存时不需要加锁。而整个程序在运行时只有一个堆，从堆中分配内存时需要加锁防止多个线程造成冲突，同时回收堆上的内存块时还需要运行可达性分析、引用计数等算法来决定内存块是否能被回收，所以从分配和回收内存的方面来看栈内存效率更高。




#### G stack的大小
`runtime/stack.go`中申明了常数`_StackMin = 2048`,代表一个G stack大小为2kb

x86_64架构下线程的默认栈大小为2M,所以G stack比线程轻量很多;

#### 分段栈和连续栈
分段栈指多个协程分配的栈空间是不连续的,多个栈空间以双向链表的形式串联起来. 此时golang把栈内存大小设置为8KB,但仍然可能出现问题;

分段栈虽然能够按需为当前 goroutine 分配内存并且及时减少内存的占用，但是它也存在一个比较大的问题：如果当前 goroutine 的栈几乎充满，那么任意的函数调用都会触发栈的扩容，当函数返回后又会触发栈的收缩，如果在一个循环中调用函数，栈的分配和释放就会造成巨大的额外开销，这被称为热分裂问题(Hot split)。

连续栈可以解决分段栈中存在的两个问题，其核心原理就是*每当程序的栈空间不足时，初始化一片比旧栈大两倍的新栈并将原栈中的所有值都迁移到新的栈中*，新的局部变量或者函数调用就有了充足的内存空间。使用连续栈机制时，栈空间不足导致的扩容会经历以下几个步骤：
1. 调用用runtime.newstack在内存空间中分配更大的栈内存空间；
2. 使用runtime.copystack将旧栈中的所有内容复制到新的栈中；
3. 将指向旧栈对应变量的指针重新指向新栈；
4. 调用runtime.stackfree销毁并回收旧栈的内存空间；

此时golang设置连续栈的大小为2KB

#### 内存管理
每个goroutine都维护着自己的栈区，栈结构是连续栈，是一块连续的内存，在goroutine的类型定义的源码里我们可以找到标记着栈区边界的stack信息，stack里记录着栈区边界的高位内存地址和低位内存地址
```go
type g struct {
 stack       stack
  ...
}

type stack struct {
 lo uintptr
 hi uintptr
}
```
栈空间在运行时中包含两个重要的全局变量，分别是 runtime.stackpool 和runtime.stackLarge，这两个变量分别表示全局的栈缓存和大栈缓存，前者可以分配小于 32KB 的内存，后者用来分配大于 32KB 的栈空间.

从调度器和内存分配的角度来看，如果运行时只使用全局变量来分配内存的话，势必会造成线程之间的锁竞争进而影响程序的执行效率，栈内存由于与线程关系比较密切，所以在每一个线程缓存 runtime.mcache 中都加入了栈缓存减少锁竞争影响。

#### 栈扩容
编译器会为函数调用插入运行时检查runtime.morestack，它会在几乎所有的函数调用之前检查当前goroutine 的栈内存是否充足，如果当前栈需要扩容，会调用runtime.newstack创建新的栈`func newstack()`; 

旧栈的大小是通过我们上面说的保存在goroutine中的stack信息里记录的栈区内存边界计算出来的，然后用旧栈两倍的大小创建新栈，创建前会检查是新栈的大小是否超过了单个栈的内存上限。

```go
oldsize := gp.stack.hi - gp.stack.lo
newsize := oldsize * 2
if newsize > maxstacksize {
  print("runtime: goroutine stack exceeds ", maxstacksize, "-byte limit\n")
  throw("stack overflow")
}
```

如果目标栈的大小没有超出程序的限制，会将 goroutine 切换至 _Gcopystack 状态并调用 runtime.copystack 开始栈的拷贝，在拷贝栈的内存之前，运行时会先通过runtime.stackalloc 函数分配新的栈空间

```go
func copystack(gp *g, newsize uintptr) {
 old := gp.stack
 used := old.hi - gp.sched.sp
  // 创建新栈
 new := stackalloc(uint32(newsize))
 ...
  // 把旧栈的内容拷贝至新栈
 memmove(unsafe.Pointer(new.hi-ncopy), unsafe.Pointer(old.hi-ncopy), ncopy)
  ...
  // 调整指针
  adjustctxt(gp, &adjinfo)
  // groutine里记录新栈的边界
  gp.stack = new
  ...
  // 释放旧栈
  stackfree(old)
}
```

新栈的初始化和数据的复制是一个比较简单的过程，整个过程中最复杂的地方是将指向源栈中内存的指针调整为指向新的栈，这一步完成后就会释放掉旧栈的内存空间了。

#### 栈缩容
在goroutine运行的过程中，如果栈区的空间使用率不超过1/4，那么在垃圾回收的时候使用runtime.shrinkstack进行栈缩容，当然进行缩容前会执行一堆前置检查，都通过了才会进行缩容
```go
func shrinkstack(gp *g) {
 ...
 oldsize := gp.stack.hi - gp.stack.lo
 newsize := oldsize / 2
 if newsize   return
 }
 avail := gp.stack.hi - gp.stack.lo
 if used := gp.stack.hi - gp.sched.sp + _StackLimit; used >= avail/4 {
  return
 }

 copystack(gp, newsize)
}
```
如果要触发栈的缩容，*新栈的大小会是原始栈的一半*，不过如果新栈的大小低于程序的最低限制 2KB，那么缩容的过程就会停止。缩容也会调用扩容时使用的 runtime.copystack 函数开辟新的栈空间，将旧栈的数据拷贝到新栈以及调整原来指针的指向。




## reference
> [深入理解GPM模型](https://www.bilibili.com/video/BV19r4y1w7Nx?p=6&spm_id_from=pageDriver)
> [Golang的协程调度器原理及GMP设计思想](https://www.kancloud.cn/aceld/golang/1958305#5GMP_305)


> [Golang源码探索(二) 协程的实现原理](https://www.cnblogs.com/zkweb/p/7815600.html)
> [G、P、M](https://github.com/friendlyhank/go-source/blob/master/runtime/golang%20pgm.md)
> [go channel](https://github.com/friendlyhank/toBeTopgopher/blob/master/golang/source/go_channel.md)
> [golang select](https://github.com/friendlyhank/toBeTopgopher/blob/master/golang/source/golang_select.md)
> [调度器](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#657-%E7%BA%BF%E7%A8%8B%E7%AE%A1%E7%90%86)