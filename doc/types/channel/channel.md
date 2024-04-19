## channel

> 以下代码 see [代码参考](./ch/main.go) ; 源码分析基于go1.20

### channel的基本结构
用dlv debug下面代码, 可以知道不管是`buffer channel` 还是 `unbuffer channel` 都是通过`runtime.makechan`来创建:

```go
ch1 := make(chan int)
ch2 := make(chan int, 3)

main.go:6   lea rax, ptr [rip+0x77f3]
main.go:6   xor ebx, ebx
main.go:6   call $runtime.makechan    // <- 在这里调用了makechan函数
main.go:6   mov qword ptr [rsp+0x40], rax // 将 RAX 寄存器中的值(即刚刚创建的 channel 的地址）存储到栈顶指针（RSP）的偏移量为 0x40 的位置上。

main.go:7   lea rax, ptr [rip+0x77e0]
main.go:7   mov ebx, 0x3				// 立即数 0x3（十进制为 3）移动到 EBX 寄存器中。这个值代表了创建 channel 时的缓冲区大小。
main.go:7   call $runtime.makechan  // <- 在这里调用了makechan函数
main.go:7   mov qword ptr [rsp+0x38], rax // 将 RAX 寄存器中的值（即刚刚创建的 channel 的地址）存储到栈顶指针（RSP）的偏移量为 0x38 的位置上。

```

在runtime包里搜makechan函数, 可以得到其声明为:`func makechan(t *chantype, size int) *hchan`, 返回的hchan结构为:

```go
type hchan struct {
	qcount   uint           // 队列中的总数据大小
	dataqsiz uint           // 循环队列的大小
	buf      unsafe.Pointer // points to an array of dataqsiz elements  环形数组
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // 发送方的索引
	recvx    uint   // 接收方的索引
	recvq    waitq  // 接收方阻塞的g队列
	sendq    waitq  // 发送方阻塞的g队列

    // lock保护hchan中的所有字段，以及该channel上被封锁的sudogs中的几个字段。在持有此锁时，不要更改另一个G的状态（特别是，不要准备好G），因为这可能会导致堆栈收缩而陷入僵局。
	lock mutex
}
```

- channel数据存储在buf中, 指向的是一个环形数组, 依靠`sendx`与`recvx`两个游标来判断写入ch的数据位置与接受数据位置.
- `recvq`与`sendq` 用来保存挂起的g的队列, 分别代表buf为空等待接收的g列表, 与buf已满等待发送的g列表
- ps.虽然大部分书上都提到并发处理不要用锁, 而是用channel. 但是其实channel底层也还是带了一把锁


### 写channel

写channel应该有四种情况:
- 写入nil chan
- unbuffer channel直接写入:  `ch2 <- 1`
- buffer channel未阻塞写入、 阻塞写入
- 已关闭的channel写入

继续使用dlv查找往channel中写入数据的处理:
```go
// 情况0: 写入nil chan
var ch3 chan int
ch3 <- 1

// ==== 汇编
xor eax, eax
lea rbx, ptr [rip+0x27f15]
nop dword ptr [rax+rax*1], eax
call $runtime.chansend1

// 情况一: 写入unbuffer ch
ch1 <- 1

// ==== 汇编
lea rbx, ptr [rip+0x27e6c]
call $runtime.chansend1

// 情况二: 写入buffer ch
ch2 <- 1

// 汇编代码
mov rax, qword ptr [rsp+0x38]
lea rbx, ptr [rip+0x27d50]
call $runtime.chansend1


// 情况三: 写入已经关闭的ch
close(ch1)
ch1 <- 1

// 汇编
mov rax, qword ptr [rsp+0x40]
lea rbx, ptr [rip+0x27f11]
nop
call $runtime.chansend1
```
可以看到三种情况下往channel中写数据都是使用`runtime.chansend1`方法. 通过阅读方法源码, 我们可以将其大致简化为以下步骤:

1. 检查 channel 是否为空
2. 非阻塞模式(select)提前判断 : 检查channel是否关闭以及是否已满，如果不满足发送条件则直接返回false
3. 加锁操作;
4. 检查 channel 是否已关闭:如果 channel 已关闭，则解锁并抛出异常 "send on closed channel"。
5. 如果存在等待接收数据的g(`recvq`队列中弹出)，则直接将数据传递给它(runtime.send方法)，跳过 channel缓冲区
6. 如果`qcount`小于`dataqsiz`,写入`buf`并返回true;
7. 非阻塞模式, 解锁并返回false
8. 阻塞模式, 推入`sendq`,挂起当前g;
9. 出现接收方后唤醒,返回true.


