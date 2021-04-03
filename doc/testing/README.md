# 测试
## 单元测试与基准测试
测试文件命名统一为`<name>_test.go`,导入： `import testing`包; 

单元测试函数命名为`Test<func name>`,func name遵循大驼峰命名;单元测试方法命名为`Test<class name>_<method name>`,method name按照原始的方法名命名(开放方法是大驼峰,私有方法小驼峰);

基准测试同单元测试,只是前缀改成`Benchmark<func name>`;

### go test命令
> 具体文档见`go help testflag`

- `go test .`: 运行当前目录下所有`*_test.go`文件里以Test开头的函数;
- `-v`: 显示测试详情;
- `-run="<regexp>"`: 执行匹配正则表达式的测试函数;
    - `go test -run=. .` 执行当前目录下所有的test函数; 
    - `-run`*参数一定要接参数*:`-run <regexp>`或者`-run="<regexp>"`;
- `-cover`: 开启测试覆盖率;
- `-race`: 开启竞争检测(测试线程安全必须要加这个参数);

### 基准测试
`-bench="<regexp>"`: 执行匹配正则表达式的基准测试;*和`-run`一样一定要接参数*

注意testing仍然会执行test测试函数,如果只做基准测试,建议使用`go test -run=none -bench=xxx`将run匹配不到任何单元测试函数;

1. 通过`-benchtime`参数可以自定义测试时间，例如`go test -v -bench=. -benchtime=5s benchmark_test.go`测试时间设置为5秒(默认测试时间1秒)
2. 在命令行中添加`-benchmem`参数以显示内存分配情况，参见下面的指令：

     "BenchmarkFoo-8   	1000000000	         0.570 ns/op	       0 B/op	       0 allocs/op"
    - BenchmarkFoo-8:代表执行BenchmarkFoo函数, 8线程;
    - 1000000000 表示测试次数;
    - 0.570 ns/op: 每一个操作耗费多少纳秒;
    - 0 B/op: 表示每次调用耗费多少字节;
    - 0 allocs/op: 每次调用有几次内存分配;

基准测试使用`ResetTimer`

#### 并行测试
使用`b.RunParallel`,注意在函数内不要使用`b.ResetTimer()`这样的全局函数:
```go
// 相当于 go func
b.RunParallel(func(pb *testing.PB) {
    // 前期处理工作
    for pb.Next() { // pb.Next 判断是否有更多的迭代要执行;
        // do something
    }
})
```
其中`pb.Next()`相当于普通基准测试的`for i := 0; i < b.N; i++`,也就是在1秒内(可以配置)由testing判断循环多少次for里面的代码;


### 代码覆盖率
`go test -coverprofile=c.out`
`go tool cover -html=c.out`

### 性能测试
`go test -bench . -cpuprofile cpu.out`
`go tool pprof cpu.out`
`web`

### 日志方法: 
|方  法|备  注|
|---|---|
|Log|测试来说，只会在失败或者设置了-test.v标志的情况下被打印出来；对于基准测试来说，为了避免 -test.v标志的值对测试的性能产生影响，格式化文本总会被打印出来。|
|Logf|格式化打印日志|
|Error|相当于在调用Log之后调用Fail(将当前测试标识为失败，但是仍继续执行该测试)|
|Errorf|相当于在调用Logf之后调用Fail|
|Fatal|相当于在调用Log之后调用FailNow(立即结束测试)|
|Fatalf	|格式化打印致命日志，同时结束测试|

#### 配合runtime与unsafe
`runtime/debug.go`中有几个测试用的函数:
- `NumGoroutine`: 返回当前存在的goroutine数;
- `NumCgoCall`: 当前进程发出的cgo调用数;
- `NumCPU`: 当前进程可用的逻辑CPU数;

### helper
go1.9加入了一个新特性，那就是Helper方法,`tesing.T`和`testing.B`中均添加了该方法. 该方法能够标记某个测试方法是一个helper函数，当一个测试包在输出测试的文件和行号信息时，将会输出调用help函数的调用者的信息，而不是输出helper函数的内部信息。举例来说：
```go
package p

import "testing"

func failure(t *testing.T) {
    t.Helper() // This call silences this function in error reports.
    t.Fatal("failure")
}

func Test(t *testing.T) {
    failure(t)
}
```
因为failure函数标记自己为helper函数，如果测试失败，即t.Fatal函数被调用时，错误信息将会输出在Test函数的位置，而不是在failure函数的位置。





## net/http/pprof 
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
