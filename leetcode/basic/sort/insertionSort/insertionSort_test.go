package insertionSort

import (
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/01base/helper"
	"testing"
)

func TestInsertionSort(t *testing.T) {
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
			InsertionSort(tt.args.arr)
			if !Helper.CheckSort(tt.args.arr) {
				t.Errorf(" array: %+v\n", tt.args.arr)
			}
		})
	}
}

func TestInsertionSort1(t *testing.T) {
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
			InsertionSort1(tt.args.arr)
			if !Helper.CheckSort(tt.args.arr) {
				t.Errorf(" array: %+v\n", tt.args.arr)
			}
		})
	}
}
