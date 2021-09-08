package _09

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
		{"test1", args{n: 2}, 1},
		{"test2", args{n: 3}, 2},
		{"test3", args{n: 4}, 3},
		{"test3", args{n: 0}, 0},
	}
}

func Test_fib(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib2(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib2(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range getTests() {
			fib(tt.args.n)
		}
	}
}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range getTests() {
			fib2(tt.args.n)
		}
	}
}
