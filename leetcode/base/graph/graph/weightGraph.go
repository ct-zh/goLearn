package graph

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// 带权图
type WeightGraph interface {
	V() int
	E() int
	AddEdge(v int, w int, weight Weight) error
	HasEdge(v int, w int) (bool, error)
	GetEdge(v int, w int) (Edge, error)
	Print()
}

// 通过文件创建带权图
func CreateWeightGraphByFile(filename string, gType int) WeightGraph {
	// 读取文件
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var g WeightGraph
	line := 1
	scan := bufio.NewScanner(f)

	for scan.Scan() {
		split := strings.Split(scan.Text(), " ")

		if line == 1 {
			vInt, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			if vInt <= 0 {
				panic("结点数不能小于等于0")
			}

			if gType == TypeWeightDense {
				g, err = NewDenseWeight(vInt, false)
				if err != nil {
					panic(err)
				}
			} else if gType == TypeWeightSparse {
				g, err = NewSparseWeight(vInt, false)
				if err != nil {
					panic(err)
				}
			} else {
				panic("当前不支持该图的类型")
			}
		} else {
			v, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			w, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			weightFloat, err := strconv.ParseFloat(split[2], 64)
			if err != nil {
				panic(err)
			}

			err = g.AddEdge(v, w, weightFloat)
			if err != nil {
				panic(err)
			}
		}
		line++
	}

	return g
}

type weightIter struct {
	g     *WeightGraph
	v     int
	index int
}

func NewWeightIter(g *WeightGraph, v int) *weightIter {
	if v < 0 || v >= (*g).V() {
		panic("迭代的起点v非法")
	}
	return &weightIter{
		g:     g,
		v:     v,
		index: -1, // 索引从-1开始, 因为每次遍历都需要调用一次next()
	}
}

// 返回图g中与v相连接的第一个顶点
func (i *weightIter) Begin() Edge {
	i.index = -1 // 每次index都是从-1开始
	return i.Next()
}

// 返回图g中与v相连接的下一个顶点
func (i *weightIter) Next() Edge {
	for i.index += 1; i.index < (*i.g).V(); i.index++ {
		hasEdge, err := (*i.g).HasEdge(i.v, i.index)
		if err != nil {
			panic(err)
		}
		if hasEdge {
			edge, err := (*i.g).GetEdge(i.v, i.index)
			if err != nil {
				panic(err)
			}
			return edge
		}
	}
	return Edge{}
}

// 查看是否已经迭代完了图G中与顶点v相连接的所有顶点
func (i *weightIter) End() bool {
	return i.index >= (*i.g).V()
}
