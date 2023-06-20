package tree

// B+ 树和 B 树有什么不同：
// 1. B+树非叶子节点上是不存储数据的，仅存储键值，而B树节点中不仅存储键值，也会存储数据。
// 数据库中页的大小是固定的，InnoDB 中页的默认大小是 16KB。
// 非叶子结点不存储数据 = 存储更多的key = 树的阶数变大、层数变小(更矮更胖) = IO次数减少 = 数据查询的效率更快。
// B+树的分叉数等于键值的数量，如果一个节点存储1000个键值，那么3层 B+ 树可以存储 1000×1000×1000=10亿个数据,
// 只需要2次磁盘IO(根结点二分查找到第二层的位置，一次io；第二层结点继续二分查找到第三层叶子结点的位置，一次io，取出数据)
//
// 2. B+树 叶子节点数据是按照顺序排列的,范围查找，排序查找，分组查找以及去重查找变得简单
// innodb中各个页之间通过双向链表连接，叶子节点中的数据通过单向链表连接的,效率更高

const M = 4
const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax
const LimitM2 = (M + 1) / 2

type Position *BPFullNode

type BPTree struct {
	keyMax int
	root   *BPFullNode
	ptr    *BPFullNode
}

type BPLeafNode struct {
	Next *BPFullNode
	data []int
}

// BPFullNode 叶子节点应该为Children为空，但leafNode中data不为空 Next一般不为空
type BPFullNode struct {
	KeyNum   int
	Key      []int
	isLeaf   bool
	Children []*BPFullNode
	leafNode *BPLeafNode
}

func MallocNewNode(isLeaf bool) *BPFullNode {
	var NewNode *BPFullNode
	if isLeaf == true {
		NewLeaf := MallocNewLeaf()
		NewNode = &BPFullNode{
			KeyNum:   0,
			Key:      make([]int, M+1), //申请M + 1是因为插入时可能暂时出现节点key大于M 的情况,待后期再分裂处理
			isLeaf:   isLeaf,
			Children: nil,
			leafNode: NewLeaf,
		}
	} else {
		NewNode = &BPFullNode{
			KeyNum:   0,
			Key:      make([]int, M+1),
			isLeaf:   isLeaf,
			Children: make([]*BPFullNode, M+1),
			leafNode: nil,
		}
	}
	for i, _ := range NewNode.Key {
		NewNode.Key[i] = IntMin
	}

	return NewNode
}
func MallocNewLeaf() *BPLeafNode {

	NewLeaf := BPLeafNode{
		Next: nil,
		data: make([]int, M+1),
	}
	for i, _ := range NewLeaf.data {
		NewLeaf.data[i] = i
	}
	return &NewLeaf
}

func (tree *BPTree) Initialize() {

	// 根结点
	T := MallocNewNode(true)
	tree.ptr = T
	tree.root = T
}

func FindMostLeft(P Position) Position {
	var Tmp Position
	Tmp = P
	if Tmp.isLeaf == true || Tmp == nil {
		return Tmp
	} else if Tmp.Children[0].isLeaf == true {
		return Tmp.Children[0]
	} else {
		for Tmp != nil && Tmp.Children[0].isLeaf != true {
			Tmp = Tmp.Children[0]
		}
	}
	return Tmp.Children[0]
}

func FindMostRight(P Position) Position {
	var Tmp Position
	Tmp = P

	if Tmp.isLeaf == true || Tmp == nil {
		return Tmp
	} else if Tmp.Children[Tmp.KeyNum-1].isLeaf == true {
		return Tmp.Children[Tmp.KeyNum-1]
	} else {
		for Tmp != nil && Tmp.Children[Tmp.KeyNum-1].isLeaf != true {
			Tmp = Tmp.Children[Tmp.KeyNum-1]
		}
	}

	return Tmp.Children[Tmp.KeyNum-1]
}

