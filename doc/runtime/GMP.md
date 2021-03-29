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

// todo netpoller


## reference
> [深入理解GPM模型](https://www.bilibili.com/video/BV19r4y1w7Nx?p=6&spm_id_from=pageDriver)
> [Golang的协程调度器原理及GMP设计思想](https://www.kancloud.cn/aceld/golang/1958305#5GMP_305)


> [Golang源码探索(二) 协程的实现原理](https://www.cnblogs.com/zkweb/p/7815600.html)
> [G、P、M](https://github.com/friendlyhank/go-source/blob/master/runtime/golang%20pgm.md)
> [go channel](https://github.com/friendlyhank/toBeTopgopher/blob/master/golang/source/go_channel.md)
> [golang select](https://github.com/friendlyhank/toBeTopgopher/blob/master/golang/source/golang_select.md)
> [调度器](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/#657-%E7%BA%BF%E7%A8%8B%E7%AE%A1%E7%90%86)