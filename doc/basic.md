# basic
## 数据类型
### go基础的数据类型
1. bool值: true, false
2. 数字类型(整型与浮点型)
3. 字符串类型
4. 指针(ptr)
5. 数组
6. strcut
7. func
8. Channel
9. slice
10. map
11. interface

> 其中值类型有: bool、数字、字符串、struct、指针、数组;
>
> 引用类型有: channel、func、slice、map、interface


### 基础数据类型的打印-sprintf占位符
1. bool值: `%t`:输出字符串 true false
2. int: `%d`:十进制；`%b`:二进制；`%x`,`%X`:十六进制；`%c`:该值对应的unicode码；
3. float:  `%e`：科学计数法; `%f`: 带小数但是不带指数,`%2.2f`带整数位两位,小数点后2位;
4. string: `%s`,`%q`。 q会输出带双括号的字符串，如果是int型的变量，则会输出带单括号的"go语法字符字面值"
5. pointer: `%p`
6. 通用结构打印: `%v`,`%+v`,`%#v`, 输出类型:`%T`


### Comma-ok断言判断interface的类型
格式为`value, ok := element.(T)`,示例如下:
```go
type t struct {
    tt int
}
a := &t{tt: 10}
get := Saver(a)
if t, ok := get.(*t); ok {
    fmt.Println(t.tt)
} else {
    fmt.Println("is not type t")
}
```
也可以写成switch版本:
```go
switch value := element.(type) {
    case string:
        // todo
    case []byte:
        // todo
    default:
        fmt.Println(reflect.TypeOf(value))
}
```

### 数字类型
#### 整型
整型有: int8、int16、int32、int64以及加上u的无符号类型(uint8、uint16、uint32、uint64); 加上u之后因为少一位符号位,所以容量要大一倍.

还有用的比较多的整型:
1. byte相当于uint8, 通常指代单个字符
2. rune相当于int32, 通常用于utf8
3. uint/int 相当于int32/uint32或者int64/uint64

#### 整型的范围
1. byte: `[0...255]` **闭区间** (2^8)
2. int64: `[-2^63...2^63]`, uint64: `[0...2^64]`

#### 浮点型
1. float32、float64: 32位/64位的浮点型数
2. complex64、complex128: 32位/64位实数+虚数


### 字符串
#### 中文字符串处理

为什么类型是int32/rune
**unicode字符集**为每一个字符分配一个唯一的id,称为码位(code point), 对于上面的「你」字,码位是20320,转换为十六进制为`4f60`, 在unicode编码中为`\u4f60`,在utf-8编码中为`&#x4F60`;

1. 测试1
   中文字符串,转换成byte,一个中文字符占用了3byte. 单个item的数据类型是byte,也就是int8:
    ```go
    s := "yes你好你好!"
    for k, b := range []byte(s) { // utf-8编码
        fmt.Printf("%d:%X ", k, b)
    }
    // 0:79 1:65 2:73 3:E4 4:BD 5:A0 6:E5 7:A5 8:BD 9:E4 10:BD 11:A0 12:E5 13:A5 14:BD 15:21
    ```


2. 测试2
   如果直接range,循环里单个item的类型是int32,如果想打印出字符串,直接将item转换为string即可`fmt.Printf("%s ", string(value))`
    ```go
    for key, value := range s {
        fmt.Printf("%d:%X ", key, value)
    }
    // 0:79 1:65 2:73 3:4F60 6:597D 9:4F60 12:597D 15:21 
    ```

3. 测试3
   在测试2中key值不是连续的,如果想获取连续的key,应该将s转换成rune:
```go
for key, value := range []rune(s) {
	fmt.Printf("%d %c \n", key, value)
}
```

4. 使用utf-8包
    ```go
    bytes := []byte("你好你好")
    for len(bytes) > 0 {
        ch, size := utf8.DecodeRune(bytes)
        bytes = bytes[size:]
        fmt.Printf("%c ", ch)
    }// 你 好 你 好
    ```

## 流程控制
1. go的switch case不需要跟break来防止程序执行到下一个case中.但是也可以使用fallthrough执行下一个case
    ```go
    switch 1 {
        case 1:
            fallthrough
        case 2:
            fmt.Println("aaa")
    }
    ```
2. 使用label跳出多层嵌套循环(慎用label)
    ```go
    MyLabel:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > 3 && j > 5 {
				break MyLabel
			}
		}
	}
    ```

