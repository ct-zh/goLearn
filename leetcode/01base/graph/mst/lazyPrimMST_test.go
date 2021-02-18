package mst

import (
	"fmt"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/leetcode/base/graph/weightGraph"
	"testing"
)

// go14 test -v -run TestNewLazyPrimMst lazyPrimMST_test.go lazyPrimMST.go
func TestNewLazyPrimMst(t *testing.T) {
	tests := []struct {
		filename  string
		graphType int
	}{
		{filename: "./testWeightG1.txt", graphType: weightGraph.TypeWeightSparse},
	}
	for _, tt := range tests {
		g := weightGraph.CreateWeightGraphByFile(tt.filename, tt.graphType)

		lazy := NewLazyPrimMst(&g)
		fmt.Printf("Graph: \"%s\"   Weight: %v \nMst: %v\n",
			tt.filename, lazy.Weight(), lazy.MstEdges())
	}
}
