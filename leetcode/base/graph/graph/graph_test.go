package graph

import (
	"testing"
)

func TestCreateGraphByFile(t *testing.T) {
	tests := []struct {
		filename string
		gType    int
	}{
		{filename: "./testG1.txt", gType: TypeSparse},
	}
	for _, tt := range tests {
		g := CreateGraphByFile(tt.filename, tt.gType)
		g.Print()
	}
}
