
## Channel简介

在阅读channel源码前，我们先复习一下channel的基础知识。

### 什么是channel

什么是Channel（通道）？简单说，channel 是多个 goroutine 之间沟通的 “管道”：一个 goroutine 可以向管道中发送数据，另一个 goroutine 可以从管道中接收数据，且发送和接收操作会**阻塞**，直到双方准备就绪。

> 严格来说，**channel（通道）** 是用于 ** goroutine 之间安全通信 ** 的核心原语，也是 Go 实现 “通过通信共享内存，而非通过共享内存通信”（Do not communicate by sharing memory; instead, share memory by communicating）这一哲学的关键工具。


### Channel 的声明与初始化

我们通过以下方式来声明一个channel：
```go
// 声明一个用于传输 T 类型数据的 channel（未初始化，nil） 
var ch chan T 

// 声明只读 channel（只能接收，不能发送） 
var ch <-chan T 

// 声明只写 channel（只能发送，不能接收） 
var ch chan<- T
```

我们通过以下方式来初始化（make）一个channel:
```go
// 1. 无缓冲 channel（容量为 0） 
ch := make(chan T) 

// 2. 有缓冲 channel（容量为 n，可存储 n 个 T 类型数据） 
ch := make(chan T, n)
// 可以通过 cap(ch)获取有缓冲的channel的容量值
```

日常开发建议直接使用make来初始化channel，因为channel 是**引用类型**，必须通过 `make` 初始化后才能使用。直接声明的channel是nil channel，无法发送 / 接收数据，会永久阻塞。

### channel 的核心操作

channel 有 3 个核心操作：**发送（send）**、**接收（receive）**、**关闭（close）**。
#### 1. 发送操作（`ch <- 数据`）

将数据写入 channel，语法为 `ch <- 数据`（箭头指向 channel，表示 “写入”）。

- 无缓冲 channel：发送会阻塞，直到有 goroutine 调用接收操作。
- 有缓冲 channel：缓冲区未满时直接写入，满时阻塞，直到有数据被接收（缓冲区腾出空间）。
- 注意：向已关闭的 channel 发送数据会触发 panic。

#### 2. 接收操作（`数据 <- ch` 或 `<-ch`）

从 channel 读取数据，有 3 种常见形式：
```go
// 形式 1：接收数据并赋值给变量 
data := <-ch // 阻塞直到有数据可用 

// 形式 2：忽略接收的数据（仅用于“消费”数据或等待发送方） 
<-ch // 阻塞直到有数据被接收 

// 形式 3：多返回值（判断 channel 是否关闭） 
data, ok := <-ch 
// 如果 ok 为 true：正常接收数据；
// 如果 ok 为 false：channel 已关闭且无剩余数据
```

还有一种常用的遍历接收的形式：`for data := range ch` 会持续从 channel 接收数据，直到 channel 被关闭且缓冲区为空，循环自动退出（无需手动判断 `ok`）。

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch) // 必须关闭，否则 range 会一直阻塞

// 遍历 channel
for num := range ch {
	fmt.Println("接收：", num)
}
```

#### 3. 关闭操作（`close(ch)`）

关闭 channel 后，无法再发送数据，但仍可接收缓冲区中剩余的数据。

- 只能关闭非 nil 且未关闭的 channel，重复关闭会触发 panic。
- 关闭只读 channel（`<-chan T`）会编译报错（只读 channel 不能发送，也不能关闭）。

### 无缓冲 vs 有缓冲 channel

#### 1. 无缓冲 channel（同步 channel）

容量为 0，发送和接收操作**必须同时就绪**，否则会阻塞（本质是 “同步通信”）。

示例：goroutine 间同步传递数据
```go
func worker(ch chan int) { 
	// 接收数据（阻塞，直到 main 发送数据） 
	data := <-ch fmt.Printf("worker 接收数据：%d\n", data) 
} 

