# go的调度模型

## base
> [源码位置](https://github.com/golang/go/blob/go1.14.15/src/runtime/runtime2.go#L395)

- G(go routine)代表协程,是go运行的最小单元;值得注意的是main函数本身也是一个G;

- M(machine)在当前版本的golang里等同于*系统进程*;
  M可以运行两种代码:
    - 原生代码(例如阻塞的syscall),不需要P;
    - go代码,即G,需要一个P;

- P(precess)代表M运行G所需要的资源;数量通过环境变量`GOMAXPROC`修改,默认等于核心数(实际无任何关联);

// todo netpoller

## reference
> [Golang源码探索(二) 协程的实现原理](https://www.cnblogs.com/zkweb/p/7815600.html)
> [G、P、M](https://github.com/friendlyhank/go-source/blob/master/runtime/golang%20pgm.md)
> [go channel](https://github.com/friendlyhank/toBeTopgopher/blob/master/golang/source/go_channel.md)
> [golang select](https://github.com/friendlyhank/toBeTopgopher/blob/master/golang/source/golang_select.md)