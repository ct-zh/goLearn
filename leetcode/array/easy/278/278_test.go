package _78

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{n: 5}, 4},
		{"test2", args{n: 1}, 1},
		{"test2", args{n: 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badVersion = tt.want
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
