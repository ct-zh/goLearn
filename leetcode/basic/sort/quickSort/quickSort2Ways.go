package quickSort

import (
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/sort/insertionSort"
	"math/rand"
	"time"
)

// 双路快排
// 双指针循环，效率比单路快排高，
func QuickSort2Ways(arr []int) {
	quickSort2Ways(0, len(arr)-1, arr)
}

func quickSort2Ways(l, r int, arr []int) {
	if l >= r {
		return
	}
	if r-l <= 15 {
		insertionSort.InsertionSort(arr[l : r+1])
		return
	}

	p := partition(l, r, arr)
	_quickSort(l, p-1, arr) // 闭区间[l...p-1] p不需要参与排序
	_quickSort(p+1, r, arr) // [p+1...r]
}

// 对arr[l...r]部分进行partition操作
// 返回p 使得arr[l...p-1] < arr[p];  arr[p+1...r] > arr[p]
func partition2Ways(l, r int, arr []int) (p int) {
	// 方法一： 一般默认将第一个值设置为pivot
	// 但是存在问题：如果第一个值本来就是这个列表的最值呢 => 举例：一个已经排好序的list
	//pivot := arr[l]

	// 方法二：随机化
	// 改成设置随机某个key为pivot
	// 这里直接交换随机某个值与第一个值，这样后面的代码与方法一保持一致
	theRand := rand.New(rand.NewSource(time.Now().Unix()))
	randKey := theRand.Int()%(r-l) + l
	arr[l], arr[randKey] = arr[randKey], arr[l]
	pivot := arr[l]

	// partition操作：
	// 使arr[l+1...i)<=pivot; arr(j...r]>=pivot;
	// 注意这个边界问题
	// http://coding.imooc.com/learn/questiondetail/4920.html
	i, j := l+1, r
	for {
		// 左游标i 寻找不符合条件 <= pivot 的值
		// 这里取等号的原因：是为了保证出现相同元素时，左右两个区间的平衡性
		// 举例：1 0 0 ... 0 0 0 0,如果区间允许等于pivot,第一个p的位置会在中间(因为相等也会产生位置交换)
		// 如果区间不允许等于pivot，第一个p的位置会在start的位置
		for i < r && arr[i] <= pivot {
			i++
		}
		// 右游标j 寻找不符合条件 >= pivot 的值
		for j > l+1 && arr[j] >= pivot {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	// 在循环过后，游标i处于第一个大于等于pivot的位置，
	// 游标j处于最后一个小于等于pivot的位置
	// 而l处于的区间是小于等于pivot的区间，因此应该与j交换位置
	arr[j], arr[l] = arr[l], arr[j]
	return p
}
