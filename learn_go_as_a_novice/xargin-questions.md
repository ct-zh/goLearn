## 问题集锦

- 有哪些情况下会有zero地址发生? 为何`map[int]struct{}`中的`struct{}`不占任何内存空间? zero地址又是什么?
    
    - 在Go语言中，空结构体的变量确实会被分配一个内存地址，这个地址是唯一的。但是，这个地址并不指向任何实际的内存位置，因为空结构体不包含任何字段，所以它不需要真正的内存空间。
    
    - 你可以使用`&`操作符获取空结构体变量的地址，但这个地址并没有实际意义，只是一个占位符。因此，即使你获取了空结构体变量的地址，也不要期望它指向真实的内存位置。这种情况下，空结构体变量的地址实际上是一个虚拟地址，用于表示这个变量在内存中的位置，但不实际分配任何内存空间。
    
    - zero地址是一个指向全局内存区域的指针,在go语言中它的表达为: `runtime.zerobase`,当你使用make、new或其他方式分配切片、映射、通道等动态数据结构时，Go运行时会分配一段内存空间，并且将其初始化为零值。这个零值内存区域的地址就由`runtime.zerobase`指针表示。

    - 空结构体指向的地址其实就是`runtime.zerobase`, 如何验证? 
        ```go
        var b = struct{}{}

        // 执行调试: dlv debug test.go
        // 断点打到申明变量b的地方 b main.main:2
        // 打印汇编: disass
        test.go:11      0x10b6f84       488d0d35600d00                  lea rcx, ptr [runtime.zerobase]
        test.go:11      0x10b6f8b       48898c2490000000                mov qword ptr [rsp+0x90], rcx
        // 可知该空结构体指向的就是runtime.zerobase
        ```

- go不是很适合做网关,看QPS.滴滴32c机器跑网关,压到几万QPS, 延迟来回几十ms.

