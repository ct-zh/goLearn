package main

import "testing"

func TestValidSubstringCount(t *testing.T) {
	tests := []struct {
		word1    string
		word2    string
		expected int64
	}{
		{"bcca", "abc", 1},
		{"abcabc", "abc", 10},
		{"abcabc", "aaabc", 0},
	}

	for _, test := range tests {
		t.Run(test.word1+"_"+test.word2, func(t *testing.T) {
			result := validSubstringCount(test.word1, test.word2)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
