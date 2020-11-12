package _67

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		nums   []int // 升序排列的有序数组
		target int
	}{
		{[]int{1, 2, 3}, 4},
		{[]int{1, 5, 9, 11}, 16},
	}
	for key, tt := range tests {
		result := twoSum2(tt.nums, tt.target)

		adder := 0
		for i := 0; i < len(result); i++ {
			adder += tt.nums[result[i]-1] // 返回的索引是从1开始的
		}
		if adder != tt.target {
			t.Errorf("[%d] 结果数组错误，result:%+v answer:%+v", key, result, tt.target)
		}
	}
}
