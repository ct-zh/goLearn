package _98_213

import "testing"

type args struct {
	nums []int
}

type tests struct {
	name string
	args args
	want int
}

func get198Tests() []tests {
	return []tests{
		{"t1", args{[]int{1, 2, 3, 1}}, 4},
		{"t2", args{[]int{2, 7, 9, 3, 1}}, 12},
		{"t3", args{[]int{0}}, 0},
		{"t4", args{[]int{2, 1, 1, 2}}, 4},
	}
}

func get213Tests() []tests {
	return []tests{
		{"t1", args{[]int{2, 3, 2}}, 3},
		{"t2", args{[]int{1, 2, 3, 1}}, 4},
		{"t3", args{[]int{0}}, 0},
		{"t4", args{[]int{2, 7, 9, 3, 1}}, 11},
		{"t5", args{[]int{1, 3, 1, 3, 100}}, 103},
	}
}

func Test_rob(t *testing.T) {
	for _, tt := range get198Tests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob2(t *testing.T) {
	for _, tt := range get213Tests() {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob2(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range get213Tests() {
			rob(tt.args.nums)
		}
	}
}

func BenchmarkRob2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range get213Tests() {
			rob2(tt.args.nums)
		}
	}
}
