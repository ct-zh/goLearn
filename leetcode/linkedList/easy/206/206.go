package _206

// linked list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 见课程5-1
// 创建3个指针：
// pre：指向前一个元素，初始为null
// cur: 指向当前的元素,初始为head节点
// next: 指向下个元素，初始为head.Next节点
// 流程：
// 1. cur节点的next更新成pre （完成反转）
// 2. pre更新成cur,cur更新成next，next更新成next.Next (所有指针往前移动了一格)
// 3. 直到cur = nil
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next // next需要保证cur不为nil

		cur.Next = pre // 反转

		pre = cur
		cur = next
	}
	return pre
}

// 简单版本
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}
