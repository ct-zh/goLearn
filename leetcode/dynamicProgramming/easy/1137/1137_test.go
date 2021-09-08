package _137

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
		{"test1", args{n: 4}, 4},
		{"test2", args{n: 25}, 1389537},
	}
}

func Test_tribonacci(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