func main() { 
	ch := make(chan int) // 初始化无缓冲 channel 
	go worker(ch) // 启动 goroutine 
	// 发送数据（阻塞，直到 worker 接收数据） 
	ch <- 100 fmt.Println("main 发送完成") 
}
```

输出：
```plaintext
worker 接收数据：100
main 发送完成
```
- 逻辑：`main` 发送 `100` 时阻塞，直到 `worker` 接收；`worker` 启动后先阻塞，直到 `main` 发送。两者同步完成后才继续执行。


#### 2. 有缓冲 channel（异步 channel）

容量为 `n`，发送操作在缓冲区未满时不会阻塞，接收操作在缓冲区非空时不会阻塞（本质是 “异步通信”，缓冲区起到 “缓冲” 作用）。

示例：有缓冲 channel 的异步通信

```go
func main() {
    ch := make(chan string, 2) // 有缓冲 channel，容量 2

    // 发送 2 个数据（缓冲区未满，不阻塞）
    ch <- "hello"
    ch <- "world"
    fmt.Println("发送 2 个数据完成，缓冲区剩余容量：", cap(ch)-len(ch)) // 0

    // 尝试发送第 3 个数据（缓冲区满，阻塞）
    // ch <- "go" // 解开注释后，main 会阻塞在这里

    // 接收 1 个数据（缓冲区非空，不阻塞）
    data := <-ch
    fmt.Printf("接收数据：%s，缓冲区剩余容量：%d\n", data, cap(ch)-len(ch)) // 1

    // 此时可以发送第 3 个数据（缓冲区有空间，不阻塞）
    ch <- "go"
    fmt.Println("发送第 3 个数据完成")

    // 关闭 channel
    close(ch)

    // 接收剩余数据
    for data := range ch { // range 会遍历 channel 直到关闭且无数据
        fmt.Printf("遍历接收：%s\n", data)
    }
}
```

输出：
```go
发送 2 个数据完成，缓冲区剩余容量： 0
接收数据：hello，缓冲区剩余容量：1
发送第 3 个数据完成
遍历接收：world
遍历接收：go
```



### channel 的核心特性

根据上面channel基础知识，我们可以总结出channel的核心特性：

- **并发安全**：channel 本身是线程安全的，多个 goroutine 同时发送 / 接收不会出现数据竞争（底层由 Go runtime 维护锁）。
-  **阻塞特性**：
    - 无缓冲 channel：发送操作会阻塞，直到有 goroutine 接收；接收操作会阻塞，直到有 goroutine 发送。
    - 有缓冲 channel：发送操作仅在缓冲区满时阻塞；接收操作仅在缓冲区空时阻塞。
- **类型安全**：channel 必须指定存储的数据类型（如 `chan int`、`chan string`），只能发送 / 接收对应类型的数据。
- **关闭机制**：可以通过 `close(ch)` 关闭 channel，关闭后不能再发送数据，但仍可接收剩余数据；接收已关闭且无数据的 channel 会返回该类型的零值，可通过多返回值判断 channel 是否关闭。


### channel 的常见用法

数据传递：一个 goroutine 生产数据，另一个消费数据。

同步等待：用无缓冲 channel 实现 “等待 goroutine 完成”（类似 `sync.WaitGroup` 的简化版）。
```go
func task(id int, done chan<- bool) {
    fmt.Printf("任务 %d 执行完成\n", id)
    done <- true // 通知主 goroutine 完成
}
```

主流程：
```go
 // 启动 3 个 goroutine
go task(1, done)
go task(2, done)
go task(3, done)

// 等待 3 个任务全部完成
for i := 0; i < 3; i++ {
	<-done // 阻塞，直到收到一个完成信号
}

fmt.Println("所有任务完成")
```


限制并发数（信号量模式）：用有缓冲 channel 作为 “信号量”，限制同时运行的 goroutine 数量。

```go
func worker(id int, sem chan struct{}) {
    defer func() { <-sem }() // 函数执行完后释放信号量（接收一个空结构体）
    
    fmt.Printf("worker %d 开始工作\n", id)
    time.Sleep(1 * time.Second) // 模拟工作
    fmt.Printf("worker %d 完成工作\n", id)
}
```

主流程：
```go
maxConcurrent := 2          // 最大并发数
sem := make(chan struct{}, maxConcurrent) // 信号量 channel

