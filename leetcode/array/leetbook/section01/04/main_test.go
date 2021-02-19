package _4

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{nums: []int{1, 1, 1, 2, 2, 5}}, 5},
		{"test2", args{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, 7},
		{"test3", args{nums: []int{1, 1, 2, 2, 5}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v. slice is %+v", got, tt.want, tt.args.nums)
			}
		})
	}
}