## 错误处理
1. 活用defer调用:
   在打开一个资源之后记得defer调用close, 防止函数提前返回,
   注意1: defer是个栈,后defer的先执行. 即使是panic也能触发defer.
   注意2: 参数在defer语句进行计算, 如下代码,i的打印会从0开始, 而不是打印100行100
    ```go
    for i:=0; i < 100;i++ {
        defer fmt.Println(i) 
    }
    ```

2. panic
> 内建函数panic停止当前Go程的正常执行。当函数F调用panic时，F的正常执行就会立刻停止。F中defer的所有函数先入后出执行后，F返回给其调用者G。G如同F一样行动，层层返回，直到该Go程中所有函数都按相反的顺序停止执行。之后，程序被终止，而错误情况会被报告，包括引发该恐慌的实参值，此终止序列称为恐慌过程。
- 停止当前函数执行
- 一直向上返回,执行每一层的defer
- 如果没有遇见recover, 则程序退出

3. recover
> 内建函数recover允许程序管理恐慌过程中的Go程。在defer的函数中，执行recover调用会取回传至panic调用的错误值，恢复正常执行，停止恐慌过程。若recover在defer的函数之外被调用，它将不会停止恐慌过程序列。在此情况下，或当该Go程不在恐慌过程中时，或提供给panic的实参为nil时，recover就会返回nil。
- 仅在defer中调用
- 获取panic的值
- 如果无法处理,可重新panic(判断panic的类型 `err, ok := r.(error)`)
```go
defer func() {
if r:=recover(); r != nil {
// panic r
}
}()
```

4. 定义UserError
   类似于继承exception
    ```go
    type UserError type {
        Err Error       // 原生error
        Message string  // 给用户的信息
        Info    string  // 记录log日志的信息
    }
    ```

## 数组、切片与集合
### 初始化与初始值
1. 数组的初始化, 三种方式
    ```go
    var a [3]int        // 初始值为0
    b := [2]int{1, 2}   // 方括号里必须是一个常数,不能是变量
    c := [...]int{0, 1, 2, 3}
    ```
2. 切片的初始化:
    ```go
    // 方法1
    var a []int     // 内容为[], 不能直接调用a[0],会报index out of range
    a = append(a, 2)    // [2]
    // 方法2
    b := []int{3}   // [3]
    // 方法3
    c := make([]int, 1) // [0]
    // 附加capacity
    d := make([]int, 1, 32)  // [0], 但是cap为32
    ```
3. 切片通过数组来初始化:
    ```go
    arr := [3]int{0, 1, 2}
    a := arr[1:2]   // [1]  这种切分都是左闭右开区间,即[1...2)
    ```
4. 集合map的初始化:
    ```go
    // 方法1, 创建nil map, 不能进行调用
    var m1 map[int]struct{}
    // 方法2 创建map[]
    m2 := map[int]struct{}{}
    // 方法3 创建map[]
    m3 := make(map[int]struct{})
    ```

5. 初始值
   int默认为0,bool默认为false, string默认为空字符串,byte默认为0,指针默认为nil,struct默认为空struct
```go
var t map[int]int
fmt.Printf("test: %+v \n", t[1])	// 0

var t2 map[int]bool
fmt.Printf("test: %+v \n", t2[1])	// false

var t3 map[int]string
fmt.Printf("test: %+v \n", t3[1])	// ""

var t4 map[int]byte
fmt.Printf("test: %+v \n", t4[1])	// 0  byte 是int8

var t5 map[int]*tData
fmt.Printf("test: %+v \n", t5[1])	// nil

var t55 map[int]tData
fmt.Printf("test: %+v \n", t55[1])	// tData{} ; type tData struct {}

var t6 map[int]struct{}
fmt.Printf("test: %+v \n", t6[1])	// {}
```

### slice的详细探讨
#### slice是array的一个view
slice本身没有数据,是对底层array的一个view. 因此如果作为参数传入函数中,在函数里修改了slice的值,函数外slice值也会跟着改变
#### slice的底层结构
如下, c在对b进行切片时,可以切到不在b中的数据
```go
a := [...]int{1, 2, 3, 4, 5, 6}
b := a[1:3] // [2,3] 
c := b[1:4] // [3, 4, 5]
```
由此我们可以探讨slice的实现方法: slice的组成有三个结构:指针ptr, 切片大小len和切片容量cap. 指针ptr指向切片在数组开头的位置,对于数组a:
```go
a := [6]int{0:0, 1:1, 2:2, 3:3, 4:4, 5:5}
```
切片b为`b := a[1:3]`, 那么切片b的ptr就指向数组a的`1`的位置,len为2,容量cap为5,这个可以打印出来:
```go
b := a[1:3] // [1,2], len(b)=2 cap(b)=5
fmt.Printf("b = %+v, len(b)= %d cap(b)= %d \n", b, len(b), cap(b))
c := b[1:4] // [2,3,4], len(b)=3 cap(b)=4
fmt.Printf("c = %+v, len(c)= %d cap(c)= %d \n", c, len(c), cap(c))
```

