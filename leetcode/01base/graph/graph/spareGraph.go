package graph

import (
	"errors"
	"fmt"
)

// 稀疏图 - 邻接表
//  遍历复杂度 O(v + e)	节点+边
type sparseGraph struct {
	n        int     // 节点数量
	m        int     // 边的数量
	directed bool    // 是否为有向图
	g        [][]int // 图的具体数据
}

// 获取一个稀疏图-邻接表
func NewSpareGraph(n int, directed bool) (*sparseGraph, error) {
	if n < 0 {
		return &sparseGraph{}, errors.New("参数非法")
	}

	var data [][]int
	for i := 0; i < n; i++ {
		data = append(data, []int{})
	}

	return &sparseGraph{
		n:        n,
		m:        0,
		directed: directed,
		g:        data,
	}, nil
}

// 获取稀疏图的节点数量
func (s *sparseGraph) V() int {
	return s.n
}

// 获取稀疏图的边数量
func (s *sparseGraph) E() int {
	return s.m
}

// 稀疏图 - 增加一条边
// 因为判断是否相邻(HasEdge)的时间复杂度为O(n),而增加一条边的时间复杂度为O(v),
// 所以为了性能考虑,新增边就不做两点是否相邻的判断了
func (s *sparseGraph) AddEdge(v int, w int) error {
	if v < 0 || w < 0 {
		return errors.New("参数非法")
	}
	if v > s.n || w > s.n {
		return errors.New("参数不能大于节点数量")
	}

	// 邻接表存在平行边与自环边
	// 平行边：两个边连接了两个相同的节点
	// 自环边：这个边是一个节点连接它自身
	s.g[v] = append(s.g[v], w)
	if !s.directed && v != w {
		s.g[w] = append(s.g[w], v)
	}

	return nil
}

// 稀疏图 - 判断两点是否相邻
func (s *sparseGraph) HasEdge(v int, w int) (bool, error) {
	if v < 0 || w < 0 {
		return false, errors.New("参数非法")
	}
	if v > s.n || w > s.n {
		return false, errors.New("参数不能大于节点数量")
	}

	for i := 0; i < len(s.g[v]); i++ {
		if s.g[v][i] == w {
			return true, nil
		}
	}
	return false, nil
}

func (s *sparseGraph) Print() {
	for key, v1 := range s.g {
		fmt.Printf("Line:%d", key)
		for _, v2 := range v1 {
			fmt.Print(" ", v2, " ")
		}
		fmt.Println()
	}
}
