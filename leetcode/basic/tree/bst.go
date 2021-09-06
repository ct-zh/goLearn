package tree

import (
	"fmt"
	"github.com/ct-zh/goLearn/leetcode/basic/queue"
)

// 二分搜索树的大部分方法都可以使用递归来实现
// 这里二分查找树的问题：非平衡二叉树

// 二分查找树
type BST struct {
	root  *bstNode // 根节点
	count int      // 节点个数
}

func (b *BST) Size() int {
	return b.count
}

func (b *BST) IsEmpty() bool {
	return b.count == 0
}

// 向二分搜索树中插入一个新的key-value数据对
func (b *BST) Insert(key int, value string) {
	b.root = b.insert(b.root, key, value)
}

// 查看key值是否存在
func (b *BST) Contain(key int) bool {
	return b.contain(b.root, key)
}

// 获取key的value; 如果不存在，则bool返回false
// 可以用 value, ok := b.Search(key); ok  的方法调用
func (b *BST) Search(key int) (string, bool) {
	return b.search(b.root, key)
}

// 先序遍历
func (b *BST) PreOrder() {
	b.preOrder(b.root)
}

// 中序遍历
func (b *BST) InOrder() {
	b.inOrder(b.root)
}

// 后续遍历
func (b *BST) PostOrder() {
	b.postOrder(b.root)
}

// 层次遍历，基于队列来实现
func (b *BST) LevelOrder() {
	if b.root == nil {
		return
	}

	q := &queue.Queue{}
	q.Push(b.root)
	for !q.IsEmpty() {
		node := q.Pop().(*bstNode)

		fmt.Println(node.key, " - ", node.value)

		if node.left != nil {
			q.Push(node.left)
		} else if node.right != nil {
			q.Push(node.right)
		}
	}
}

// 获取最小的key
func (b *BST) Minimum() int {
	if b.count == 0 {
		panic("没有数据")
	}
	n := b.minimum(b.root)
	return n.key
}

// 获取最大的key
func (b *BST) Maximum() int {
	if b.count == 0 {
		panic("没有数据")
	}
	n := b.maximum(b.root)
	return n.key
}

// 从二分搜索树中删除最小值所在节点 问题：如果根节点是最小值就无法删除
func (b *BST) RemoveMin() {
	if b.root != nil {
		b.removeMin(b.root)
	}
}

// 删除最大值所在的节点 问题：如果根节点是最大值就无法删除
func (b *BST) RemoveMax() {
	if b.root != nil {
		b.removeMax(b.root)
	}
}

func (b *BST) Remove(key int) {
	b.root = b.remove(b.root, key)
}
