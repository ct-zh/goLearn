package graph

// 最小生成树问题：（只针对带权图连通图）
// 存在一棵树 连接了图里的所有节点,所有边的权值相加是最小的

// 切分 cut：把图的所有结点分为两个部分，称为一个切分
// 如果一个边的两个端点，属于切分不同的两边，这个边称为横切边 Crossing Edge
// 切分定理 cut property: 给定任意的切分，横切边中权值最小的边必然属于最小生成树

// lazy prim的时间复杂度为 O(ELogE)

// lazy prim 流程
// 把节点0作为切分的一部分，剩下的节点作为切分的另外一部分，
// 这样节点0的所有边都是横切边
// 将所有边写入最小堆中（lazy的原因：写入堆中的是所有边，而不是横切边。
// 因为后面可能出现情况：如节点0，1，2作为一个切分(横切边[0,1],[0,2])，最小堆中写入了边[1,2]，
// 这个边不是横切边，不应该加入堆中，但是lazyPrim则忽略了这点）
// 最小堆的顶点是这些边中权值最小的边，根据切分定理，此边必然属于最小生成树
// 假设从最小堆里取出的是[0,7]这条边，那么将节点7也划分到切分中
// 这样节点7对应的边也将压入堆中，再从最小堆中取出权值最小的边
type LazyPrimMst struct {
	g         *WeightGraph // 图
	pq        *minHeap     // 最小堆
	marked    []bool       // 表示该点是否被标记了，根据true和false将图划分为两个切分
	mst       []Edge       // 最小生成树 v-1个边
	mstWeight float64      // 最小生成树的权值
}

// Lazy prim 算法辅助函数
// 遍历节点v所有横切边写入队列pq
func (l *LazyPrimMst) visit(v int) {
	if l.marked[v] { // 如果该点已经判断过了，直接返回
		return
	}
	l.marked[v] = true // 代表已经遍历过了
	adj := NewWeightIter(l.g, v)

	for i := adj.Begin(); !adj.End(); i = adj.Next() {
		// 遍历v的所有边，如果存在边i还未标记，则写入队列pq中
		if !l.marked[i.Other(v)] {
			l.pq.Insert(i) // insert 的type是Edge
		}
	}
}

func NewLazyPrimMst(g *WeightGraph) *LazyPrimMst {
	l := LazyPrimMst{
		g:      g,
		pq:     NewMinHeap((*g).E()), // 最差的情况下所有数据都要进入堆中
		marked: make([]bool, (*g).V()),
		mst:    []Edge{},
	}

	for i := 0; i < (*g).V(); i++ { // marked默认全部填充false
		l.marked[i] = false
	}

	// lazy prim开始
	// 遍历所有的结点，记到marked
	l.visit(0) // 从 0 开始寻找mst
	for {
		if l.pq.IsEmpty() {
			break
		}

		e := l.pq.ExtractMin() // 获取优先队列里最小的边

		// 下面这个判断说明 这两个点在一个切分里面
		// 说明 取出来的边不是横切边，跳过
		if l.marked[e.V()] == l.marked[e.W()] {
			continue
		}

		// 将该边写入最小生成树
		l.mst = append(l.mst, e)

		// 找到该边属于另一个切分的节点，继续遍历
		if !l.marked[e.V()] {
			l.visit(e.V())
		} else {
			l.visit(e.W())
		}
	}

	// 计算最小生成树的权值
	l.mstWeight = l.mst[0].weight.(float64)
	for i := 0; i < len(l.mst); i++ {
		l.mstWeight += l.mst[i].weight.(float64)
	}

	return &l
}

func (l *LazyPrimMst) MstEdges() []Edge {
	return l.mst
}

func (l *LazyPrimMst) Weight() Weight {
	return l.mstWeight
}
