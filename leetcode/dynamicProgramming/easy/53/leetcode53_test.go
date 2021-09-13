package leetcode53

import "testing"

type args struct {
	nums []int
}

type tests struct {
	name string
	args args
	want int
}

func getTests() []tests {
	return []tests{
		{
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
		{
			args: args{[]int{1}},
			want: 1,
		},
		{
			args: args{[]int{0}},
			want: 0,
		},
	}
}

func Test_maxSubArray(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubArray2(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
