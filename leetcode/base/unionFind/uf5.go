package unionFind

import "errors"

// quick union 并查集4
// 路径压缩
type uf5 struct {
	parent []int // parent[i]表示第i个元素所指向的父节点
	count  int
	rank   []int // rank[i]表示以i为根的集合所表示的树的层数
	// 在后续的代码中, 我们并不会维护rank的语意, 也就是rank的值在路径压缩的过程中, 有可能不在是树的层数值
	// 这也是我们的rank不叫height或者depth的原因, 他只是作为比较的一个标准
	// 事实上，这正是我们将这个变量叫做rank而不是叫诸如depth或者height的原因。
	// 因为这个rank只是我们做的一个标志当前节点排名的一个数字，
	// 当我们引入了路径压缩以后，维护这个深度的真实值相对困难一些，
	// 而且实践告诉我们，我们其实不需要真正维持这个值是真实的深度值，
	// 我们依然可以以这个rank值作为后续union过程的参考。
	// 因为根据我们的路径压缩的过程，rank高的节点虽然被抬了上来，
	// 但是整体上，我们的并查集从任意一个叶子节点出发向根节点前进，
	// 依然是一个rank逐渐增高的过程。
	// 也就是说，这个rank值在经过路径压缩以后，虽然不是真正的深度值，
	// 但仍然可以胜任，作为union时的参考。
}

func NewUf5(n int) *uf5 {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &uf5{
		parent: parent,
		count:  n,
		rank:   rank,
	}
}

// 查找过程，查找元素p所对应的集合编号
// o(h)的复杂度，h为树的高度
func (u *uf5) Find(p int) (int, error) {
	if p < 0 || p > u.count {
		return 0, errors.New("参数非法")
	}

	for {
		if u.parent[p] == p {
			break
		}
		p = u.parent[p]
		// path compression 1
		// 路径压缩1 将父节点指向父节点的父节点
		u.parent[p] = u.parent[u.parent[p]]
	}

	// path compression 2, 递归算法  todo 未验证
	//var err error
	//if p != u.parent[p] {
	//	u.parent[p], err = u.Find(u.parent[p])
	//	if err != nil {
	//		return 0, err
	//	}
	//	return u.parent[p], nil
	//}

	return p, nil
}

// 查找两点是否相邻
// 复杂度o(h) h为树的高度
func (u *uf5) IsConnected(p int, q int) (bool, error) {
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
func (u *uf5) UnionElements(p int, q int) error {
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
		u.rank[pRoot] += 1 // 因为只有这种情况会增加rank的深度
	}

	return nil
}
