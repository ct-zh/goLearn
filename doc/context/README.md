# doc translate

Package context defines the Context type, 
which carries deadlines, cancelation signals, 
and other request-scoped values across API boundaries and between processes.

> context包定义了 超时、取消信号和 api边界以及两个进程之间 请求值 的 上下文类型

Incoming requests to a server should create a Context, 
and outgoing calls to servers should accept a Context. 
The chain of function calls between them must propagate the Context, 
optionally replacing it with a derived Context created using WithCancel,
 WithDeadline, WithTimeout, or WithValue. When a Context is canceled, 
 all Contexts derived from it are also canceled.
 
> 每个发向服务器的请求都需要创建context，并且对服务器的外部调用需要接受context
> 函数之间的调用链需要传递这个context，也可以用派生的context替换他，
> 例如WithCancel,WithDeadline,WithTimeout,或者 WithValue
> 当一个context失效，所有派生context都应失效.

The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent) 
and return a derived Context (the child) and a CancelFunc. 
Calling the CancelFunc cancels the child and its children, 
removes the parent's reference to the child, and stops any associated timers. 
Failing to call the CancelFunc leaks the child and its children until the parent is canceled or the timer fires. 
The go vet tool checks that CancelFuncs are used on all control-flow paths.

> WithCancel, WithDeadline, 和 WithTimeout这三个函数会携带他们的父类context
> 并且返回派生的context以及失效函数
> 调用失效函数来无效化子类以及子类的子类，移除父类对子类的连接并且停止所有相关连的计时器
> 调用失效函数失败会导致该函数以及其子类的泄漏，直到该函数的父类失效或者计时器记数完成
> go的vet工具能检查  失效函数是否用于所有的控制流路径

Programs that use Contexts should follow these rules to keep interfaces consistent 
across packages and enable static analysis tools to check context propagation:
Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. 
The Context should be the first parameter, typically named ctx:

> 使用context的程序需要遵循包的接口一致规则(duck typing ?) 并且启用静态分析工具来检查程序这些内容:
> 不要将context保存在一个结构体中，相反，应该将context显式地传递给每一个需要他的函数
> context需要作为第一个个参数，并且被命名为ctx：

```
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
```


Do not pass a nil Context, even if a function permits it. 
Pass context.TODO if you are unsure about which Context to use.
Use context Values only for request-scoped data that transits processes and APIs, 
not for passing optional parameters to functions.
The same Context may be passed to functions running in different goroutines; 
Contexts are safe for simultaneous use by multiple goroutines.

> 即使函数允许，也不要传递空的context. 如果你不确信要使用哪个context,请使用TODO
> 仅对传输进程和API的请求范围内的数据使用context，而不用于向函数传递可选参数
> 同一个context可能运行在多个协程上面；context在多协程同步运行的环境中是安全的


See [blog](https://blog.golang.org/context) for example code for a server that uses Contexts.


# 了解context
1. context 几乎成为了并发控制和超时控制的标准做法。


# 参考资料
> 深度解密Go语言之context https://zhuanlan.zhihu.com/p/68792989
> go程序包源码解读——golang.org/x/net/contex https://studygolang.com/articles/5131
>
> Golang 如何正确使用 Context https://studygolang.com/articles/23247?fr=sidebar