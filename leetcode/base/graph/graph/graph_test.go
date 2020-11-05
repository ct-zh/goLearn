package graph

import (
	"fmt"
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

func TestCreateWeightGraphByFile(t *testing.T) {
	tests := []struct {
		filename string
		gType    int
	}{
		{filename: "./testWeightG1.txt", gType: TypeWeightDense},
	}
	for _, tt := range tests {
		g := CreateWeightGraphByFile(tt.filename, tt.gType)
		g.Print()
	}
}

func TestNewIterator(t *testing.T) {
	tests := []struct {
		filename string
		gType    int
		v        int
	}{
		{filename: "./testG1.txt", gType: TypeSparse, v: 0},
	}
	for key, tt := range tests {
		g := CreateGraphByFile(tt.filename, tt.gType)
		iter := NewIterator(&g, tt.v)
		for i := iter.Begin(); !iter.End(); i = iter.Next() {
			fmt.Printf("Key: %d  I: %d\n", key, i)
		}
	}
}
