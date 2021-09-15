package _1maxSubArr

import "testing"

type args struct {
	nums []int
}

type tests struct {
	name string
	args args
	want int
}

func getTests() []tests {
	return []tests{
		{
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
		{
			args: args{[]int{1}},
			want: 1,
		},
		{
			args: args{[]int{0}},
			want: 0,
		},
	}
}

// dp
func Test_maxSubArray(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 分治法
func Test_maxSubArray2(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray2(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 暴力算法
func Test_maxSubArrayForceBetter(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayForceBetter(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// kadane算法
func Test_maxSubArrayKadane(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArrayKadane(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// dp
func Benchmark_maxSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range getTests() {
			maxSubArray(tt.args.nums)
		}
	}
}

// 分治
func Benchmark_maxSubArray2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range getTests() {
			maxSubArray2(tt.args.nums)
		}
	}
}

// 暴力算法
func Benchmark_maxSubArrayForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range getTests() {
			maxSubArrayForceBetter(tt.args.nums)
		}
	}
}

// kadane算法
func Benchmark_maxSubArrayKadane(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range getTests() {
			maxSubArrayKadane(tt.args.nums)
		}
	}
}
