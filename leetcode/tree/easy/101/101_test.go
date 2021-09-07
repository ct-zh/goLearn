package _01

import "testing"
import . "github.com/ct-zh/goLearn/leetcode/basic/tree"

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{root: CreateTreeByArr([]string{"1", "2", "2", "3", "4", "4", "3"})},
			want: true,
		},

		{
			name: "test2",
			args: args{root: CreateTreeByArr([]string{"1", "2", "2", "null", "3", "null", "3"})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
