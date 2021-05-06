package leetcode35

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	b.Run("searchInsert1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			searchInsert([]int{1, 3, 5, 6}, 7)
		}
	})

	b.Run("searchInsert2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			searchInsert1([]int{1, 3, 5, 6}, 7)
		}
	})
}
