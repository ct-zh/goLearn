package mst

import "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/graph/weightGraph"

// 最小索引堆
// 可以用于prim算法
type indexMinHeap struct {
	data     []weightGraph.Edge // 最小索引堆中的数据, 带权图
	indexes  []int              // 最小索引堆中的索引,从1开始, indexes[x] = i 表示索引i在x的位置
	reverse  []int              // 最小索引堆中的反向索引, reverse[i] = x 表示索引i在x的位置
	count    int                // 总数
	capacity int                // 容量
}

// 创建一个空的最小索引堆
// capacity: 最小索引堆的容量
func NewIndexMinHeap(capacity int) *indexMinHeap {
	i := &indexMinHeap{
		data:     make([]weightGraph.Edge, capacity+1),
		indexes:  make([]int, capacity+1),
		reverse:  make([]int, capacity+1),
		count:    0,
		capacity: capacity,
	}
	for m := 0; m <= capacity; m++ {
		i.reverse[m] = 0
	}

	return i
}

func (i *indexMinHeap) getVal(k int) float64 {
	return i.data[i.indexes[k]].Wt().(float64)
}

// 索引堆中, 数据之间的比较根据data的大小进行比较, 但实际操作的是索引
func (i *indexMinHeap) shiftUp(k int) {
	for {
		if k <= 1 || i.getVal(k/2) <= i.getVal(k) {
			break
		}
		i.indexes[k/2], i.indexes[k] = i.indexes[k], i.indexes[k/2]
		i.reverse[i.indexes[k/2]] = k / 2
		i.reverse[i.indexes[k]] = k
		k /= 2
	}
}

// 索引堆中, 数据之间的比较根据data的大小进行比较, 但实际操作的是索引
func (i *indexMinHeap) shiftDown(k int) {
	for {
		if 2*k > i.count {
			break
		}
		j := 2 * k
		if j+1 <= i.count && i.getVal(j) > i.getVal(j+1) {
			j += 1
		}

		if i.getVal(k) <= i.getVal(j) {
			break
		}
		i.indexes[k], i.indexes[j] = i.indexes[j], i.indexes[k]
		i.reverse[i.indexes[k]] = k
		i.reverse[i.indexes[j]] = j
		k = j
	}
}

func (i *indexMinHeap) Size() int {
	return i.count
}

func (i *indexMinHeap) IsEmpty() bool {
	return i.count == 0
}

// 向最小索引堆中插入一个新的元素, 新元素的索引为i, 元素为item
// 传入的i对用户而言,是从0索引的
func (i *indexMinHeap) Insert(index int, item weightGraph.Edge) {
	if i.count >= i.capacity {
		panic("堆满了")
	}
	if index+1 < 1 && index+1 > i.capacity {
		panic("index不能小于0，大于等于堆的容量")
	}

	index += 1
	i.data[index] = item
	i.indexes[i.count+1] = index // count+1:1  index:1 写入索引之后再count++
	i.reverse[index] = i.count + 1
	i.count++
	i.shiftUp(i.count)
}

// 从最小索引堆中取出堆顶元素, 即索引堆中所存储的最小数据
func (i *indexMinHeap) ExtractMin() weightGraph.Edge {
	if i.count <= 0 {
		panic("堆是空的")
	}
	ret := i.data[i.indexes[1]]
	i.indexes[1], i.indexes[i.count] = i.indexes[i.count], i.indexes[1]
	i.reverse[i.indexes[1]] = 1
	i.reverse[i.indexes[i.count]] = 0
	i.count--
	i.shiftDown(1)
	return ret
}

// 从最小索引堆中取出堆顶元素的索引
func (i *indexMinHeap) ExtractMinIndex() int {
	if i.count <= 0 {
		panic("堆是空的")
	}
	ret := i.indexes[1] - 1
	i.indexes[1], i.indexes[i.count] = i.indexes[i.count], i.indexes[1]
	i.reverse[i.indexes[1]] = 1
	i.reverse[i.indexes[i.count]] = 0
	i.count--
	i.shiftDown(1)
	return ret
}

// 获取最小索引堆中的堆顶元素
func (i *indexMinHeap) GetMin() weightGraph.Edge {
	if i.count <= 0 {
		panic("堆是空的")
	}
	return i.data[i.indexes[1]]
}

// 获取最小索引堆中的堆顶元素的索引
func (i *indexMinHeap) GetMinIndex() int {
	if i.count <= 0 {
		panic("堆是空的")
	}
	return i.indexes[1] - 1
}

// 看索引i所在的位置是否存在元素
func (i *indexMinHeap) Contain(index int) bool {
	return i.reverse[index+1] != 0
}

// 获取最小索引堆中索引为i的元素
func (i *indexMinHeap) GetItem(index int) weightGraph.Edge {
	if !i.Contain(index) {
		return weightGraph.Edge{}
	}
	return i.data[index+1]
}

// 将最小索引堆中索引为index的元素修改为item
func (i *indexMinHeap) Change(index int, item weightGraph.Edge) {
	if !i.Contain(index) {
		return
	}

	index += 1
	i.data[index] = item
	i.shiftUp(i.reverse[index])
	i.shiftDown(i.reverse[index])
}
