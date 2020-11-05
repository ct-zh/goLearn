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
