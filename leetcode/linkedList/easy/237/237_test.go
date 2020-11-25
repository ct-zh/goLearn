package _237

import (
	"fmt"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				// 这里其实并不是头节点为1的链表，而是给的一个指针片段，从1开始的；
				head: createLinkedList([]int{1, 2, 3, 4}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("before")
			tt.args.head.print()
			fmt.Println()
			deleteNode(tt.args.head)
			fmt.Println("after")
			tt.args.head.print()
			fmt.Println()
		})
	}
}
