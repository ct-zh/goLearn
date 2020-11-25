package insertionSort

// 插入排序
// 将 n + 1 插入 [0, n] 的子序列中
// 最差O(n^2) 近乎有序的数组则是无限接近与On
type InsertionSort struct {
	Arr   map[int]int
	Count int
}

// 插入排序 版本1
// 这个版本的排序对比选择排序，因为会提前break掉，所以看起来性能要优于选择排序
// 实际上测试出来的这个插入排序效率要低于选择排序，因为插入排序在每次元素做比较时都会swap一次，产生了内存操作
func (i *InsertionSort) Do() {
	// 因为第一个元素arr[0]不需要做判断， 所以循环是从1开始的
	for m := 1; m < i.Count; m++ {
		// 寻找元素 arr[m] 合适的插入位置

		// 每次与前面一个数做比较，因为最后是与arr[0]做比较，所以这里 n > 0就可以了
		for n := m; n > 0; n-- {
			if i.Arr[n] < i.Arr[n-1] { // 每次与前一个元素比较
				i.Arr[n], i.Arr[n-1] = i.Arr[n-1], i.Arr[n]
			} else {
				break // 插入排序相对选择排序的优点：可能会提前结束
			}
		}

		// 可以简写为：
		//for n := m; n > 0 && i.Arr[n] < i.Arr[n-1]; n-- {
		//	i.swap(n, n-1)
		//}
	}
}

// 插入排序优化版，可以少做一次数据交换
// 优化后的插排比选择排要快很多， 尤其在数组中重复数据较多的情况下
func (i *InsertionSort) DoBetter() {
	for index := 1; index < i.Count; index++ {
		saver := i.Arr[index]
		n := index
		for ; n > 0 && saver < i.Arr[n-1]; n-- {
			i.Arr[n] = i.Arr[n-1]
		}
		i.Arr[n] = saver
	}
}
