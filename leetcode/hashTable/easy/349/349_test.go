package _349

import (
	"testing"
)

func Test_intersection(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		result map[int]int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2},
			map[int]int{2: 1},
		},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4},
			map[int]int{9: 1, 4: 1}},
	}
	for key, tt := range tests {
		result := intersection2(tt.nums1, tt.nums2)
		//if reflect.DeepEqual(result, tt.result) {
		//	t.Errorf("[%d] 结果错误，result:%+v answer:%+v", key, result, tt.result)
		//}

		// 先判断是否存在重复数值
		if len(result) != len(tt.result) {
			t.Errorf("[%d] 结果错误，result:%+v answer:%+v", key, result, tt.result)
		}

		for _, i := range result {
			if tt.result[i] == 0 {
				t.Errorf("[%d] 结果错误，result:%+v answer:%+v", key, result, tt.result)
			}
		}
	}
}
