package graph

import (
	"testing"
)

func TestNewPath(t *testing.T) {
	tests := []struct {
		filename  string
		gType     int
		startNode int
	}{
		{filename: "./testG1.txt", gType: TypeSparse, startNode: 0},
	}
	for _, tt := range tests {
		g := CreateGraphByFile(tt.filename, tt.gType)
		p := NewPath(&g, tt.startNode)
		p.ShowPath(3)
	}
}
