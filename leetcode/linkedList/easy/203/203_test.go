package _203

import (
	"fmt"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{
			head: create([]int{1, 2, 6, 3, 4, 5, 6}),
			val:  6,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("before: \n")
			tt.args.head.print()
			got := removeElements(tt.args.head, tt.args.val)
			fmt.Printf("after: \n")
			got.print()
		})
	}
}
