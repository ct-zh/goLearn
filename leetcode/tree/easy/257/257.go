package _257

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	var result []string
	left := binaryTreePaths(root.Left)
	for i := 0; i < len(left); i++ {
		result = append(result, strconv.Itoa(root.Val)+"->"+left[i])
	}

	right := binaryTreePaths(root.Right)
	for i := 0; i < len(right); i++ {
		result = append(result, strconv.Itoa(root.Val)+"->"+right[i])
	}

	return result
}
