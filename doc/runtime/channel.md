## channel
### channel用法
- [创建channel](./channel/basic/main.go)
- [select配合channel](./channel/select/select.go)
- [如何优雅关闭channel、channel的关闭原则](./channel/closeChan/closeChan.go)

### channel的结构
> 从这里开始 start from `src/runtime/chan.go`
```go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // 循环队列的大小
	buf      unsafe.Pointer // 指向dataqsiz元素的数组
	elemsize uint16
	closed   uint32
	elemtype *_type // 元素类型
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	// 锁保护了hchan中的所有字段，以及sudogs中被此通道阻塞的几个字段。
    // 
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
    // 在持有这个锁时不要更改另一个G的状态（特别是，不要准备好一个G），因为这会导致堆栈收缩而死锁。
	lock mutex
}
```
同步/异步(buffed channel/unbuffed channel)的区别,在于是否有缓存`buf`;

> sudog: 表示等待列表中的g，例如用于在信道上发送/接收。见`waitq`类型;

### 初始化makechan
> 源码位置:`src/runtime/chan.go:makechan()`


### 发送
> 源码位置:`src/runtime/chan.go` `chansend()` `send()`



### 关闭
> 源码位置:`src/runtime/chan.go` `closechan()`

### 接收
> 源码位置:`src/runtime/chan.go` `chanrecv()` `recv()`
