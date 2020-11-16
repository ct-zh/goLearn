package linkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 根据数组创建一个链表
// arr: 需要转换成链表的数组
// n： 数组的长度
func CreateLinkedList(arr []int, n int) *ListNode {
	if n == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}

	curNode := head
	for i := 1; i < n; i++ {
		curNode.Next = &ListNode{Val: arr[i]}
		curNode = curNode.Next
	}

	return head
}

func (l *ListNode) Print() {
	curNode := l
	for curNode != nil {
		fmt.Printf("%d -> ", curNode.Val)
		curNode = curNode.Next
	}
	fmt.Printf("NULL\n")
}

// 删除元素
func remove(head *ListNode, val int) *ListNode {
	// 如果删除的是头节点
	for head != nil && head.Val == val {
		delNode := head
		head = delNode.Next
		delNode.Next = nil
	}
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			delNode := cur.Next
			cur.Next = delNode.Next
			delNode.Next = nil
		} else {
			cur = cur.Next
		}
	}

	return head
}

// 使用虚拟节点dummy来减少判断
func removeBetter(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}

	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			delNode := cur.Next
			cur.Next = delNode.Next
			delNode.Next = nil
		} else {
			cur = cur.Next
		}
	}

	retNode := dummyHead.Next
	dummyHead.Next = nil // 删除虚拟节点dummy

	return retNode
}
