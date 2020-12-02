package selectionSort

// 选择排序 时间复杂度 O(n^2)
// 每次在子list中找出最小值，与子list的第一个值交换
func selectionSort(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		min := i                     // 当前子区间的第一个值i
		for j := i + 1; j < l; j++ { // 从i+1开始获取最小的值
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min] // 将最小值min与i做交换
	}
}