- channel能实现无锁化设计吗? 
    
    - 不能. 其实无锁化设计是有成本的,无锁化设计会存在一个 *忙等待* 一般会消耗cpu,(可以从pprof看到经常有runtime.schedul或者runtime.runnable之类函数占用很高, 其实就是无锁化设计在忙等待自旋)
    
    - 环形队列常用无锁化处理atomic.CAS

    - Go channel中的实现中使用了mutex,这个mutex和标准库中的Mutex有什么不同？ [Go运行时中的 Mutex](https://colobu.com/2020/12/06/mutex-in-go-runtime/)


- 单个四叉堆timer的loop check时间间隔是多少?
    - 算堆顶的到期时间, 然后sleep
    - 为啥是四叉堆? 空间局部性好一些, N叉堆通常是2的N次方

- 计算密集型
    - ffmpeg
    - bcrypt
    - json.Unmarshal 
    - 文本/数据处理相关

- timeproc会因goroutine的公平性而延迟执行吗? 会

- Go的map不会缩容, 所有手动缩容需要把map a手动放到map b. 
    - 在生产环境有影响.
    - 暂时没有什么好的解决办法


### 并发编程

- 当p的数量发生变化时，已有的pool.Local会跟着变化么？ - 会

- 为什么pool的shared通过双端链表能实现无锁
    - 队列头队列尾 CAS

- 为什么semaRoot是个treap而不是一个链表或者队列？这个地址是semaphore的地址？每个锁上不是一个独立的semaphore吗?
    - 用链表查询是线性复杂度 O(n) 
    - 用数组可以用二分，但是增删麻烦
    - 地址就是sema的地址 
    - 每个锁上都是独立的semaphore, 但是同一个锁可能有多个goroutine去lock unlock;

- 不是按等待时间来唤醒竞争的goroutine么？
    - a. 非饥饿模式：最新进自旋的 goroutine 优先级最高 
    - b. 饥饿模式：排队

- 如果sync.Mutex锁己经由一个协程持有，然后另外一个goroutine在加锁时，会有自旋过程吗？ - 会

- 如何在Go语言里写单例模式? - 参考 sync.Once

- sync.Map 设计 read 和 dirty 的目的是什么？为什么分开设计？
    - map + lock 多核扩展性差
    - sync.Map 在读多写少的情况下基本上不需要加锁 dirty读不加锁

### 框架

- 企业级框架应该集大成, 以单库维护1依赖版本, 避免业务出问题;

- 日志差距; 日志过多影响速度, 如果用异步, 非sugar模式的zap可以快很多; 也要看硬盘的水平; 没日志比有日志快20~30%
    - 打的不好可能成为p0事故; 在for循环中打印;
    - sla 值  接口速度承诺
    - 交易类的日志可能都需要同步, 不能异步打; 而且要保留一定时间的交易计算流程日志;

- 书单 building microservices 微服务设计

- https://microservices.io/

- 如何做热更新
    - 把预设的模式编码好, 根据配置做修改, 例如govaluate 不太灵活
    - gopherlua gopherjs - 不稳定
    - go-plugin 没人用
    - wasm
    - go做热更新还是不方便
    - 配置文件热更 viper 双buffer

- 依赖注入
    - 自动做变量绑定, 少些一些代码

- 分布式接口限流
    - 注册中心存储特定的api, 对其访问量限制
    - 每个服务订阅自己, 知道自己的实例数
    - 容器自适应限流: 监控系统, 本地有个agent, 业务可以查出来metrics sdk性能指标, 参考的sentinel 代码


- 改造gorm logger下的trace方法, 找到慢查询的sql; (公司框架已实现)

### 框架底层

- `type HandlerFunc func(wr, *r)`与`type HandlerFunc = func(wr, *r)`有何区别? (注意第二个开头关键字为type而不是var)
    - 第一个是`type alias`, 第二个是`type equality`; 
    - 前者相当于定义类型的别名, 但是这个别名是一个新的、独立的类型; 在类型检查中,`HandleFunc`与`func(wr, *r)`不是同一个类型;你还可以专门给`HandlerFunc`定义新的方法, 使之与`func(wr, *r)`不同
    - 后者则是两边完全等价;如果声明一个变量`var b HandlerFunc`, 打印其类型`fmt.Print("%T", b)`得到的是`func(wr, *r)`而不是`HandlerFunc`;也因如此无法给`HandlerFunc`定义方法;

- 认证逻辑是放在request判断中比较好还是放在中间件中判断比较好?

- `graphql`与`restful api`的区别
    - 数据获取方式
        - REST：每个资源都有一个唯一的URL。例如，获取用户信息可能需要GET /users/{id}。
        - GraphQL：所有数据通过一个单一的端点（通常是/graphql）获取，客户端通过查询来指定需要的数据结构。
    - 灵活性
        - REST：服务器决定了API的结构，客户端只能请求预定义的数据结构。获取复杂数据可能需要多次请求。
        - GraphQL：客户端可以灵活地指定所需的数据结构和字段，避免了过多或不足的数据传输。单个请求可以获取复杂的关联数据。
    - 数据效率
        - REST：可能会出现过多或不足的数据传输，导致效率低下。例如，要获取嵌套数据，可能需要多次请求。
        - GraphQL：客户端指定精确的数据需求，避免了数据冗余和不足，提高了效率。
    - 所以
        - GraphQL 更适合需要灵活数据查询和复杂关系的应用，特别是在客户端需求多变的情况下。
        - REST 更适合简单、资源明确的应用，容易理解和实现。
    - 为什么现在广泛使用`RESTFUL API`而不是`graphql`? 
        - `graphql`通过单一节点可以查询所有信息, 权限控制不好做, 甚至可能有安全问题;
        - `graphql`的一次查询, 可能只获取1条数据, 也可能获取1000条数据; 具体的数据信息是由客户端决定的, 这就意味每次请求的成本不同, 无法做精确的限流;
        - `graphql`目前最广泛的应用场景可能是公司后台的运营数据查询系统; 

- 一般的测试套路
    - unit test -> pkg、 func
    - 模块测试 -> 脚本、CI
    - 集成测试 -> 通常是从线上采集下来全流程的case, 在CI流程里跑; (流量回放, 线上数据脱敏后将其改造成特定的测试集, 每次上线都要跑一次)
    - QA负责维护自动化系统;

- validator 不能分辨零与默认值
    - 语言特性无法区分;
    - 只能改造成指针类型, 用nil来判断; 但是数据量太大会有gc问题;

- grpc如何做validator与interception?
    - grpc本身支持interception
    - pb里面也可以写校验规则, 或者把unmarshal的内容再validator一次

### 业务分层

### 优雅代码

### 性能调优
- 看一下tcp reader goroutine / writer goroutine 

- goroutine占用的内存指标: stack_inuse_bytes; 与heap内存在不同的统计指标里;

- 如何做简单压测?
    - 固定QPS压测; 首先对自己的api的QPS有一定预估, 如至少能压到6000? 然后找对应工具去压;
    - 极限QPS压测; 