// FindSibling  寻找一个兄弟节点，其存储的关键字未满，若左右都满返回nil
func FindSibling(Parent Position, i int) Position {
	var Sibling Position
	var upperLimit int
	upperLimit = M
	Sibling = nil
	if i == 0 {
		if Parent.Children[1].KeyNum < upperLimit {

			Sibling = Parent.Children[1]
		}
	} else if Parent.Children[i-1].KeyNum < upperLimit {
		Sibling = Parent.Children[i-1]
	} else if i+1 < Parent.KeyNum && Parent.Children[i+1].KeyNum < upperLimit {
		Sibling = Parent.Children[i+1]
	}
	return Sibling
}

// FindSiblingKeyNumM2 查找兄弟节点，其关键字数大于M/2 ;没有返回nil j用来标识是左兄还是右兄
func FindSiblingKeyNumM2(Parent Position, i int, j *int) Position {
	var lowerLimit int
	var Sibling Position
	Sibling = nil

	lowerLimit = LimitM2

	if i == 0 {
		if Parent.Children[1].KeyNum > lowerLimit {
			Sibling = Parent.Children[1]
			*j = 1
		}
	} else {
		if Parent.Children[i-1].KeyNum > lowerLimit {
			Sibling = Parent.Children[i-1]
			*j = i - 1
		} else if i+1 < Parent.KeyNum && Parent.Children[i+1].KeyNum > lowerLimit {
			Sibling = Parent.Children[i+1]
			*j = i + 1
		}

	}
	return Sibling
}

// InsertElement 当要对X插入data的时候，i是X在Parent的位置，insertIndex是data要插入的位置，j可由查找得到
//
//	当要对Parent插入X节点的时候，posAtParent是要插入的位置，Key和j的值没有用
func (tree *BPTree) InsertElement(isData bool, Parent Position, X Position, Key int, posAtParent int, insertIndex int, data int) Position {

	var k int
	if isData {
		// 插入data
		k = X.KeyNum - 1
		for k >= insertIndex {
			X.Key[k+1] = X.Key[k]
			X.leafNode.data[k+1] = X.leafNode.data[k]
			k--
		}

		X.Key[insertIndex] = Key
		X.leafNode.data[insertIndex] = data
		if Parent != nil {
			Parent.Key[posAtParent] = X.Key[0] //可能min_key 已发生改变
		}

		X.KeyNum++

	} else {
		// 插入节点
		// 对树叶节点进行连接
		if X.isLeaf == true {
			if posAtParent > 0 {
				Parent.Children[posAtParent-1].leafNode.Next = X
			}
			X.leafNode.Next = Parent.Children[posAtParent]
			//更新叶子指针
			if X.Key[0] <= tree.ptr.Key[0] {
				tree.ptr = X
			}
		}

		k = Parent.KeyNum - 1
		for k >= posAtParent { //插入节点时key也要对应的插入
			Parent.Children[k+1] = Parent.Children[k]
			Parent.Key[k+1] = Parent.Key[k]
			k--
		}
		Parent.Key[posAtParent] = X.Key[0]
		Parent.Children[posAtParent] = X
		Parent.KeyNum++
	}

	return X
}

// RemoveElement 两个参数X posAtParent 有些重复 posAtParent可以通过X的最小关键字查找得到
func (tree *BPTree) RemoveElement(isData bool, Parent Position, X Position, posAtParent int, deleteIndex int) Position {

	var k, keyNum int

	if isData {
		keyNum = X.KeyNum
		// 删除key
		k = deleteIndex + 1
		for k < keyNum {
			X.Key[k-1] = X.Key[k]
			X.leafNode.data[k-1] = X.leafNode.data[k]
			k++
		}

		X.Key[keyNum-1] = IntMin
		X.leafNode.data[keyNum-1] = IntMin
		Parent.Key[posAtParent] = X.Key[0]
		X.KeyNum--
	} else {
		// 删除节点
		// 修改树叶节点的链接
		if X.isLeaf == true && posAtParent > 0 {
			Parent.Children[posAtParent-1].leafNode.Next = Parent.Children[posAtParent+1]
		}

		keyNum = Parent.KeyNum
		k = posAtParent + 1
		for k < keyNum {
			Parent.Children[k-1] = Parent.Children[k]
			Parent.Key[k-1] = Parent.Key[k]
			k++
		}

		if X.Key[0] == tree.ptr.Key[0] { // refresh ptr
			tree.ptr = Parent.Children[0]
		}
		Parent.Children[Parent.KeyNum-1] = nil
		Parent.Key[Parent.KeyNum-1] = IntMin

		Parent.KeyNum--

	}
	return X
}

