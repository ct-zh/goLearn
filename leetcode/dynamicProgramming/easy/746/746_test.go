package _46

import "testing"

type args struct {
	cost []int
}

type tests struct {
	name string
	args args
	want int
}

func getTests() []tests {
	return []tests{
		{"t1", args{cost: []int{10, 15, 20}}, 15},
		{"t2", args{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
	}
}

func Test_minCostClimbingStairs(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostClimbingStairsDynamic(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairsDynamic(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
