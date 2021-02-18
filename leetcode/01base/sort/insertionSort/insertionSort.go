package insertionSort

// 插入排序
// 将 n + 1 插入 [0, n] 的子序列中
// 最差O(n^2) 近乎有序的数组则是无限接近与On

// 插入排序 版本1
// 这个版本的排序对比选择排序，因为会提前break掉，所以看起来性能要优于选择排序
// 实际上测试出来的这个插入排序效率要低于选择排序，因为插入排序在每次元素做比较时都会swap一次，产生了内存操作
func InsertionSort1(arr []int) {

	// 因为第一个元素arr[0]前面没有元素，不需要做判断，所以循环是从1开始的
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- { // 每次与前面一个数做比较；最后与arr[0]做比较，所以这里 n > 0
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break // 插入排序相对选择排序的优点：可能会提前结束
			}
			// 可以简写为：
			//for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			//	arr[j], arr[j-1] = arr[j-1], arr[j]
			//}
		}
	}
}

// 插入排序优化版，可以少做一次数据交换
// 优化后的插排比选择排要快很多，尤其在数组中重复数据较多的情况下
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		saver := arr[i]
		j := i
		for ; j > 0 && saver < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = saver
	}
}
