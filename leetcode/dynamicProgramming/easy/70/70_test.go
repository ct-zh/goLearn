package _70

import "testing"

type args struct {
	n int
}

type tests struct {
	name string
	args args
	want int
}

func getTests() []tests {
	return []tests{
		{"t1", args{n: 2}, 2},
		{"t2", args{n: 3}, 3},
	}
}

func Test_climbStairs(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs2(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs2(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_climbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range getTests() {
			climbStairs(tt.args.n)
		}
	}
}

func Benchmark_climbStairs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range getTests() {
			climbStairs2(tt.args.n)
		}
	}
}
