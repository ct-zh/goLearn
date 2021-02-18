package quickSort

import (
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{arr: Helper.GenerateRandArr(1000, 0, 9999)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			if !Helper.CheckSort(tt.args.arr) {
				t.Errorf(" wrong sort: %+v", tt.args.arr)
			}
		})
	}
}

func TestQuickSort2Ways(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{arr: Helper.GenerateRandArr(1000, 0, 9999)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			if !Helper.CheckSort(tt.args.arr) {
				t.Errorf(" wrong sort: %+v", tt.args.arr)
			}
		})
	}
}

func TestQuickSort3Ways(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"t1", args{arr: Helper.GenerateRandArr(1000, 0, 9999)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr)
			if !Helper.CheckSort(tt.args.arr) {
				t.Errorf(" wrong sort: %+v", tt.args.arr)
			}
		})
	}
}
