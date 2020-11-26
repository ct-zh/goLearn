package _350

import (
	"fmt"
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
	for _, tt := range tests {
		result := intersect(tt.nums1, tt.nums2)
		fmt.Println(result)
	}
}
