package _16

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{
			nums:   []int{-1, 2, 1, -4},
			target: 1,
		}, want: 2},
		{name: "test2", args: args{
			nums:   []int{1, 1, 1, 1},
			target: 0,
		}, want: 3},
		{name: "test3", args: args{
			nums:   []int{1, 1, -1, -1, 3},
			target: -1,
		}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
