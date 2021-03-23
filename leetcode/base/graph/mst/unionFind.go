package mst

import "errors"

// quick union 并查集
type unionFind struct {
	parent []int
	count  int
	rank   []int
}

func NewUnionFind(n int) *unionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &unionFind{
		parent: parent,
		count:  n,
		rank:   rank,
	}
}

// 查找过程，查找元素p所对应的集合编号
// o(h)的复杂度，h为树的高度
func (u *unionFind) Find(p int) (int, error) {
	if p < 0 || p > u.count {
		return 0, errors.New("参数非法")
	}

	for {
		if u.parent[p] == p {
			break
		}
		p = u.parent[p]
		u.parent[p] = u.parent[u.parent[p]]
	}

	return p, nil
}

// 查找两点是否相邻
// 复杂度o(h) h为树的高度
func (u *unionFind) IsConnected(p int, q int) (bool, error) {
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
func (u *unionFind) UnionElements(p int, q int) error {
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
	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = qRoot
	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else {
		u.parent[qRoot] = pRoot
		u.rank[pRoot] += 1	// 因为只有这种情况会增加rank的深度
	}

	return nil
}
