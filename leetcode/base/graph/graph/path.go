package graph

import (
	"fmt"
	queue2 "github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/queue"
)

// 从s开始寻路
type path struct {
	g       *Graph
	s       int
	visited []bool // 记录该点是否已经访问过了
	from    []int  // 记录某个点是从哪个点过来的，默认值为-1
	ord     []int  // 记录从s到某个节点的最短路径
}

func NewPath(g *Graph, s int) *path {
	if s < 0 || s >= (*g).V() {
		panic("path:s参数非法")
	}

	p := &path{
		g:       g,
		s:       s,
		visited: make([]bool, (*g).V()),
		from:    make([]int, (*g).V()),
		ord:     make([]int, (*g).V()),
	}

	for i := 0; i < (*p.g).V(); i++ {
		p.visited[i] = false
		p.from[i] = -1
		p.ord[i] = -1
	}

	// todo: 执行深度优先遍历还是广度优先遍历，不同遍历有不同的功能

	// 深度寻路
	//p.dfs(s)

	// 广度找最短路径
	p.bfs(s)

	return p
}

// 深度优先遍历： deep first search
func (p *path) dfs(v int) {
	p.visited[v] = true

	i := NewIterator(p.g, v)
	for n := i.Begin(); !i.End(); n = i.Next() {
		if !p.visited[n] {
			p.from[n] = v // 访问n节点的路径 是从v节点过来的
			p.dfs(n)
		}
	}
}

// 广度优先遍历 breadth first search
func (p *path) bfs(s int) {
	queue := queue2.Queue{}
	queue.Push(s)

	p.visited[s] = true
	p.ord[s] = 0 // v点到v点的距离为0

	for {
		if queue.Len() <= 0 {
			break
		}

		// 先取出队列中的第一个元素
		v := queue.Pop()

		// 获取到该元素邻接的节点，压入队列，并记录相关信息
		adj := NewIterator(p.g, v.(int))
		for i := adj.Begin(); !adj.End(); i = adj.Next() {
			if !p.visited[i] {
				queue.Push(i)
				p.visited[i] = true           // 节点i与s是连通的
				p.from[i] = v.(int)           // 节点i是来源于v
				p.ord[i] = p.ord[v.(int)] + 1 // 节点i与s的距离是 节点v与s的距离+1
			}
		}

	}

}

// 节点s到节点w是否存在路径
func (p *path) HasPath(w int) bool {
	if w < 0 || w >= (*p.g).V() {
		panic("参数非法")
	}
	return p.visited[w]
}

// 获取节点s到节点w的路径
func (p *path) getPath(w int) []int {
	var s []int

	// 从w往回推到s
	m := w
	for {
		if m == -1 {
			break
		}
		s = append(s, m)
		m = p.from[m]
	}

	fmt.Printf("%+v\n", s)

	// 转置数组，即是s到w的路径
	var back []int
	for i := len(s) - 1; i >= 0; i-- {
		back = append(back, s[i])
	}

	return back
}

// 打印s到w的路径
func (p *path) ShowPath(w int) {
	list := p.getPath(w)
	for _, v := range list {
		fmt.Println(v)
	}
}

// 获取s到w的最短路径
func (p *path) Length(w int) int {
	return p.ord[w]
}
