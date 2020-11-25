package _344

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s      []byte
		result []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{
			s:      []byte{'h', 'e', 'l', 'l', 'o'},
			result: []byte{'o', 'l', 'l', 'e', 'h'}}},
		{name: "2", args: args{
			s:      []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			result: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			for i := 0; i < len(tt.args.s); i++ {
				if tt.args.s[i] != tt.args.result[i] {
					t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.args.result)
					break
				}
			}
		})
	}
}
