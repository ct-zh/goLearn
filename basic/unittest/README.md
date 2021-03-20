# 测试
## 单元测试
这里介绍几个常用的参数：
- bench regexp 执行相应的 benchmarks，例如 -bench=.；
- cover 开启测试覆盖率；
- run regexp 只运行 regexp 匹配的函数，例如 -run=Array 那么就执行包含有 Array 开头的函数；
- v 显示测试的详细命令。

单元测试框架提供的日志方法: 
|方  法|备  注|
|---|---|
|Log|测试来说，只会在失败或者设置了-test.v标志的情况下被打印出来；对于基准测试来说，为了避免 -test.v标志的值对测试的性能产生影响，格式化文本总会被打印出来。|
|Logf|格式化打印日志|
|Error|相当于在调用Log之后调用Fail(将当前测试标识为失败，但是仍继续执行该测试)|
|Errorf|相当于在调用Logf之后调用Fail|
|Fatal|相当于在调用Log之后调用FailNow(立即结束测试)|
|Fatalf	|格式化打印致命日志，同时结束测试|

文件命名：`name_test.go`,导入： `import testing`

### 代码覆盖率
`go test -coverprofile=c.out`
`go tool cover -html=c.out`

### 性能测试
`go test -bench . -cpuprofile cpu.out`
`go tool pprof cpu.out`
`web`


## 基准测试
`go test -v -bench=. benchmark_test.go`

1. 通过`-benchtime`参数可以自定义测试时间，例如`go test -v -bench=. -benchtime=5s benchmark_test.go`
2. 在命令行中添加`-benchmem`参数以显示内存分配情况，参见下面的指令：
    ```bash
    $ go test -v -bench=Alloc -benchmem benchmark_test.go
    goos: linux
    goarch: amd64
    Benchmark_Alloc-4 20000000 109 ns/op 16 B/op 2 allocs/op
    PASS
    ok          command-line-arguments        2.311s
    ```
    代码说明如下：
    - 第 1 行的代码中-bench后添加了 Alloc，指定只测试 Benchmark_Alloc() 函数。
    - 第 4 行代码的“16 B/op”表示每一次调用需要分配 16 个字节，“2 allocs/op”表示每一次调用有两次分配。

# net/http/pprof 
> http 服务器性能检测
`go tool pprof `分析性能


## go mock
[demo](./gomock_demo/t1.go)

- [Go Mock (gomock)简明教程](https://geektutu.com/post/quick-gomock.html)
- [使用Golang的官方mock工具--gomock](https://www.jianshu.com/p/598a11bbdafb)
- [go mock github](https://github.com/golang/mock)

## go convey
[demo](./goconvey_demo/t3_test.go)

- [优雅的单元测试](https://studygolang.com/articles/1513)
- [官方 wiki](https://github.com/smartystreets/goconvey/wiki/Assertions)
- [docs](https://gowalker.org/github.com/smartystreets/goconvey)

## reference
- [go test完全攻略](http://c.biancheng.net/view/124.html)
- [使用go做测试](https://zhuanlan.zhihu.com/p/168539526)
