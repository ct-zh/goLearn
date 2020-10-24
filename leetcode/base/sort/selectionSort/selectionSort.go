package selectionSort

import "fmt"

// 选择排序 时间复杂度 O(n^2)
type SelectionSort struct {
	Arr     map[int]int               // 需要排序的数组
	N       int                       // 数组元素总数
	Compare func(k1 int, k2 int) bool // 比较函数
}

// 进行选择排序
func (s *SelectionSort) Do() {
	for index := 0; index < s.N; index++ {
		minIndex := index // 注意这里的minIndex从index开始递增,而不是从0开始递增

		for i := index + 1; i < s.N; i++ {
			// 寻找 [i, n) 区间里的最小值
			if s.Arr[i] < s.Arr[minIndex] {
				minIndex = i
			}
		}

		// swap 交换两个值
		s.Arr[index], s.Arr[minIndex] = s.Arr[minIndex], s.Arr[index]
	}
}

// 打印数组列表
func (s *SelectionSort) P() {
	for k, v := range s.Arr {
		fmt.Printf("Key-Value: %d - %+v \n", k, v)
	}
	fmt.Println("============")
}
