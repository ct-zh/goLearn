package _330

import "testing"

func Test_minPatches(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{
			nums: []int{1, 3},
			n:    6,
		}, 1},
		{"t2", args{
			nums: []int{1, 5, 10},
			n:    20,
		}, 2},
		{"t3", args{
			nums: []int{1, 2, 2},
			n:    5,
		}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPatches(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("minPatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
