package quickSort

import (
	"math/rand"
	"time"
)

// 快速排序
// 最优情况O(nlogn)
// 近似有序的数组排序，即最差情况，退化为O(n^2)
type QuickSort struct {
	Arr map[int]int
	N   int
}

// 普通快排+随机化快速排序
func (q *QuickSort) Do() {
	q.Sort(0, q.N-1)
}

// 双路快排
func (q *QuickSort) Do2() {
	q.Sort2(0, q.N-1)
}

// 三路快排
func (q *QuickSort) Do3() {
	q.Sort3(0, q.N-1)
}

func (q *QuickSort) Sort(start int, end int) {
	// 优化1,数据量少时走插入排序
	//if end - start <= 15 {
	//	insertionSort
	//}
	if start >= end {
		return
	}

	p := q.partition(start, end)
	q.Sort(start, p-1)
	q.Sort(p+1, end)
}

func (q *QuickSort) Sort2(start int, end int) {
	if start >= end {
		return
	}

	p := q.partition2(start, end)
	q.Sort(start, p-1)
	q.Sort(p+1, end)
}

func (q *QuickSort) Sort3(start int, end int) {
	if start >= end {
		return
	}

	lt, gt := q.partition3(start, end)
	q.Sort(start, lt-1)
	q.Sort(gt, end)
}

// 对arr[l...r]部分进行partition操作
// 返回p 使得arr[l...p-1] < arr[p];  arr[p+1...r] > arr[p]
func (q *QuickSort) partition(start int, end int) int {
	// 方法一： 一般默认将第一个值设置为flag
	// 但是存在问题：如果第一个值本来就是这个列表的最值呢 => 举例：一个已经排好序的list
	//flag := q.Arr[start]

	// 方法二：改成设置随机某个key为flag
	// 这里直接交换随机某个值与第一个值，这样就不用修改后面的代码了
	theRand := rand.New(rand.NewSource(time.Now().Unix()))
	randKey := theRand.Int()%(end-start) + start
	q.Arr[start], q.Arr[randKey] = q.Arr[randKey], q.Arr[start]
	flag := q.Arr[start]

	// 使arr[start+1...p]<flag; arr[p+1...end]>=flag;
	// 注意这里的区间[p+1...end]是大于等于flag的
	p := start
	for i := start + 1; i <= end; i++ {
		if q.Arr[i] < flag {
			q.Arr[p+1], q.Arr[i] = q.Arr[i], q.Arr[p+1]
			p++
		}
	}

	// 最后再交换p的值与start的值
	q.Arr[start], q.Arr[p] = q.Arr[p], q.Arr[start]

	return p
}

// partition 方法二： 双路快排
// 返回p 使得 arr[start...p-1] <= arr[p]; arr[p+1...r] >= arr[p]
// 双路快排处理元素正好等于arr[p]时需要额外注意
func (q *QuickSort) partition2(start int, end int) int {
	// 取锚定点：某个随机值
	theRand := rand.New(rand.NewSource(time.Now().Unix()))
	randKey := theRand.Int()%(end-start) + start
	q.Arr[start], q.Arr[randKey] = q.Arr[randKey], q.Arr[start]
	pivot := q.Arr[start]

	// [start+1...i) <= pivot; pivot <= (j...end]
	// 注意这个边界问题
	// http://coding.imooc.com/learn/questiondetail/4920.html
	i := start + 1
	j := end
	for {
		for { // 左游标i 寻找不符合条件 <= pivot 的值
			// 这里取等号的原因：是为了保证出现相同元素时，左右两个区间的平衡性
			// 举例：1 0 0 ... 0 0 0 0,如果区间允许等于pivot,第一个p的位置会在中间(因为相等也会产生位置交换)
			// 如果区间不允许等于pivot，第一个p的位置会在start的位置
			if i > end || q.Arr[i] >= pivot {
				break
			}
			i++
		}
		for { // 右游标j 寻找不符合条件 >= pivot 的值
			if j < start+1 || q.Arr[j] <= pivot {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		q.Arr[i], q.Arr[j] = q.Arr[j], q.Arr[i]
		i++
		j--
	}
	// 在循环过后，游标i处于第一个大于等于pivot的位置，
	// 游标j处于最后一个小于等于pivot的位置
	// 而start处于的区间是小于等于pivot的区间，因此应该与j交换位置
	q.Arr[j], q.Arr[start] = q.Arr[start], q.Arr[j]
	return j
}

// partition 方法三：三路快排
// 从上面的快排到双路快排，可以发现存在大量元素等于p时会影响排序效率
// 将list划分成[start...lt-1] < pivot; [lt...gt-1] = pivot; [gt...end]>pivot;
// 因为是从start+1开始判断，最后将start与lt进行交换，所以小于pivot的部分是start到lt-1
func (q *QuickSort) partition3(start int, end int) (lt int, gt int) {
	// 随机化快排，找到锚定点pivot
	r := rand.New(rand.NewSource(time.Now().Unix()))
	randKey := r.Int()%(end-start) + start
	q.Arr[start], q.Arr[randKey] = q.Arr[randKey], q.Arr[start]
	pivot := q.Arr[start]

	lt = start
	gt = end + 1
	index := start + 1 // 从start+1开始判断
	for {
		if index >= gt { // 所有元素判断完毕
			break
		}
		if q.Arr[index] == pivot {
			index++
		} else if q.Arr[index] < pivot {
			q.Arr[index], q.Arr[lt+1] = q.Arr[lt+1], q.Arr[index]
			index++
			lt++
		} else { // q.Arr[index] > pivot
			q.Arr[index], q.Arr[gt-1] = q.Arr[gt-1], q.Arr[index]
			gt--
		}
	}

	q.Arr[lt], q.Arr[start] = q.Arr[start], q.Arr[lt]
	return
}
