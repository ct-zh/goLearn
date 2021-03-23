package unionFind

// 并查集通用接口
type UnionFind interface {
	Find(p int) (int, error)
	IsConnected(p int, q int) (bool, error)
	UnionElements(p int, q int) error
}

// 并查集
// 并查集到底是什么 ？
// 并查集是一种树型的数据结构，用于处理一些不相交的集合的 **合并与查询** 问题
type unionFind1 struct {
	Count int
	Id    map[int]int // 第一版Union-Find本质就是一个数组
}

// 构造函数  n: 元素个数
func NewUnionFind(n int) *unionFind1 {
	id := make(map[int]int) // 结构：key: id  Value:指向的id

	for i := 0; i < n; i++ { // 每个id初始都是指向自己
		id[i] = i
	}
	return &unionFind1{
		Count: n,
		Id:    id,
	}
}

// 查找 时间复杂度: O(1)
func (u *unionFind1) Find(p int) (int, error) {
	return u.Id[p], nil
}

// 判断p q 两点是否相连
func (u *unionFind1) IsConnected(p int, q int) (bool, error) {
	f1, err := u.Find(p)
	if err != nil {
		return false, err
	}
	f2, err := u.Find(q)
	if err != nil {
		return false, err
	}
	return f1 == f2, nil
}

// 连接pq两点 时间复杂度: O(n)
func (u *unionFind1) UnionElements(p int, q int) error {
	pId, err := u.Find(p)
	if err != nil {
		return err
	}
	qId, err := u.Find(q)
	if err != nil {
		return err
	}
	if pId == qId {
		return nil
	}

	for i := 0; i < u.Count; i++ {
		if u.Id[i] == pId {
			u.Id[i] = qId
		}
	}
	return nil
}
