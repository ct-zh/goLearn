package selectionSort

// 选择排序 时间复杂度 O(n^2)
// 每次在子list [i+1, len) 中找出小于i的最小值min，与i交换
func SelectionSort(arr []int) {
	l := len(arr)
	for i := 0; i < l; i++ {
		min := i // 当前子区间的第一个值i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] { // 找出[i+1, l)这个区间的最小值
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min] // 将最小值min与i做交换
	}
}
