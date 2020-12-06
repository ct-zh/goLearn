package _447

import "testing"

func Test_numberOfBoomerangs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{points: [][]int{
			{0, 0},
			{1, 0},
			{2, 0},
		}}, 2},
		{"t2", args{points: [][]int{
			{1, 1},
			{2, 2},
			{3, 3},
		}}, 2},
		{"t3", args{points: [][]int{
			{1, 1},
		}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBoomerangs(tt.args.points); got != tt.want {
				t.Errorf("numberOfBoomerangs() = %v, want %v", got, tt.want)
			}
		})
	}
}
