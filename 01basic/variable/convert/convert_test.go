package convert

import "testing"

func TestToBin(t *testing.T) {
	// 走到这个目录下执行test

	type args struct {
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{"test1", args{n: 5}, "101"},
		{"test2", args{n: 13}, "1101"},
		{"test3", args{n: 0}, ""},
		{"test4", args{n: 6}, "110"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ToBin(tt.args.n); gotResult != tt.wantResult {
				t.Errorf("ToBin() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
