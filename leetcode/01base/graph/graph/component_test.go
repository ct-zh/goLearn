package graph

import (
	"fmt"
	"testing"
)

func TestNewComponent(t *testing.T) {
	tests := []struct {
		filename string
		gType    int
	}{
		{filename: "./testG1.txt", gType: TypeSparse},
	}
	for _, tt := range tests {
		g := CreateGraphByFile(tt.filename, tt.gType)
		component := NewComponent(&g)
		fmt.Println("图  testG1 的连通分量是", component.CCount())
	}
}
