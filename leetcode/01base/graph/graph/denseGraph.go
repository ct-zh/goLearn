package graph

import (
	"errors"
	"fmt"
)

// 稠密图 - 邻接矩阵
// 遍历复杂度 O(v^2)	节点的平方
type denseGraph struct {
	n        int                  // 节点数量
	m        int                  // 边的数量
	directed bool                 // 是否为有向图
	g        map[int]map[int]bool //图的具体数据
}

// 获取一个稠密图-邻接矩阵
func NewDenseGraph(n int, directed bool) (*denseGraph, error) {
	if n < 0 {
		return &denseGraph{}, errors.New("参数非法")
	}

	// 矩阵内容全部填充false
	data := make(map[int]map[int]bool)
	for i := 0; i < n; i++ {
		data[i] = make(map[int]bool)
		for j := 0; j < n; j++ {
			data[i][j] = false
		}
	}

	return &denseGraph{
		n:        n,
		m:        0,
		directed: directed,
		g:        data,
	}, nil
}

// 获取稠密图的节点数量
func (d *denseGraph) V() int {
	return d.n
}

// 获取稠密图的边数量
func (d *denseGraph) E() int {
	return d.m
}

// 给稠密图添加一条边
func (d *denseGraph) AddEdge(v int, w int) error {
	if v < 0 || w < 0 {
		return errors.New("参数非法")
	}
	if v > d.n || w > d.n {
		return errors.New("参数不能大于节点数量")
	}

	has, err := d.HasEdge(v, w)
	if err != nil {
		return err
	}
	if has {
		return nil
	}

	d.g[v][w] = true
	if !d.directed {
		d.g[w][v] = true
	}

	d.m++

	return nil
}

// 稠密图 - 检测是否有从 v 到 w 的边
func (d *denseGraph) HasEdge(v int, w int) (bool, error) {
	if v < 0 || w < 0 {
		return false, errors.New("参数非法")
	}
	if v > d.n || w > d.n {
		return false, errors.New("参数不能大于节点数量")
	}

	return d.g[v][w], nil
}

func (d *denseGraph) Print() {
	for k, v := range d.g {
		fmt.Printf("Line: %d  ", k)
		for _, v2 := range v {
			fmt.Print(" ", v2, " ")
		}
		fmt.Println()
	}
}
