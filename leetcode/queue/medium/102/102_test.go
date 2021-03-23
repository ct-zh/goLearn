package _102

import (
	"reflect"
	"strconv"
	"testing"
)

func CreateByArr(arr []string) *TreeNode {
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

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes [][]int
	}{
		{"t1",
			args{root: CreateByArr([]string{
				"3", "9", "20", "null", "null", "15", "7",
			})},
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := levelOrder(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("levelOrder() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
