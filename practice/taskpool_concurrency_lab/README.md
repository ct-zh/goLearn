# Task Pool Concurrency Lab

这是一个偏原理训练型的 Go 并发练习。

目标是实现一个“带背压、可优雅关闭、支持取消和超时、可观测基础统计”的任务执行池。这个题目表面上像普通 worker pool，但真正要练的是：

- goroutine 生命周期管理
- 有界队列和背压
- `context.Context` 的取消传播
- `channel` 的关闭责任边界
- `Submit` / `Stop` 并发下的竞态处理
- 并发测试的可验证写法

只允许使用 Go 标准库。

补充约定：

- 生产代码使用标准库实现
- 单元测试允许使用 `github.com/stretchr/testify/assert`

## 背景设定

你在实现一个本地调度组件。上游不断提交任务，执行池负责并发消费。

系统约束：

- worker 数量固定
- 队列容量有限，不能无限缓存任务
- `Submit` 在队列满时需要阻塞等待
- `Submit` 在等待期间必须响应 `ctx.Done()`
- `Stop` 后不再接收新任务
- `Stop` 需要尽量等已入队任务执行完
- 基础统计需要可并发读取

## 推荐语义

你可以调整设计，但如果要改，建议先写在代码注释里说明。

### `Submit(ctx, task)`

- pool 已停止：返回错误
- 队列未满：成功入队
- 队列已满：阻塞等待
- 等待期间 `ctx.Done()`：返回错误
- `task == nil`：返回错误

### `Stop(ctx)`

- 幂等
- 停止接收新任务
- 等待已入队任务尽量执行完
- 若等待期间 `ctx.Done()`，返回错误

### `Stats()`

建议至少包含：

- `Submitted`
- `Started`
- `Succeeded`
- `Failed`
- `Rejected`

## 分阶段任务

### 阶段 1（约 30 分钟）

实现最小可运行骨架：

- `New(workerN, queueSize int) (*Pool, error)`
- `Submit(ctx context.Context, task Task) error`
- 固定数量 worker 启动
- 有界任务队列
- 队列满时 `Submit` 可被 `ctx` 取消

建议测试：

1. 提交任务后会被执行
2. 多个任务可以被并发 worker 消费
3. 队列满时 `Submit` 阻塞，并在 `context` 超时后返回
4. 非法参数返回错误

### 阶段 2（约 30 分钟）

实现优雅关闭：

- `Stop(ctx)` 实现
- stop 幂等
- stop 后拒绝新任务
- 已入队任务尽量执行完

### 阶段 3（约 30 分钟）

实现统计：

- 线程安全统计
- 成功 / 失败 / 拒绝计数
- 统计测试

### 阶段 4（约 30 分钟）

提升项，二选一或都做：

- 为任务执行增加 pool 级别取消上下文
- 增加 `TrySubmit(task)` 非阻塞提交

## Review 标准

我在 review 时会重点看：

- 并发语义是否先定义清楚
- channel 关闭责任是否单一
- 是否存在 goroutine 泄漏
- 是否有 `Submit` / `Stop` 的竞态
- 测试是否验证了边界，而不是只测 happy path
- API 是否克制，没有过早抽象

## 复盘文档

第一阶段的详细问题复盘已单独整理在：

- `stage1_review_notes.md`

建议在进入第 2 阶段前先通读这份文档，重点看：

- `Submit` 的阻塞与取消语义
- 并发测试中“提交完成”和“执行完成”的区别
- barrier 的正确使用方式

## 当前目录说明

- `taskpool/pool.go`
  - 第 1 阶段基础骨架
- `taskpool/pool_test.go`
  - 测试模板和占位说明

你现在要做的是：先完成阶段 1，再把代码给我 review。
