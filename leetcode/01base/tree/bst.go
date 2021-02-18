package tree

import (
	"fmt"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/queue"
)

// 二分搜索树的大部分方法都可以使用递归来实现
// 这里二分查找树的问题：非平衡二叉树

// 二分查找树
type BST struct {
	root  *bstNode // 根节点
	count int      // 节点个数
}

// 二分搜索树中的节点为私有的结构体, 外界不需要了解二分搜索树节点的具体实现
// 使用链表实现
type bstNode struct {
	key   int      // 数据在二叉树里的值，用于在二叉树里排序
	value string   // 数据的内容，不一定是string类型：node是私有结构，不能返回node，所以用value字段来存储返回的内容
	left  *bstNode // 左孩子
	right *bstNode // 右孩子
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

// 向以node为根的二分搜索树中, 插入节点(key, value), 使用递归算法
// 返回插入新节点后的二分搜索树的根
func (b *BST) insert(node2 *bstNode, key int, value string) *bstNode {
	if node2 == nil {
		b.count++
		return &bstNode{key: key, value: value}
	}
	if key == node2.key { // 查找到 则直接更新value
		node2.value = value
	} else if key < node2.key { // 小于根节点的值，往左子树查找
		node2.left = b.insert(node2.left, key, value)
	} else { // key > node2.key	大于key值，则往右子树查找
		node2.right = b.insert(node2.right, key, value)
	}

	return node2
}

// 查看key值是否存在
func (b *BST) Contain(key int) bool {
	return b.contain(b.root, key)
}

// 查看以node为根的二分搜索树中是否包含了以键值为key的节点，使用递归算法
func (b *BST) contain(n *bstNode, key int) bool {
	if n == nil {
		return false
	}

	if key == n.key {
		return true
	} else if key < n.key {
		return b.contain(n.left, key)
	} else {
		return b.contain(n.right, key)
	}
}

// 获取key的value; 如果不存在，则bool返回false
// 可以用 value, ok := b.Search(key); ok  的方法调用
func (b *BST) Search(key int) (string, bool) {
	return b.search(b.root, key)
}

func (b *BST) search(n *bstNode, key int) (string, bool) {
	if n == nil {
		return "", false
	}

	if key == n.key {
		return n.value, true
	} else if key < n.key {
		return b.search(n.left, key)
	} else {
		return b.search(n.right, key)
	}
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

// 先序遍历
func (b *BST) preOrder(n *bstNode) {
	if n == nil {
		return
	}

	fmt.Println(n.key, " - ", n.value)
	b.preOrder(n.left)
	b.preOrder(n.right)
}

// 中序遍历
func (b *BST) inOrder(n *bstNode) {
	if n == nil {
		return
	}

	b.inOrder(n.left)
	fmt.Println(n.key, " - ", n.value)
	b.inOrder(n.right)
}

// 后续遍历
func (b *BST) postOrder(n *bstNode) {
	if n == nil {
		return
	}

	b.postOrder(n.left)
	b.postOrder(n.right)
	fmt.Println(n.key, " - ", n.value)
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

func (b *BST) minimum(n *bstNode) *bstNode {
	if n.left == nil {
		return n
	}
	return b.minimum(n.left)
}

// 获取最大的key
func (b *BST) Maximum() int {
	if b.count == 0 {
		panic("没有数据")
	}
	n := b.maximum(b.root)
	return n.key
}

func (b *BST) maximum(n *bstNode) *bstNode {
	if n.right == nil {
		return n
	}
	return b.maximum(n.right)
}

// 从二分搜索树中删除最小值所在节点 问题：如果根节点是最小值就无法删除
func (b *BST) RemoveMin() {
	if b.root != nil {
		b.removeMin(b.root)
	}
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (b *BST) removeMin(n *bstNode) *bstNode {
	// 左节点是空，说明该节点已经是最小值了
	// 将右子树提取出来直接接到上面一层的左子树
	if n.left == nil {
		b.count--
		return n.right
	}

	n.left = b.removeMin(n.left)
	return n
}

// 删除最大值所在的节点 问题：如果根节点是最大值就无法删除
func (b *BST) RemoveMax() {
	if b.root != nil {
		b.removeMax(b.root)
	}
}

func (b *BST) removeMax(n *bstNode) *bstNode {
	if n.right == nil {
		b.count--
		return n.left
	}
	n.right = b.removeMax(n.right)
	return n
}

func (b *BST) Remove(key int) {
	b.root = b.remove(b.root, key)
}

// 删除掉以node为根的二分搜索树中键值为key的节点, 递归算法
// 返回删除节点后新的二分搜索树的根
func (b *BST) remove(n *bstNode, key int) *bstNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = b.remove(n.left, key)
		return n
	} else if key > n.key {
		n.right = b.remove(n.right, key)
		return n
	} else { // 删除逻辑
		// 待删除节点左子树为空的情况
		if n.left == nil {
			b.count--
			return n.right
		}
		// 待删除节点右子树为空的情况
		if n.right == nil {
			b.count--
			return n.left
		}

		// 待删除节点左右子树均不为空的情况

		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := b.minimum(n.right)
		b.count++

		successor.right = b.removeMin(n.right)
		successor.left = n.left
		b.count--
		return successor
	}
}
