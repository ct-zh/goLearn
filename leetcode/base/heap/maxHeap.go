package heap

import "fmt"

// 最大堆：指每个子节点都小于等于父节点, 并且每次获取数据只能获取到根节点（堆的性质）
// 这里堆对应的二叉树默认是从1开始计数（即数组的起始key是1）
type maxHeap struct {
	data  map[int]int // 一棵完全二叉树
	count int
}

func (m *maxHeap) Size() int {
	return m.count
}

func (m *maxHeap) IsEmpty() bool {
	return m.count == 0
}

// 插入
func (m *maxHeap) Insert(i int) {
	// 因为索引是从1开始的，所以这里应该是给key=count+1赋值
	m.data[m.count+1] = i

	m.count++
	m.shiftUp(m.count) // 因为上面已经count++ 所以这里shiftUp的值就是数组的最后一个值
}

// 推出最大值
func (m *maxHeap) ExtractMax() (int, error) {
	if m.count == 0 {
		return 0, fmt.Errorf("当前堆数据为空！")
	}
	ref := m.data[1]

	m.data[1], m.data[m.count] = m.data[m.count], m.data[1]
	m.count--
	m.shiftDown(1)

	return ref, nil
}

// 整理插入的数据，往上检索二叉树，保证符合最大堆的定义
func (m *maxHeap) shiftUp(k int) {
	for k > 1 && m.data[k] > m.data[k/2] {
		m.data[k], m.data[k/2] = m.data[k/2], m.data[k]
		k /= 2
	}
}

// 整理丢掉根节点后的数据，往下检索二叉树，保证符合最大堆的定义
func (m *maxHeap) shiftDown(k int) {
	for k < m.count/2 {
		j := 2 * k
		if j+1 <= m.count && m.data[j+1] > m.data[j] {
			j += 1
		}
		if m.data[k] > m.data[j] {
			break
		}
		// todo: 优化： 如插入排序一样，可以不用一次一次重复赋值，而是记一个额外的变量，只有最后一次再做交换
		m.data[k], m.data[j] = m.data[j], m.data[k]
		k = j
	}
}

func NewMaxHeap() *maxHeap {
	return &maxHeap{
		data:  map[int]int{0: 0}, // 从1开始索引的Map
		count: 0,
	}
}

// Heapify 过程
// arr 生成最大堆的数组
// n arr的大小
func NewMaxHeapByArr(arr []int, n int) *maxHeap {
	data := make(map[int]int)
	for i := 0; i < n; i++ {
		data[i+1] = arr[i] // 从1开始
	}

	heap := maxHeap{
		data:  data,
		count: len(data),
	}

	// heapify 流程
	// 从 n/2 开始循环这棵树的每一棵子树
	// 为什么从n/2开始？ n/2 是这棵树的最后一个 最小子树（只有一个父节点两个子节点） 的父节点，
	// heapify的流程是从最后一个 最小子树 开始，每个子树进行一次shiftDown操作
	// 为什么是shiftDown操作？ 因为要从最后一个子树开始筛选出最大值，慢慢将最大值交换到index=1的位置
	// 如果是shiftUp操作
	for i := n / 2; i > -1; i-- {
		heap.shiftDown(i)
	}

	return &heap
}

func HeapSort(arr []int, n int) {

	// 创建maxHeap的方法1。可以直接遍历执行Insert函数
	//heap := NewMaxHeap()
	//for _, v := range arr {
	//	heap.Insert(v)
	//}

	// 创建maxHeap的方法2。可以先构建一个二叉树然后重复执行shiftDown
	heap := NewMaxHeapByArr(arr, n)

	// 结论：
	// 将n个元素逐个插入到一个空堆中，算法复杂度是O(nlogn)
	// heapify 的过程（即方法2）算法复杂度是O9n0

	var sort []int
	for {
		item, err := heap.ExtractMax()
		if err != nil {
			break
		}
		sort = append(sort, item)
	}

}
