package mergeSort

import (
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/sort/insertionSort"
	"math"
)

// 归并排序
type MergeSort struct {
	Arr map[int]int
	N   int
}

// 自顶向下的归并排序
func (m *MergeSort) Do() {
	m.mergeSort(0, m.N-1)
}

// 递归使用归并排序， 对 Arr[l...r]的范围进行排序
func (m *MergeSort) mergeSort(start int, end int) {
	// 优化2 在数据量比较小的时候可以使用插排 todo: 思考插排和归并排序的性能对比
	if end-start <= 15 {
		i := insertionSort.InsertionSort{
			Arr:   m.Arr,
			Count: len(m.Arr),
		}
		i.Do()
		m.Arr = i.Arr
		return
	}
	if start >= end { // 跳出递归
		return
	}

	//mid := (start + end) / 2 // 历史上曾发生过 start + end int溢出的bug
	mid := (end-start)/2 + start // 所以改成end - start 的方式取中值
	m.mergeSort(start, mid)
	m.mergeSort(mid+1, end) // 闭区间，所以是由mid+1开始

	// 优化1: 如果[start, mid] 区间排序后的最大值小于[mid+1, end]区间的最小值，说明不需要排序了
	// 举例近乎有序的数组,可以测试一下
	if m.Arr[mid] > m.Arr[mid+1] {
		m.merge(start, mid, end)
	}
}

// 将arr[start,mid]和arr[mid+1,end] 两部分进行归并
func (m *MergeSort) merge(start int, mid int, end int) {
	aux := make(map[int]int)
	for i := start; i <= end; i++ {
		aux[i-start] = m.Arr[i] // 因为aux是从0开始的，而Arr[i]是从l开始的，所以key值要减去l偏移量
	}

	cursorL := start   // [start,mid] 的游标
	cursorR := mid + 1 // [mid+1,end] 的游标

	// 要考虑到游标越界的情况
	for k := start; k <= end; k++ {
		if cursorL > mid {
			m.Arr[k] = aux[cursorR-start]
			cursorR++
		} else if cursorR > end {
			m.Arr[k] = aux[cursorL-start]
			cursorL++
		} else if aux[cursorL-start] < aux[cursorR-start] {
			m.Arr[k] = aux[cursorL-start]
			cursorL++
		} else {
			m.Arr[k] = aux[cursorR-start]
			cursorR++
		}
	}
}

// 自底向上的归并排序
// 性能对比自顶向下的归并排序速度可能会稍慢一些 todo:为什么？
// 但是可以看到代码里面没有用到数组的key值，也就是说这个算法可以用在链表上 todo: 实现一次
func (m *MergeSort) Do2() {
	for size := 1; size <= m.N; size += size { // 1 2 4 8 ...
		for i := 0; i+size < m.N; i += size + size { // 每次都是两个模块做比较，所以自增 2size
			// 注意处理越界问题：
			// 对于一个模块来说,内部已经是排好序了,所以每次都要保证是两个模块进行比较,即保证这两个模块的边界
			// 一个是第一个模块的endKey: i+size 需要小于等于Count  一个是第二个模块的endKey，不能大于count-1

			// 对  arr[i...i+size-1] 和 arr[i+size...i+2*size-1]进行归并
			m.merge(i, i+size-1, int(math.Min(float64(i+size+size-1), float64(m.N-1))))
		}
	}
}
