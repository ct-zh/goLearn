package graph

import (
	"fmt"
	"testing"
)

// go14 test -v -run TestNewLazyPrimMst lazyPrimMST_test.go lazyPrimMST.go
func TestNewLazyPrimMst(t *testing.T) {
	tests := []struct {
		filename  string
		graphType int
	}{
		{filename: "./testWeightG1.txt", graphType: TypeWeightSparse},
	}
	for _, tt := range tests {
		g := CreateWeightGraphByFile(tt.filename, tt.graphType)

		lazy := NewLazyPrimMst(&g)
		fmt.Printf("Graph: \"%s\"   Weight: %v \nMst: %v\n",
			tt.filename, lazy.Weight(), lazy.MstEdges())
	}
}
