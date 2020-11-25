package unionFind

import "errors"

// quick union 并查集2
// 跟第一个并查集不同的地方是：
// 使用一个数组构建一棵指向父节点的树，parent[i]表示第i个元素所指向的父节点
type uf2 struct {
	parent []int
	count  int
}

func NewUf2(n int) *uf2 {
	parent := make([]int, n)
	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &uf2{
		parent: parent,
		count:  n,
	}
}

// 查找过程，查找元素p所对应的集合编号
// o(h)的复杂度，h为树的高度
func (u *uf2) Find(p int) (int, error) {
	if p < 0 || p > u.count {
		return 0, errors.New("参数非法")
	}
	for {
		if u.parent[p] == p {
			break
		}
		p = u.parent[p]
	}

	return p, nil
}

// 查找两点是否相邻
// 复杂度o(h) h为树的高度
func (u *uf2) IsConnected(p int, q int) (bool, error) {
	pRoot, err := u.Find(p)
	if err != nil {
		return false, err
	}
	qRoot, err := u.Find(q)
	if err != nil {
		return false, err
	}

	return pRoot == qRoot, nil
}

// 合并p q 两个元素所属的集合
// o(h) 复杂度 h为树的高度
func (u *uf2) UnionElements(p int, q int) error {
	pRoot, err := u.Find(p)
	if err != nil {
		return err
	}
	qRoot, err := u.Find(q)
	if err != nil {
		return err
	}

	if pRoot == qRoot {
		return nil
	}

	// 直接将一个树的根节点指向另外一个树的根节点
	u.parent[pRoot] = qRoot

	return nil
}
