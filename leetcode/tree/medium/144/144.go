package _144

import (
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/stack"
)

func preorderTraversal(root *TreeNode) []int {
	arr := []int{}
	if root != nil {
		_traversal(root, &arr)
	}
	return arr
}

func _traversal(root *TreeNode, arr *[]int) {
	if root != nil {
		*arr = append(*arr, root.Val)
		_traversal(root.Left, arr)
		_traversal(root.Right, arr)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type command struct {
	s    string
	node *TreeNode
}

func (c *command) cmd(s string, node *TreeNode) {
	c.s = s
	c.node = node
}

const (
	cmdP = "print"
	cmdG = "go"
)

// 运用栈模拟递归,完成先序遍历
func preorderTraversal1(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	stack := stack.NewStack(100)
	stack.Push(command{
		s:    cmdG,
		node: root,
	})

	for !stack.IsEmpty() {
		cmd := stack.Pop().(command)

		if cmd.s == cmdP {
			res = append(res, cmd.node.Val)
		} else {
			if cmd.s != cmdG {
				panic("error")
			}

			if cmd.node.Right != nil {
				stack.Push(command{
					s:    cmdG,
					node: cmd.node.Right,
				})
			}

			if cmd.node.Left != nil {
				stack.Push(command{
					s:    cmdG,
					node: cmd.node.Left,
				})
			}

			stack.Push(command{
				s:    cmdP,
				node: cmd.node,
			})
		}
	}

	return res
}
