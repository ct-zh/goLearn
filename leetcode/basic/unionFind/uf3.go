package unionFind

import "errors"

// quick union 并查集3
// 在uf2的基础上增加了union方法的 size优化
type uf3 struct {
	parent []int // parent[i]表示第i个元素所指向的父节点
	count  int
	sz     []int // sz[i]表示以i为根的集合中元素个数
}

func NewUf3(n int) *uf3 {
	parent := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		sz[i] = 1
	}
	return &uf3{
		parent: parent,
		count:  n,
		sz:     sz,
	}
}

// 查找过程，查找元素p所对应的集合编号
// o(h)的复杂度，h为树的高度
func (u *uf3) Find(p int) (int, error) {
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
func (u *uf3) IsConnected(p int, q int) (bool, error) {
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
func (u *uf3) UnionElements(p int, q int) error {
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

	// 根据两个元素所在树的元素个数不同判断合并方向
	// 将元素个数少的集合合并到元素个数多的集合上
	if u.sz[pRoot] > u.sz[qRoot] {
		u.parent[qRoot] = pRoot
		u.sz[qRoot] += u.sz[pRoot]
	} else {
		u.parent[pRoot] = qRoot
		u.sz[qRoot] += u.sz[pRoot]
	}

	return nil
}