#### slice的append操作
使用上面的数组`a`与切片`b`、`c`,进行以下操作:
```go
b = append(b, 10)   // [1,2,10] len(b)=3 cap(b)=5
fmt.Printf("b = %+v, len(b)= %d cap(b)= %d \n", b, len(b), cap(b))
```
1. 已知数组a的cap为6, 切片b进行append操作后,b的长度加1,cap不变,原数组a变成`[0,1,2,10,4,5]`,切片b变成`[1, 2, 10]`,切片c变成`[2,10,4]`

```go
c = append(c, 11)   // [2,10,4,11], len(b)=4 cap(b)=4
fmt.Printf("c = %+v, len(c)= %d cap(c)= %d \n", c, len(c), cap(c))
```
2. 再对c进行append,c的长度加1,cap不变,原数组a变成`[0,1,2,10,4,11]`,切片c变成`[2,10,4,11]`. 此时切片c已经达到数组的末尾

```go
d := append(c, 12)  // [2,10,4,11,12], len(b)=5 cap(b)=8
fmt.Printf("d = %+v, len(d)= %d cap(d)= %d \n", d, len(d), cap(d))
fmt.Println("a=", a)    // [0, 1, 2, 10, 4, 11]
```
3. 再对c进行append变成d, d的长度为c+1, cap为**c的两倍**,原数组a不变,**切片c不变**

上面三个步骤分别对不同情况下的slice进行append, 可知: 如果进行append操作时`len<cap`, 则直接在原数组上进行更改,如果`len>=cap`,则先会对原数组进行拷贝生成一个新的数组,然后对这个数组进行扩容,容量为slice的len的两倍. 此时对于切片d来说,底层的数组已经和切片a切片b不是同一个了

另外,由于slice是值传递的关系,append必须接受返回值,即只能写`a = append(a, 10)`,不能写成`append(a, 10)`


#### slice的copy
`copy(new, old)`, 将old的内容copy到new上面, 因为是copy,底层数组也是复制了一份.
```go
s1 := []int{0, 1, 2, 3, 4}
s2 := make([]int, 6, 32)
copy(s2, s1)
s2[5] = 5       // s2: [0 1 2 3 4 5] s1: [0 1 2 3 4]
```

如果复制到某个切片上, 这个切片长度不够,则只会根据new的len复制对应的数据(和cap无关)
```go
s3 := make([]int, 2, 3)
copy(s3, s1)    // s3: [0 1] 
```

#### slice的delete
1. 删除头元素: `s2[1:]`, 删除尾元素:`s2[:len(s2)-1]`

2. 直接利用切片的功能:
    ```go
    s1 := []int{0, 1, 2, 3, 4}
    s2 := append(s1[:1], s1[2:]...)
    // s1 [0 2 3 4 4] len=5 cap=5 
    // s2 [0 2 3 4] len=4 cap=5 
    ```
    1. s1的数据发生了变化, s1和s2的cap没有变化;
    2. 这里有个合并两个slice的写法: `merge := append(s1, s2...)`


### map的详细探讨
1. map是基于哈希表实现的,内部是没有顺序的;
2. `map[int]struct{}`空struct占用的内存是最少;

#### 判断是否存在
map: `_, ok := m[1]`,判断ok的值;
#### make的问题
make的第二个参数对map无效
```go
m1 := make(map[int]int, 2)
fmt.Println(len(m1))        // 0 
a := make([]int, 3)
fmt.Println(len(a))         // 3
```

### 综合讨论问题
#### 调用问题
调用未声明的部分时:
```go
a := map[int]int{}
fmt.Println(a[2])   // 0
b := []int{}
fmt.Println(b[2])   // panic
c := [3]int{0, 1, 2}
fmt.Println(c[5])   // panic
```

## 文件操作
go语言对文件操作的包一般是`os`,`ioutil`,`bufio`

1. 打开文件: `os.Open()` 或者 `ioutil.ReadFile()`
   > os.Open拿到的是一个File struct, 而ReadFile是直接给的Byte数组。大文件不推荐使用ReadFile，而是使用bufio的形式

