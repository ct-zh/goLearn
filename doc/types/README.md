# go语言的类型
> 此文章受限于笔者的经验与理解,不能保证其准确性时效性.

## 基本类型
### 前言
在go1.4或之前,go还是基于c实现的.基本类型申明在[runtime.h](https://github.com/golang/go/blob/go1.4/src/runtime/runtime.h)上,网上很多文章都是基于这个分析的;

go1.5开始,golang实现了`自举`, 自己实现了编译器,所以从1.5开始golang的编译器与运行完全用go写了(带一点汇编),c语言不再参与实施;

> [types/type.go(1.14)](https://github.com/golang/go/tree/go1.14.15/src/go/types/type.go)

> [cmd/compile/internal/types/type.go(1.14)](https://github.com/golang/go/blob/go1.14.15/src/cmd/compile/internal/types/type.go)

// todo


### 基本类型的大小
基本类型有:整型、浮点型、字符串、bool;
派生类型有:指针、array、struct、channel、func、slice、map、interface;

基本类型的大小:
```go
// 基本类型大小(单位: 字节)
fmt.Println(unsafe.Sizeof(1))		// int整型 8 bytes 
// int8(byte)/int16/int32(rune)/int64大小分别为(1/2/4/8)字节
fmt.Println(unsafe.Sizeof(1.1111))	// 浮点型 8 bytes
fmt.Println(unsafe.Sizeof('a'))		// 字节 4 bytes
fmt.Println(unsafe.Sizeof("a"))		// 字符串 16 bytes
fmt.Println(unsafe.Sizeof(true))	// bool 1byte

// 派生类型
fmt.Println(unsafe.Sizeof(&a))			// 指针 8
fmt.Println(unsafe.Sizeof([4]int{1}))	// 数组是一段连续的内存空间,根据容量(cap)的大小决定
fmt.Println(unsafe.Sizeof(struct {}{}))	// 空struct不占空间, 非空struct按照struct的字段决定大小

a := make(chan int)
fmt.Println(unsafe.Sizeof(a))			// unbuffer channel: 8 bytes
b := make(chan int, 10)
fmt.Println(unsafe.Sizeof(b))			// buffer channel:  8 bytes
fmt.Println(unsafe.Sizeof(func() {}))	// func 8 bytes

fmt.Println(unsafe.Sizeof([]int{}))		// slice: 24bytes
fmt.Println(unsafe.Sizeof(make(map[string]string)))	// map: 8 bytes

var eface interface{}
fmt.Println(unsafe.Sizeof(eface))		// interface 16 bytes
```

## int
// todo


#### int到底是32位还是64位/取int的最大值
```go
tpm1 := uint(0)     // 0
// 对uint(0)取反,即可获得uint最大值
tpm2 := ^uint(0)    // 18446744073709551615, 可知是 2^64 即uint与int为64位
// 因为int的二进制第一位是符号位,所以对uint最大值右移一位,即 11...111 => 011...111 即int的最大值
tpm3 := ^uint(0) >> 1   // 9223372036854775807
```
见注释,对uint(0)取反即可得到当前最大的uint值,即可得uint的位数;

> [golang 里面为什么要设计 int 这样一个数据类型？](https://www.v2ex.com/t/744921)
>  int与机器的字长一致,可以保证最大的运行效率;在不关心数值范围的场景下使用int;而int32与int64一般用于编解码、底层硬件相关或者数值范围敏感的场景.


## string
string在`src/builtin/builtin`中定义为:
> string是8位字节的所有字符串的集合，通常但不一定代表UTF-8编码的文本。字符串可以为空，但不能为nil。字符串类型的值是不可变的。(这里是说string底层指向的字符串是不可更改的. 其实string的指针可以指向其他字符串内存空间.)

在`src/runtime/string.go`中定义为:
```go
type stringStruct struct {
	str unsafe.Pointer
	len int
}
```



#### string与[]byte的区别
string底层是一个指向字符串的地址与字符串的长度,而`[]byte`是一个切片slice,底层是由一个指向array的数组、长度len与数组容量cap这三个元素构成;

string底层指向的字符串是不可更改的,每次更改字符串就需要重新分配一次内存;而`[]byte`底层数组如果cap足够,更改是不需要重新分配内存的,只有当cap不够了才需要重新申请一个array



## interface
接口有两种底层结构,一种是空接口:eface;一种是有方法的接口:iface;

> eface源码见src/runtime/runtime2.go
eface的定义是:
```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```

> iface源码见src/runtime/runtime2.go
iface的定义是:
```go
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
```
eface有`_type`字段;而iface有内容更丰富的`itab`字段,其中就包含了`_type`字段;

注意,`[]interface{}`类型是数组,`*interface{}`类型是指针,这两个的类型都不是接口,没有接口相关特性,这很容易搞混(坑).见下面例子:

#### []interface
> 原文见: https://github.com/golang/go/wiki/InterfaceSlice
以下语句将会报错:
```go
dataSlice := []int{1, 2, 3, 4, 5}
var iFaceSlice []interface{}
iFaceSlice = dataSlice
```
interface不应该什么类型都能表示吗?为什么这里无法将dataSlice赋给interface切片呢?

因为`[]interface{}`的类型不是interface,而是slice;在slice中,每个interface类型占两个字,而int类型只占一个字,它们的底层结构是不相同的;

正确的写法应该是:
```go
// method 1, 直接赋给interface
var iFaceSlice interface{}
iFaceSlice = dataSlice

// method 2, for循环赋值
iFaces := make([]interface{}, len(dataSlice))
for k, v := range dataSlice {
    iFaces[k] = v
}
```

#### 复制接口内容
对于结构体,一般有两种初始化方法:`a := foo{}`和`b := &foo{}`

对于方法`a := foo{}`,如果我们想复制接口内容,直接赋值就可以了: `c = a`

但是对于`b := &foo{}`,由于是引用类型,复制的是指针地址,所以不能直接赋值;

```go
type User interface {
	Name() string
	SetName(name string)
}

type Admin struct {
	name string
}

func (a *Admin) Name() string {
	return a.name
}

func (a *Admin) SetName(name string) {
	a.name = name
}
```
见上面这个例子, 因为Admin结构体的方法需要`*Admin`才能调用,(*存在指针方法的结构体,只能初始化成引用类型*),所以会出现这种情况:
```go
var user1 User
user1 = &Admin{name:"user1"}
fmt.Printf("User1's name: %s\n", user1.Name())
// User1's name: user1

var user2 User
user2 = user1
user2.SetName("user2")
fmt.Printf("User1's name: %s\n", user1.Name())
// User1's name: user2
fmt.Printf("User2's name: %s\n", user2.Name())
// User2's name: user2
```

那么如何*值复制*user1呢?

##### 方法一,解引用
```go
var user3 User
// 先转换user1的类型为Admin
padmin := user1.(*Admin)
// 再取出user1底层的数据(解引用)
admin := *padmin
// 将其再赋给*Admin
user3 = &admin

user3.SetName("user3")
fmt.Printf("User3's name: %s\n", user3.Name())
// User3's name: user3
fmt.Printf("User1's name: %s\n", user1.Name())
// User1's name: user1
```

##### 方法二,reflect
```go
var user5 User

// 如果user1是指针类型
if reflect.TypeOf(user1).Kind() == reflect.Ptr {
	user5 = reflect.New(reflect.ValueOf(user1).Elem().Type()).Interface().(User)
} else {
	// 如果user1是值类型
	user5 = reflect.New(reflect.TypeOf(user1)).Elem().Interface().(User)
}
user5.SetName("uaaaa")
fmt.Printf("User5's name: %s\n", user5.Name())
// User5's name uaaaa
fmt.Printf("User1's name: %s\n", user1.Name())
// User1's name: user1
```



## nil

> nil is a predeclared identifier representing the zero value for a pointer, channel, func, interface, map, or slice type.
> 
> nil是一个预先声明的标识符，表示指针、通道、函数、接口、映射或切片类型的零值。

nil并不是go的关键字之一,甚至你可以设定一个变量名字就要`nil`(但是最好不要这样做);

### golang大坑: 两个为nil的变量未必相等
先看一段代码:
```go
type MagicError struct{}

func (MagicError) Error() string {
	return "[Magic]"
}
func Generate() *MagicError {
	return nil
}

func Test() error {
	return Generate()
}
func main() {
	fmt.Printf("%T %+v \n", Test(), Test())          // *main.MagicError nil
	fmt.Printf("%T %+v \n", Generate(), Generate())  // *main.MagicError nil
    fmt.Println(Generate() == nil)                   // true
	fmt.Println(Test() == Generate())                // true
	fmt.Println(Test() == nil)                       // false
}
```
main函数里,为什么`Generate()`与`Test()`返回的数据都是类型为`*main.MagicError`,值为`nil`, 为何Generate就能等于nil,而Test不等于nil呢?

> 上面这段代码有bug, Test函数应该这样写:
```go
func Test() error{
	err := Generate()
	if err != nil {
		return err
	}
	return nil
}
```

这是golang的一个经典问题了,想要搞清楚这个问题,必须先知道interface的实现,因为`Test()`方法返回的error是interface类型;我们知道,interface底层,无论是iface还是eface,都存在type与data,*类型为interface的变量判断是否等于nil,必须type和data都为nil*

上面代码中`Generate()`返回的类型是`*MagicError`的指针,指向的是nil,因此等于nil; 但是`Test()`返回的是类型为`*MagicError`的interface,因为type不等于nil,所以Test返回的值不等于nil;

这里提一种写法: `(*interface{})(nil)`, 意思是将nil转换成interface类型的指针;得到的结果仅仅是空接口类型指针指向无效的地址,这样写的作用是*强调val虽然是无效的数据,但是它是有类型`*interface{}`的*;

用这种写法很容易证明上面那个问题:
```go
a := (*interface{})(nil)
fmt.Println(a == nil)	// true

b := (*MagicError)(nil)
fmt.Println(b == nil)	// true

var c interface{}
c = b					// fmt.Println(c) => nil
fmt.Println(c == nil)	// false
```
当b转换为interface类型后,就不等于nil了


## struct
// todo


#### struct的内存对齐
> [内存对齐](https://zhuanlan.zhihu.com/p/53413177)
> [Go struct 内存对齐](https://geektutu.com/post/hpg-struct-alignment.html)

## func
// todo


## map
// todo https://www.cnblogs.com/-lee/p/12807063.html

1. map是hash table实现,无序;
2. map*不是线程安全的*; (golang 1.9 在sync包里实现了并发安全的map)
3. hash冲突常用*线性探测*或者*拉链法*

   开放定址（线性探测）和拉链的优缺点
    - 拉链法比线性探测处理简单
    - 线性探测查找是会被拉链法会更消耗时间
    - 线性探测会更加容易导致扩容，而拉链不会
    - 拉链存储了指针，所以空间上会比线性探测占用多一点
    - 拉链是动态申请存储空间的，所以更适合链长不确定的

### 源码分析
位置: https://github.com/golang/go/tree/go1.14.15/src/runtime/map.go#L149

map同样也是数组存储的的，每个数组下标处存储的是一个bucket,这个bucket的类型见下面代码，每个bucket中可以存储8个kv键值对，当每个bucket存储的kv对到达8个之后，会通过overflow指针指向一个新的bucket，从而形成一个链表,看bmap的结构，我想大家应该很纳闷，没看见kv的结构和overflow指针啊，事实上，这两个结构体并没有显示定义，是通过指针运算进行访问的。

```go
// A bucket for a Go map.
type bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
    // tophash通常包含此bucket中每个键的哈希值的顶部字节。如果tophash[0]<minTopHash，则tophash[0]是一个bucket撤离状态。
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt elems.
    // 接着是bucketCnt键，然后是bucketCnt元素。
    // 
	// NOTE: packing all the keys together and then all the elems together makes the
	// code a bit more complicated than alternating key/elem/key/elem/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
    // 注意：将所有keys打包在一起，然后将所有元素打包在一起，这使得代码比交替的key/elem/key/elem/要复杂一些。。。但它允许我们消除需要的填充，例如map[int64]int8。后跟溢出指针。
}
```
我们能得到bucket中存储的kv是这样的，tophash用来快速查找key值是否在该bucket中，而不同每次都通过真值进行比较；还有kv的存放，为什么不是k1v1，k2v2..... 而是k1k2...v1v2...，我们看上面的注释说的`map[int64]int8`,key是int64（8个字节），value是int8（一个字节），kv的长度不同，如果按照kv格式存放，则考虑内存对齐v也会占用int64，而按照后者存储时，8个v刚好占用一个int64,从这个就可以看出go的map设计之巧妙。

最后我们分析一下go的整体内存结构，阅读一下map存储的源码，如下图所示，当往map中存储一个kv对时，通过k获取hash值，hash值的低八位和bucket数组长度取余，定位到在数组中的那个下标，hash值的高八位存储在bucket中的tophash中，用来快速判断key是否存在，key和value的具体值则通过指针运算存储，当一个bucket满时，通过overfolw指针链接到下一个bucket。


## array
> cmd/compile/internal/types/type.go:Array
> go/types/type.go:Array
数组是一块固定大小的连续的内存空间。
```go
type Array struct {
	Elem  *Type //元素类型
	Bound int64 //元素的个数(长度)
}
```
> 如果数组使用`[...]int{1,2,3}`初始化,则会调用:`cmd/compile/internal/gc/typecheck.go:typecheckcomplit`来计算长度;

### 新建array
```go
// cmd/compile/internal/types/type.go:473
// NewArray returns a new fixed-length array Type.
func NewArray(elem *Type, bound int64) *Type {
	if bound < 0 {
		Fatalf("NewArray: invalid bound %v", bound)
	}
	t := New(TARRAY)
	t.Extra = &Array{Elem: elem, Bound: bound}
	t.SetNotInHeap(elem.NotInHeap())
	return t
}
```
可以看到在新建array时就会判断array是分配在堆上还是分配在栈上;


## slice
### 引用类型的坑
slice是引用类型, 初学者可能会碰到一个坑,例如:
```go
func foo(t []int) {
	t[0] = 99
}
a := []int{1}
foo(a)
fmt.Println(a)	// [99]
```
因为是引用类型,底层数组里0的位置修改成了99,所以`a[0]`就变成了99;

但是如果进行append操作:
```go
func foo(t []int) {
	t[0] = 99
	t = append(t, 100)
	t[0] = 101
}
a := []int{1}
foo(a)
fmt.Println(a)	// [99]
```
因为append触发了扩容操作, 因此foo函数对应的局部变量t底层的array已经变成了另外一个array;所以只有还未扩容前的改动生效了;

### 如果是N维slice呢?
例子1:
```go
i := make([][]int, 3)
fmt.Println(i)		// [[] [] []]
fmt.Println(i[0] == nil)	// true	
var i2 [][]int
fmt.Println(i2)		// []
```
1. 使用var申明的是nil切片;
2. N维数组使用make只会初始化最外面那一层,里面的slice仍然是nil;

例子2:
```go
func foo(i [][]int) {
	i[0] = append(i[0], 2)
	i = append(i, []int{10})
	i[0] = append(i[0], 3)
	fmt.Println("foo", i)	// foo [[0 2 3] [0] [0] [10]]
}
i := make([][]int, 3)
for key := range i {
	i[key] = make([]int, 1)
}
fmt.Println(i)	// [[0] [0] [0]]
foo(i)
fmt.Println(i)	// [[0, 2] [0] [0]]
```
1. 当执行foo这一行`i[0] = append(i[0], 2)`时: 如下图,i和`foo i`两个切片底层都是指向的array数组,而array数组里存的slice又保存了分别指向3个数组的指针;此时array1发生扩容,地址发生了变化,array第一个slice指向的地址发生变化,外部变量i同样发生变化;

2. 当执行`i = append(i, []int{10})`这一行时,foo i发生扩容操作,此时foo i与i不在指向同一个数组了,因此后面的改动也不会对i生效了;

<img src="./slice.png" width=200px>


### 源码分析
> 源码分析 [slice](./slice.md)



## map、array和slice的区别与注意点
// todo


## reference
- > [go源码](https://github.com/golang/go)
- > [深入解析 Go 中 Slice 底层实现](https://halfrost.com/go_slice/#toc-0)
- > [Go之读懂interface的底层设计](https://zhuanlan.zhihu.com/p/109964497)
- > [详解interface和nil](https://blog.csdn.net/kai_ding/article/details/41322473)

