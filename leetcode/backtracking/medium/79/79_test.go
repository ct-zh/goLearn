package _79

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{board: board, word:  "ABCCED",}, true},
		{"t2", args{board: board, word:  "SEE",}, true},
		{"t2", args{board: board, word:  "ABCB",}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist2(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}