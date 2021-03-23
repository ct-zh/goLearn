package quickSort

import (
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/sort/insertionSort"
	"math/rand"
	"time"
)

// 三路快排
// 双指针分别指向大于p与小于p的区间，过滤了元素等于p的情况
func QuickSort3Ways(arr []int) {
	quickSort3Ways(0, len(arr)-1, arr)
}

func quickSort3Ways(l, r int, arr []int) {
	if l >= r {
		return
	}
	if r-l <= 15 {
		insertionSort.InsertionSort(arr[l : r+1])
		return
	}

	lt, gt := partition3Ways(l, r, arr)
	quickSort3Ways(l, lt-1, arr) // 闭区间[l...p-1] p不需要参与排序
	quickSort3Ways(gt, r, arr)   // [p+1...r]
}

// partition 方法三：三路快排
// 从上面的快排到双路快排，可以发现存在大量元素等于p时会影响排序效率
// 将list划分成[l...lt-1] < pivot; [lt...gt-1] = pivot; [gt...r]>pivot;
// 因为是从l+1开始判断，最后将l与lt进行交换，所以小于pivot的部分是l到lt-1
func partition3Ways(l, r int, arr []int) (lt int, gt int) {
	// 随机化
	theRand := rand.New(rand.NewSource(time.Now().Unix()))
	randKey := theRand.Int()%(r-l) + l
	arr[l], arr[randKey] = arr[randKey], arr[l]
	pivot := arr[l]

	// partition操作：
	// 循环区间[l+1, r];lt、gt开始都在边界外，即l与r+1, 而i从第一个元素也就是l+1开始循环
	lt, gt, i := l, r+1, l+1
	for i < gt {
		if arr[i] == pivot {
			i++
		} else if arr[i] < pivot {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			i++
			lt++
		} else { // 因为gt指向的元素是未走入过循环的，所以i不需要++
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		}
	}
	arr[lt], arr[l] = arr[l], arr[lt]

	return
}