// MoveElement Src和Dst是两个相邻的节点，posAtParent是Src在Parent中的位置；
// 将Src的元素移动到Dst中 ,eNum是移动元素的个数
func (tree *BPTree) MoveElement(src Position, dst Position, parent Position, posAtParent int, eNum int) Position {
	var TmpKey, data int
	var Child Position
	var j int
	var srcInFront bool

	srcInFront = false

	if src.Key[0] < dst.Key[0] {
		srcInFront = true
	}
	j = 0
	// 节点Src在Dst前面
	if srcInFront {
		if src.isLeaf == false {
			for j < eNum {
				Child = src.Children[src.KeyNum-1]
				tree.RemoveElement(false, src, Child, src.KeyNum-1, IntMin)      //每删除一个节点keyNum也自动减少1 队尾删
				tree.InsertElement(false, dst, Child, IntMin, 0, IntMin, IntMin) //队头加
				j++
			}
		} else {
			for j < eNum {
				TmpKey = src.Key[src.KeyNum-1]
				data = src.leafNode.data[src.KeyNum-1]
				tree.RemoveElement(true, parent, src, posAtParent, src.KeyNum-1)
				tree.InsertElement(true, parent, dst, TmpKey, posAtParent+1, 0, data)
				j++
			}

		}

		parent.Key[posAtParent+1] = dst.Key[0]
		// 将树叶节点重新连接
		if src.KeyNum > 0 {
			FindMostRight(src).leafNode.Next = FindMostLeft(dst) //似乎不需要重连，src的最右本身就是dst最左的上一元素
		} else {
			if src.isLeaf == true {
				parent.Children[posAtParent-1].leafNode.Next = dst
			}
			//  此种情况肯定是merge merge中有实现先移动再删除操作
			//tree.RemoveElement(false ,parent.parent，parent ,parentIndex,Int_Min )
		}
	} else {
		if src.isLeaf == false {
			for j < eNum {
				Child = src.Children[0]
				tree.RemoveElement(false, src, Child, 0, IntMin) //从src的队头删
				tree.InsertElement(false, dst, Child, IntMin, dst.KeyNum, IntMin, IntMin)
				j++
			}

		} else {
			for j < eNum {
				TmpKey = src.Key[0]
				data = src.leafNode.data[0]
				tree.RemoveElement(true, parent, src, posAtParent, 0)
				tree.InsertElement(true, parent, dst, TmpKey, posAtParent-1, dst.KeyNum, data)
				j++
			}

		}

		parent.Key[posAtParent] = src.Key[0]
		if src.KeyNum > 0 {
			FindMostRight(dst).leafNode.Next = FindMostLeft(src)
		} else {
			if src.isLeaf == true {
				dst.leafNode.Next = src.leafNode.Next
			}
			//tree.RemoveElement(false ,parent.parent，parent ,parentIndex,Int_Min )
		}
	}

	return parent
}

