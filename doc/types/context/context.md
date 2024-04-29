# context包

## 了解context
> context定义上下文类型，它跨API边界和进程之间传递截止日期、取消信号和其他请求范围的值。

context是并发控制和超时控制的标准做法。[简单的demo](./simple/main.go)

context的使用原则:
- 不要把Context存在一个结构体当中，显式地传入函数。Context变量需要作为第一个参数使用，一般命名为ctx;
- 即使方法允许，也不要传入一个nil的Context，如果你不确定你要用什么Context的时候传一个context.TODO;
- 使用context的Value相关方法只应该用于在程序和接口中传递的和请求相关的元数据(不变的数据)，不要用它来传递一些可选、可变的参数;
- 同样的Context可以用来传递到不同的goroutine中，Context在多个goroutine中是安全;

^ _ ^ 是不是看完以上内容可能还一头雾水 ? 我们下面直接看源码, 看能不能找到答案:


## 源码解析
> 见src/context/context.go

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool) // 该方法返回一个time和标识是否已设置deadline的bool值，如果没有设置deadline，则ok == false，此时deadline为一个初始值的time.Time值

    Done() <-chan struct{} // // 当timeout或者调用cancel方法时，将会close掉该chan
    
    Err() error
    
    Value(key any) any
}
```

context分为多个基础类型, 皆实现了Context这个接口:
- `emptyCtx` 空的ctx, 来源context.TODO与context.Background

  ```go
  var (
  	background = new(emptyCtx)
  	todo       = new(emptyCtx)
  )
  // 实现了Context interface的全部方法, 返回皆是nil
  type emptyCtx int
  ```

- `valueCtx`: 在emptyCtx上包了一层,带有key/value的context,来源为context.WithValue;

  ```go
  type valueCtx struct {
  	Context
  	key, val any
  }
  // 额外提供了String方法与Value方法
  
  // 返回context名+key+val组合的字符串
  func (c *valueCtx) String() string
  
  // 如果key=c.key则返回c.val, 否则去父级的Context递归查找
  func (c *valueCtx) Value(key any) any
  ```

- `cancelCtx`: 同样是在emptyCtx上包了一层,带有cancel函数的context;来源为context.WithCancel;

  ```go
  type cancelCtx struct {
  	Context
  
  	mu       sync.Mutex					// 加锁防止并发
  	done     atomic.Value          // 类型为chan struct{}  懒加载, 第一个调用Done方法的会初始化该值
  
  	children map[canceler]struct{}
  	err      error
  	cause    error
  }
  ```

- `timerCtx`: 带有定时器的context,在cancelCtx上包了一层, 来源为context.WithTimeout与contxt.WithDeadline;

  ```go
  type timerCtx struct {
  	*cancelCtx
  	timer *time.Timer
  
  	deadline time.Time
  }
  ```


### context公共方法
context包一共有四个with初始化方法:
- `WithCancel` 返回ctx与一个cancel函数，调用这个函数则可以主动停止goroutine;
- `WithValue` 可以设置一个key/value的键值对，在下游任何一个嵌套的context中通过key获取value。但是不建议使用这种来做goroutine之间的通信;并且建议使用自定义类型作为key来防止冲突;
- `WithTimeout` 函数可以设置一个time.Duration，到了这个时间则会cancel这个context;
- `WithDeadline` 函数跟WithTimeout很相近，只是WithDeadline设置的是一个时间点, timeout是倒计时;
这四个方法的[demo见](./with/with.go)

> *值得注意的是*: 不要在WithValue里面存储可变的value, 这个value应该是伴随着这个请求恒定不变的值;
> 哪些是不变的值?如:
>
>   - requestId
>   - userId
> 哪些显然是会变化的值?如:
>   - db连接,或者其他会close的连接
>   - auth认证有关的内容
> 总之,尽量不要用Context.Value来保存/传递数据,而是存储*常量*

还有两个获取context的方法:`TODO`和`Background`,这两个方法返回的都是`emptyCtx`,`emptyCtx`从不取消，没有值，也没有截止日期; 当你的ctx不知道传什么的时候就可以传这两个.

两个error变量:
- `Canceled`: context canceled时返回的error;
- `DeadlineExceeded`: context deadline过期时的error;


### 私有方法
#### cancelCtx

`WithCancel`方法会创建一个`cancelCtx`. 下面是结构

```go
type cancelCtx struct {
	Context // 保存parent context

	mu       sync.Mutex            // 保护下列字段防止并发操作
	done     chan struct{}         // 用来标识是否已被cancel; 当外部触发cancel、或者父Context的channel关闭时，此done也会关闭
	children map[canceler]struct{} // 保存它的所有子canceler, 第一次调用cancel函数时设置为nil
	err      error                 // 调用过cancel后就不为nil了
}

