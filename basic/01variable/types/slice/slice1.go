package slice

// 三种复制slice方法的性能测试;
// 结果: 使用copy速度最快,foo2和foo3差不多
//

func foo1() []int {
	a := []int{1, 2, 3, 4, 5, 6}
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func foo2() []int {
	a := []int{1, 2, 3, 4, 5, 6}
	b := append([]int{}, a...)
	return b
}

func foo3() []int {
	a := []int{1, 2, 3, 4, 5, 6}
	b := append(a[:0:0], a...)
	return b
}
