package main

import "testing"

func Test_sortColors(t *testing.T) {
	tests := []struct {
		nums   []int
		result []int
	}{
		{nums: []int{2, 0, 2, 1, 1, 0}, result: []int{0, 0, 1, 1, 2, 2}},
	}
	for key, tt := range tests {
		sortColors2(tt.nums)
		for i := 0; i < len(tt.nums); i++ {
			if tt.nums[i] != tt.result[i] {
				t.Errorf("[%d]结果错误：result: %+v answer: %+v \n", key, tt.nums, tt.result)
			}
		}
	}
}
