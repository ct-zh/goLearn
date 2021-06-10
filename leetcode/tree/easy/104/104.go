package _104

import Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/basic/helper"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil { // 递归终止条件
		return 0
	}
	return Helper.MaxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
