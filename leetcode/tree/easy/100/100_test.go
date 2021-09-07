package _00

import "testing"
import . "github.com/ct-zh/goLearn/leetcode/basic/tree"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				p: CreateTreeByArr([]string{"1", "2", "3"}),
				q: CreateTreeByArr([]string{"1", "2", "3"}),
			},
			want: true,
		},

		{
			name: "test2",
			args: args{
				p: CreateTreeByArr([]string{"1", "2"}),
				q: CreateTreeByArr([]string{"1", "null", "2"}),
			},
			want: false,
		},

		{
			name: "test3",
			args: args{
				p: CreateTreeByArr([]string{"1", "2", "1"}),
				q: CreateTreeByArr([]string{"1", "1", "2"}),
			},
			want: false,
		},

		{
			name: "test4",
			args: args{
				p: CreateTreeByArr([]string{"10", "5", "15"}),
				q: CreateTreeByArr([]string{"10", "5", "null", "null", "15"}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
