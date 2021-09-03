package _76

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1. 笨方法，直接遍历转成切片，然后取中间节点；
// 2. 快慢指针

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for slow != nil && slow.Next != nil {
		fast = fast.Next
		slow = slow.Next.Next
	}
	return fast
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNodeByArg(a []int) *ListNode {
	head := &ListNode{}
	current := head
	for k, val := range a {
		current.Val = val
		if k < len(a)-1 {
			newNode := &ListNode{}
			current.Next = newNode
			current = newNode
		}
	}
	return head
}
