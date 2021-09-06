package shortest

import (
	"github.com/ct-zh/goLearn/leetcode/basic/graph/weightGraph"
	"github.com/ct-zh/goLearn/leetcode/basic/stack"
)

// 负权边和Bellman-Ford算法
// 拥有负权环的图，没有最短路径（因为可以不停地在负权环里面转，每转一次权值就小一些）
// 负权环有可能在两个节点内产生,如 [1,2,4] [2,1,-5], 在节点1，2组成的环中每转一次权值小1

// bellman ford算法
// 前提：图中不能有负权环  bellman ford 可以判断图中是否有负权环
// 复杂度 O(EV)

// 基本思想：
// 如果一个图没有负权环
// 从一点到另外一点的最短路径，最多经过所有的V个顶点，有V-1条边
// 否则，存在顶点经过了两次，即存在负权环

// 流程
// 对所有的点进行一次松弛操作，找到从原点开始，经过一条边最短的路径
// 对所有的点再进行松弛操作，找到从原点为止，经过两条边的最短路径

// 对一个点的一次松弛操作，就是找到经过这个点的另外一条路径，多一条边，权值更小
// 如果一个图没有负权环，从一点到另外一点的最短路径，最多经过所有的V个顶点，有V-1条边
// 那就对所有点进行V-1次松弛操作

// 对所有点进行V-1次松弛操作，理论上就找到了从原点到其他所有点的最短路径
// 如果还可以松弛，说明原图中存在负权环

// bellman ford 还存在一种优化结构：
// 使用队列数据结构 queue-based bellman ford算法
type bellmanFord struct {
	g                *weightGraph.WeightGraph
	s                int                // 起点
	distTo           []float64          // s到某个节点的最小权值
	from             []weightGraph.Edge // s到某个节点从哪个边过来的
	hasNegativeCycle bool               // 是否存在负权环
}

// 再进行一次松弛操作，判断是否存在负权环
func (b *bellmanFord) detectNegativeCycle() bool {
	for i := 0; i < (*b.g).V(); i++ {
		adj := weightGraph.NewWeightIter(b.g, i)
		for e := adj.Begin(); !adj.End(); e = adj.Next() {
			if b.from[e.W()].IsNull() ||
				b.distTo[e.V()]+e.Wt().(float64) < b.distTo[e.W()] {
				return true
			}
		}
	}

	return false
}

func NewBellmanFord(g *weightGraph.WeightGraph, s int) *bellmanFord {
	b := &bellmanFord{
		g:                g,
		s:                s,
		distTo:           make([]float64, (*g).V()),
		from:             make([]weightGraph.Edge, (*g).V()),
		hasNegativeCycle: false,
	}

	for i := 0; i < (*g).V(); i++ {
		b.from = append(b.from, weightGraph.Edge{})
	}

	// bellman ford
	b.distTo[s] = 0
	for pass := 1; pass < (*b.g).V(); pass++ {

		// Relaxation 松弛操作
		for i := 0; i < (*b.g).V(); i++ {
			adj := weightGraph.NewWeightIter(b.g, i)
			for e := adj.Begin(); !adj.End(); e = adj.Next() {

				// 如果w点还未到达过，或者
				// 绕道到i这个节点 再去e的终点w 比 直接去w 距离还要小（负权边）
				if b.from[e.W()].IsNull() ||
					b.distTo[e.V()]+e.Wt().(float64) < b.distTo[e.W()] {

					b.distTo[e.W()] = b.distTo[e.V()] + e.Wt().(float64)
					b.from[e.W()] = e
				}
			}
		}
	}

	// 做完所有的松弛操作后再来判断
	b.hasNegativeCycle = b.detectNegativeCycle()

	return b
}

func (b *bellmanFord) HasNegativeCycle() bool {
	return b.hasNegativeCycle
}

func (b *bellmanFord) ShortestPathTo(w int) float64 {
	if w < 0 || w >= (*b.g).V() {
		panic("参数非法")
	}
	if b.hasNegativeCycle {
		return -1
	}
	return b.distTo[w]
}

func (b *bellmanFord) HasPathTo(w int) bool {
	if w < 0 || w >= (*b.g).V() {
		panic("参数非法")
	}
	return !b.from[w].IsNull()
}

func (b *bellmanFord) ShortestPath(w int, vec *[]weightGraph.Edge) {
	if w < 0 || w >= (*b.g).V() {
		panic("参数非法")
	}
	if b.hasNegativeCycle {
		return
	}

	var s stack.Stack
	e := b.from[w]
	for {
		if e.V() == e.W() {
			break
		}
		s.Push(e)
		e = b.from[e.V()]
	}

	for {
		if s.IsEmpty() {
			break
		}

		*vec = append(*vec, e)
		s.Pop()
	}

}

func (b *bellmanFord) ShowPath(w int) {

}
