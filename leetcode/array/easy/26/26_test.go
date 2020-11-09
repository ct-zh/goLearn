package _6

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		nums   []int
		result int
		answer []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
	}
	for key, tt := range tests {
		result := removeDuplicates(tt.nums)
		if result != tt.result {
			t.Errorf("[%d] 结果错误，result:%+v answer:%+v", key, result, tt.result)
		}

		for i := 0; i < result; i++ {
			if tt.answer[i] != tt.nums[i] {
				t.Errorf("[%d] 结果数组错误，result:%+v answer:%+v", key, tt.nums, tt.answer)
			}
		}
	}
}
