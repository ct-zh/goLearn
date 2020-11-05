package graph

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// 图 - 实现接口
type Graph interface {
	V() int
	E() int
	AddEdge(v int, w int) error
	HasEdge(v int, w int) (bool, error)
	Print()
}

// 稠密图 - 邻接矩阵
const TypeDense = 1

// 稀疏图 - 邻接表
const TypeSparse = 2

// 带权稠密图
const TypeWeightDense = 3

// 带权稀疏图
const TypeWeightSparse = 4

// 通过文件创建一个图的实例
// filename: 文件路径
// gType: 图的类型，见上面的type常量
func CreateGraphByFile(filename string, gType int) Graph {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	if fileInfo.IsDir() {
		panic("读取文件失败")
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(f)

	var g Graph
	line := 1
	for scan.Scan() {
		split := strings.Split(scan.Text(), " ")
		if line == 1 { // 第一行是节点个数n和边的个数m
			n, err := strconv.Atoi(split[0])
			if n <= 0 {
				panic("行数非法")
			}
			if err != nil {
				panic(err)
			}
			if gType == TypeDense {
				g, err = NewDenseGraph(n, false)
				if err != nil {
					panic(err)
				}
			} else if gType == TypeSparse {
				g, err = NewSpareGraph(n, false)
				if err != nil {
					panic(err)
				}
			} else {
				panic("暂不支持的图的类型")
			}
		} else { // 其他行的内容则是各个边的信息
			split1, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			split2, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			err = g.AddEdge(split1, split2)
			if err != nil {
				panic(err)
			}
		}

		line++
	}

	return g
}

// 图的迭代器
type iterator struct {
	g     *Graph
	v     int
	index int
}

// g: 迭代的图； v: 迭代开始的起点
func NewIterator(g *Graph, v int) *iterator {
	if v < 0 || v >= (*g).V() {
		panic("迭代的起点v非法")
	}
	return &iterator{
		g:     g,
		v:     v,
		index: -1, // 索引从-1开始, 因为每次遍历都需要调用一次next()
	}
}

// 返回图g中与v相连接的第一个顶点
func (i *iterator) Begin() int {
	i.index = -1 // 每次index都是从-1开始
	return i.Next()
}

// 返回图g中与v相连接的下一个顶点
func (i *iterator) Next() int {
	for i.index += 1; i.index < (*i.g).V(); i.index++ {
		hasEdge, err := (*i.g).HasEdge(i.v, i.index)
		if err != nil {
			panic(err)
		}
		if hasEdge {
			return i.index
		}
	}
	return -1
}

// 查看是否已经迭代完了图G中与顶点v相连接的所有顶点
func (i *iterator) End() bool {
	return i.index >= (*i.g).V()
}
