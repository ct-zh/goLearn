# pprof 
## 快速使用
- 压测生成文件：`go test -run=^$ -bench=^{压测函数名}$ -cpuprofile=cpu.prof/-memprofile=mem.prof`；
- `go tool pprof`


## 用法
- web服务器

    引入`import _ "net/http/pprof"`, 然后使用`http://localhost:port/debug/pprof/`直接看到当前web服务的状态;

- 服务进程

    同样引入包net/http/pprof，然后在开启另外一个goroutine来开启端口监听:
    ```go
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil)) 
    }()
    ```

- 程序

    使用`runtime/pprof`包; StartCPUProfile和StopCPUProfile;运行的时候加个参数`--cpuprofile=fabonacci.prof`生成对应的pprof文件;再使用`go tool pprof`分析;

### top命令
topN命令的输出结果默认按flat(flat%)从大到小的顺序输出：

- flat列的值表示函数自身代码在数据采样过程的执行时长；
- flat%列的值表示函数自身代码在数据采样过程的执行时长占总采样执行时长的比例；
- sum%列的值是当前行flat%值与排在该值前面所有行的flat%值的累加和。以第三行的sum%值75.00%为例，该值由前三行flat%累加而得，即75.00% = 16.67% + 20.83% + 37.50%；
- cum列的值表示函数自身在数据采样过程中出现的时长，这个时长是其自身代码执行时长以及其等待其调用的函数返回所用时长的总和。越是接近函数调用栈底层的代码，其cum列的值越大；
- cum%列的值表示该函数cum值占总采样时长的百分比。比如：runtime.findrunnable函数的cum值为 130ms，总采样时长为 240ms，则其```cum%值为两者的比值百分化后的值(130.0/240，再百分化)。

命令行交互模式也支持按cum值从大到小排序输出采样结果：`top -cum`

### list命令
我们可以通过list命令列出函数对应的源码

在展开源码的同时，pprof还列出了代码中对应行的消耗时长（基于采样数据）。我们可以选择耗时较长的函数，做进一步的向下展开，这个过程类似一个对代码进行向下钻取的过程，直到找到令我们满意的结果（某个导致性能瓶颈的函数中的某段代码


## pprof测试优化实例
http://imooc.com/read/87/article/2440#anchor_5

### step0
1. 执行`step0/demo_test.go`, 看到性能如下：

   `// BenchmarkHi-8   	  280515	      3717 ns/op	    3044 B/op	      49 allocs/op`

2. 执行`go test -v -run=^$ -bench=^BenchmarkHi$ -benchtime=2s -cpuprofile=cpu.prof`

3. 分析prof：`go tool pprof step0.test cpu.prof`

    - 使用top -cum命令看到handleHi函数消耗cpu最多: ` 1.97s 77.56%  ../step0.handleHi`
    - `list handleHi`命令,看到regexp.MatchString函数消耗了1.45s（总共2s的测试时间）: 

    ```
    1.45s     22:	if match, _ := regexp.MatchString(`^\w*$`, r.FormValue("color")); !match {
    ```

4. 优化, 优化代码在`step1/demo.go`
```go
func handleHi(w http.ResponseWriter, r *http.Request) {
    match, _ := regexp.MatchString(`^\w*$`, r.FormValue("color"));

    // ...
}

// 优化为
var rxOptionalID = regexp.MustCompile(`^\d*$`)
func handleHi(w http.ResponseWriter, r *http.Request) {
	rxOptionalID.MatchString(r.FormValue("color"))

    // ...
}
```

5. 执行`step1/demo_test.go`, 看到性能如下,速度对比step0快了7倍:

    `// BenchmarkHi-8   	 2061201ww	       499 ns/op	    3044 B/op	      49 allocs/op`;

### step1
1. 执行`step1/demo_test.go`; `go test -v -run=^$ -bench=^BenchmarkHi$ -benchtime=2s -memprofile=mem.prof` 

2. 执行`go tool pprof step1.test mem.prof`分析pprof
> 在go tool pprof的输出中有一行为Type: alloc_space。这行的含义是当前 pprof 将呈现的是程序运行期间的所有内存分配的采样数据（即使该分配的内存在最后一次采样时已经被释放）; 我们还可以让 pprof 将Type切换为inuse_space，这个类型表示的是内存数据采样结束时依然在用的内存。切换命令:`sample_index = inuse_space`

    `783.05MB 38.15% 38.15%  2052.44MB   100%  ...handleHi`
    
    ```
    362.52MB     1.48GB     30:	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
    420.52MB   449.52MB     31:		"'>Welcome!</h1>You are visitor number " + fmt.Sprint(visitNum) + "!"))
    ```

3. 优化写入方法,优化代码在step2
    ```go
    w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome!</h1>You are visitor number " + fmt.Sprint(visitNum) + "!"))

    // 优化为
    fmt.Fprintf(w, "<html><h1 stype='color: %s'>Welcome!</h1>You are visitor number %d!",
		r.FormValue("color"), visitNum)
    ```

4. 测试结果: `go test -v -run=^$ -bench=^BenchmarkHi$ -benchtime=2s -memprofile=mem.prof`

    可以看到:
    ```
    60MB  4.59%  4.59%  1308.38MB   100%  .../step2.handleHi

    (pprof) list handleHi
       .     1.22GB     29:	fmt.Fprintf(w, "<html><h1 stype='color: %s'>Welcome!</h1>You are visitor number %d!",
    60MB       60MB     30:		r.FormValue("color"), visitNum)
    ```
    从783.05MB优化到60MB了

### step2
1. 将`fmt.Fprintf`替换为`sync.Pool`;
2. 再执行pprof,可以看到`0     0%   100%     2.12GB   100%  .../step3.handleHi` 变成了0;

### 测试并发竞争
- 测试Parallel基准测试函数: `go test -bench=Parallel -blockprofile=block.prof`
- 分析prof文件: `go tool pprof step3.test block.prof`
- 发现handleHi函数在prof里面没有任何占用,说明并发竞争没有问题;

### 归纳总结
源码见`step0/demo.go`, 这个代码有两个优化点:
- `regexp.MatchString`对每个请求,每次需要重新编译正则表达式再去匹配str; 可以改成全局`regexp.MustCompile`,直接匹配str,只需做一次编译操作;
- `w.Write`方法优化,先是使用`fmt.Fprintf`优化一次; 后面使用`sync.Pool`几乎不消耗内存;



 
## 使用pprof定位线上超时问题

首先登陆服务机器,查看pprof分配的端口;

```go
ps -ef |grep "service_name"
// 得到pid

netstat -nultp |grep service_pid
// 这里应该能有两个进程,其中一个是服务web端口,另一个应该就是pprof端口
```

登陆可以直接连接到pprof的机器,如果机器有go环境,则可以直接采样.

如果机器没有go环境,看是否能将请求转发到服务机器上,这里使用一个proxy程序将服务机器的pprof端口转发到灰度机器的某端口上,再从测试环境存在go服务的机器上开启采样

```shell
./proxy -dst 10.99.34.158:41595 -src 0.0.0.0:10801

// 然后本地采样
go tool pprof -http=:8000 -seconds=60 http://10.100.97.2:10800/debug/pprof/profile
```

