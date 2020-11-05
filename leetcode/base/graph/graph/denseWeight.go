package graph

import (
	"errors"
	"fmt"
)

// 带权稠密图 - 邻接矩阵
// 遍历复杂度 O(v^2)	节点的平方
type denseWeight struct {
	n        int                   // 节点数量
	m        int                   // 边的数量
	directed bool                  // 是否为有向图
	g        map[int]map[int]*Edge // 图的具体数据
}

// 获取一个稠密图-邻接矩阵
func NewDenseWeight(n int, directed bool) (*denseWeight, error) {
	if n < 0 {
		return &denseWeight{}, errors.New("参数非法")
	}

	// 矩阵内容全部填充nil
	data := make(map[int]map[int]*Edge)
	for i := 0; i < n; i++ {
		data[i] = make(map[int]*Edge)
		for j := 0; j < n; j++ {
			data[i][j] = nil
		}
	}

	return &denseWeight{
		n:        n,
		m:        0,
		directed: directed,
		g:        data,
	}, nil
}

// 获取稠密图的节点数量
func (d *denseWeight) V() int {
	return d.n
}

// 获取稠密图的边数量
func (d *denseWeight) E() int {
	return d.m
}

// 给稠密图添加一条边
func (d *denseWeight) AddEdge(v int, w int, weight Weight) error {
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

	// 带权图如果存在了这条边，则删掉旧的边再存储一条新边
	if has {
		d.g[v][w] = nil
		if !d.directed {
			d.g[w][v] = nil
		}
	}

	d.g[v][w] = &Edge{
		a:      v,
		b:      w,
		weight: weight,
	}
	if !d.directed {
		d.g[w][v] = &Edge{
			a:      w,
			b:      v,
			weight: weight,
		}
	}
	d.m++

	return nil
}

// 稠密图 - 检测是否有从 v 到 w 的边
func (d *denseWeight) HasEdge(v int, w int) (bool, error) {
	if v < 0 || w < 0 {
		return false, errors.New("参数非法")
	}
	if v > d.n || w > d.n {
		return false, errors.New("参数不能大于节点数量")
	}

	return d.g[v][w] != nil, nil
}

func (d *denseWeight) Print() {
	for k, v := range d.g {
		fmt.Printf("Line: %d  ", k)
		for _, v2 := range v {
			fmt.Print(" ", v2, " ")
		}
		fmt.Println()
	}
}
