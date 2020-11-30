package _104

import (
	"testing"
)

func CreateTree(arr []int) *TreeNode {
	root := &TreeNode{
		Val:   arr[0],
		Left:  nil,
		Right: nil,
	}

	for i := 1; i < len(arr); i++ {

	}
	return root
}

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{root: CreateTree([]int{3, 9, 20, -1, -1, 15, 7})}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
