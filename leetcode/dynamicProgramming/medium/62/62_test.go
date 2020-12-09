package _62

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{
			m: 3,
			n: 7,
		}, 28},
		{"t2", args{
			m: 3,
			n: 2,
		}, 3},
		{"t3", args{
			m: 7,
			n: 3,
		}, 28},
		{"t4", args{
			m: 3,
			n: 3,
		}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths1(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
