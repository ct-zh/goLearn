package shortest

import (
	"github.com/ct-zh/goLearn/leetcode/basic/graph/weightGraph"
	"github.com/ct-zh/goLearn/leetcode/basic/stack"
)

// 有权图的最短路径问题：
// 寻找最短路径树,单源最短路径问题
// 算法：
// 1. dijkstra算法,前提：无无负权边； 有向图无向图均可；复杂度 O(ElogV)
// 2. Bellman-Form算法,前提：无负权环；有向图；复杂度O(VE)
// 3. 利用拓扑排序,前提有向无环图DAG，复杂度O(V+E)

// 所有对最短路径算法
// Floyed算法，处理无负权环的图，复杂度O(V^3)

// 最长路径算法
// 最长路径不能有正权环
// 无权图的最长路径问题是指数级的
// 对于有权图，不能使用Dijkstra算法求最长路径问题
// 可以使用bellman-form算法

// 松弛操作Relaxation => 是最短路径求解的核心
// 当我们到达一个结点的时候我们就要尝试一下从这个结点到它的结点的路径长度
// 是否比之前求到的不经过这个结点而到其他结点所得到的路径要更短一些
// 如果更短一些，那么我们就要更新一下从原始结点到这个结点相应的信息

// dijkstra 单源最短路径算法
// 前提:图中不能有负权边
// 复杂度 O(E log(V))

// 从结点0出发，到其他结点的路径，其中最小的路径，必是结点0到该点的最小路径
// 如 [0,2,2],[0,1,5],[0,3,6]这三个路径中最小的路径为[0,2],消耗为2
// 那么0到2的最小路径必然是走[0,2]这条边

// 确定新的节点的最短路径后，进行relaxation
// 首先找到此时还没有找到最短路径的顶点中，现在存的那个最短的路径能抵达的顶点是谁

type dijkstra struct {
	g      *weightGraph.WeightGraph
	s      int                // 源，也就是起点
	distTo []float64          // 源点s到其他点的距离
	marked []bool             // 某点是否已经找到了最短路径
	from   []weightGraph.Edge // 最短路径
}

func NewDijkstra(g *weightGraph.WeightGraph, s int) *dijkstra {
	d := dijkstra{
		g:      g,
		s:      s,
		distTo: make([]float64, (*g).V()),
		marked: make([]bool, (*g).V()),
		from:   make([]weightGraph.Edge, (*g).V()),
	}

	for i := 0; i < (*g).V(); i++ {
		d.distTo[i] = 0
		d.marked[i] = false
		d.from = append(d.from, weightGraph.Edge{})
	}

	ipq := NewIndexMinHeap((*g).V())

	// dijkstra
	d.distTo[s] = 0
	d.marked[s] = true
	ipq.Insert(s, d.from[s])
	for {
		if ipq.IsEmpty() {
			break
		}

		// 找到当前节点中离源s最近的节点
		v := ipq.ExtractMinIndex()

		// distTo[v]就是s到v的最短距离
		d.marked[v] = true

		// Relaxation 松弛操作
		adj := weightGraph.NewWeightIter(g, v)
		for e := adj.Begin(); !adj.End(); adj.Next() {
			w := e.Other(v)

			if !d.marked[w] { // v的相邻端点w 是否还未找到从s到w的最短路径
				//  from[w]为空，说明w没有被访问过，需要进行松弛操作
				// 如果从s到v的距离加上[v,w]这条边的权值小于 从s到w的距离，说明也要进行松弛操作
				if d.from[w].IsNull() ||
					d.distTo[v]+e.Wt().(float64) < d.distTo[w] {

					d.distTo[w] = d.distTo[v] + e.Wt().(float64)
					d.from[w] = e // 从s到达w点是从e到达的

					// 如果最小索引堆中已经包含了w节点，则更新其权值
					if ipq.Contain(w) {
						ipq.Change(w, d.from[w])
					} else {
						ipq.Insert(w, d.from[w])
					}
				}
			}
		}

	}

	return &d
}

func (d *dijkstra) ShortestPathTo(w int) float64 {
	return d.distTo[w]
}

func (d *dijkstra) HasPathTo(w int) bool {
	return d.marked[w]
}

func (d *dijkstra) ShortestPath(w int, vec *[]weightGraph.Edge) {
	var s stack.Stack
	e := d.from[w]
	for {
		if e.V() == e.W() {
			break
		}
		s.Push(e)
		e = d.from[e.V()]
	}

	for {
		if s.IsEmpty() {
			break
		}

		*vec = append(*vec, e)
		s.Pop()
	}

}

func (d *dijkstra) ShowPath(w int) {

}
