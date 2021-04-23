# unsafe

>  瞎研究瞎写的，仅做记录之用

不要在生产环境使用unsafe包，unsafe包很多操作不可移植并且会打破Go的类型与内存安全限制。

## unsafe基本操作

`unsafe.Pointer`表示任意类型且可寻址的指针值；

包含四种核心操作：

- 任何类型的指针值都可以转换为 Pointer
- Pointer 可以转换为任何类型的指针值
- uintptr 可以转换为 Pointer
- Pointer 可以转换为 uintptr

三个函数

- `unsafe.Sizeof` 获取变量的内存大小
- `unsafe.Alignof`获取变量内存对齐方式
- `unsafe.Offsetof`获取变量的偏移量


### 内存分配规则

```go
// 结构体
// 分别打印每个字段的Sizeof、Alignof、Offsetof
type foo struct {	// Sizeof=32 Alignof=8
	a bool				// 	1		1		0
	b byte				//  1		1		1
	c []int				//  24	8		8
}
```

- struct是连续分配的内存；这里内存分配器给strut分配的单元大小是8字节;
- `[]int`占用24字节的原因是：指针8字节+Len/Cap都是int64，各8字节=24字节；
- foo的Alignof=8，代表foo内存分配单元是8字节;
- 所以这里a、b两个字段占了第一个单元，c占一个单元，一共32字节；
- 字段c的偏移量是8，也能证明上述观点;

如果改一下字段顺序:

```go
// 分别打印每个字段的Sizeof、Alignof、Offsetof
type foo struct {	// Sizeof=40 Alignof=8
	a bool					// 1  1  0
	c []int					// 24 8  8
	b byte					// 1  1  32
}
```

  同样8字节一个单元，a占1个、c占3个、d占1个，比上面多了8字节的空间;

大概总结一下struct的优化点：

- 将常用字段放在前面，减少内存偏移计算操作；
- 小字段放前面；

### 如何拿到切片底层数组

#### 前置知识

使用reflect与`unsafe.Pointer`组合；SliceHeader是reflect包里定义的一个结构体，用于反射切片的。

```go
a := make([]int, 5, 10)
d := (*reflect.SliceHeader)(unsafe.Pointer(&a))
```

使用这种写法： `(*type)(unsafe.Pointer(ptr))`可以拿到某些数据的底层数据；

#### 拿到切片的底层数组

```go
a := make([]int, 5, 10)	// 切片内容：0 0 0 0 0 
d := (*reflect.SliceHeader)(unsafe.Pointer(&a)) // d 结构为 Data,Len,Cap
d2 := (*[10]int)(unsafe.Pointer(d.Data))			// 取出数组的地址，赋给一个长度为10的数组， d2即为切片a底层的数组，修改d2的数据a也会跟着变
```

#### 进一步研究：子切片

对于上面切片a的子切片c：`c := a[3:6]`，与a的底层数组进行对比:

```go
d := (*reflect.SliceHeader)(unsafe.Pointer(&a))	// &{Data:824634966016 Len:5 Cap:10}
c2 := (*reflect.SliceHeader)(unsafe.Pointer(&c))// &{Data:824634966040 Len:3 Cap:7}
```

发现两个切片底层数组的地址不一样了：

- 计算`c2.Data - d.Data`，结果是24
- 已知a与c的差距为3个元素，这里可以证明`24/3=8`
- 而int64占64/8=8个字节
- 说明数组确实是连续的内存块，子切片的实现方式是：start：改变底层数组的偏移量；end：改变len与cap；

### 使用Offsetof获取struct的字段

```go
f := &foo{Name: "张三", Age:  18}	// 连续内存段struct
p := unsafe.Pointer(f)
name := (*string)(p)			// 因为name字段偏移量为0，直接将该段地址转成*string
fmt.Println(*name)				// 打印出：张三
age := (*int)(unsafe.Pointer(uintptr(p) + unsafe.Offsetof(f.Age))) // 加上age的偏移量
fmt.Println(*age)					// 打印出： 18
```



  