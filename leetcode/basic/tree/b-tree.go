package tree

import "errors"

// todo
// b-tree B树, 一棵多路平衡查找树
// M阶 代表一个树节点最多有多少个查找路径，当M=2则是2叉树,M=3则是3叉
// 1. 每个结点关键字个数大于等于ceil(m/2)-1小于等于M-1,根结点可以只有1个关键字;
// 2. 所有节点关键字是按递增次序排列,并遵循左小右大原则,每个关键字的左子树中的所有关键字都小于它,而右子树中的所有关键字都大于它;
// 3. 所有叶子结点都位于同一层,或者说根结点到每个叶子结点的长度都相同;

var (
	mErr        = errors.New("m数值非法")
	notFoundErr = errors.New("未找到对应的数据")
)

type bTree struct {
	root  *bNode
	count int
	m     int // 代表最多有m个查找路径
}

type bNode struct {
	parent *bNode              // 父节点指针
	keyNum int                 // 关键字个数
	key    []int               // 关键字向量
	child  []*bNode            // 子树的指针向量
	data   map[int]interface{} // 存储的数据,key为bNode.key
}

// 新建一个空的BTree;
// 参数m代表m路b-tree
func NewBTree(m int) (*bTree, error) {
	if m < 2 {
		return nil, mErr
	}
	return &bTree{m: m}, nil
}

func (t *bTree) Size() int {
	return t.count
}

func (t *bTree) IsEmpty() bool {
	return t.count == 0
}

func (t *bTree) Search(key int) (interface{}, error) {
	node, index := t.traverse(t.root, key)
	if node == nil {
		return nil, notFoundErr
	}
	return node.data[index], nil
}

// 插入数据
func (t *bTree) Insert(key int) {
	panic("")
}

func (t *bTree) Remove() {
	panic("")
}
