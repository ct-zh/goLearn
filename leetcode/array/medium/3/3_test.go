package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s      string
		result int
	}{
		{s: "abcabcbb", result: 3},
		{s: "bbbbb", result: 1},
		{s: "pwwkew", result: 3},
	}
	for key, tt := range tests {
		result := lengthOfLongestSubstring(tt.s)
		if tt.result != result {
			t.Errorf("[%d]结果错误：result: %+v answer: %+v \n",
				key, result, tt.result)
		}
	}
}
