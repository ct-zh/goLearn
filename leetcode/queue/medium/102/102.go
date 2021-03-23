package _102

import "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/queue"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	t     *TreeNode // 节点
	level int       // 这个节点在第几层
}

// 使用队列对二叉树进行层序遍历
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	q := queue.Queue{}
	q.Push(pair{
		t:     root,
		level: 0,
	})

	for !q.IsEmpty() {
		// 先获取队列当前最顶层的节点
		pop := q.Pop()
		node := pop.(pair).t
		level := pop.(pair).level

		// 说明该节点在一个新的层中
		// level = 0 时， len(res) = 0， 说明需要新增一行
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)

		if node.Left != nil {
			q.Push(pair{
				t:     node.Left,
				level: level + 1,
			})
		}
		if node.Right != nil {
			q.Push(pair{
				t:     node.Right,
				level: level + 1,
			})
		}
	}

	return
}
