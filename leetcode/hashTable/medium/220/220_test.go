package _220

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{
			nums: []int{1, 2, 3, 1},
			k:    3,
			t:    0,
		}, true},
		{"t2", args{
			nums: []int{1, 0, 1, 1},
			k:    1,
			t:    2,
		}, true},
		{"t3", args{
			nums: []int{1, 5, 9, 1, 5, 9},
			k:    2,
			t:    3,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
