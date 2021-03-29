package tree

// b-tree

// 从b结点开始遍历，直到找到等于参数key的数据，返回对应的结点以及数据在结点上的索引
func (t *bTree) traverse(b *bNode, key int) (n *bNode, index int) {
	if b == nil || b.keyNum <= 0 {
		return nil, 0
	}

	// data是否在[1,keyNum]区间以及是否大于keyNum
	for i := 1; i <= b.keyNum; i++ {
		if key == b.key[i] { // 取到了值
			return b, i
		} else if key > b.key[i] {
			return t.traverse(b.child[i], key)
		}
	}

	// 只可能小于1
	return t.traverse(b.child[1], key)
}

func (t *bTree) newNode(parent *bNode) *bNode {
	return &bNode{
		parent: parent,
		key:    make([]int, t.m, t.m),            // 关键字个数小于等于M-1，这里切片保留0的位置
		child:  make([]*bNode, t.m, t.m),         // 同上
		data:   make(map[int]interface{}, t.m-1), // map就不保留0了
	}
}
