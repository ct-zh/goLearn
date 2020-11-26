package tree

import (
	"fmt"
	Helper "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/helper"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/queue"
	"math"
	"strconv"
)

// 基于leetcode中大部分和树有关题目的结构

// 链表形式的树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将数组转换成一棵树, 这个数组包含空节点null
// 如 []string{
//		"3", "9", "20", "null", "null", "15", "7",
//	} 转换成：
// 		3
//    9   20
//       15 7
func CreateTreeByArr(arr []string) *TreeNode {
	if len(arr) <= 0 {
		return nil
	}

	i1, _ := strconv.Atoi(arr[0])
	root := TreeNode{
		Val: i1,
	}

	// 辅助列表list：将生成的节点指针存入列表，索引从1开始
	// 这样相当于一棵满二叉树，可以获取某个节点的左右子树或者根节点
	list := make([]*TreeNode, len(arr)+1)
	list[0] = nil
	list[1] = &root

	// 0 1 2

	// 类似于层次遍历
	for i := 1; i < len(arr); i++ {
		if arr[i] == "null" {
			// 当前节点设置为nil
			list[i+1] = nil
		} else {
			v, _ := strconv.Atoi(arr[i])
			list[i+1] = &TreeNode{
				Val: v,
			}
		}

		// 判断是左子树还是右子树
		if (i+1)%2 == 0 { // 左子树
			list[(i+1)/2].Left = list[i+1]
		} else { // 右子树
			list[(i+1)/2].Right = list[i+1]
		}
	}

	return &root
}

// 将树作为一个string数组输出,默认树是一棵满二叉树，缺失的部分使用"null"代替
func (t *TreeNode) ToArr() (res []string) {
	cur := t

	// 使用队列来进行遍历：
	q := queue.Queue{}
	q.Push(cur)

	for !q.IsEmpty() {
		pop := q.Pop()
		if pop == nil {
			res = append(res, "null")
		} else {
			cur = pop.(*TreeNode)
			if cur.Left != nil {
				q.Push(cur.Left)
			} else {
				q.Push(nil)
			}
			if cur.Right != nil {
				q.Push(cur.Right)
			} else {
				q.Push(nil)
			}
			res = append(res, strconv.Itoa(cur.Val))
		}
	}
	return
}

// 求树的最大深度
func (t *TreeNode) MaxDeep() int {
	deep := 1
	return max(t, deep)
}

// 递归求最大深度
func max(node *TreeNode, deep int) int {
	if node.Left == nil && node.Right == nil {
		return deep
	}

	leftDeep := 0
	rightDeep := 0
	if node.Left != nil {
		leftDeep = max(node.Left, deep+1)
	}
	if node.Right != nil {
		rightDeep = max(node.Right, deep+1)
	}
	return Helper.MaxInt(leftDeep, rightDeep)
}

// 打印一棵树
func (t *TreeNode) Print() {
	list := t.ToArr()
	for deep := 1; deep <= t.MaxDeep(); deep++ {
		floorStart := math.Pow(2.0, float64(deep-1))
		floorEnd := math.Pow(2.0, float64(deep)) - 1
		//fmt.Printf("deep: %d start: %.f end: %.f\n", deep, floorStart, floorEnd)
		for i := floorStart; i <= floorEnd; i++ {
			fmt.Printf("%s ", list[int(i)-1])
		}
		fmt.Println()
	}
}
