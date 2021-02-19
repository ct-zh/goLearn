package mergeSort

import (
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/01base/helper"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{Helper.GenerateRandArr(10, 0, 99)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.args.arr)
			if !Helper.CheckSort(tt.args.arr) {
				t.Errorf(" array: %+v\n", tt.args.arr)
			}
		})
	}
}

func TestMergeSort1(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{Helper.GenerateRandArr(10, 0, 99)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort1(tt.args.arr)
			if !Helper.CheckSort(tt.args.arr) {
				t.Errorf(" array: %+v\n", tt.args.arr)
			}
		})
	}
}