// SplitNode i为节点X的位置
func (tree *BPTree) SplitNode(Parent Position, beSplitedNode Position, i int) Position {
	var j, k, keyNum int
	var NewNode Position

	if beSplitedNode.isLeaf == true {
		NewNode = MallocNewNode(true)
	} else {
		NewNode = MallocNewNode(false)
	}

	k = 0
	j = beSplitedNode.KeyNum / 2
	keyNum = beSplitedNode.KeyNum
	for j < keyNum {
		if beSplitedNode.isLeaf == false { //Internal node
			NewNode.Children[k] = beSplitedNode.Children[j]
			beSplitedNode.Children[j] = nil
		} else {
			NewNode.leafNode.data[k] = beSplitedNode.leafNode.data[j]
			beSplitedNode.leafNode.data[j] = IntMin
		}
		NewNode.Key[k] = beSplitedNode.Key[j]
		beSplitedNode.Key[j] = IntMin
		NewNode.KeyNum++
		beSplitedNode.KeyNum--
		j++
		k++
	}

	if Parent != nil {
		tree.InsertElement(false, Parent, NewNode, IntMin, i+1, IntMin, IntMin)
		// parent > limit 时的递归split recurvie中实现
	} else {
		// 如果是X是根，那么创建新的根并返回
		Parent = MallocNewNode(false)
		tree.InsertElement(false, Parent, beSplitedNode, IntMin, 0, IntMin, IntMin)
		tree.InsertElement(false, Parent, NewNode, IntMin, 1, IntMin, IntMin)
		tree.root = Parent
		return Parent
	}

	return beSplitedNode
	// 为什么返回一个X一个Parent?
}

// MergeNode 合并节点,X少于M/2关键字，S有大于或等于M/2个关键字
func (tree *BPTree) MergeNode(Parent Position, X Position, S Position, i int) Position {
	var Limit int

	// S的关键字数目大于M/2
	if S.KeyNum > LimitM2 {
		// 从S中移动一个元素到X中
		tree.MoveElement(S, X, Parent, i, 1)
	} else {
		// 将X全部元素移动到S中，并把X删除
		Limit = X.KeyNum
		tree.MoveElement(X, S, Parent, i, Limit) //最多时S恰好MAX MoveElement已考虑了parent.key的索引更新
		tree.RemoveElement(false, Parent, X, i, IntMin)
	}
	return Parent
}

func (tree *BPTree) RecursiveInsert(beInsertedElement Position, Key int, posAtParent int, Parent Position, data int) (Position, bool) {
	var InsertIndex, upperLimit int
	var Sibling Position
	var result bool
	result = true
	// 查找分支
	InsertIndex = 0
	for InsertIndex < beInsertedElement.KeyNum && Key >= beInsertedElement.Key[InsertIndex] {
		// 重复值不插入
		if Key == beInsertedElement.Key[InsertIndex] {
			return beInsertedElement, false
		}
		InsertIndex++
	}
	//key必须大于被插入节点的最小元素，才能插入到此节点，故需回退一步
	if InsertIndex != 0 && beInsertedElement.isLeaf == false {
		InsertIndex--
	}

	// 树叶
	if beInsertedElement.isLeaf == true {
		beInsertedElement = tree.InsertElement(true, Parent, beInsertedElement, Key, posAtParent, InsertIndex, data) //返回叶子节点
		// 内部节点
	} else {
		beInsertedElement.Children[InsertIndex], result = tree.RecursiveInsert(beInsertedElement.Children[InsertIndex], Key, InsertIndex, beInsertedElement, data)
		//更新parent发生在split时
	}
	// 调整节点

	upperLimit = M
	if beInsertedElement.KeyNum > upperLimit {
		// 根
		if Parent == nil {
			// 分裂节点
			beInsertedElement = tree.SplitNode(Parent, beInsertedElement, posAtParent)
		} else {
			Sibling = FindSibling(Parent, posAtParent)
			if Sibling != nil {
				// 将T的一个元素（Key或者Child）移动的Sibing中
				tree.MoveElement(beInsertedElement, Sibling, Parent, posAtParent, 1)
			} else {
				// 分裂节点
				beInsertedElement = tree.SplitNode(Parent, beInsertedElement, posAtParent)
			}
		}

	}
	if Parent != nil {
		Parent.Key[posAtParent] = beInsertedElement.Key[0]
	}

	return beInsertedElement, result
}

// Insert 插入
func (tree *BPTree) Insert(Key int, data int) (Position, bool) {
	return tree.RecursiveInsert(tree.root, Key, 0, nil, data) //从根节点开始插入
}

