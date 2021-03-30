# sync包
## 用法解析
> sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步*使用channel通信*更好一些。

go run -race 查看冲突

### type
- Locker接口: 代表一个可以加锁解锁的对象;
- Once: once.Do(func())无论do被调用几次,里面的函数都只会执行一次;[例子](./once/once.go)
- Mutex: 提供线程安全的互斥锁;
- RWMutex: 读写互斥锁,可以只加写锁;
- Cond: 实现多个线程wait单个线程通知或者广播;[例子](./cond/simple/simple.go)
- WaitGroup: 父线程用来等待子线程的运行结束;
- Pool: 多线程共享的item池;[例子](./pool/simple/simple.go)

- [使用sync.Pool来复用对象](https://geektutu.com/post/hpg-sync-pool.html)


### atomic
atomic包的方法保证了对几种基础类型提供原子操作,在多线程的情况下保证数据的一致性;atomic提供了五类原子操作分别是:
- Add, 增加和减少
- CompareAndSwap, 比较并交换
- Swap, 交换
- Load , 读取
- Store, 存储


## 源码分析
### Locker
> src/sync/mutex.go:Locker
接口类型,代表锁; mutex、rwmutex和cond都是它的实现;

### sync
> see https://colobu.com/2018/12/18/dive-into-sync-mutex/
sync的公平处理与饥饿机制

> 互斥锁有两种状态：正常状态和饥饿状态。
>
> 在正常状态下，所有等待锁的goroutine按照FIFO顺序等待。唤醒的goroutine不会直接拥有锁，而是会和新请求锁的goroutine竞争锁的拥有。新请求锁的goroutine具有优势：它正在CPU上执行，而且可能有好几个，所以刚刚唤醒的goroutine有很大可能在锁竞争中失败。在这种情况下，这个被唤醒的goroutine会加入到等待队列的前面。 如果一个等待的goroutine超过1ms没有获取锁，那么它将会把锁转变为饥饿模式。
>
> 在饥饿模式下，锁的所有权将从unlock的gorutine直接交给交给等待队列中的第一个。新来的goroutine将不会尝试去获得锁，即使锁看起来是unlock状态, 也不会去尝试自旋操作，而是放在等待队列的尾部。
>
> 如果一个等待的goroutine获取了锁，并且满足一以下其中的任何一个条件：(1)它是队列中的最后一个；(2)它等待的时候小于1ms。它会将锁的状态转换为正常状态。
>
> 正常状态有很好的性能表现，饥饿模式也是非常重要的，因为它能阻止尾部延迟的现象。



### Cond
> src/sync/cond.go


### Pool 

#### reference
- https://studygolang.com/pkgdoc
- [Golang Sync.Pool浅析](https://segmentfault.com/a/1190000019973632)
- [Golang 的 sync.Pool设计思路与原理](https://blog.csdn.net/u010853261/article/details/90647884)
- [官方包sync.Pool的实现原理和适用场景](https://blog.csdn.net/yongjian_lian/article/details/42058893?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com)
- [理解 Sync.Pool 的设计](https://juejin.cn/post/6844903864634720263)



