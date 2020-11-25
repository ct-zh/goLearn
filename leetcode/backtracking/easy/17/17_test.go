package _17

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{"test1", args{digits: "23"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := letterCombinations(tt.args.digits)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("letterCombinations() = %v, want %v", gotRes, tt.wantRes)
			}
			fmt.Println(gotRes)
		})
	}
}