func (tree *BPTree) RecursiveRemove(beRemovedElement Position, Key int, posAtParent int, Parent Position) (Position, bool) {

	var deleteIndex int
	var Sibling Position
	var NeedAdjust bool
	var result bool
	Sibling = nil

	// 查找分支   TODO查找函数可以在参考这里的代码 或者实现一个递归遍历
	deleteIndex = 0
	for deleteIndex < beRemovedElement.KeyNum && Key >= beRemovedElement.Key[deleteIndex] {
		if Key == beRemovedElement.Key[deleteIndex] {
			break
		}
		deleteIndex++
	}

	if beRemovedElement.isLeaf == true {
		// 没找到
		if Key != beRemovedElement.Key[deleteIndex] || deleteIndex == beRemovedElement.KeyNum {
			return beRemovedElement, false
		}
	} else {
		if deleteIndex == beRemovedElement.KeyNum || Key < beRemovedElement.Key[deleteIndex] {
			deleteIndex-- //准备到下层节点查找
		}
	}

	// 树叶
	if beRemovedElement.isLeaf == true {
		beRemovedElement = tree.RemoveElement(true, Parent, beRemovedElement, posAtParent, deleteIndex)
	} else {
		beRemovedElement.Children[deleteIndex], result = tree.RecursiveRemove(beRemovedElement.Children[deleteIndex], Key, deleteIndex, beRemovedElement)
	}

	NeedAdjust = false
	//有子节点的root节点，当keyNum小于2时
	if Parent == nil && beRemovedElement.isLeaf == false && beRemovedElement.KeyNum < 2 {
		NeedAdjust = true
	} else if Parent != nil && beRemovedElement.isLeaf == false && beRemovedElement.KeyNum < LimitM2 {
		// 除根外，所有中间节点的儿子数不在[M/2]到M之间时。(符号[]表示向上取整)
		NeedAdjust = true
	} else if Parent != nil && beRemovedElement.isLeaf == true && beRemovedElement.KeyNum < LimitM2 {
		// （非根）树叶中关键字的个数不在[M/2]到M之间时
		NeedAdjust = true
	}

	// 调整节点
	if NeedAdjust {
		// 根
		if Parent == nil {
			if beRemovedElement.isLeaf == false && beRemovedElement.KeyNum < 2 {
				//树根的更新操作 树高度减一
				beRemovedElement = beRemovedElement.Children[0]
				tree.root = beRemovedElement.Children[0]
				return beRemovedElement, true
			}

		} else {
			// 查找兄弟节点，其关键字数目大于M/2
			Sibling = FindSiblingKeyNumM2(Parent, posAtParent, &deleteIndex)
			if Sibling != nil {
				tree.MoveElement(Sibling, beRemovedElement, Parent, deleteIndex, 1)
			} else {
				if posAtParent == 0 {
					Sibling = Parent.Children[1]
				} else {
					Sibling = Parent.Children[posAtParent-1]
				}

				Parent = tree.MergeNode(Parent, beRemovedElement, Sibling, posAtParent)
				//Merge中已考虑空节点的删除
				beRemovedElement = Parent.Children[posAtParent]
			}
		}

	}

	return beRemovedElement, result
}

// Remove 删除
func (tree *BPTree) Remove(Key int) (Position, bool) {
	return tree.RecursiveRemove(tree.root, Key, 0, nil)
}

func (tree *BPTree) FindData(key int) (int, bool) {
	var currentNode *BPFullNode
	var index int
	currentNode = tree.root
	for index < currentNode.KeyNum {
		index = 0
		for key >= currentNode.Key[index] && index < currentNode.KeyNum {
			index++
		}
		if index == 0 {
			return IntMin, false
		} else {
			index--
			if currentNode.isLeaf == false {
				currentNode = currentNode.Children[index]
			} else {
				if key == currentNode.Key[index] {
					return currentNode.leafNode.data[index], true
				} else {
					return IntMin, false
				}
			}
		}

	}
	return IntMin, false
}
