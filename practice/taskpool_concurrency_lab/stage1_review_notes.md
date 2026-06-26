# Stage 1 Review Notes

这份文档把第一阶段所有反复暴露的问题单独整理出来，目标不是记录“最终代码长什么样”，而是记录：

- 当时错在哪里
- 为什么会错
- 以后怎么一眼识别
- 更稳的替代思路是什么

这些问题集中暴露了两个短板：

- 对 channel 阻塞语义和取消语义的组合不够熟
- 对并发测试中的“要等什么”不够敏感

## 1. 先判断 channel 状态再发送

- 现象：并发提交时可能卡住，且卡住后不能及时响应 `ctx.Done()`
- 问题代码：

```go
if len(p.tasks) < p.queueSize {
    p.tasks <- task
    return nil
}
```

- 为什么错：
  - `len(ch)` 和 `ch <- v` 不是原子操作
  - 多个 goroutine 可能同时判断“还有空间”
  - 真正发送时，只有一部分能成功，其余会阻塞在发送上
  - 一旦阻塞点落在 `p.tasks <- task` 上，就脱离了 `ctx.Done()` 的控制
- 正确原则：
  - 发送成功和取消通知必须在同一个阻塞点竞争
- 识别信号：
  - 看到 `len(ch)` / `cap(ch)` 预检查后立刻发送
- 替代思路：
  - 用 `select` 同时等待 `p.tasks <- task` 和 `ctx.Done()`

## 2. 用轮询模拟 channel 背压

- 现象：实现能跑，但语义不干净，延迟和测试都依赖时序
- 问题代码：

```go
ticker := time.NewTicker(100 * time.Millisecond)
for {
    select {
    case <-ticker.C:
        // 再检查一次是否有空位
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

- 为什么错：
  - 背压本来就可以用 channel 的阻塞语义直接表达
  - 用 ticker 轮询只是绕远路
  - 还会引入额外延迟和偶发时序差异
- 正确原则：
  - 能用 channel 阻塞表达的等待，不要退化成 `Ticker` 或 `Sleep`
- 识别信号：
  - 并发控制代码里出现“周期性检查状态”
- 替代思路：
  - 直接让 goroutine 阻塞在真正的同步点上

## 3. `Submit` 的 `ctx` 只适合控制入队，不适合直接当执行上下文

- 现象：一开始容易纠结“worker 执行 task 时传什么 `ctx`”
- 典型困惑：

```go
func (p *Pool) Submit(ctx context.Context, task Task) error
```

- 为什么会困惑：
  - 这个 `ctx` 的职责是控制“提交过程”
  - 不是控制“任务生命周期”
  - 如果直接把它传给 worker，语义会混乱
- 正确原则：
  - 入队取消和执行取消是两套语义
- 第一阶段可接受做法：
  - worker 先用 `context.Background()` 执行 task
- 后续演进方向：
  - 第 4 阶段再引入 pool 级别执行上下文

## 4. 测试“队列满时取消”时，没有稳定构造出满队列

- 现象：目标是测“queue full 时取消”，但实际可能根本没进入阻塞路径
- 问题代码：

```go
for i := 0; i < queueSize; i++ {
    _ = p.Submit(ctx, occupyTask)
}
```

- 为什么错：
  - 只提交 `queueSize` 个任务不代表队列会持续是满的
  - worker 可能已经提前把任务取走了
  - 后续 `Submit` 可能根本不阻塞
- 正确原则：
  - 并发边界测试必须先稳定构造前置状态，再触发被测行为
- 识别信号：
  - 测试依赖“大概率已经满了”“应该还没消费完”
- 替代思路：
  - 先控制任务不结束，让 `workerN + queueSize` 个位置都占住，再测额外 `Submit`

## 5. 在 goroutine 里断言，但测试本身不等待它结束

- 现象：测试看起来通过了，但断言可能根本没稳定执行
- 问题代码：

```go
go func() {
    err := p.Submit(doneCtx, task)
    assert.Equal(t, err, context.Canceled)
}()
cancel()
```

- 为什么错：
  - 主测试线程没有等待这个 goroutine 完成
  - 测试函数可能先返回
  - 结果是断言不稳定，甚至未真正执行完
- 正确原则：
  - 被测逻辑放进 goroutine 后，测试必须显式回收结果
- 识别信号：
  - 有 goroutine，但没有 `WaitGroup`、结果 channel、done channel
- 替代思路：
  - 用 `errCh`、`cancelCh`、`WaitGroup` 等把结果收回主测试线程再断言

## 6. “提交完成”和“执行完成”不是同一个事件

- 现象：以为等一个全局 `WaitGroup` 就够了，实际上经常只等对了一半
- 容易写出的代码：

```go
go func() {
    _ = p.Submit(ctx, task)
    wg.Done()
}()