### 读channel
读channel应该有如下几种情况:
- 读取nil chan
- 读取一个空的channel/读unbuffer channel
- 读取一个已经关闭的channel, 并且该channel内没有数据
- 读取一个写满了,并且sendq存在挂起g的buffer channel
- 读取一个未写满或者刚写满的buffer channel

同样使用dlv查找读channel的汇编代码:
```go
// 读取unbuffer ch
<-ch1

mov rax, qword ptr [rsp+0x48]
xor ebx, ebx
call $runtime.chanrecv1

// 读取buffer ch
<-ch2

mov rax, qword ptr [rsp+0x10]
xor ebx, ebx
call $runtime.chanrecv1

// 读取已经关闭的ch
close(ch1)
<-ch1

xor ebx, ebx
call $runtime.chanrecv1

// 读取nil chan
<-ch3    

xor eax, eax
xor ebx, ebx
call $runtime.chanrecv1
```
可以发现也都是在`runtime.chanrecv1`方法中进行读取操作.recv函数大致步骤如下:

1. 存在两个返回值, 其中第一个基本默认为true, 而第二个参数`received`表示是否成功接收到了数据;
2. 检查 channel 是否为空;
3. 非阻塞模式判断 channel 状态 : 
   - 通过原子操作检查 channel 是否关闭以及是否为空。
   - 如果 channel 已经关闭并且为空，则表示没有数据可接收，返回 `selected: true` 和 `received: false`。
4. 加锁
5. 如果 channel 已经关闭，并且缓冲区为空，则表示没有数据可接收，解锁并返回 `received: false`
6. 查看`sendq`是否有阻塞g, 有则直接从里面拿数据;  // todo 测试一下写阻塞ch后第一次recev是否拿到的是阻塞的g数据, 而不是buf里某个g的数据
7. `qcount`大于0, 从buff中拿数据, 解锁, 返回received=true;
8. 非阻塞模式, 直接返回received=false;
9. 阻塞模式, gopark



### 关闭channel

dlv查看关闭channel调用的函数:

```go
close(ch1)
// mov rax, qword ptr [rsp+0x48]
// call $runtime.closechan
```

函数为`runtime.closechan`, 查看其实现:

1. 判断chan是否为nil,是则返回panic close of nil channel
2. 加锁
3. 判断是否已经关闭了,如果是,则返回panic:close of closed channel
4. 将 channel 的 `closed` 标志设置为 1，表示 channel 已关闭
5. 遍历`recvq`,
   -  将g取出并标记失败
   - 如果程序使用了堆栈内存存储接收数据，则将其清空 (`typedmemclr`)
   - 将程序添加到 `glist` 列表中
6. 遍历sendq
   - 依次将这些程序从等待队列中取出并标记为发送失败 (`success` 设置为 `false`)。
   - 清空程序可能用于发送数据的内存 (`sg.elem`)。
   - 将程序添加到 `glist` 列表中
7. 解锁
8. 遍历 `glist` 列表，将所有等待的程序标记为可运行状态 (`goready`)



**需要注意的点:**

- 关闭 channel 是一个不可逆的操作，一旦关闭，就不能再向 channel 发送数据。
- 关闭 channel 会唤醒所有等待接收或发送数据的程序，这些程序可能会由于 channel 关闭而抛出异常。



**glist 列表:**

- `glist` 列表用来临时存储所有因 channel 关闭而需要唤醒的程序。
- 在释放 channel 锁之后再唤醒这些程序，可以避免出现竞争条件。

### 源码解析
#### makechan



#### chansend1
`chansend1`实际调用的是`chansend`, chansend函数声明为: `func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool`其参数含义为:

- c: 指向 channel 的指针
- ep: 要发送的数据的指针
- block: 是否阻塞等待发送成功, 在`c <- 1`情况下默认为true, 在select里默认为false
- callerpc: 调用方的PC寄存器地址, 这个使用stubs的通用方法`getcallerpc()`获取
- 返回值true=发送成功, false=发送失败 (非阻塞模式下 channel 满或者 channel 已关闭)


代码逻辑: 
首先检查 channel 是否为空:如果为空，并且是非阻塞模式，则直接返回 false。如果为空，并且是阻塞模式，则挂起当前线程。
```go
if c == nil {
	if !block {
		return false
	}
	gopark(nil, nil, waitReasonChanSendNilChan, traceEvGoStop, 2)
	throw("unreachable")
}
```

#### chanrecv
函数声明:`func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool)`