type canceler interface {
	cancel(removeFromParent bool, err, cause error)
	Done() <-chan struct{}
}
```

使用WithCancel初始化cancelCtx, 会返回一个闭包类型的cancel函数:

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := withCancel(parent)
	return c, func() { c.cancel(true, Canceled, nil) }
}
```

其中`withCancel`方法会初始化`cancelCtxs结构体, 并调用`propagateCancel`, 这个方法作用是:当parent调用cancel时, 安排其children也执行cancel; 而闭包cancel函数调用的是context包内置的cancel方法; 下面我们来分析这两个函数


##### propagateCancel

先看该函数第一部分:

```go
func propagateCancel(parent Context, child canceler) {
	done := parent.Done()		// 获取parent.Done方法; 其中cancelCtx与timerCtx会返回chan; 
	if done == nil { // emptyCtx与valueCtx会返回nil, 直接退出
		return
	}

	select {
	case <-done:	// done 已经closed 直接cancel
		child.cancel(false, parent.Err(), Cause(parent))
		return
	default:	// default不会阻塞select, 继续往下执行
	}
	// ......
}
```

- 首先获取父context的Done方法, 这个方法会返回一个`chan struct{};`
- 如果返回的是nil, 说明该父context永远不会取消(没有cancel属性) 直接退出;
- 监听父context的Done chan, 如果已经取消(父context  close Done chan), 则当前context执行cancel函数, 否则继续往下执行;

```go
func propagateCancel(parent Context, child canceler) {
  	// ......
  if p, ok := parentCancelCtx(parent); ok {	// 判断parent ctx是否是原生的cancelCtx
		// ......
	} else {	// 非原生cancelCtx, 需要开启一个goroutine去监听结束事件;
		// .....
	}
}
```

- 首先执行`parentCancelCtx`函数并判断结果,  该函数的流程是:  (question: 这个函数的作用是什么?)

  ```go
  func parentCancelCtx(parent Context) (*cancelCtx, bool) {
  	done := parent.Done()
  	if done == closedchan || done == nil {
  		return nil, false
  	}
  	p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)	// 从父context的value中取出值并将其类型转成cancelCtx
  	if !ok {
  		return nil, false
  	}
  	pdone, _ := p.done.Load().(chan struct{})
  	if pdone != done {
  		return nil, false
  	}
  	return p, true
  }
  ```

  - `p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)` 从父context的value中取出值并将其类型转成cancelCtx; 

    - 这里有两个疑问: `&cancelCtxKey` 是什么? cancelCtx调用Value方法的逻辑是什么?

    - 解释: Context全局变量:`var cancelCtxKey int`, 用`&cancelCtxKey`获取该变量的指针, 将其保存在cancelCtx的value中, 结合cancelCtx的Value方法(如下), 当key为cancelCtxKey的地址时, 就将cancelCtx自己返回;

      ```go
      func (c *cancelCtx) Value(key any) any {
      	if key == &cancelCtxKey {
      		return c
      	}
      	// ..... 
      }
      ```

    - 在上述代码中, `parent.Value(&cancelCtxKey)`, 如果parent是cancelCtx, 相当于返回parent自己; 再将其转换为*cancelCtx类型;

  - `pdone, _ := p.done.Load().(chan struct{})` 对上面取出的cancelCtx(也就是parent)执行done.Load()方法, 得到一个`chan struct{}`

  - 如果这个pdone与`parent.Done`是同一个chan, 则返回`p`与ok

- 所以由上可知`parentCancelCtx`函数的作用是:

  -  判断parent是不是CancelCtx, 如果重写了Value函数, 或者非cancelCtx调用了Value, 不会返回canceCtx类型的parent;
  - pdone是从done.Load中取出来的,有可能Done函数经过重写, 其类型或者值不一定与pdone是同一个chan stuct{}
  - 所以该函数的作用是, 证明parent是原生的cancelCtx, 没有经过重写或者Done函数返回的chan与done保存的chan是同一个chan;


- 如果pdone与parent.Done是同一个chan,，执行以下逻辑：

  ```go
  func propagateCancel(parent Context, child canceler) {
    	// ......
    if p, ok := parentCancelCtx(parent); ok {	// 判断parent ctx是否是原生的cancelCtx
  		p.mu.Lock()
  		if p.err != nil {
  			child.cancel(false, p.err, p.cause) 			// parent has already been canceled
  		} else {
  			if p.children == nil {
  				p.children = make(map[canceler]struct{})
  			}
  			p.children[child] = struct{}{}	// parent是原生cancelCtx, 直接把当前的ctx追加到parent的map中
  		}
  		p.mu.Unlock()
  	} else {	// 非原生cancelCtx, 需要开启一个goroutine去监听结束事件;
  		// ......
  	}
  }
  ```

  - 父ctx加锁
  - 如果父ctx的err不为空, 说明父ctx已取消, 调用child.cancel来取消子ctx;
  - 否则, 将子ctx存入父ctx的childre map中

- 如果parent不是原生的cancelCtx，则：

  ```go
  func propagateCancel(parent Context, child canceler) {
    	// ......
    if p, ok := parentCancelCtx(parent); ok {	// 判断parent ctx是否是原生的cancelCtx
  		// ......
  	} else {	// 非原生cancelCtx, 需要开启一个goroutine去监听结束事件;
  		goroutines.Add(1)
  		go func() {
  			select {
  			case <-parent.Done():	// 手动监听parent的Done事件
  				child.cancel(false, parent.Err(), Cause(parent))		// 再手动cancel当前ctx
  			case <-child.Done():
  			}
  		}()
  	}
  }
  ```

  - Context包全局的goroutines计数器+1
  - 开启 goroutine ，select监听父ctx与子ctx, 父ctx如果chan关闭,则子ctx执行cancel; 如果子ctx关闭, 则直接退出g;

>  总结`propagateCancel`函数的作用是: 当parent取消时, childre也执行取消; 
>
> - 所以需要判断parent当前是否已经取消了, 如果已取消, 当前ctx也执行取消;
>
> - 如果parent是正儿八经的cancelCtx, 则只需要将当前ctx追加到parent的children map里就行了. 在parent ctx调用自身的cancel()方法时会自动遍历children map去执行子ctx的cancel()方法
> - 如果parent不是正儿八经的cancelCtx, 那只能由这个函数开一个全局的g去监听parent与自身ctx的的Done了, 当parentDone, 则调用当前ctx的cancel; 如果是当前ctx Done, 则退出协程.

##### cancelCtx.Done()

公开的Done方法如下, 用于*初始化并返回一个chan strcut{}*, 该chan在cancelCtx被取消时关闭;

```go
func (c *cancelCtx) Done() <-chan struct{} {
	d := c.done.Load()
	if d != nil {
		return d.(chan struct{})
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	d = c.done.Load()
	if d == nil {
		d = make(chan struct{})
		c.done.Store(d)
	}
	return d.(chan struct{})
}
```

- 先检查done保存的值是否为nil, done是`atomic.Value`类型, 支持并发操作;
- 如果done不为nil, 说明已经初始化结束, 直接返回chan;
- 如果为nil, 需要加锁防止并发store. 新建一个chan保存到done字段中, 并返回;

##### cancelCtx.cancel()

这是cancelCtx最主要的cancel方法, 用于取消上下文; 

```go
func (c *cancelCtx) cancel(removeFromParent bool, err, cause error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
	if cause == nil {
		cause = err
	}
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return // already canceled
	}
	c.err = err
  c.cause = cause
  // ......
}
```

- 先是一些校验, 然后加锁, 给err与cause赋值; 其中err表示context的错误, cause表示context被取消的原因;

```go
func (c *cancelCtx) cancel(removeFromParent bool, err, cause error) {
  // .......
	d, _ := c.done.Load().(chan struct{})
  if d == nil {
    c.done.Store(closedchan)
  } else {
    close(d)
  }
  for child := range c.children {
    child.cancel(false, err, cause)	// 这里该g会一直持有父ctx的锁, 去遍历子ctx的cancel
  }
  c.children = nil
  c.mu.Unlock()
	if removeFromParent {
		removeChild(c.Context, c)
	}
}
```

- 加载chan, 如果chan为空, 直接写入固定的closedchan;否则关闭chan; 
- 遍历子ctx, 全部调用cancel
- 清空cancel、解锁;
- 如果`removeFromParent` 为true,则从父ctx的children中移除(delete)当前ctx;
  - 查看cancel调用位置, 只有`WithCancel`与`WithCancelCause`两个公共方法给对外提供的闭包函数中, removeFromParent=true;  其他所有调用都是false
  - 也就说**只有外部主动调用Cancel才会从父ctx的children中摘除当前ctx**;

##### cancel的运行流程

- 外部调用`WithCancel`或者`WithCancelCause`两个函数获得初始化的`cancelCtx`与`cancelFn`;
  - cancelCtx.Context 赋值其parent context;
  - 同时执行`propagateCancel`函数, 保证其函数cancel时, 子ctx都会cancel;

- 外部使用ctx.Done 监听结束事件; 此时会初始化cancelCtx.done; 如果外部没有调用Done方法, 则cancelCtx.done不会初始化;
- 外部调用cancelFn, 如果cancelCtx done未初始化则直接写一个close chan, 否则close chan; 然后递归调用子ctx的cancel.



#### timerCtx

先看结构, 可以发现timerCtx是继承自cancelCtx, 它有一个Deadline方法获取过期时间:

```go
type timerCtx struct {
	*cancelCtx
	timer *time.Timer
	deadline time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}
