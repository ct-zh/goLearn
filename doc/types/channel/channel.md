## channel

### channel的基本结构
用dlv debug下面代码([代码参考](./ch/main.go)), 可以知道不管是`buffer channel` 还是 `unbuffer channel` 都是通过`runtime.makechan`来创建:

```go
ch1 := make(chan int)
ch2 := make(chan int, 3)

main.go:6   lea rax, ptr [rip+0x77f3]
main.go:6   xor ebx, ebx
main.go:6   call $runtime.makechan
main.go:6   mov qword ptr [rsp+0x40], rax

main.go:7   lea rax, ptr [rip+0x77e0]
main.go:7   mov ebx, 0x3
main.go:7   call $runtime.makechan
main.go:7   mov qword ptr [rsp+0x38], rax

```

在runtime包里搜makechan函数, 可以得到其声明为:`func makechan(t *chantype, size int) *hchan`, 返回的hchan结构为:

```go
type hchan struct {
	qcount   uint           // 队列中的总数据
	dataqsiz uint           // 循环队列的大小
	buf      unsafe.Pointer // points to an array of dataqsiz elements  环形数组
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // 发送方的索引
	recvx    uint   // 接收方的索引
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters


    // lock保护hchan中的所有字段，以及该channel上被封锁的sudogs中的几个字段。在持有此锁时，不要更改另一个G的状态（特别是，不要准备好G），因为这可能会导致堆栈收缩而陷入僵局。
	lock mutex
}
```

- channel数据存储在buf中, 指向的是一个环形数组, 依靠`sendx`与`recvx`两个游标来判断写入ch的数据位置与接受数据位置.
- `recvq`与`sendq` 用来保存挂起的g的队列, 分别代表buf为空等待接收的g列表, 与buf已满等待发送的g列表
- 虽然大部分书上都提到并发处理不要用锁, 而是用channel. 但是其实channel底层也还是带了一把锁



### 写channel

继续使用dlv查找往channel中写入数据的处理, 往channel中写数据使用的是`runtime.chansend1`方法
```go
// go代码
ch2 <- 1

// 汇编代码
mov rax, qword ptr [rsp+0x38]
lea rbx, ptr [rip+0x27d50]
call $runtime.chansend1
```

