package _1maxSubArr

import "testing"

func getTests2() []tests {
	return []tests{
		{"t1", args{[]int{1, -2, 3, -2}}, 3},
		{"t2", args{[]int{5, -3, 5}}, 10},
		{"t3", args{[]int{3, -1, 2, -1}}, 4},
		{"t4", args{[]int{3, -2, 2, -3}}, 3},
		{"t5", args{[]int{-2, -3, -1}}, -1},
		{"t6", args{[]int{1}}, 1},
		{"t7", args{[]int{-2}}, -2},
		{"t8", args{[]int{3, 1, 3, 2, 6}}, 15},
	}
}

func Test_maxSubarraySumCircular(t *testing.T) {
	for _, tt := range getTests2() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySumCircular(tt.args.nums); got != tt.want {
				t.Errorf("maxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
