package _01

import (
	. "github.com/ct-zh/goLearn/leetcode/basic/tree"
)

// 与上一题一模一样

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return _isSymmetric(root.Left, root.Right)
}

func _isSymmetric(p, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	} else if p == nil && q == nil {
		return true
	} else if p.Val != q.Val {
		return false
	}
	if _isSymmetric(p.Left, q.Right) && _isSymmetric(p.Right, q.Left) {
		return true
	}
	return false
}
