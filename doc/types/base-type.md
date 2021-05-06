# 基本类型
### 前言
在go1.4或之前,go还是基于c实现的.基本类型申明在[runtime.h](https://github.com/golang/go/blob/go1.4/src/runtime/runtime.h)上,网上很多文章都是基于这个分析的;

go1.5开始,golang实现了`自举`, 自己实现了编译器,所以从1.5开始golang的编译器与运行完全用go写了(带一点汇编),c语言不再参与实施;

> [types/type.go(1.14)](https://github.com/golang/go/tree/go1.14.15/src/go/types/type.go)
> 
> [cmd/compile/internal/types/type.go(1.14)](https://github.com/golang/go/blob/go1.14.15/src/cmd/compile/internal/types/type.go)

### 基本类型的大小
基本类型有:整型、浮点型、字符串、bool;
派生类型有:指针、array、struct、channel、func、slice、map、interface;

基本类型的大小:
```go
// 基本类型大小(单位: 字节)

// int8(byte)/int16/int32(rune)/int64大小分别为(1/2/4/8)字节
fmt.Println(unsafe.Sizeof(1))		// int整型 8 bytes 

fmt.Println(unsafe.Sizeof(1.1111))	// 浮点型 8 bytes
fmt.Println(unsafe.Sizeof('a'))		  // 字节 4 bytes
fmt.Println(unsafe.Sizeof("a"))		  // 字符串 16 bytes
fmt.Println(unsafe.Sizeof(true))	  // bool 1byte

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

### int到底是32位还是64位/取int的最大值
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


### string
string在`src/builtin/builtin`中定义为:
> string是8位字节的所有字符串的集合，通常但不一定代表UTF-8编码的文本。字符串可以为空，但不能为nil。字符串类型的值是不可变的。

在`src/runtime/string.go`中定义为:
```go
type stringStruct struct {
	str unsafe.Pointer	// 指向底层字符串的指针
	len int							// 字符串的长度
}
```

> 与切片的结构体相比，字符串只少了一个表示容量的 `Cap` 字段，而正是因为切片在 Go 语言的运行时表示与字符串高度相似，所以我们经常会说字符串是一个只读的切片类型。

#### string的性能问题
string因为只读的特性，在对string变量进行操作时每次都会触发内存分配操作。大性能消耗尤其出现在string与`[]byte`的转换场景，建议在开发中多多关注。
> [string性能影响的一个例子](https://github.com/ct-zh/goLearn/blob/master/doc/testing/pprof/README.md#%E4%BC%98%E5%8C%96%E5%AE%9E%E8%B7%B5)

#### string与[]byte的区别

string底层是一个指向字符串的地址与字符串的长度,而`[]byte`是一个切片slice,底层是由一个指向array的数组、长度len与数组容量cap这三个元素构成;

string底层指向的字符串是不可更改的,每次更改字符串就需要重新分配一次内存;而`[]byte`底层数组如果cap足够,更改是不需要重新分配内存的,只有当cap不够了才需要重新申请一个array

### nil
> nil是一个预先声明的标识符，表示指针、通道、函数、接口、映射或切片类型的零值。

nil并不是go的关键字之一,甚至你可以设定一个变量名字就要`nil`(但是最好不要这样做);

> [golang大坑: 两个为nil的变量未必相等](https://github.com/ct-zh/goLearn/blob/master/doc/types/nil/nil.go)

## reference
- > [go源码](https://github.com/golang/go)
- > [深入解析 Go 中 Slice 底层实现](https://halfrost.com/go_slice/#toc-0)
- > [Go之读懂interface的底层设计](https://zhuanlan.zhihu.com/p/109964497)
- > [详解interface和nil](https://blog.csdn.net/kai_ding/article/details/41322473)

