package _33

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:       1,
			sc:       1,
			newColor: 2,
		}, [][]int{
			{2, 2, 2},
			{2, 2, 0},
			{2, 0, 1},
		}},

		{
			"test2", args{
				image: [][]int{
					{0, 0, 0},
					{0, 0, 0},
				},
				sr:       0,
				sc:       0,
				newColor: 2,
			},
			[][]int{
				{2, 2, 2},
				{2, 2, 2},
			},
		},

		{
			"test3", args{
				image: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0}},
				sr:       1,
				sc:       1,
				newColor: 2,
			}, [][]int{
				{0, 0, 0},
				{0, 2, 0},
				{0, 0, 0},
			},
		},

		{
			"test4", args{
				image: [][]int{
					{0, 0, 0},
					{0, 1, 1},
				},
				sr:       1,
				sc:       1,
				newColor: 1,
			}, [][]int{
				{0, 0, 0},
				{0, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
