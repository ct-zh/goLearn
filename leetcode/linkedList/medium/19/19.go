package _19

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func (l *ListNode) print() {
	cur := l
	for cur.Next != nil {
		fmt.Printf(" %d -> ", cur.Val)
		cur = cur.Next
	}
	fmt.Printf(" %d -> NULL\n", cur.Val)
}

// 思路： 双指针
// p1 指向虚拟头节点（头节点前的空节点），p2指向p1往后的n个节点,往后开始遍历，当p2=nil的时候p1就是待删除的节点
// 设置虚拟头节点的目的是当需要删除的节点是头节点时方便操作
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	p1, p2 := dummy, head
	count := 0
	for count < n { // 因为p2初始节点是head，相当于已经走了一格了,删除的其实是p1的下一个节点
		if p2.Next == nil && count < n {
			return nil
		}
		p2 = p2.Next
		count++
	}

	for p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p1.Next = p1.Next.Next

	return dummy.Next
}
