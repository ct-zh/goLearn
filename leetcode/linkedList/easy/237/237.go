package _237

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
	fmt.Printf(" %d -> NULL", cur.Val)
}

// 题目中给了当前节点必不是末尾节点
// 因为不知道前面节点的地址，所以只能够修改当前节点的val值
// 流程：
// 1. 将当前节点的val修改成下一个节点的val
// 2. 删除下一个节点
func deleteNode(node *ListNode) {

	// 注意考虑边界问题
	if node == nil || node.Next == nil {
		return
	}

	node.Val, node.Next = node.Next.Val, node.Next.Next
}
