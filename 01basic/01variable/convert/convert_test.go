package convert

import (
	"fmt"
	"testing"
)

func TestToBin(t *testing.T) {
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

func TestIntToString(t *testing.T) {
	tests := []struct {
		i int
		r string
	}{
		{5, "5"},
	}
	for _, tt := range tests {
		result := IntToString(tt.i)
		if result != tt.r {
			t.Errorf("func test err! int: %d answer: %s result:%s", tt.i, tt.r, result)
		}
	}
}

func TestIntToString2(t *testing.T) {
	tests := []struct {
		i int
		r string
	}{
		{5, "5"},
		{6, "6"},
	}
	for _, tt := range tests {
		result := IntToString2(tt.i)
		if result != tt.r {
			t.Errorf("func test err! int: %+v answer: %+v result:%+v", tt.i, tt.r, result)
		}
	}
}

func TestStringSplit(t *testing.T) {
	tests := []struct {
		s string
		r []string
	}{
		{"abc", []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		result := StringSplit(tt.s)
		fmt.Sprintf("%+v\n", result)
		//for k, item := range tt.r {
		//	if item != result[k] {
		//		t.Errorf("func test err! int: %+v answer: %+v result:%+v", tt.s, tt.r, result)
		//	}
		//}
	}
}
