package _235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：利用二叉搜索树的特性
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || q == nil || p == nil {
		return nil
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}
