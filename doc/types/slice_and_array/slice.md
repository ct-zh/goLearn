# slice，go里面的切片
## 数据结构
> slice源码: src/runtime/slice.go
```go
type slice struct {
	array unsafe.Pointer	// 指针
	len   int							// 当前切片长度
	cap   int							// 容量
}
```



### nil切片与空切片
```go
var a []int     // nil切片
b := []int{}    // 空切片
```
区别:
- nil切片slice.array值为nil,而空切片slice.array指向一个空数组;
- `a == nil`, `b != nil`

两者都可以通过`append`方法添加数据;


## slice扩容
> 扩容见src/runtime/slice.go:growslice方法

> 关于源码中出现的参数`raceenabled`与`msanenabled`
> - `raceenabled`参数代表是否启用数据竞争检测; 在`go build`或者`go run`中加入`-race`参数就代表该选项为`true`
> - `msanenabled`参数: go1.6新增的参数,类似上面的`-race`,这个参数为`-msan`,并且仅在 linux/amd64上可用;作用是将调用插入到C/C++内存清理程序;这对于测试包含可疑 C 或 C++ 代码的程序很有用。在使用新的指针规则测试 cgo 代码时，您可能想尝试一下.

```go
// growtslice处理附加期间的切片增长。
// et: 数据类型; old: 需要扩容的切片; cap: 目标申请的容量;
func growslice(et *_type, old slice, cap int) slice {
}
```

### 扩容策略:
> *注意：扩容针对的是slice的容量，而不是针对原来数组的长度。*

```go
// newcap 是最终扩容的大小
newcap := old.cap
doublecap := newcap + newcap
if cap > doublecap {        // 1. 如果扩容大小 大于 旧容量的两倍
    newcap = cap            // 最终扩容就是cap的大小
} else {
    if old.len < 1024 {         // 2. 如果旧长度小于1024
        newcap = doublecap      // 最终扩容为当前容量的两倍
    } else {
        for 0 < newcap && newcap < cap {    // 3. 从旧容量开始循环增加1/4, 直到大于等于cap
            newcap += newcap / 4
        }
        if newcap <= 0 {        // 4. 如果溢出, 则新容量直接就是cap
            newcap = cap
        }
    }
}
// ...
```
可归纳为:
- 1. 如果新容量大于2倍旧容量，最终容量就是新申请的容量;否则234;
- 2. 如果旧切片的长度小于1024，则最终容量就是旧容量的两倍;
- 3. 如果旧切片长度大于等于1024，则最终容量从旧容量开始循环增加自身的1/4,直到最终容量大于等于新申请的容量;
- 4. 如果最终容量计算值溢出，则最终容量就是新申请容量

### 内存对齐:
`growslice()`方法继续往下, 上面计算出了需要扩容的大小,在申请空间前进行内存对齐;根据`et.size`参数,当数组中元素所占的字节大小为1、8或者2的倍数时,对应相应的内存空间计算;

> uintptr是用来保存指针的类型; 整型,可以足够保存指针的值的范围,在32平台下为4字节,在64位平台下是8字节;

```go    
// lenmem: 旧slice长度
// newlenmem: 新容量
// capmem: 
var lenmem, newlenmem, capmem uintptr
switch {
case et.size == 1:
   
case et.size == sys.PtrSize:
    
case isPowerOfTwo(et.size):
    
default:
    lenmem = uintptr(old.len) * et.size
    newlenmem = uintptr(cap) * et.size
    capmem, overflow = math.MulUintptr(et.size, uintptr(newcap))
    capmem = roundupsize(capmem)
    newcap = int(capmem / et.size)
}

```


### 给底层数组分配空间
`growslice()`继续往下, 申请一个新的数组, 将旧数组按照大小移动到新开辟的内存中:

```go
var p unsafe.Pointer
if et.ptrdata == 0 {    
    // 根据cap的大小申请一个新空间
    p = mallocgc(capmem, nil, false)
    // append调用的growslice()方法将从old.len覆盖到cap
    // 只清除不会被覆盖的部分。
    memclrNoHeapPointers(add(p, newlenmem), capmem-newlenmem)
} else {
    p = mallocgc(capmem, et, true)
    // 如果长度大小大于0(也就是已经申请过空间的变量) 并且写屏障开启了
    if lenmem > 0 && writeBarrier.enabled {
        bulkBarrierPreWriteSrcOnly(uintptr(p), uintptr(old.array), lenmem)
    }
}
// 将旧数组移动到新开辟的内存中, 大小为lenmem
memmove(p, old.array, lenmem)

// 最后返回新的slice, 数组为p, 长度为old.len, 容量为新申请的容量newcap
return slice{p, old.len, newcap}
```


## slice拷贝
> 源码 src/runtime/slice.go:slicecopy 方法

```go
// 将fm拷贝到to, 并返回被复制的元素个数;
func slicecopy(to, fm slice, width uintptr) int {
	if fm.len == 0 || to.len == 0 {
		return 0
	}

	n := fm.len
	if to.len < n { // n为 {to和fm中长度短的那个} 的长度
		n = to.len
	}

	if width == 0 { // 直接返回短的那个的长度
		return n
	}

	// ... 
	size := uintptr(n) * width  // 最小的数据长度 * 单个数据大小 = 内存大小
	if size == 1 {
        // 如果只有一个元素, 指针直接转换
		*(*byte)(to.array) = *(*byte)(fm.array)
	} else {    // 把size个bytes从fm.array地址开始,拷贝到to.array的地址
		memmove(to.array, fm.array, size)
	}
	return n
}
```

## slice常见坑
### slice是引用类型而不是值类型
slice是引用类型, 初学者可能会碰到一个坑,例如:
```go
func foo(t []int) {
	t[0] = 99
}
a := []int{1}
foo(a) // a = [99]	
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

### N维slice触发扩容呢
1. N维slice使用make只会初始化最外面那一层,里面的slice仍然是nil;
2. N维slice底层仍然是数组，数组内容是slice；

### 切片长度的问题 

思考一下，为什么最后`a[5]`结果为3？

```go
a := make([]int, 4, 8)		// 0 0 0 0 
b := append(a, 4, 4, 4)		// 0 0 0 0 4 4 4
a = append(a, 3, 3, 3)		// 0 0 0 0 3 3 3
fmt.Println(a[5])	// 3
```

## 问题
```go
// 最后输出的结果是什么？ 为什么？
str := "Go语言"
runeStr := []rune("Go语言")
fmt.Println(len(str))       // ?
fmt.Println(len(runeStr))   // ?
fmt.Println(unsafe.Sizeof(str))     // ?
fmt.Println(unsafe.Sizeof(runeStr)) // ?
```

## reference

> [go slice源码](https://github.com/golang/go/blob/go1.14.15/src/runtime/slice.go)
> [深入解析 Go 中 Slice 底层实现](https://halfrost.com/go_slice/#toc-0)
> [slice切片源码解析](https://github.com/friendlyhank/toBeTopgopher/blob/master/golang/source/golang%E4%B9%8Bslice%E5%88%87%E7%89%87%E6%BA%90%E7%A0%81%E8%A7%A3%E6%9E%90.md)