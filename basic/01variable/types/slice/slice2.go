package slice

// 性能陷阱，返回切片的局部不会释放之前切片底层数组的内存
// 如下， 如果传入的i是一个超大的切片，并且之后都不会使用了
// 使用copy可以让gc早日回收i的内存

func lastNumsBySlice(i []int) []int {
	return i[len(i)-2:]
}

func lastNumsByCopy(i []int) []int {
	tmp := make([]int, 2)
	copy(tmp, i[len(i)-2:])
	return tmp
}
