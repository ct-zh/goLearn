package weightGraph

import "testing"

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
