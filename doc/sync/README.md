# sync包
## 用法解析
> https://studygolang.com/pkgdoc

sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步*使用channel通信*更好一些。

### type
- Locker接口: 代表一个可以加锁解锁的对象;
- Once: once.Do(func())无论do被调用几次,里面的函数都只会执行一次;[例子](./once/once.go)
- Mutex: 提供线程安全的互斥锁;
- RWMutex: 读写互斥锁,可以只加写锁;
- Cond: 实现多个线程wait单个线程通知或者广播;[例子](./cond/simple/simple.go)
- WaitGroup: 父线程用来等待子线程的运行结束;
- Pool: 多线程共享的item池;[例子](./pool/simple/simple.go)


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


### Cond
> src/sync/cond.go


### Pool 

#### see
- [Golang Sync.Pool浅析](https://segmentfault.com/a/1190000019973632)
- [Golang 的 sync.Pool设计思路与原理](https://blog.csdn.net/u010853261/article/details/90647884)
- [官方包sync.Pool的实现原理和适用场景](https://blog.csdn.net/yongjian_lian/article/details/42058893?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com)
- [理解 Sync.Pool 的设计](https://juejin.cn/post/6844903864634720263)



