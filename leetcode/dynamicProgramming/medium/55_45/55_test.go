package leetcode55_45

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{[]int{2, 3, 1, 1, 4}}, true},
		{"t2", args{[]int{3, 2, 1, 0, 4}}, false},
		{"t3", args{[]int{0}}, true},
		{"t4", args{[]int{2, 0, 0}}, true},
		{"t5", args{[]int{3, 0, 8, 2, 0, 0, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
