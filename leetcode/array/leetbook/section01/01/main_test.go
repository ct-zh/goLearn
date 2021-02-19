package leetbookSection1

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name   string
		args   args
		result []int
	}{
		{
			name:   "1",
			args:   args{nums: []int{0, 1, 0, 3, 12}},
			result: []int{1, 3, 12, 0, 0},
		},
		{
			name:   "2",
			args:   args{nums: []int{1}},
			result: []int{1},
		},
		{
			name:   "3",
			args:   args{nums: []int{-959151711, 623836953, 209446690, -1950418142, 0 - 1626162038}},
			result: []int{-959151711, 623836953, 209446690, -1950418142, -1626162038, 0},
		},
	}
	for _, tt := range tests {
		moveZeroes(tt.args.nums)
		for key, value := range tt.args.nums {
			if value != tt.result[key] {
				t.Errorf("test error! result: %+v right answer: %+v", tt.args.nums, tt.result)
				t.FailNow()
			}
		}
	}
}
