package quickSort

// 快速排序：对数组做partition分割，再分别对两个子数组做partition分割，如此循环
// 最优情况O(nlogn)
// 近似有序的数组排序，即最差情况，退化为O(n^2)

import (
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/01base/sort/insertionSort"
	"math/rand"
	"time"
)

// 普通快排+随机化
func QuickSort(arr []int) {
	_quickSort(0, len(arr)-1, arr)
}

func _quickSort(l, r int, arr []int) {
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
func partition(l, r int, arr []int) (p int) {
	// 方法一： 一般默认将第一个值设置为flag
	// 但是存在问题：如果第一个值本来就是这个列表的最值呢 => 举例：一个已经排好序的list
	//flag := arr[l]

	// 方法二：随机化
	// 改成设置随机某个key为flag
	// 这里直接交换随机某个值与第一个值，这样后面的代码与方法一保持一致
	theRand := rand.New(rand.NewSource(time.Now().Unix()))
	randKey := theRand.Int()%(r-l) + l
	arr[l], arr[randKey] = arr[randKey], arr[l]
	flag := arr[l]

	// partition操作：
	// 使arr[l+1...p]<flag; arr[p+1...r]>=flag;
	// 注意这里的区间[p+1...r]是大于等于flag的
	p = l                         // p指向小于flag的最后一个值，从l开始
	for i := l + 1; i <= r; i++ { // l是用来对flag， 所以从l+1开始; 因为是[l...r]闭区间，所以结束位置在i<=r;
		if arr[i] < flag {
			arr[p+1], arr[i] = arr[i], arr[p+1] // 将小于flag的值与p+1进行交换
			p++
		}
	}
	arr[l], arr[p] = arr[p], arr[l]

	return p
}
