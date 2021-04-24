package offer10I

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{n: 2}, want: 1},
		{name: "", args: args{n: 40}, want: 102334155},
		{name: "", args: args{n: 41}, want: 165580141},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{n: 2}, want: 1},
		{name: "", args: args{n: 40}, want: 102334155},
		{name: "", args: args{n: 41}, want: 165580141},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib2(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
