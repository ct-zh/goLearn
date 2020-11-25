package _24

import (
	"fmt"
	"testing"
)

func createLinkedList(arr []int) *ListNode {
	l := len(arr)
	if l <= 0 {
		return nil
	}

	head := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	cur := head
	for i := 1; i < l; i++ {
		cur.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		cur = cur.Next
	}

	return head
}

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{head: createLinkedList([]int{1, 2, 3, 4})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("before")
			tt.args.head.print()
			fmt.Println()
			got := swapPairs(tt.args.head)
			fmt.Println("after")
			got.print()
			fmt.Println()
		})
	}
}
