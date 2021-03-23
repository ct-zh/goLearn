package knapsack

import "testing"

func Test_knapsack(t *testing.T) {
	w := []int{
		0: 1,
		1: 2,
		2: 3,
	}
	v := []int{
		0: 6,
		1: 10,
		2: 12,
	}
	type args struct {
		w []int
		v []int
		C int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{
			w: w,
			v: v,
			C: 5,
		}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knapsack(tt.args.w, tt.args.v, tt.args.C); got != tt.want {
				t.Errorf("knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}
