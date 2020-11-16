package _24

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	cur := l
	for cur.Next != nil {
		fmt.Printf(" %d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf(" %d -> NULL", cur.Val)
}

// 思路：虚拟头节点
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}

	prev := dummyNode
	for prev.Next != nil && prev.Next.Next != nil {
		node1 := prev.Next
		node2 := node1.Next
		next := node2.Next

		node1.Next, node2.Next, prev.Next = next, node1, node2

		prev = node1
	}

	return dummyNode.Next
}
