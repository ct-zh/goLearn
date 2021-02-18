package main

import (
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/sort/mergeSort"
	"testing"
)

func TestSortFunc(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name     string
		args     args
		sortFunc func([]int)
	}{
		{"t1",
			args{arr: Helper.GenerateRandArr(1000, 0, 9999)},
			mergeSort.MergeSort},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sortFunc(tt.args.arr)
			if !Helper.CheckSort(tt.args.arr) {
				t.Errorf(" wrong sort: %+v", tt.args.arr)
			}
		})
	}
}
