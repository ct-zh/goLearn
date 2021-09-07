package _00

import . "github.com/ct-zh/goLearn/leetcode/basic/tree"

/**
 * Definition for a binary tree node.
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return _isSameTree(p, q)
}

func _isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}
	if _isSameTree(p.Left, q.Left) && _isSameTree(p.Right, q.Right) {
		return true
	}
	return false
}