2. 获取文件内容：`bufio.NewScanner`
   > NewScanner需要传入一个 io.Reader,os.Open返回的File可以传进去，返回一个scanner，可以用for range遍历出来
    ```go
    for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
    ```

3. 写入文件：
    1. `io.WriteString`
    2. `ioutil.WriteFile`
    3. `File(Write,WriteString)`
    4. `bufio.NewWriter`

## 协程 coroutine
- 轻量级“线程”
- *非抢占式*多任务处理,由协程主动交出控制权
- 编译器/解释器/虚拟机层面的多任务
- 多个协程可能在一个或者多个线程上运行(由调度器决定)

> Subroutines are special case of more general program components, called coroutines. In contrast to the unsymmetric. 子程序是协程的一个特例 - Donnald Knuth "The Art of Computer Programming. Vol1"

### 线程与协程
线程是: main函数调用某个函数(这里称之为doWork),doWork执行完之后再把控制权交换给main函数,main函数继续运行

协程则在main函数与doWork之间有个双向的通道(*其实main函数也是一个goroutine*), 控制权也可以双向流通; 而main与doWork可能在一个线程内执行,也有可能在不同线程内执行.

### 其他语言的协程
- C++: Boost.Coroutine
- jave: 不支持
- python: 3.5前使用yield关键字,3.5之后async def 对协程原生支持


### goroutine
1. 调度器:负责调度协程
2. 任何函数加上 go 就能送给调度器运行
3. 不需要在定义时需要区分异步函数(python需要在def前声明async)
4. 调度器会在合适的点进行切换
5. 使用-race来检测数据访问的冲突

#### 匿名函数变量作用域的问题
如下例子, 两个匿名函数的i是不一样的,正确的用法是第二种,将i以值传递的方式传入匿名函数中;
```go
for i := 0; i < 10; i++ {
    go func() {
        fmt.Printf("i: %d address %p \n", i, &i)
    }()
    go func(i int) {
        fmt.Printf("i: %d address %p \n", i, &i)
    }(i)
}
```
上面第一种用法直接调用了函数外的变量,会有不可见的后果:
```go
var a [10]int
for i := 0; i < 10; i++ {
    go func() {
        a[i]++
        runtime.Gosched()
    }()
}
time.Sleep(time.Second)
fmt.Println(a)
```
i在循环结束时的值为10,但是`a[10]`会超出数组范围.可以使用-race命令来查看数据访问的冲突

#### gorutine是非抢占式的多任务处理
以下代码会死锁,因为对应协程执行到for循环之后永远不会交出进程的控制权.
```go
go func(i int) {
    for {
        a[i]++
    }
}(i)
```
或者主动交出控制权`runtime.Gosched`, 但是其实很少使用这种方式:
```go
go func(i int) {
    for {
        a[i]++
        runtime.Gosched()
    }
}(i)
```

#### goroutine可能交出控制权的点
1. I/O: print系列函数, 文件io
2. select
3. channel
4. 等待锁
5. 函数调用(不一定,看调度器)
6. `runtime.Goshed()`

*只是参考,不能保证切换,不能保证在其他地方不切换*

## 编译
### 编译其他平台的二进制文件
1. 通过以下linux命令,先了解目标平台的架构
    1. `uname -a`: 操作系统
    2. `uname -m`: 架构
    3. `arch`:
    4. `file /bin/cat`
    5. 获取系统类型： `lsb_release -d`

2. 增加编译参数:
    1. 编译Windows平台的64位可执行程序
    ```
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go
    ```
    2. 编译mac平台64位可执行程序
    ```
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build test.go
    ```
    3. 编译linux平台
    ```
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build test.go
    ```

   参数GOODS可以是: darwin,freebsd, linux, windows
   参数GOARCH参数可以是: 386,amd64,arm

3. build参数
- -o: 输出的二进制文件名称, 用来代替默认包名;
- -v: 编译时显示包名;
- -p n: 开启并发编译，默认情况下该值为CPU逻辑核数
- -a: 强制重新构建
- -n: 打印编译时会用到的所有命令，但不真正执行
- -x: 打印编译时会用到的所有命令
- -race: 开启竞态检测
- -gcflags: 


## 编码转换：
下载两个包:  编码转换`gopm get -g -v golang.org/x/text`；自动检测编码：`gopm get -g -v golang.org/x/net/html`
```go

// 自动检测编码
func determineEncoding(r io.Reader) encoding.Encoding {
    bytes, err := bufio.NewReader(r).Peek(1024)
    if err != {
        panic(err)
    }
    e, _, _ := charset.DetermineEncoding(bytes, "")
    return e
}

// gbk转utf8
utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
```

