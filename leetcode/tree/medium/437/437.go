package _437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 在以root为根即诶单的二叉树中寻找和为sum的路径
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	// 寻找包含root并且和为sum的路径
	res := findPath(root, sum)

	// 寻找不包含root并且和为sum的路径
	res += pathSum(root.Left, sum)
	res += pathSum(root.Right, sum)

	return res
}

// 在以root为根的节点中找包含node的路径和为sum
// 返回路径个数
func findPath(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	res := 0
	if node.Val == sum {
		res += 1
	} // 可能存在负数 需要继续运行

	res += findPath(node.Left, sum-node.Val)
	res += findPath(node.Right, sum-node.Val)

	return res
}
