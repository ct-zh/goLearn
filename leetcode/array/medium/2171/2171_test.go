package l2171

import "testing"

type args struct {
	beans []int
}

var tests = []struct {
	name string
	args args
	want int64
}{
	{
		name: "test1",
		args: args{
			beans: []int{4, 1, 6, 5},
		},
		want: 4,
	},
	{
		name: "test2",
		args: args{
			beans: []int{2, 10, 3, 2},
		},
		want: 7,
	},
}

func Test_minimumRemoval(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRemoval(tt.args.beans); got != tt.want {
				t.Errorf("minimumRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumRemovalv2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRemovalV2(tt.args.beans); got != tt.want {
				t.Errorf("minimumRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}
