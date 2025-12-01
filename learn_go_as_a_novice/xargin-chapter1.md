# section01-06
## section01 调度

- 部分低延迟系统还是不能用go，得用c、c++、rust, 如:网关；
- PHP-FPM 是多进程模型，FPM 内单线程执行；PHP 底层是 C 语言实现，整套系统难精通。遇到 PHP 底层的 bug，束手无策；

### 个人发展

- 多写代码，多写代码，多写代码，思考别人的设计和自己的设计；
- 锻炼口才和演讲能力，逼着自己做技术分享；
- 读好书，建立完善的知识体系；
- 常关注信息源，如：Github Trending、reddit、medium、hacker news，morning paper(作者不干了)，acm.org，oreily，国外的领域相关大会(如 OSDI， SOSP，VLDB)论文，国际一流公司的技术博客，YouTube 上的国外工程师 演讲

- [工程师应该怎样去学习](https://xargin.com/how-to-learn/)

  

### 理解可执行文件

// todo 找时间学习汇编，补上相关内容

- 可执行文件在不同操作系统上规范不一样；
- 游戏：人力资源机器；
- 王爽老师的汇编教程；


### 调度循环
- q本地队列以外还做了一个runnext结构，他是一个指针，但是可以理解为一个高优先级的队列；（对于）

- 一个很有名的理论： 刚创建/调度的线程/协程有很大概率再次调度；

- 本地队列是一个无锁队列；

- 解决局部性的问题，生产/消费端尽可能减少加锁；

- 队列为什么是数组？因为数组是连续的，而且本地队列是固定了大小不需要考虑扩容的问题；

- p: proccess,虚拟处理器；但是按照曹大的思路，这更像一个token（请参考令牌桶算法），对于gomaxprocs其实就是这个桶的尺寸，m拥有了token才有运行权限；

- 本地队列循环，每调度60次，第61次会去全局队列拿一个g来执行；


#### 无聊的输出问题
下面两段代码分别输出什么? [运行一下](./codes/print/main.go)

```go
func A() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
	ch := make(chan int)
	<-ch
}

func B() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}

```

考察的是runnext, 目前版本的go都会先返回9, 然后返回0-8. 但是在go1.10版本以前, `time.Sleep`里面会开一个协程, 挤掉了9的runnext, 导致正常输出0-9.


## section02 语法背后
### 业务场景

- 场景一, 以下两段代码, 同事说代码一比代码二速度快. 如何打他的脸?
    ```go
    type person struct {
        age int
    }
    
    // 代码1
    var a = &person{111}
    println(a)
    
    // 代码2
    var b = person{111}
	var a = &b
	println(a)
    ```

- 场景二 // todo 补充

### 编译与反编译工具
在研究任何关于go*语言级别*的问题时, 没有任何一个问题是不能通过`go tool compile -S`解决的.如果不行的话就使用`go tool objdump`.

> compile是将go文件编译为`.o`文件, 再link后编译为可执行文件, objdump是将可执行文件反编译为汇编

#### go tool compile -S

关于`go tool compile -S`, 可以看[这个例子](./codes/hello/hello.go), 代码如下:

```go
// go tool compile -S ./hello.go|grep "hello.go:6"
func main() {
	var a = "hello world"
	var b = []byte(a) // runtime.stringtoslicebyte(SB)
	println(b)
}

// 输出
LEAQ    main..autotmp_2+40(SP), AX          // 将栈指针(SP)偏移40个字节处的值加载到寄存器AX中
LEAQ    go:string."hello world"(SB), BX     // 将字符串变量"hello world"的地址加载到寄存器BX中,这里的SB是静态基址寄存器，用于表示全局符号表。
MOVL    $11, CX  // 将11（字符串长度）加载到寄存器CX中
PCDATA  $1, $0   // PCDATA是 go gc辅助生成的调试内容, 可以无视. 实际汇编不会有这些内容
CALL    runtime.stringtoslicebyte(SB) // 调用这个函数, 将SB寄存器中的内容传入
MOVQ    AX, main.b.ptr+72(SP) // 将寄存器AX的值（即临时变量的地址）存储到栈指针（SP）偏移72个字节处的位置。这是字节切片的指针。
MOVQ    BX, main.b.len+24(SP) // 将寄存器BX的值（即字符串常量"hello world"的地址）存储到栈指针（SP）偏移24个字节处的位置。这是字节切片的长度。
MOVQ    CX, main.b.cap+32(SP) // 将寄存器CX的值（即字符串长度）存储到栈指针（SP）偏移32个字节处的位置。这是字节切片的容量。
```

- 在上面的代码中SB（Static Base）寄存器是一个专门用于访问静态数据的寄存器。
    - 它通常用于引用全局变量、全局函数和静态数据。
    - 在Go语言的编译器中，SB通常指向程序的静态基址，它可以帮助汇编器定位全局符号表中的变量或函数。

- SP（Stack Pointer）寄存器是用于指向栈顶的寄存器。
    - 栈是一种数据结构，遵循后进先出（LIFO）的原则。在函数调用和返回时，栈被用来保存函数的局部变量、参数以及函数调用的返回地址等信息。
    - 在汇编语言中，SP寄存器指向当前栈顶的位置。当需要在栈上分配内存或者从栈上弹出数据时，SP寄存器的值会相应地增加或减少。

- AX、BX、CX等是通用目的寄存器，它们用于存储临时数据和地址。

// todo 补充更详细的汇编内容

##### go tool compile -S 报错: could not import fmt

see https://github.com/golang/go/issues/56776

- 使用 `go build -gcflags=-S`, 加上grep 找到对应的汇编内容; `go build -gcflags=MYPKG=-S`

- 


##### 业务场景一实践操作步骤
同样我们也可以使用`go tool complie -S`来判断业务场景一中的代码

```
❯❯❯ cat -n 1.go
     1	package main
     2
     3	type person struct {
     4		age int
     5	}
     6
     7	func main() {
     8		var a = &person{111}
     9		println(a)
    10	}
```

我们要看第 8 行编译后变成啥了：

```
❯❯❯ go tool compile -S 1.go | grep "1.go:8"
	0x001d 00029 (1.go:8)	PCDATA	$0, $0
	0x001d 00029 (1.go:8)	PCDATA	$1, $0
	0x001d 00029 (1.go:8)	MOVQ	$0, ""..autotmp_2+8(SP)
	0x0026 00038 (1.go:8)	MOVQ	$111, ""..autotmp_2+8(SP)
```

两行的版本：

```
 ❯❯❯ cat -n 2.go
     1	package main
     2
     3	type person struct {
     4		age int
     5	}
     6
     7	func main() {
     8		var b = person{111}
     9		var a = &b
    10		println(a)
    11	}
    12
```

我们要看第 8 和第 9 行：

```
❯❯❯ go tool compile -S 2.go | grep -E '(2.go:8|2.go:9)'
	0x001d 00029 (2.go:8)	PCDATA	$0, $0
	0x001d 00029 (2.go:8)	PCDATA	$1, $0
	0x001d 00029 (2.go:8)	MOVQ	$0, "".b+8(SP)
	0x0026 00038 (2.go:8)	MOVQ	$111, "".b+8(SP)
```

可以看到，这里的一行版本的代码和两行版本的代码最终编译出的结果是完全一致的，没有任何区别。


#### go tool objdump

拿「使用`go tool objdump`来寻找make的实现」举例. 

首先参考spec.可以确定make一共有三种用法: 初始化slice、初始化map与初始化channel. 如果我们想研究完整的make实现, 需要探讨这三种用法.

> 代码的确定性与非确定性:我们引用一个函数,返回一个指针和一个error, 当error不为空时, 我们不能依赖指针的值, 这叫确定性. 如果我们没有判断error的值就直接使用指针的值, 则被称为非确定性. 非确定性经典错误如 undefined .

##### 观察 make
参考代码在[这里](./codes/make_and_new/make.go)

```go
❯❯❯ cat -n make.go
     1	package main
     2
     3	func main() {
     4		// make slice
     5		// 空间开的比较大，是为了让这个 slice 分配在堆上，栈上的 slice 结果不太一样
     6		var sl = make([]int, 100000)
     7		println(sl)
     8
     9		// make channel
    10		var ch = make(chan int, 5)
    11		println(ch)
    12
    13		// make map
    14		var m = make(map[int]int, 22)
    15		println(m)
    16	}
```

```shell
go build make.go && go tool objdump ./make | grep -E "make.go:6|make.go:10|make.go:14"
```

这里使用 go tool compile -S 也是可以的。

##### 观察 new(**输出内容难以读懂，不推荐**)

```
❯❯❯ cat -n new.go
     1	package main
     2
     3	type person struct{ age int }
     4
     5	func main() {
     6		var a = new(int)
     7		var b = new(person)
     8		var c = new(chan int)
     9		var d = new(map[int]int)
    10
    11		println(a, b, c, d)
    12	}
```

```
❯❯❯ go build -gcflags="-N -l" new.go && go tool objdump new | grep -E "new.go:6|new.go:7|new.go:8|new.go:9"
  new.go:6		0x1051591		48c744241000000000	MOVQ $0x0, 0x10(SP)
  new.go:6		0x105159a		488d442410		LEAQ 0x10(SP), AX
  new.go:6		0x105159f		4889442430		MOVQ AX, 0x30(SP)
  new.go:7		0x10515a4		48c744240800000000	MOVQ $0x0, 0x8(SP)
  new.go:7		0x10515ad		488d442408		LEAQ 0x8(SP), AX
  new.go:7		0x10515b2		4889442428		MOVQ AX, 0x28(SP)
  new.go:8		0x10515b7		48c744244000000000	MOVQ $0x0, 0x40(SP)
  new.go:8		0x10515c0		488d442440		LEAQ 0x40(SP), AX
  new.go:8		0x10515c5		4889442420		MOVQ AX, 0x20(SP)
  new.go:9		0x10515ca		48c744243800000000	MOVQ $0x0, 0x38(SP)
  new.go:9		0x10515d3		488d442438		LEAQ 0x38(SP), AX
  new.go:9		0x10515d8		4889442418		MOVQ AX, 0x18(SP)
```

被优化搞得面目全非了。

```
❯❯❯ go tool compile -N -S  new.go | grep -E "new.go:6|new.go:7|new.go:8|new.go:9"
	0x0021 00033 (new.go:6)	MOVQ	$0, ""..autotmp_4+16(SP)
	0x002a 00042 (new.go:6)	LEAQ	""..autotmp_4+16(SP), AX
	0x002f 00047 (new.go:6)	MOVQ	AX, "".a+48(SP)
	0x0034 00052 (new.go:7)	MOVQ	$0, ""..autotmp_5+8(SP)
	0x003d 00061 (new.go:7)	LEAQ	""..autotmp_5+8(SP), AX
	0x0042 00066 (new.go:7)	MOVQ	AX, "".b+40(SP)
	0x0047 00071 (new.go:8)	MOVQ	$0, ""..autotmp_6+64(SP)
	0x0050 00080 (new.go:8)	LEAQ	""..autotmp_6+64(SP), AX
	0x0055 00085 (new.go:8)	MOVQ	AX, "".c+32(SP)
	0x005a 00090 (new.go:9)	MOVQ	$0, ""..autotmp_7+56(SP)
	0x0063 00099 (new.go:9)	LEAQ	""..autotmp_7+56(SP), AX
	0x0068 00104 (new.go:9)	MOVQ	AX, "".d+24(SP)
```

可见就是生成了一些临时变量。


### 调试工具、语法分析实现

> - 阅读: [Table of Contents · Crafting Interpreters](https://craftinginterpreters.com/contents.html)
> - 阅读: [Writing An Interpreter In Go | Thorsten Ball (interpreterbook.com)](https://interpreterbook.com/)
> - 阅读: [Writing A Compiler In Go | Thorsten Ball (compilerbook.com)](https://compilerbook.com/)

调试工具建议参考我写的[dlv入门源码调试](https://github.com/ct-zh/goLearn/blob/master/doc/01basic/%E6%BA%90%E7%A0%81%E8%B0%83%E8%AF%95.md)

可以使用dlv来对一些经典问题进行调试,如

- go是如何启动协程的? 用dlv调试下面代码
    ```go
    func main(){
        go func() {
            fmt.Println(5)
        }()
        time.Sleep(time.Second*3)
    }
    ```

- 管道的收发: 
    ```go
    func main() {
        var a = make(chan int, 1)
        a <- 666
    
        x := <-a
        println(x)
    }
    ```

- 管道配合select:
    ```go
    func main() {
        var ch1 = make(chan int, 1)
        select {
            case <-ch1:
            default:
        }
    }
    ```

- 找出panic位置: 写已经关闭的管道、关闭nil管道、关闭已经关闭的管道


#### 语法分析 parser, 一些使用场景

- ast可以作为一个规则引擎去匹配相应的逻辑

    例如用户级别规则, 我们可以大概用这段代码来实现:
    ```go
    func getUserLevel(u *user.UserProfile) level {
        if u.invest > 1000 && u.post > 1000 {
            return highLevel // 用户充值大于1000并且发帖数大于1000, 返回 高级会员
        } else if u.invest > 100 && u.post > 500 {
            return middleLevel // 用户充值大于100并且发帖数大于500, 返回 中级会员
        } else if u.post > 100 {
            return juniorLevel // 用户发帖数大于100, 返回 初级会员
        } else {
            return newLevel // 返回新手
        }
    }
    ```

    实际开发中可能规则会随时发生变动, 所以我们可以将参数配置化, 例如`invest gt xx, post gt xx`之类的. 或者可以使用ast树,将一段配置翻译成代码.


- swagger使用ast树, 根据方法的注释生成对应的文档. 

    > [活文档](https://www.amazon.com/Living-Documentation-Cyrille-Martraire/dp/0134689321) 即文档的变化速度与软件设计和开发的速度相同.提倡在代码中编写注释, 通过程序将其扫描生成文档.

- 使用parser将thrift切换到grpc

- 使用parser进行sql审计的工作, 拿到即将执行的sql. 如google的 vitess.

- One sql to rule them all: elasticsearch、mysql、kv、others. 通过语法分析将sql语句转换为es、kv或者其他系统的语句.

- 如何在交付二进制的同时让我们的go模块具备一定的扩展性?
    - rpc: 定义扩展的api, 运行时调用一下; 有性能问题.
    - go-plugin: go内置的插件功能, 但是 不同版本编译不兼容.
    - repl: gopher-lua或者gojs之类的, 相当于内置一个go虚拟机来执行lua或者js代码. 不建议使用,需要编译原理,而且不能保证稳定性
    - wasm: 目前go版本不算很完善, 后续可以关注一下go-wasm


### 函数调用规约

函数调用规约(Calling Convention)是一个重要的基础概念，它规定了程序执行过程中函数的调用者（caller）和被调用者（callee）之间如何传递参数以及如何恢复栈平衡之间的约定。

函数调用规约主要包括以下几个方面：

- 参数传递方式：参数是通过寄存器传递还是通过栈传递？如果通过寄存器传递，哪些寄存器用于传递参数？
- 返回值传递方式：返回值是通过寄存器传递还是通过栈传递？如果通过寄存器传递，哪个寄存器用于传递返回值？
- 函数名修饰：编译器在编译阶段如何创建函数名唯一标识符，以便在链接阶段进行函数定位。
- 栈平衡：在函数调用前后，谁负责保存和恢复调用者的寄存器值？

函数调用规约对于程序的性能和可移植性至关重要。不同的函数调用规约可能会导致不同的性能表现，并且在不同的平台上可能无法兼容。例如go语言在调用cgo时就必须遵守c语言的函数调用规约.


#### Go语言的函数调用规约
Go语言的函数调用规约在 Go 1.17 之前和之后有所不同。

**Go 1.17 之前**

在 Go 1.17 之前，Go 语言的函数调用规约采用栈传递机制，即：

* 所有参数都通过栈传递。
* 返回值也通过栈传递。

具体来说，在调用函数之前，调用方会先为参数和返回值在栈上分配空间。然后，将参数值从调用者的寄存器复制到栈上。函数调用完成后，被调用方会将返回值复制到栈上。最后，调用方会从栈上获取返回值并销毁被调用方的栈空间。

这种调用规约的优点是简单易实现，并且在大多数情况下性能足够。但是，它也存在一些缺点：

* 对于较大的参数，栈传递可能会导致较高的性能开销。
* 栈传递可能会导致额外的寄存器保存和恢复操作。

**Go 1.17 及之后**

为了提高函数调用的性能，Go 1.17 引入了一套基于寄存器传参的调用规约。该调用规约的主要特点如下：

* 对于 9 个以内的参数，使用寄存器传递。
* 对于 9 个以上的参数，仍然使用栈传递。
* 函数调用的返回值，9 个以内通过寄存器传递回 caller，9 个以外在栈上传递。

具体来说，在调用函数之前，调用方会先将 9 个以内的参数值复制到特定的寄存器中。然后，进行函数调用。函数调用完成后，被调用方会将返回值复制到特定的寄存器中。最后，调用方会从特定的寄存器中获取返回值。

这种调用规约可以有效地减少栈操作，从而提高函数调用的性能。但是，它也增加了编译器的复杂性，并且可能导致与 Go 1.17 之前的版本不兼容。

下表总结了 Go 1.17 之前和之后函数调用规约的参数传递方式：

| 参数个数 | Go 1.17 之前 | Go 1.17 及之后 |
|---|---|---|
| ≤ 9 | 栈传递 | 寄存器传递 |
| ≥ 10 | 栈传递 | 栈传递 |

**函数名修饰**

在 Go 语言中，函数名使用 Go 包名、函数名和函数类型作为唯一标识符。例如，math.Abs 标识符表示 math 包中的 Abs 函数。

**栈平衡**

在 Go 语言中，***调用者***负责保存和恢复调用者的寄存器值。


## section03 内置数据结构

下面列出来一些Go内置数据结构, 源码分析时可以着重看一下.

- runtime: channel、timer、semaphore、map、iface、eface、slice、string

- netpoll: netpoll related...

- sync: mutex、cond、pool、once、map、waitgroup

- memory: allocation related..., gc related...

- container: heap, list, ring

- os: os related...

- context: context

具体解析channel、timer和map三个数据结构, 可以看一下下面三篇文章 todo

- channel 


## section04 系统调用 (没看懂, 再看一遍)

操作系统是资源的管理器, 对硬件层面的资源进行抽象; 
- 磁盘抽象 -> 文件夹
- cpu抽象 -> 时间片
- 内存抽象 -> 虚拟内存

cpu分级保护域 protect ring  // todo 完善

寄存器registers: cpu内部特殊存储单元; 可以记内存地址;

系统调用的调用规约, 参考 https://chromium.googlesource.com/chromiumos/docs/+/master/constants/syscalls.md#naming

linux内核系统调用代码参考搜索 https://codebrowser.dev/

* 常见系统调用 如
    - os.GetPid()  -> getpid()
    - println()  ->  write(2, "xxx", 3)
    - startm -> newm -> newosproc


strace(linux) dtruss(maxos)工具  查看二进制文件的一些系统调用, 一些行为.  // todo实验
    - `-f` 多线程查看?
    - `-c` 查看系统调用统计


系统调用分类: 
- (sys)阻塞系统调用与(sysnb)非阻塞系统调用.

// todo 答疑环节

系统调用更加深入了解可以参考 https://man7.org/

#### VDSO 优化概述
> from gemini

VDSO（**Virtual Dynamic Shared Object**），即**虚拟动态共享对象**，是一种 Linux 内核机制，用于减少用户态和内核态之间的切换次数，从而提高系统性能。

在传统的系统调用机制中，当用户程序需要调用系统服务时，需要将控制权切换到内核态，执行相应的内核函数。这种模式会导致大量的用户态和内核态切换，从而降低系统性能。

VDSO 优化通过将一些常用的系统函数映射到用户空间的内存中来解决这个问题。这样，用户程序可以直接调用这些函数，而无需切换到内核态。

VDSO 的优势在于：

* 减少了用户态和内核态之间的切换次数，从而提高了系统性能。
* 提高了系统安全性，因为用户程序无法直接访问内核空间。

VDSO 的劣势在于：

* 增加了一定的内存开销，因为需要为 VDSO 分配内存空间。
* 需要额外的代码来维护 VDSO，增加了系统的复杂性。

##### VDSO 优化工作原理

VDSO 优化工作原理如下：

1. 内核将一些常用的系统函数映射到用户空间的内存中，形成 VDSO。
2. 用户程序通过符号链接或直接映射的方式访问 VDSO。
3. 当用户程序需要调用系统服务时，首先会尝试在 VDSO 中找到相应的函数。
4. 如果在 VDSO 中找到了相应的函数，则直接调用该函数，而无需切换到内核态。
5. 如果在 VDSO 中没有找到相应的函数，则会像传统的系统调用机制那样，将控制权切换到内核态，执行相应的内核函数。

##### VDSO 优化效果

VDSO 优化可以显著提高系统性能。例如，在一些基准测试中，VDSO 优化可以使系统调用速度提高 20% 以上。

##### VDSO 优化适用场景

VDSO 优化适用于以下场景：

* 需要频繁进行系统调用的应用程序。
* 对系统性能要求较高的应用程序。

##### VDSO 优化注意事项

在进行 VDSO 优化时，需要注意以下几点：

* VDSO 的大小需要控制在合理范围内，以免造成内存开销过大。
* 需要定期维护 VDSO，以确保其与内核版本兼容。




## section05 内存管理与垃圾回收

## section06 并发编程

### 并发内置的数据结构

- `sync.Once` 保证只运行一次.  常用运用在Close函数中保证只Close一次;    // todo 完善源码解析
- `sync.Pool` 在两种场景使用: inuse_objects过多, 导致 gc mark消耗了大量cpu;或者rss占用过高.在生命周期开始时Get, 在请求结束时Put. // todo 完善源码解析
    - fasthttp中有大量应用
    - `sync.Pool` 有no copy特性, 不能进行拷贝
    - syncPool 源码解析

- `semaphore` 锁实现的基础. // todo 完善源码解析
- `sync.Mutex`  // todo 完善源码解析
- `sync.Map` 
- `sync.WaitGroup`

