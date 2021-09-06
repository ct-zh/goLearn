package mst

import "github.com/ct-zh/goLearn/leetcode/basic/graph/weightGraph"

// prim 的时间复杂度为O(ELogV) 相比lazy prim要更快些

// prim 流程：
// 设置一个最小索引堆
// 把节点0作为切分的一部分，剩下的节点作为切分的另外一部分，
// 将节点0的所有横切边的另一个顶点与对应的权值写入最小索引堆
// 获取到最小索引堆里的最小值，将其对应的顶点加入切分，例如这里是节点7
// 同样的，将节点7对应的横切边里另外一个顶点与对应的权值写入最小索引堆
// *注意，这里如果遇到已经在堆的顶点，例如顶点6,存在[0,6],[6,7]两条横切边
// 如果[6,7]这条边权值小于[0,6]，则更新掉索引堆里6对应的权值
// 以上操作重复循环，直到所有点都划入新的切分中，即可得到最小生成树
type PrimMst struct {
	g         *weightGraph.WeightGraph // 图
	ipq       *indexMinHeap            // 最小索引堆
	edgeTo    []weightGraph.Edge       // 存储和每个节点相邻的最短的横切边
	marked    []bool                   // 表示该点是否被标记了，根据true和false将图划分为两个切分
	mst       []weightGraph.Edge       // 最小生成树 v-1个边
	mstWeight float64                  // 最小生成树的权值
}

// prim 算法辅助函数
// 遍历节点v所有横切边写入队列pq
func (l *PrimMst) visit(v int) {
	if l.marked[v] { // 如果该点已经判断过了，直接返回
		return
	}
	l.marked[v] = true // 代表已经遍历过了
	adj := weightGraph.NewWeightIter(l.g, v)

	for e := adj.Begin(); !adj.End(); e = adj.Next() {
		// 遍历v的所有边，如果存在边i还未标记，则写入队列pq中
		w := e.Other(v)
		if !l.marked[w] {
			if l.edgeTo[w].Wt().(float64) == 0 {
				l.ipq.Insert(w, e)
				l.edgeTo[w] = e
			} else if e.Wt().(float64) < l.edgeTo[w].Wt().(float64) {
				l.edgeTo[w] = e
				l.ipq.Change(w, e)
			}
		}
	}
}

func NewPrimMst(g *weightGraph.WeightGraph) *PrimMst {
	l := PrimMst{
		g:      g,
		ipq:    NewIndexMinHeap((*g).E()), // 最差的情况下所有数据都要进入堆中
		marked: make([]bool, (*g).V()),
		edgeTo: make([]weightGraph.Edge, (*g).V()),
		mst:    []weightGraph.Edge{},
	}

	for i := 0; i < (*g).V(); i++ { // marked默认全部填充false
		l.marked[i] = false
		l.edgeTo[i] = weightGraph.Edge{}
	}

	// lazy prim开始
	// 遍历所有的结点，记到marked
	l.visit(0) // 从 0 开始寻找mst
	for {
		if l.ipq.IsEmpty() {
			break
		}

		v := l.ipq.ExtractMinIndex() // 获取优先队列里最小的边
		if l.edgeTo[v].Wt().(float64) <= 0 {
			panic("该横切边不存在")
		}
		l.mst = append(l.mst, l.edgeTo[v])
		l.visit(v)
	}

	// 计算最小生成树的权值
	l.mstWeight = l.mst[0].Wt().(float64)
	for i := 0; i < len(l.mst); i++ {
		l.mstWeight += l.mst[i].Wt().(float64)
	}

	return &l
}

func (l *PrimMst) MstEdges() []weightGraph.Edge {
	return l.mst
}

func (l *PrimMst) Weight() weightGraph.Weight {
	return l.mstWeight
}
