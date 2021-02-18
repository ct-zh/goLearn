package graph

type componentGraph struct {
	g       *Graph // 需要遍历的图的指针
	visited []bool // 记录哪些节点已经访问过了
	ccount  int    // 连通分量 componentCount
	id      []int  // 判读节点是否相连接
}

// 深度优先遍历得到图的连通分量
// 连通分量：一个图中存在多少个未相连的节点集合。比如如果图中所有节点都相邻，则该图的连通分量就是为0。
func NewComponent(g *Graph) *componentGraph {
	// 初始化结构， visited的大小为图的节点数量
	c := componentGraph{
		g:       g,
		visited: make([]bool, (*g).V()),
		ccount:  0,
		id:      make([]int, (*g).V()),
	}

	// 先默认所有节点数量都未访问过，设置为false
	for i := 0; i < (*c.g).V(); i++ {
		c.visited[i] = false
	}

	// 开始遍历
	for i := 0; i < (*c.g).V(); i++ {
		// i1 在dfs中遍历了从i1开始所有相邻的节点，将visited置为true，此时连通分量为0
		// 后面如果发现i2的visited为false,则说明与i2相邻的所有节点都与i1不相邻，继续dfs，此时连通分量为1
		// 以此类推，可以将一个图中所有未连通的节点都遍历出来
		if !c.visited[i] {
			c.dfs(i)
			c.ccount++
		}
	}
	return &c
}

// 深度优先遍历： deep first search
// todo: 使用深度优先遍历检查有向图是否存在环
func (c *componentGraph) dfs(v int) {
	c.visited[v] = true

	// 记录图的id值,等于当前的连通分量
	// 如果节点的连通分量相同，说明这两个节点在同一个dfs内被循环到
	// 也就证明这两个点相邻
	c.id[v] = c.ccount

	i := NewIterator(c.g, v)
	for n := i.Begin(); !i.End(); n = i.Next() {
		if !c.visited[n] {
			c.dfs(n)
		}
	}
}

// 获取图的连通分量
func (c *componentGraph) CCount() int {
	return c.ccount
}

// 两个节点是否相连
func (c *componentGraph) IsConnected(v int, w int) bool {
	if v < 0 || v >= (*c.g).V() {
		panic("参数非法")
	}
	if w < 0 || w >= (*c.g).V() {
		panic("参数非法")
	}
	return c.id[v] == c.id[w]
}
