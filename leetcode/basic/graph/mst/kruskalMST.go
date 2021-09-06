package mst

import (
	"github.com/ct-zh/goLearn/leetcode/basic/graph/weightGraph"
)

// 补充： Vyssotsky's Algorithm 算法
// 将边逐渐添加到生成树中，一旦形成环，删除环中权值最大的边

// kruskal 算法效率低于prim算法，但是实现简单，可用于比较小的图

// kruskal 寻找最小生成树 流程：
// 1. 将所有边按照权值从小到大排序（最小堆，这里的复杂度就已经达到O(ELogE)了）
// 2. 从最小堆中推出最小边，如果该边不会导致树形成一个环，那该边必定在最小生成树上
// 3. 环的判断：将生成树写入并查集中，将新边与生成树做union操作，如果root相同，说明会形成一个环
// 4. 重复推出数据，直到生成树的边的数量等于V - 1 (结点数量-1)
// 复杂度 排序的复杂度是O(ELogE), 推出来数据并判断的复杂度是O(ELogV) 如果没提前break掉的话
// 所以总的复杂度是O(ELogE)+O(ELogV)
type kruskalMST struct {
	mst       []weightGraph.Edge // 最小生成树
	mstWeight float64
}

func NewKruskalMST(g *weightGraph.WeightGraph) *kruskalMST {
	k := &kruskalMST{
		mst:       make([]weightGraph.Edge, (*g).E()),
		mstWeight: 0.0,
	}

	// 使用堆排序，将所有边写入最小堆中
	pq := NewMinHeap((*g).E())
	for i := 0; i < (*g).V(); i++ {
		adj := weightGraph.NewWeightIter(g, i)
		for e := adj.Begin(); !adj.End(); e = adj.Next() {

			// 因为每条边连接两个顶点，这个循环又是按照点来遍历的
			// 如果全部insert，则每个边会出现两次，整个堆的大小为2E
			// 所以我们只写入V < W 的边   (无向图的问题)
			if e.V() < e.W() {
				pq.Insert(e)
			}
		}
	}

	uf := NewUnionFind((*g).V())
	for {
		// 最小生成树的长度大于等于 V - 1 了，则直接跳出循环
		if pq.IsEmpty() || len(k.mst) >= (*g).V()-1 {
			break
		}

		e := pq.ExtractMin()

		// 判断是否成环
		isConn, err := uf.IsConnected(e.V(), e.W())
		if err != nil {
			panic(err)
		}
		if isConn {
			continue
		}

		// 将e写入最小生成树
		k.mst = append(k.mst, e)

		err = uf.UnionElements(e.V(), e.W())
		if err != nil {
			panic(err)
		}
	}

	k.mstWeight = k.mst[0].Wt().(float64)
	for i := 1; i < len(k.mst); i++ {
		k.mstWeight += k.mst[i].Wt().(float64)
	}

	return k
}

func (k *kruskalMST) MesEdges() []weightGraph.Edge {
	return k.mst
}

func (k *kruskalMST) Result() float64 {
	return k.mstWeight
}
