package mergeSort

import (
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/01base/helper"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/01base/sort/insertionSort"
)

// 归并排序

// 自顶向下的归并排序
func MergeSort(arr []int) {
	_mergeSort(0, len(arr)-1, arr)
}

// // 递归使用归并排序， 对 Arr[l...r]的范围进行排序
func _mergeSort(l, r int, arr []int) {
	if r-l <= 15 { // 优化2 在数据量比较小的时候可以使用插排
		insertionSort.InsertionSort(arr[l : r+1])
		return
	}
	if l >= r {
		return
	}

	// 历史上曾发生过 r + l int溢出的bug, 所以取中间值都使用减法
	mid := (r-l)/2 + l
	_mergeSort(l, mid, arr)   // [l...mid]
	_mergeSort(mid+1, r, arr) // [mid+1..r]	两个都是闭区间

	// 优化1: 如果[start, mid] 区间排序后的最大值小于[mid+1, end]区间的最小值，说明不需要排序了
	if arr[mid] > arr[mid+1] {
		merge(l, mid, r, arr)
	}
}

// 自底向上的归并排序
// 也就是第一步就将数组分成最小块，然后依次进行归并操作
// 性能对比自顶向下的归并排序速度可能会稍慢一些 为什么?
// 但是可以看到代码里面没有用到数组的key值，也就是说这个算法可以用在链表上
func MergeSort1(arr []int) {
	for size := 1; size <= len(arr); size += size { // 模块大小 1 2 4 8 ... 直到和数组长度相等，即完成所有归并
		for i := 0; i+size < len(arr); i += size + size { // 每次都是两个模块做比较，所以自增 2size
			// 注意处理越界问题：
			// 对于一个模块来说,内部已经是排好序了,所以每次都是两个模块进行比较,我们需要保证这两个模块的边界
			// 第一个模块[i, i+size-1],需要保证: i+size 小于等于Count (在for循环的条件里保证了)
			// 第二个模块[i+size, i+size+size-1] ，需要保证 endKey 不能大于count-1
			// (因为第二个模块不需要是完整的长度为size的模块，允许长度不足size，所以不放在for循环的条件里限制死)
			merge(i, i+size-1,
				Helper.MinInt(i+size+size-1, len(arr)-1), // 保证 endKey 不能大于count-1
				arr)
		}
	}
}

// 将arr[start,mid]和arr[mid+1,end] 两部分进行归并
func merge(l, mid, r int, arr []int) {
	aux := make([]int, len(arr))
	for i := l; i <= r; i++ { // 将[l...r]拷贝到aux中, 因为是闭区间所以i<=r
		aux[i-l] = arr[i] // 注意aux是从0开始的，而此部分arr是从l开始的
	}

	cursorL, cursorR := l, mid+1 // 两个游标，分别代表[l...mid]区间与[mid+1...r]区间

	for j := l; j <= r; j++ { // 因为是闭区间所以j<=r
		if cursorL > mid { // 游标L越界
			arr[j] = aux[cursorR-l]
			cursorR++
		} else if cursorR > r { // 游标R越界
			arr[j] = aux[cursorL-l]
			cursorL++
		} else if aux[cursorL-l] < aux[cursorR-l] { // 游标L所指的值 < 游标R所指的值
			arr[j] = aux[cursorL-l]
			cursorL++
		} else { // 游标L所指的值 > 游标R所指的值
			arr[j] = aux[cursorR-l]
			cursorR++
		}
	}
}
