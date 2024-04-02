package generics

// 泛型的使用场景：当开发者编写了重复的代码，而这些代码唯一的不同是使用了不同类型。
// 最经典的例子：Add函数

func AddInt(a, b int) int {
	return a + b
}

func AddFloat(a, b float64) float64 {
	return a + b
}

func AddUint32(a, b uint32) uint32 {
	return a + b
}

// Add 泛型版本
// 函数测试 执行
// go test -v --run=TestAdd .
func Add[T int64 | int | float64 | uint32](a T, b T) T {
	return a + b
}

// Comparable
// go语言的泛型是基于合约的泛型。
// Go语言中的合约称为“类型约束”。类型约束可以用于定义类型参数的约束
// 声明方法如下
type Comparable interface {
	~int | ~int64 | ~float64 | ~string
}

// MyList 在该结构体中，类型T的真实类型为接口Comparable声明的类型
type MyList[T Comparable] struct {
	data []T
}

// Append
// 测试执行
// go test -v --run=TestMyList_Append .
func (m *MyList[T]) Append(item T) {
	m.data = append(m.data, item)
}