```

初始化函数WithDeadline与WithTimeout都是指向同一个函数:

```go
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
  if parent == nil {
		panic("cannot create context from nil parent")
	}
  if cur, ok := parent.Deadline(); ok && cur.Before(d) {	// 如果父ctx的过期时间比当前早,则返回WithCancel包装的ctx
		return WithCancel(parent)
	}
  c := &timerCtx{	// 新建timerCtx
		cancelCtx: newCancelCtx(parent),
		deadline:  d,
	}
  propagateCancel(parent, c)
  dur := time.Until(d)	// 获取当前到d的时间
	if dur <= 0 {	// 小于等于0则直接cancel
		c.cancel(true, DeadlineExceeded, nil) // deadline has already passed
		return c, func() { c.cancel(false, Canceled, nil) }
	}
  c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {	// 正常注册timer
		c.timer = time.AfterFunc(dur, func() {
			c.cancel(true, DeadlineExceeded, nil)
		})
	}
	return c, func() { c.cancel(true, Canceled, nil) }
}
```

- 如果父ctx的过期时间比当前早,则返回WithCancel包装的ctx; 因为父ctx先到时间, 先调用cancel遍历子ctx调用cancel; 所以子ctx不需要再计时了;
- 对于已经过期的timerCtx, 直接调用cancel; 
- 对于未过期的timerCtx, 注册timer, 延迟执行cancel; 返回ctx与cancelFn
- 对于timerCtx, 有两种cancel调用, 一种是超时`DeadlineExceeded`, 一种是外部主动调用cancelFn的`Canceled`

##### timer.Ctx.cancel

timerCtx的cancel方法比较简单, 调用cancelCtx的cancel即可, 代码如下:

```go
func (c *timerCtx) cancel(removeFromParent bool, err, cause error) {
	c.cancelCtx.cancel(false, err, cause)
	if removeFromParent {
		removeChild(c.cancelCtx.Context, c)
	}
	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	c.mu.Unlock()
}
```



#### valueCtx

先看结构

```go
type valueCtx struct {
	Context
	key, val any
}
func (c *valueCtx) Value(key any) any {
	if c.key == key {
		return c.val
	}
	return value(c.Context, key)
}
```

在emptyCtx的基础上增加了key、val两个字段, 用于存储数据. `valueCtx`重写了Value方法, 递归查找与key相等的value值. 

使用valueCtx需要注意几点:

- key需要是可比较的(Comparable), 如果不可比较, 就无法判断Value函数中传入的key与valueCtx.key是否相等了;
- key的值最好是独一无二的值,防止产生冲突; 可以从两点入手:
  -  使type独一无二, 可以自定义一个type:  `type MyCtxValue int` 这种
  - 使value独一无二, 比较简单的方法就是传入地址, 如`ctx.Value(&cancelCtxKey)`

- key-val的值最好是元数据, 即在某个场景下描述某个数据的值. 例如: 当次请求的ip地址; 当次请求的uid(前提是这个uid固定不变); 



#### emptyCtx

先看结构:

```go
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}
func (*emptyCtx) Done() <-chan struct{} {
	return nil
}
func (*emptyCtx) Err() error {
	return nil
}
func (*emptyCtx) Value(key any) any {
	return nil
}
func (e *emptyCtx) String() string {
	// ......
}
```

emptyCtx实现了Context这个接口的所有方法. 除此之外没有任何功能;



##### 解答

上面我们分析了context这个包的源码, 从中我们应该可以得出一些有关文章最开始提出的那几条规则的结论:

- 不要把Context存在一个结构体当中，显式地传入函数。Context变量需要作为第一个参数使用，一般命名为ctx;

  >  因为Context是一个树形结构, 表示当前逻辑的调用链; 如果在结构体中存储context并使用,则生成的就不是树形结构了; 

- 即使方法允许，也不要传入一个nil的Context，如果你不确定你要用什么Context的时候传一个context.TODO;

  > 1. 传入nil  context会报panic;
  > 2. 对于所有With方法, 都需要传入父ctx, 所以根节点必然是不需要父ctx的 `context.TODO()`或者`context.Background()`
  > 3. 也是因为根节点必然是一个`emptyCtx`, 该ctx实现了Context接口的所有方法; 所以其他种类的ctx都不需要再去实现这些方法了. 调用对应的方法会往上溯源到emptyCtx的对应方法;

- 使用context的Value相关方法只应该用于在程序和接口中传递的和请求相关的元数据(不变的数据)，不要用它来传递一些可选、可变的参数;

  > valueCtx的设计思想就是保存不变的数据; 因此都没有提供写操作;

- 同样的Context可以用来传递到不同的goroutine中，Context在多个goroutine中是安全;

  >  对于存在并发操作的timerCtx与cancelCtx, 内部有锁保证并发安全;
  >
  > 对于valueCtx, 因此我们需要保证保存的val是不变的, 只有读操作没有写操作;





## context的问题

因为所有goroutine都是从g0开始往下开的一个树形结构, 当ctx作为参数传递时同样在context内部也形成了一个与goroutine相同的树形结构. 某个节点ctx取消时, 可以传导到所有的子节点.但是问题是ctx必须侵入代码, 也就是说必须要写select ctx来配合ctx的取消逻辑.

## reference
- [深度解密Go语言之context](https://zhuanlan.zhihu.com/p/68792989)
- [go程序包源码解读——golang.org/x/net/context](https://studygolang.com/articles/5131)
- [Golang 如何正确使用 Context](https://studygolang.com/articles/23247?fr=sidebar)