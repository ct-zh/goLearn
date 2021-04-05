# 测试
## 测试流程
1. 基本接口的单元测试, 保证代码没有逻辑性的bug;    
2. 代码基准测试;
3. 使用pprof性能测试;
4. 使用GDB、Delve进行debug;
5. jenkins ci;
6. Jmeter进行压测;


## 1. 单元测试与基准测试
单元测试分为包内测试与包外测试: 
- 包内测试直接在包内写test函数,特点是可以访问到包内所有变量, 更为直接地构造测试数据和实施测试逻辑, 测试覆盖率高;缺点是入侵性大,需要经常维护;
- 包外测试就是正常的`*_test.go`文件测试;主要是对包暴露出的api进行测试;*在日常测试中优先考虑包外测试*

### 测试注意事项
> see https://www.imooc.com/read/87/article/2438#anchor_1
- 测试用的外部资源放在`testdata`文件夹下(go会自动忽略该目录)
- 使用`*.golden`文件来对输出数据做比较
- 使用`-cover`测试代码覆盖率
- 使用`go vet`/`go tool vet`测试代码问题
- 使用`-race`测试代码冲突


### 命名规范
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

注意testing仍然会执行test测试函数,如果只做基准测试,建议使用`go test -run=^$ -bench=xxx`将run匹配不到任何单元测试函数;

1. 通过`-benchtime`参数可以自定义测试时间，例如`go test -v -bench=. -benchtime=5s benchmark_test.go`测试时间设置为5秒(默认测试时间1秒)
2. 在命令行中添加`-benchmem`参数(或者在代码中添加`b.ReportAllocs()`)以显示内存分配情况，参见下面的指令:

     "BenchmarkFoo-8   	1000000000	         0.570 ns/op	       0 B/op	       0 allocs/op"
    - BenchmarkFoo-8:代表执行BenchmarkFoo函数, 8线程;
    - 1000000000 表示测试次数;
    - 0.570 ns/op: 每一个操作耗费多少纳秒;
    - 0 B/op: 表示每次调用耗费多少字节;
    - 0 allocs/op: 每次调用有几次内存分配;

### 排除干扰因素
当某个测试有部分代码不想计算在计时器内 (比如把数据库初始化时间排除在外), 使用`b.ResetTimer()`或者`b.StopTimer()`和`b.StartTimer()`来控制基准测试的计时器;

> *不要在for循环里面使用ResetTimer!!*


### 并行测试
通常用来测试多协程代码;使用`b.RunParallel`,注意在`runParallel`函数内不要使用`b.ResetTimer()`这样的全局函数:
```go
func BenchmarkHiParallel(b *testing.B) {
    r, err := http.ReadRequest(bufio.NewReader(strings.NewReader("GET /hi HTTP/1.0\r\n\r\n")))
    if err != nil {
            b.Fatal(err)
    }
    b.ResetTimer()      // 重置测试计时器

    b.RunParallel(func(pb *testing.PB) {
        rw := httptest.NewRecorder()
        for pb.Next() {
            handleHi(rw, r) // 测试handle
        }
    })
}
```
其中`pb.Next()`相当于普通基准测试的`for i := 0; i < b.N; i++`,也就是在1秒内(可以配置)由testing判断循环多少次for里面的代码;

- 测试Parallel基准测试函数: `go test -bench=Parallel -blockprofile=block.prof`
- 分析prof文件: `go tool pprof step3.test block.prof`: `top`查看占用情况; `list handleHi`查看handleHi函数情况
- 发现handleHi函数在prof里面没有任何占用,说明并发竞争没有问题;


### 基准测试结果对比
使用`benchstat`: 
```bash
go test -run=NONE -bench . strcat_test.go > new.txt
benchstat old.txt new.txt
```



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



## 2. pprof 性能测试
- [性能检测优化实例](./pprof/demo/README.md)

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
