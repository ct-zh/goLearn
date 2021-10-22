# array
> cmd/compile/internal/types/type.go:Array
> go/types/type.go:Array

数组是一块固定大小的连续的内存空间。大多数array操作在都会转换成直接读写内存，在中间代码生成期间，编译器还会插入运行时方法调用防止发生越界错误。
```go
type Array struct {
	Elem  *Type //元素类型
	Bound int64 //元素的个数(长度)
}
```

> 这片文章[go里面的数组](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array/)讲得很清楚了

## 很蠢的一个题目
```go
// 下面生成的数组内部数据是什么样？
arr := [...]int{1, 2, 5: 4, 6}
```

