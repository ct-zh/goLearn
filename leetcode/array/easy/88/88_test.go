package _88

import (
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3,
			nums2: []int{2, 5, 6}, n: 3,
			result: []int{1, 2, 2, 3, 5, 6}},
	}
	for key, tt := range tests {
		merge2(tt.nums1, tt.m, tt.nums2, tt.n)
		for i := 0; i < len(tt.nums1); i++ {
			if tt.nums1[i] != tt.result[i] {
				t.Errorf("[%d]错误，result：%+v answer:%+v",
					key, tt.nums1, tt.result)
				break
			}
		}
	}
}
