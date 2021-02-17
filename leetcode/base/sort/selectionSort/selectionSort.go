package selectionSort

// 选择排序 时间复杂度 O(n^2)
// 每次在子list中找出最小值，与子list的第一个值交换
type SelectionSort struct {
	Arr   map[int]int // 需要排序的数组
	Count int         // 数组元素总数
}

// 进行选择排序
func (s *SelectionSort) Do() {
	for index := 0; index < s.Count; index++ {
		minIndex := index // 注意这里的minIndex从index开始递增,而不是从0开始递增

		for i := index + 1; i < s.Count; i++ {
			// 寻找 [i, Count) 区间里的最小值
			if s.Arr[i] < s.Arr[minIndex] {
				minIndex = i
			}
		}

		// swap 交换两个值
		s.Arr[index], s.Arr[minIndex] = s.Arr[minIndex], s.Arr[index]
	}
}