// 启动 5 个任务
for i := 1; i <= 5; i++ {
	sem <- struct{}{} // 获取信号量（发送空结构体），缓冲区满时阻塞
	go worker(i, sem)
}

// 等待所有任务完成（这里简化，实际可结合 sync.WaitGroup）
time.Sleep(3 * time.Second)
fmt.Println("所有任务执行完毕")
```


### channel常见的坑

1. **nil channel 的坑**：未初始化的 channel（nil）执行发送 / 接收操作会**永久阻塞**，需避免。
2. **关闭 channel 的原则**：
    - 通常由**发送方**关闭 channel（因为发送方知道何时不再发送数据）。
    - 接收方不应关闭 channel（可能导致发送方 panic）。
3. **单向 channel 的作用**：通过 `<-chan T`（只读）和 `chan<- T`（只写）限制 channel 的操作权限，提高代码安全性（如函数参数传递时，避免误操作）。


#### goroutine 泄漏（最常见、最隐蔽）

goroutine 泄漏是 channel 使用中**最高发的问题**：当一个 goroutine 因等待 channel 的发送 / 接收而阻塞，但永远没有其他 goroutine 满足其条件时，该 goroutine 会一直占用资源，无法被 GC 回收。

##### 常见场景 1：发送方无人接收，且 channel 无缓冲 / 缓冲区满
```go
func leak() {
	ch := make(chan int) // 无缓冲 channel

	// 启动 goroutine 发送数据，但无人接收
	go func() {
		fmt.Println("goroutine 开始发送数据")
		ch <- 100 // 阻塞：无接收方，永远无法发送成功
		fmt.Println("goroutine 发送完成") // 永远不会执行
	}()

	// 主 goroutine 不接收，直接退出（或做其他事）
	time.Sleep(1 * time.Second)
	fmt.Println("main 退出")
}
```

**原因**：无缓冲 channel 的发送操作需要接收方就绪，否则阻塞。若主 goroutine 不接收，发送 goroutine 会永久阻塞，无法退出。

##### 常见场景 2：接收方无人发送，且 channel 未关闭
```go
func leak2() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2 // 缓冲区已满

	// 启动 goroutine 接收数据，但无更多发送
	go func() {
		fmt.Println("接收：", <-ch) // 接收 1
		fmt.Println("接收：", <-ch) // 接收 2
		fmt.Println("等待下一个数据...")
		<-ch // 阻塞：缓冲区空，且无发送方，永远无法接收
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("main 退出")
}
```
**原因**：接收方已消费完所有数据，但 channel 未关闭，后续接收操作会永久阻塞，导致 goroutine 泄漏。
##### 解决方案：
1. **明确关闭 channel**：发送方在所有数据发送完毕后关闭 channel，接收方通过 `range` 或 `ok` 判断退出（避免无数据时阻塞）。
2. **使用带超时的接收**：通过 `time.After` 避免永久阻塞（适用于不确定是否有数据的场景）。
3. **使用 context 取消**：通过 `context.Context` 控制 goroutine 退出（推荐，支持主动取消）。
```go
go func(ctx context.Context) {
	select {
	case ch <- 100:
		fmt.Println("发送成功")
	case <-ctx.Done(): // 接收取消信号
		fmt.Println("goroutine 被取消，退出")
		return
	}
}(ctx)
```



##### 场景3  range 遍历未关闭的 channel（永久阻塞）

`for data := range ch` 会持续接收 channel 数据，直到 channel 被关闭且缓冲区为空。若 channel 未关闭，且无更多数据发送，range 会永久阻塞，导致 goroutine 泄漏。

```go
// 启动 goroutine 遍历 channel，但 channel 未关闭
go func() {
	for num := range ch {
		fmt.Println("接收：", num) // 接收 1、2 后，阻塞在 range
	}
	fmt.Println("遍历结束") // 永远不会执行
}()
```
**原因**：range 遍历的终止条件是「channel 关闭 + 缓冲区空」，若只发送数据不关闭，遍历会一直等待新数据，导致阻塞。

##### 解决方案：

1. **发送方发送完数据后主动关闭 channel**（推荐，明确语义）。
2. **若无法关闭（如多发送方），用 context 控制遍历退出**。



#### 关闭已关闭的 channel（触发 panic）

重复关闭 channel 会直接导致 panic，且这种错误在并发场景中容易隐蔽（比如多个 goroutine 同时尝试关闭同一个 channel）。
##### 解决方案：
1. **单一关闭者原则**：明确由「发送方」关闭 channel（因为发送方知道何时停止发送），且确保只有一个发送方（或通过同步控制唯一关闭动作）。
2. **用 sync.Once 保证唯一关闭**：若必须多个 goroutine 可能触发关闭，用 `sync.Once` 确保关闭操作只执行一次。

#### 向已关闭的 channel 发送数据（触发 panic）

关闭 channel 后，发送操作会直接 panic，但接收操作仍可读取剩余数据（或零值），这个区别容易被忽略。
##### 解决方案：

1. **发送前判断 channel 是否关闭**：通过 `select` + 「默认分支」避免发送到已关闭 channel（但需注意竞态条件）。
2. **避免在接收方关闭 channel**：关闭动作应由发送方主导，接收方只负责接收，不主动关闭。

```go
func safeSend(ch chan int, data int) bool {
	select {
	case ch <- data:
		return true
	default:
		// 若 channel 已关闭或缓冲区满，直接返回 false（非阻塞）
		// 注意：这种方式无法区分「已关闭」和「缓冲区满」，需结合场景
		return false
	}
}
```

注：若需要明确判断 channel 是否关闭，只能通过接收操作的 `ok` 值（发送操作无直接判断方式），因此更推荐「单一关闭者」原则从源头避免。


#### 无缓冲 channel 导致的死锁（同步阻塞不匹配）

无缓冲 channel 的发送和接收必须「同时就绪」，否则会阻塞。若多个 goroutine 相互等待对方的发送 / 接收，会导致死锁。
##### 典型场景：goroutine 相互等待
```go
func deadlock() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// goroutine A：等待 ch1 接收，再向 ch2 发送
	go func() {
		data := <-ch1 // 阻塞：等待 goroutine B 发送 ch1
		ch2 <- data + 1 // 若 goroutine B 未接收 ch2，也会阻塞
	}()

	// goroutine B：等待 ch2 接收，再向 ch1 发送
	go func() {
		data := <-ch2 // 阻塞：等待 goroutine A 发送 ch2
		ch1 <- data + 1 // 若 goroutine A 未接收 ch1，也会阻塞
	}()

	// 主 goroutine 等待，导致死锁：两个 goroutine 相互阻塞
	time.Sleep(1 * time.Second)
}
```

**原因**：goroutine A 等待 ch1 的接收，而 ch1 的发送需要 goroutine B 完成，但 goroutine B 又在等待 ch2 的接收，形成循环依赖，最终所有 goroutine 阻塞，程序死锁。

##### 解决方案：

1. **确保发送 / 接收顺序匹配**：避免相互依赖的阻塞，可调整逻辑让一个 goroutine 先发送，另一个先接收。
2. **使用有缓冲 channel 打破同步依赖**：通过缓冲区暂存数据，避免必须同时就绪。


#### 总结：channel 避坑核心原则

1. **单一关闭者**：由发送方（或唯一控制方）关闭 channel，避免重复关闭、向已关闭 channel 发送。
2. **避免永久阻塞**：用 `context`、超时、关闭 channel 等方式，确保每个发送 / 接收操作都有退出路径。
3. **明确 channel 生命周期**：知道 channel 何时发送完毕、何时关闭，避免 goroutine 因等待未关闭的 channel 而泄漏。
4. **合理设计缓冲区**：根据业务场景选择无缓冲 / 有缓冲，缓冲区容量匹配生产消费速度。
5. **慎用单向 channel**：仅在函数参数 / 返回值中用于权限控制，避免手动声明单向 channel 变量。



