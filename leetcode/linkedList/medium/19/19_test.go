package _19

import (
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{
			head: createLinkedList([]int{1, 2, 3, 4, 5}),
			n:    2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.head.print()
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			got.print()
		})
	}
}
