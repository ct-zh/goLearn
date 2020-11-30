package _18

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{"t1", args{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
		}, [][]int{
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1},
		}},
		{"t1", args{
			nums:   []int{0, 0, 0, 0},
			target: 0,
		}, [][]int{
			{0, 0, 0, 0},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("fourSum() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
