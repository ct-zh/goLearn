package _7

import "testing"

func Test_removeElement(t *testing.T) {
	tests := []struct {
		nums   []int
		val    int
		result int
		answer []int
	}{
		{[]int{2}, 3, 1, []int{2}},
		{[]int{2}, 2, 0, []int{}},
		{[]int{1, 2, 3}, 4, 3, []int{1, 2, 3}},
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
	}
	for key, tt := range tests {
		result := removeElement2(tt.nums, tt.val)
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