wg.Wait()
assert.Equal(t, testVar, taskNum)
```

- 为什么错：
  - `Submit` 返回，只代表“提交结束”
  - 它不等于“task 已执行完成”
  - 反过来，只等 task 完成，也不等于每个提交协程都已经返回并检查过错误
- 正确原则：
  - 提交路径和执行路径要分开建模
- 识别信号：
  - 一个同步原语同时承担“等提交”和“等执行”两个职责
- 替代思路：
  - 提交结果单独收集
  - 任务完成单独等待

## 7. 收集了 `Submit` 错误，但没有确定性收满结果

- 现象：看似收集了所有提交错误，实际上可能漏检
- 问题代码：

```go
for i := 0; i < taskNum; i++ {
    select {
    case err := <-errCh:
        assert.Nil(t, err)
    default:
    }
}
```

- 为什么错：
  - `default` 会跳过“此刻还没写入 channel 的结果”
  - 于是测试不能保证真收到了 `taskNum` 个提交结果
- 正确原则：
  - 既然决定收集异步结果，就要确定性收满
- 识别信号：
  - 结果数量已知，却用了 `default` 非阻塞读取
- 替代思路：
  - 明确读取 `taskNum` 次，或者用额外同步保证所有发送已完成

## 8. 测试名写的是“并发消费”，但测试本身只证明了“最终消费”

- 现象：测试名很强，验证内容很弱
- 问题代码形态：

```go
task := func(ctx context.Context) error {
    atomic.AddInt64(&n, 1)
    return nil
}
```

- 为什么不够：
  - 这只能证明“任务最终跑过”
  - 即使只有一个 worker，也可能很快串行执行完，测试照样通过
- 正确原则：
  - 测试名和测试真正证明的性质必须一致
- 替代思路：
  - 要么把名字改弱一点
  - 要么显式证明“多个 worker 同时进入执行阶段”

## 9. barrier 设计不当，导致测试死锁

- 现象：测试卡住不返回
- 问题代码：

```go
task := func(ctx context.Context) error {
    startCh <- struct{}{}
    <-release
    return nil
}

for i := 0; i < workerN; i++ {
    <-startCh
}
close(release)
```

- 为什么错：
  - 如果所有 task 都会往 `startCh` 写
  - 但主测试只消费前 `workerN` 次
  - 后续 task 会把 `startCh` 塞满并卡死 worker
- 正确原则：
  - barrier 通道只服务于“证明并发”的那一小批任务
- 识别信号：
  - 所有任务共享 barrier 写入，但测试只消费一部分信号
- 替代思路：
  - 只让前 `workerN` 个任务参与 barrier
  - 或者 barrier 用完后继续把信号排空

## 10. worker 读取任务通道时，要考虑通道关闭后的零值

- 现象：第一阶段未必暴露，但第二阶段实现 `Stop` 时很容易踩坑
- 问题代码：

```go
case t := <-p.tasks:
    t(ctx)
```

- 为什么错：
  - `tasks` 被关闭后，读取会得到零值
  - 对 `Task` 来说，零值就是 `nil`
  - 直接调用会 panic
- 正确原则：
  - 从任务通道读取时，如果未来可能关闭通道，就要接 `ok`
- 替代思路：
  - `t, ok := <-p.tasks`
  - `!ok` 时退出 worker

## 11. 第一阶段最终收口时，哪些问题算已解决

当前第一阶段已经解决或收敛的点：

- 参数校验
- 固定 worker 启动
- `Submit` 的阻塞与取消语义
- `nil task` 拒绝
- 队列满时取消测试
- 用 barrier 证明至少有 `workerN` 个 worker 同时进入执行阶段
- worker 从关闭任务通道读取时的基本防御

## 12. 第一阶段后续最该记住的几条原则

### 并发实现

- “先判断再操作”在并发下通常危险
- 成功路径和取消路径要放进同一个阻塞点建模
- 能用 channel 直接表达同步，就不要轮询共享状态

### 并发测试

- 先定义你要证明的性质，再设计同步点
- 提交完成和执行完成分开等待
- 测试必须稳定构造前置状态，不能依赖“差不多”
- barrier 很有用，但只让需要参与证明的那批任务使用

### 学习重点

你这轮最需要补的不是 Go 语法，而是两件事：

1. channel / select / context 三者一起出现时的语义建模
2. 并发测试中“事件边界”和“同步对象”之间的对应关系

如果这两块练熟了，后面的 `Stop(ctx)`、统计、错误归集会顺很多。
