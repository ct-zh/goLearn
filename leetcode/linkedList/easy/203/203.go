package _203

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func create(arr []int) *ListNode {
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

func (l *ListNode) print() {
	cur := l
	for cur.Next != nil {
		fmt.Printf(" %d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf(" %d -> NULL\n", cur.Val)
}

// 思路 虚拟节点
func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}

	// 检查下一个节点是否等于val
	cur := dummyNode
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyNode.Next
}
