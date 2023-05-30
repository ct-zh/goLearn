# context包
> 受限于作者经验与理解,本文不保证时效性准确性;

## 了解context
> context定义上下文类型，它跨API边界和进程之间传递截止日期、取消信号和其他请求范围的值。

context是并发控制和超时控制的标准做法。[简单的demo](./simple/main.go)

context的使用原则:
- 不要把Context存在一个结构体当中，显式地传入函数。Context变量需要作为第一个参数使用，一般命名为ctx;
- 即使方法允许，也不要传入一个nil的Context，如果你不确定你要用什么Context的时候传一个context.TODO;
- 使用context的Value相关方法只应该用于在程序和接口中传递的和请求相关的元数据，不要用它来传递一些可选的参数;
- 同样的Context可以用来传递到不同的goroutine中，Context在多个goroutine中是安全;


## 源码解析
> 见src/context/context.go

### context对外的方法
context包一共有四个with方法:
- `WithCancel` 返回一个cancel函数，调用这个函数则可以主动停止goroutine;
- `WithValue` WithValue可以设置一个key/value的键值对，可以在下游任何一个嵌套的context中通过key获取value。但是不建议使用这种来做goroutine之间的通信;
- `WithTimeout` 函数可以设置一个time.Duration，到了这个时间则会cancel这个context;
- `WithDeadline` WithDeadline函数跟WithTimeout很相近，只是WithDeadline设置的是一个时间点;
这四个方法的[demo见](./with/with.go)

> *值得注意的是*: 不要在WithValue里面存储可变的value, 这个value应该是伴随着这个请求恒定不变的值;
> 哪些是不变的值?如:
>   - requestId
>   - userId
> 哪些显然是会变化的值?如:
>   - db连接,或者其他会close的连接
>   - auth认证有关的内容
> 总之,尽量不要用Context.Value

还有两个获取context的方法:`TODO`和`Background`,这两个方法返回的都是`emptyCtx`,`emptyCtx`从不取消，没有值，也没有截止日期;

两个error变量:
- `Canceled`: context canceled时返回的error;
- `DeadlineExceeded`: context deadline过期时的error;

以及最重要的context接口:
```go
type Context interface {
    // 该方法返回一个time和标识是否已设置deadline的bool值，如果没有设置deadline，则ok == false，此时deadline为一个初始值的time.Time值
    Deadline() (deadline time.Time, ok bool)
    
    // 当timeout或者调用cancel方法时，将会close掉该chan
    Done() <-chan struct{}

    Err() error

    Value(key interface{}) interface{}
}
```

### 私有方法
#### cancelCtx
`WithCancel`方法会创建一个`cancelCtx`
```go
type cancelCtx struct {
	Context // 保存parent context

	mu       sync.Mutex            // 保护下列字段
	done     chan struct{}         // 用来标识是否已被cancel; 当外部触发cancel、或者父Context的channel关闭时，此done也会关闭
	children map[canceler]struct{} // 保存它的所有子canceler, 第一次调用cancel函数时设置为nil
	err      error                 // 调用过cancel后就不为nil了
}
```

`cancelCtx`的主要方法:
- `Done`: Done函数返回一个chan struct{}的channel，用来判断context是否已经被close;
- `cancel`: `WithCancel`方法暴露到外面的cancel函数, 调用cancel就会走到这个逻辑;会关闭`done`,并且将所有的children ctx的done关闭;

#### timerCtx
使用`WithTimeout`和`WithDeadline`会创建一个`timerCtx`, 并注册一个函数来定时cancel:
```go
c.timer = time.AfterFunc(dur, func() {
    c.cancel(true, DeadlineExceeded)
})
```
`timerCtx`内部有一个`cancelCtx`,在cancel的时候直接调用`cancelCtx`的cancel方法,并关闭定时器;

#### valueCtx
使用`WithValue`会创建一个`valueCtx`, 比Context多key与value参数;


## 文档翻译doc translate

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



## reference
- [深度解密Go语言之context](https://zhuanlan.zhihu.com/p/68792989)
- [go程序包源码解读——golang.org/x/net/context](https://studygolang.com/articles/5131)
- [Golang 如何正确使用 Context](https://studygolang.com/articles/23247?fr=sidebar)