package _76

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{NewListNodeByArg([]int{1, 2, 3, 4, 5})},
			3,
		},
		{
			"test2",
			args{NewListNodeByArg([]int{1, 2, 3, 4, 5, 6})},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); got.Val != tt.want {
				t.Errorf("middleNode() = %v, want %v", got.Val, tt.want)
			}
		})
	}
}

func Test_NewListNodeByArg(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	n := NewListNodeByArg(a)
	var m []int
	for n != nil {
		m = append(m, n.Val)
		n = n.Next
	}
	if !reflect.DeepEqual(a, m) {
		t.Fatalf("error: 原始数据: %+v 转换后的数据: %+v", a, m)
	}
}
