package _89

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
	k    int
}

type tests struct {
	name string
	args args
	want []int
}

func getTests() []tests {
	return []tests{
		{"t1", args{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			// 7, 1, 2, 3, 4, 5, 6
			// 6, 7, 1, 2, 3, 4, 5
			// 5, 6, 7, 1, 2, 3, 4
			k: 3,
		}, []int{5, 6, 7, 1, 2, 3, 4}},
		{"t2", args{
			nums: []int{-1, -100, 3, 99},
			k:    2,
		}, []int{3, 99, -1, -100}},
		{"t3", args{
			nums: []int{-1},
			k:    2,
		}, []int{-1}},
		{"t4", args{
			nums: []int{1, 2},
			k:    3,
		}, []int{2, 1}},
		{"t5", args{
			nums: []int{-1, -100, 3, 99},
			k:    3,
		}, []int{-100, 3, 99, -1}},
		{"t6", args{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    4,
		}, []int{4, 5, 6, 7, 1, 2, 3}},
		{"t7", args{
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    2,
		}, []int{5, 6, 1, 2, 3, 4}},
	}
}

func Test_rotate(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Logf("[%s]error answer ! result is %v but the answer is %v", tt.name, tt.args.nums, tt.want)
			}
		})
	}
}

func Test_rotate2(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			rotate2(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Logf("[%s]error answer ! result is %v but the answer is %v", tt.name, tt.args.nums, tt.want)
			}
		})
	}
}

func Test_rotate3(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			rotate3(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Logf("[%s]error answer ! result is %v but the answer is %v", tt.name, tt.args.nums, tt.want)
			}
		})
	}
}

func Test_rotate4(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			rotate4(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Logf("[%s]error answer ! result is %v but the answer is %v", tt.name, tt.args.nums, tt.want)
			}
		})
	}
}
